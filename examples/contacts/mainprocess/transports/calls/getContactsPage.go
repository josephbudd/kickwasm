package calls

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/behavior/repoi"
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
)

//RendererToMainProcessGetContactsPageParams is the param the renderer sends to the main process to update a contact.
type RendererToMainProcessGetContactsPageParams struct {
	SortedIndex uint64
	PageSize    uint64
	State       uint64
}

// MainProcessToRendererGetContactsPageParams is the params the main process sends back to the renderer after updating a contact.
type MainProcessToRendererGetContactsPageParams struct {
	SortedIndex  uint64
	Records      []*records.ContactRecord
	RecordCount  uint64
	State        uint64
	Error        bool
	ErrorMessage string
}

// newGetContactsPageLPC is the constructor for the GetContactsPage local procedure call.
// It should only receive the repos that are needed. In this case the contact repo.
// Param contactRepo repoi.ContactRepoI is the contact repo needed to update a contact record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetContactsPageLPC(contactRepo repoi.ContactRepoI, rendererSendPayload func(payload []byte) error) *LPC {
	return newLPC(
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetContactsPage(params, call, contactRepo)
		},
		rendererReceiveAndDispatchGetContactsPage,
		rendererSendPayload,
	)
}

// mainProcessReceiveGetContactsPage is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererGetContactsPageParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param contactRepo is the contact repo.
// The func is simple:
// 1.  Unmarshall the params. Call back any errors.
// 2.a Get the records from the repo. Call back any errors.
// 2.b Sort the records.
// 2.c Collect the requested records.
// 3.  Call the renderer back with
//     * the same State
//     * the same SortedIndex
//     * the requested records.
func mainProcessReceiveGetContactsPage(params []byte, callBackToRenderer func(params []byte), contactRepo repoi.ContactRepoI) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessGetContactsPageParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		log.Println("mainProcessReceiveGetContactsPage error is ", err.Error())
		message := fmt.Sprintf("mainProcessGetContactsPage: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactsPageParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2.a Get the records from the repo.
	rr, err := contactRepo.GetContacts()
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContactsPage: contactRepo.GetContacts(): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactsPageParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2.b Sort the records.
	tracker := make(map[string]*records.ContactRecord)
	keys := make([]string, 0, 5)
	for _, r := range rr {
		if _, ok := tracker[r.State]; !ok {
			tracker[r.State] = r
			keys = append(keys, r.State)
		}
	}
	sort.Strings(keys)
	l := uint64(len(keys))
	sorted := make([]*records.ContactRecord, l, l)
	for i, state := range keys {
		sorted[i] = tracker[state]
	}
	// 2.c Collect the requested records.
	start := rxparams.SortedIndex
	end := start + rxparams.PageSize
	if start > l {
		start = 0
		end = 0
	} else {
		if end > l {
			end = l
		}
	}
	requested := sorted[start:end]
	// 3. Call the renderer back with the correct records.
	txparams := &MainProcessToRendererGetContactsPageParams{
		SortedIndex: rxparams.SortedIndex,
		State:       rxparams.State,
		Records:     requested,
		RecordCount: uint64(len(sorted)),
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

// rendererReceiveAndDispatchGetContactsPage is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a MainProcessToRendererGetContactsPageParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageParams.
// 2. Dispatch the *MainProcessToRendererGetContactsPageParams.
func rendererReceiveAndDispatchGetContactsPage(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageParams.
	rxparams := &MainProcessToRendererGetContactsPageParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContactsPage defined about.
		rxparams = &MainProcessToRendererGetContactsPageParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *MainProcessToRendererGetContactsPageParams to the renderer panel callers that want to handle the GetContactsPage call backs.
	dispatch(rxparams)
}

/*

	So here is some renderer code.
	This is some code for a panel's caller file.

	import 	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/calls"

	func (caller *Caller) setCallBacks() {
		caller.connection.GetContactsPage.AddCallBack(caller.updateContactCB)
	}

	// updateContact calls the main process GetContactsPage procedure.
	func (caller *Caller) updateContact(id uint64) {
		params := &calls.RendererToMainProcessGetContactsPageParams{
			ID: id,
		}
		if err := caller.connection.GetContactsPage.CallMainProcess(params); err != nil {
			caller.tools.Error(err.Error())
		}
	}

	// updateContactCB handles a call back from the main process.
	// This func is simple:
	// Use switch params.(type) to update the *calls.MainProcessToRendererGetContactsPageParams.
	// 1. Process the params.
	func (caller *Caller) updateContactCB(params interface{}) {
		switch params.(type) {
		case *calls.MainProcessToRendererGetContactsPageParams:
			if params.Error {
				caller.tools.Error(params.ErrorMessage)
				return
			}
			// No errors so show the contact record.
			caller.presenter.showContact(params.Record)
		default:
			// default should only happen during development.
			// It means that the mainprocess func "mainProcessReceiveGetContactsPage" passed the wrong type of param to callBackToRenderer.
			caller.tools.Error("Wrong param type send from mainProcessReceiveGetContactsPage")
		}
	}

*/
