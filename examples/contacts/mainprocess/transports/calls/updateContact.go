package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/behavior/repoi"
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
)

//RendererToMainProcessUpdateContactParams is the param the renderer sends to the main process to update a contact.
type RendererToMainProcessUpdateContactParams struct {
	Record *records.ContactRecord
}

// MainProcessToRendererUpdateContactParams is the params the main process sends back to the renderer after updating a contact.
type MainProcessToRendererUpdateContactParams struct {
	Record       *records.ContactRecord
	Error        bool
	ErrorMessage string
}

// newUpdateContactLPC is the constructor for the UpdateContact local procedure call.
// It should only receive the repos that are needed. In this case the contact repo.
// Param contactRepo repoi.ContactRepoI is the contact repo needed to update a contact record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newUpdateContactLPC(contactRepo repoi.ContactRepoI, rendererSendPayload func(payload []byte) error) *LPC {
	return newLPC(
		func(params []byte, call func([]byte)) {
			mainProcessReceiveUpdateContact(params, call, contactRepo)
		},
		rendererReceiveAndDispatchUpdateContact,
		rendererSendPayload,
	)
}

// mainProcessReceiveUpdateContact is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererUpdateContactParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param contactRepo is the contact repo.
// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Update the contact from the repo. Call back any errors or not found.
// 3. Call the renderer back with the contact record.
func mainProcessReceiveUpdateContact(params []byte, callBackToRenderer func(params []byte), contactRepo repoi.ContactRepoI) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessUpdateContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		log.Println("mainProcessReceiveUpdateContact error is ", err.Error())
		message := fmt.Sprintf("mainProcessUpdateContact: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererUpdateContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2. Update the contact from the repo.
	err := contactRepo.UpdateContact(rxparams.Record)
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessUpdateContact: contactRepo.UpdateContact(rxparams.ID): error is %s\n", err.Error())
		txparams := &MainProcessToRendererUpdateContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the contact record.
	txparams := &MainProcessToRendererUpdateContactParams{
		Record: rxparams.Record,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

// rendererReceiveAndDispatchUpdateContact is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a MainProcessToRendererUpdateContactParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *MainProcessToRendererUpdateContactParams.
// 2. Dispatch the *MainProcessToRendererUpdateContactParams.
func rendererReceiveAndDispatchUpdateContact(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererUpdateContactParams.
	rxparams := &MainProcessToRendererUpdateContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveUpdateContact defined about.
		rxparams = &MainProcessToRendererUpdateContactParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *MainProcessToRendererUpdateContactParams to the renderer panel callers that want to handle the UpdateContact call backs.
	dispatch(rxparams)
}

/*

	So here is some renderer code.
	This is some code for a panel's caller file.

	import 	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/calls"

	func (caller *Caller) setCallBacks() {
		caller.connection.UpdateContact.AddCallBack(caller.updateContactCB)
	}

	// updateContact calls the main process UpdateContact procedure.
	func (caller *Caller) updateContact(id uint64) {
		params := &calls.RendererToMainProcessUpdateContactParams{
			ID: id,
		}
		if err := caller.connection.UpdateContact.CallMainProcess(params); err != nil {
			caller.tools.Error(err.Error())
		}
	}

	// updateContactCB handles a call back from the main process.
	// This func is simple:
	// Use switch params.(type) to update the *calls.MainProcessToRendererUpdateContactParams.
	// 1. Process the params.
	func (caller *Caller) updateContactCB(params interface{}) {
		switch params.(type) {
		case *calls.MainProcessToRendererUpdateContactParams:
			if params.Error {
				caller.tools.Error(params.ErrorMessage)
				return
			}
			// No errors so show the contact record.
			caller.presenter.showContact(params.Record)
		default:
			// default should only happen during development.
			// It means that the mainprocess func "mainProcessReceiveUpdateContact" passed the wrong type of param to callBackToRenderer.
			caller.tools.Error("Wrong param type send from mainProcessReceiveUpdateContact")
		}
	}

*/
