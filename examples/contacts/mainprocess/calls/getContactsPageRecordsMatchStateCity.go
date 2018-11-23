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

func newGetContactsPageRecordsMatchStateCityCall(contactStorer storer.ContactStorer) *calling.MainProcess {
	return calling.NewMainProcess(
		callids.GetContactsPageStatesCallID,
		func(params []byte, call func([]byte)) {
			mainProcessGetContactsPageRecordsMatchStateCity(params, call, contactStorer)
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
func mainProcessGetContactsPageRecordsMatchStateCity(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &types.RendererToMainProcessGetContactsPageRecordsMatchStateCityParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContactsPageRecordsMatchStateCity: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{
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
		message := fmt.Sprintf("mainProcessGetContactsPageRecordsMatchStateCity: contactStorer.GetContacts(): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2.b Match and Sort the records.
	tracker := make(map[string][]*types.ContactRecord)
	keys := make([]string, 0, len(rr))
	for _, r := range rr {
		if r.State == rxparams.StateMatch && r.City == rxparams.CityMatch {
			if _, ok := tracker[r.State]; !ok {
				tracker[r.Name] = make([]*types.ContactRecord, 0, 10)
				keys = append(keys, r.Name)
			}
			tracker[r.Name] = append(tracker[r.Name], r)
		}
	}
	sort.Strings(keys)
	lrr := uint64(len(rr))
	sorted := make([]*types.ContactRecord, 0, lrr)
	for _, name := range keys {
		for _, r := range tracker[name] {
			sorted = append(sorted, r)
		}
	}
	// 2.c Collect the requested records.
	ls := uint64(len(sorted))
	start := rxparams.SortedIndex
	end := start + rxparams.PageSize
	if start >= ls {
		start = 0
		end = 0
	} else {
		if end > ls {
			end = ls
		}
	}
	requested := sorted[start:end]
	// 3. Call the renderer back with the correct records.
	txparams := &types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{
		SortedIndex: rxparams.SortedIndex,
		State:       rxparams.State,
		Records:     requested,
		RecordCount: lrr,
		StateMatch:  rxparams.StateMatch,
		CityMatch:   rxparams.CityMatch,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
