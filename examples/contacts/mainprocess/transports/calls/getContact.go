package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/behavior/repoi"
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
)

//RendererToMainProcessGetContactParams is the param the renderer sends to the main process to update a contact.
type RendererToMainProcessGetContactParams struct {
	ID    uint64
	State uint64
}

// MainProcessToRendererGetContactParams is the params the main process sends back to the renderer after updating a contact.
type MainProcessToRendererGetContactParams struct {
	Record       *records.ContactRecord
	State        uint64
	Error        bool
	ErrorMessage string
}

// newGetContactLPC is the constructor for the GetContact local procedure call.
// It should only receive the repos that are needed. In this case the contact repo.
// Param contactRepo repoi.ContactRepoI is the contact repo needed to update a contact record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetContactLPC(contactRepo repoi.ContactRepoI, rendererSendPayload func(payload []byte) error) *LPC {
	return newLPC(
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetContact(params, call, contactRepo)
		},
		rendererReceiveAndDispatchGetContact,
		rendererSendPayload,
	)
}

// mainProcessReceiveGetContact is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererGetContactParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param contactRepo is the contact repo.
// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Get the contact from the repo. Call back any errors or not found.
// 3. Call the renderer back with the contact record.
func mainProcessReceiveGetContact(params []byte, callBackToRenderer func(params []byte), contactRepo repoi.ContactRepoI) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessGetContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		log.Println("mainProcessReceiveGetContact error is ", err.Error())
		message := fmt.Sprintf("mainProcessGetContact: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2. Get the contact from the repo.
	contact, err := contactRepo.GetContact(rxparams.ID)
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContact: contactRepo.GetContact(rxparams.ID): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the contact record.
	txparams := &MainProcessToRendererGetContactParams{
		Record: contact,
		State:  rxparams.State,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

// rendererReceiveAndDispatchGetContact is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a MainProcessToRendererGetContactParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *MainProcessToRendererGetContactParams.
// 2. Dispatch the *MainProcessToRendererGetContactParams.
func rendererReceiveAndDispatchGetContact(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactParams.
	rxparams := &MainProcessToRendererGetContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContact defined about.
		rxparams = &MainProcessToRendererGetContactParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *MainProcessToRendererGetContactParams to the renderer panel callers that want to handle the GetContact call backs.
	dispatch(rxparams)
}

/*

	So here is some renderer code.
	This is some code for a panel's caller file.

	import 	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/calls"

	func (caller *Caller) setCallBacks() {
		caller.connection.GetContact.AddCallBack(caller.getContactCB)
	}

	// GetContact

	func (caller *Caller) getContact(id uint64) {
		params := &calls.RendererToMainProcessGetContactParams{
			ID:    id,
			State: controler.contactEditSelectID,
		}
		caller.connection.GetContact.CallMainProcess(params)
	}

	func (caller *Caller) getContactCB(params interface{}) {
		switch params := params.(type) {
		case *calls.MainProcessToRendererGetContactParams:
			state := params.State & controler.contactEditSelectID
			if state == controler.contactEditSelectID {
				if params.Error {
					caller.tools.Error(params.ErrorMessage)
				}
				// no error so let the edit panel handle the call back.
			}
		}
	}

*/
