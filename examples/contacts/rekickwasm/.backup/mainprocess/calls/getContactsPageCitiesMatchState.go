package calls

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newGetContactsPageCitiesMatchStateCall(contactStorer storer.ContactStorer) *calling.MainProcess {
	return calling.NewMainProcess(
		callids.GetContactsPageCitiesMatchStateCallID,
		func(params []byte, call func([]byte)) {
			mainProcessGetContactsPageCitiesMatchState(params, call, contactStorer)
		},
	)
}

// The func is simple:
// 1.  Unmarshall the params. Call back any errors.
// 2.a Get the records from the repo. Call back any errors.
// 2.b Sort the records.
// 2.c Collect the requested records.
// 3.  Call the renderer back with
//     * the same State
//     * the same SortedIndex
//     * the requested records.
func mainProcessGetContactsPageCitiesMatchState(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &types.RendererToMainProcessGetContactsPageCitiesMatchStateParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContactsPageCitiesMatchState: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetContactsPageCitiesMatchStateParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2.a Get the records from the repo.
	rr, err := contactStorer.GetContacts()
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContactsPageCitiesMatchState: contactStorer.GetContacts(): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetContactsPageCitiesMatchStateParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2.b Match and Sort the records.
	tracker := make(map[string]*types.ContactRecord)
	keys := make([]string, 0, 5)
	for _, r := range rr {
		if r.State == rxparams.StateMatch {
			if _, ok := tracker[r.City]; !ok {
				tracker[r.City] = r
				keys = append(keys, r.City)
			}
		}
	}
	sort.Strings(keys)
	l := uint64(len(keys))
	sorted := make([]*types.ContactRecord, l, l)
	for i, state := range keys {
		sorted[i] = tracker[state]
	}
	// 2.c Collect the requested records.
	start := rxparams.SortedIndex
	end := start + rxparams.PageSize
	if start >= l {
		start = 0
		end = 0
	} else {
		if end > l {
			end = l
		}
	}
	requested := sorted[start:end]
	// 3. Call the renderer back with the correct records.
	txparams := &types.MainProcessToRendererGetContactsPageCitiesMatchStateParams{
		SortedIndex: rxparams.SortedIndex,
		State:       rxparams.State,
		Records:     requested,
		RecordCount: uint64(len(rr)),
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
