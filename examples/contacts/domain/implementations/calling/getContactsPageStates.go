package calling

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

// GetContactsPageStatesCallID is the GetContactsPageStates call id.
var GetContactsPageStatesCallID = nextCallID()

//RendererToMainProcessGetContactsPageStatesParams is the param the renderer sends to the main process for a page of unique states.
type RendererToMainProcessGetContactsPageStatesParams struct {
	SortedIndex uint64
	PageSize    uint64
	State       uint64
}

// MainProcessToRendererGetContactsPageStatesParams is the params the main process sends back to the renderer with a page of unique states
type MainProcessToRendererGetContactsPageStatesParams struct {
	SortedIndex  uint64
	Records      []*types.ContactRecord
	RecordCount  uint64
	State        uint64
	Error        bool
	ErrorMessage string
}

// newGetContactsPageStatesCall is the constructor for the GetContactsPageStates local procedure call.
// It should only receive the repos that are needed. In this case the contact repo.
// Param contactStorer storer.ContactStorer is the contact repo needed to update a contact record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetContactsPageStatesCall(contactStorer storer.ContactStorer, rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		GetContactsPageStatesCallID,
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetContactsPageStates(params, call, contactStorer)
		},
		rendererReceiveAndDispatchGetContactsPageStates,
		rendererSendPayload,
	)
}

// mainProcessReceiveGetContactsPageStates is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererGetContactsPageStatesParams
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
func mainProcessReceiveGetContactsPageStates(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessGetContactsPageStatesParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		log.Println("mainProcessReceiveGetContactsPageStates error is ", err.Error())
		message := fmt.Sprintf("mainProcessGetContactsPageStates: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactsPageStatesParams{
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
		message := fmt.Sprintf("mainProcessGetContactsPageStates: contactStorer.GetContacts(): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactsPageStatesParams{
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
	txparams := &MainProcessToRendererGetContactsPageStatesParams{
		SortedIndex: rxparams.SortedIndex,
		State:       rxparams.State,
		Records:     requested,
		RecordCount: uint64(len(rr)),
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

// rendererReceiveAndDispatchGetContactsPageStates is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a MainProcessToRendererGetContactsPageStatesParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageStatesParams.
// 2. Dispatch the *MainProcessToRendererGetContactsPageStatesParams.
func rendererReceiveAndDispatchGetContactsPageStates(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageStatesParams.
	rxparams := &MainProcessToRendererGetContactsPageStatesParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContactsPageStates defined about.
		rxparams = &MainProcessToRendererGetContactsPageStatesParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *MainProcessToRendererGetContactsPageStatesParams to the renderer panel callers that want to handle the GetContactsPageStates call backs.
	dispatch(rxparams)
}

/*

	See the renderer code at github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/panels/EditButton/EditContactSelectPanel/caller.go

*/
