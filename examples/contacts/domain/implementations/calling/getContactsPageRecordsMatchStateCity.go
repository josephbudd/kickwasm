package calling

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

// GetContactsPageRecordsMatchStateCityCallID is the GetContactsPageRecordsMatchStateCity call id.
var GetContactsPageRecordsMatchStateCityCallID = nextCallID()

//RendererToMainProcessGetContactsPageRecordsMatchStateCityParams is the param the renderer sends to the main process for a page of unique states.
type RendererToMainProcessGetContactsPageRecordsMatchStateCityParams struct {
	SortedIndex uint64
	PageSize    uint64
	State       uint64
	StateMatch  string
	CityMatch   string
}

// MainProcessToRendererGetContactsPageRecordsMatchStateCityParams is the params the main process sends back to the renderer with a page of unique states
type MainProcessToRendererGetContactsPageRecordsMatchStateCityParams struct {
	SortedIndex  uint64
	Records      []*types.ContactRecord
	RecordCount  uint64
	State        uint64
	Error        bool
	ErrorMessage string
}

// newGetContactsPageRecordsMatchStateCityCall is the constructor for the GetContactsPageRecordsMatchStateCity local procedure call.
// It should only receive the repos that are needed. In this case the contact repo.
// Param contactStorer storer.ContactStorer is the contact repo needed to update a contact record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetContactsPageRecordsMatchStateCityCall(contactStorer storer.ContactStorer, rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		GetContactsPageRecordsMatchStateCityCallID,
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetContactsPageRecordsMatchStateCity(params, call, contactStorer)
		},
		rendererReceiveAndDispatchGetContactsPageRecordsMatchStateCity,
		rendererSendPayload,
	)
}

// mainProcessReceiveGetContactsPageRecordsMatchStateCity is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererGetContactsPageRecordsMatchStateCityParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param contactStorer is the contact repo.
// The func is simple:
// 1.  Unmarshall the params. Call back any errors.
// 2.a Get the records from the repo. Call back any errors.
// 2.b Sort the records.
// 2.c Collect the requested records.
// 3.  Call the renderer back with
//     * the same State
//     * the same SortedIndex
//     * the requested records.
func mainProcessReceiveGetContactsPageRecordsMatchStateCity(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessGetContactsPageRecordsMatchStateCityParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		log.Println("mainProcessReceiveGetContactsPageRecordsMatchStateCity error is ", err.Error())
		message := fmt.Sprintf("mainProcessGetContactsPageRecordsMatchStateCity: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	log.Printf("mainProcessReceiveGetContactsPageRecordsMatchStateCity:\n%s", string(params))
	// 2.a Get the records from the repo.
	rr, err := contactStorer.GetContacts()
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContactsPageRecordsMatchStateCity: contactStorer.GetContacts(): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{
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
	txparams := &MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{
		SortedIndex: rxparams.SortedIndex,
		State:       rxparams.State,
		Records:     requested,
		RecordCount: lrr,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

// rendererReceiveAndDispatchGetContactsPageRecordsMatchStateCity is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a MainProcessToRendererGetContactsPageRecordsMatchStateCityParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageRecordsMatchStateCityParams.
// 2. Dispatch the *MainProcessToRendererGetContactsPageRecordsMatchStateCityParams.
func rendererReceiveAndDispatchGetContactsPageRecordsMatchStateCity(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageRecordsMatchStateCityParams.
	rxparams := &MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContactsPageRecordsMatchStateCity defined about.
		rxparams = &MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *MainProcessToRendererGetContactsPageRecordsMatchStateCityParams to the renderer panel callers that want to handle the GetContactsPageRecordsMatchStateCity call backs.
	dispatch(rxparams)
}

/*

	See the renderer code in github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/panels/EditButton/EditContactSelectPanel/caller.go

*/
