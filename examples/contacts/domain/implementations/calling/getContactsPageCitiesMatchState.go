package calling

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

// GetContactsPageCitiesMatchStateCallID is the GetContactsPageCitiesMatchStateCall call id.
var GetContactsPageCitiesMatchStateCallID = nextCallID()

//RendererToMainProcessGetContactsPageCitiesMatchStateParams is the param the renderer sends to the main process for a page of unique states.
type RendererToMainProcessGetContactsPageCitiesMatchStateParams struct {
	SortedIndex uint64
	PageSize    uint64
	State       uint64
	StateMatch  string
}

// MainProcessToRendererGetContactsPageCitiesMatchStateParams is the params the main process sends back to the renderer with a page of unique states
type MainProcessToRendererGetContactsPageCitiesMatchStateParams struct {
	SortedIndex  uint64
	Records      []*types.ContactRecord
	RecordCount  uint64
	State        uint64
	Error        bool
	ErrorMessage string
}

// newGetContactsPageCitiesMatchStateCall is the constructor for the GetContactsPageCitiesMatchState local procedure call.
// It should only receive the repos that are needed. In this case the contact repo.
// Param contactStorer storer.ContactStorer is the contact repo needed to update a contact record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetContactsPageCitiesMatchStateCall(contactStorer storer.ContactStorer, rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		GetContactsPageCitiesMatchStateCallID,
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetContactsPageCitiesMatchState(params, call, contactStorer)
		},
		rendererReceiveAndDispatchGetContactsPageCitiesMatchState,
		rendererSendPayload,
	)
}

// mainProcessReceiveGetContactsPageCitiesMatchState is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererGetContactsPageCitiesMatchStateParams
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
func mainProcessReceiveGetContactsPageCitiesMatchState(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessGetContactsPageCitiesMatchStateParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		log.Println("mainProcessReceiveGetContactsPageCitiesMatchState error is ", err.Error())
		message := fmt.Sprintf("mainProcessGetContactsPageCitiesMatchState: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactsPageCitiesMatchStateParams{
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
		txparams := &MainProcessToRendererGetContactsPageCitiesMatchStateParams{
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
	txparams := &MainProcessToRendererGetContactsPageCitiesMatchStateParams{
		SortedIndex: rxparams.SortedIndex,
		State:       rxparams.State,
		Records:     requested,
		RecordCount: uint64(len(rr)),
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

// rendererReceiveAndDispatchGetContactsPageCitiesMatchState is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a MainProcessToRendererGetContactsPageCitiesMatchStateParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageCitiesMatchStateParams.
// 2. Dispatch the *MainProcessToRendererGetContactsPageCitiesMatchStateParams.
func rendererReceiveAndDispatchGetContactsPageCitiesMatchState(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageCitiesMatchStateParams.
	rxparams := &MainProcessToRendererGetContactsPageCitiesMatchStateParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContactsPageCitiesMatchState defined about.
		rxparams = &MainProcessToRendererGetContactsPageCitiesMatchStateParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *MainProcessToRendererGetContactsPageCitiesMatchStateParams to the renderer panel callers that want to handle the GetContactsPageCitiesMatchState call backs.
	dispatch(rxparams)
}

/*

	See the renderer code in github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/panels/EditButton/EditContactSelectPanel/caller.go

*/
