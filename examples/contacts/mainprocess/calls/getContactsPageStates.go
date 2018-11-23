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

func newGetContactsPageStatesCall(contactStorer storer.ContactStorer) *calling.MainProcess {
	return calling.NewMainProcess(
		callids.GetContactsPageStatesCallID,
		func(params []byte, call func([]byte)) {
			mainProcessGetContactsPageStates(params, call, contactStorer)
		},
	)
}

// The func is simple:
// 1.  Unmarshall the params. Call back any errors.
// 2.a Get the records from the store. Call back any errors.
// 2.b Sort the records.
// 2.c Collect the requested records beginning with SortedIndex.
// 3.  Call the renderer back with
//     * the same State,
//     * the same SortedIndex,
//     * the total record count,
//     * the requested records.
func mainProcessGetContactsPageStates(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &types.RendererToMainProcessGetContactsPageStatesParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContactsPageStates: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetContactsPageStatesParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2.a Get the records from the store.
	rr, err := contactStorer.GetContacts()
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContactsPageStates: contactStorer.GetContacts(): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetContactsPageStatesParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2.b Sort the records.
	tracker := make(map[string]*types.ContactRecord)
	keys := make([]string, 0, 5)
	for _, r := range rr {
		if _, ok := tracker[r.State]; !ok {
			tracker[r.State] = r
			keys = append(keys, r.State)
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
	txparams := &types.MainProcessToRendererGetContactsPageStatesParams{
		SortedIndex: rxparams.SortedIndex,
		State:       rxparams.State,
		Records:     requested,
		RecordCount: uint64(len(rr)),
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
