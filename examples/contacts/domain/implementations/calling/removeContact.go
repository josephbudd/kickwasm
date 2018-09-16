package calling

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
)

// RemoveContactCallID is the RemoveContact call id.
var RemoveContactCallID = nextCallID()

//RendererToMainProcessRemoveContactParams is the param the renderer sends to the main process to remove a contact.
type RendererToMainProcessRemoveContactParams struct {
	ID uint64
}

// MainProcessToRendererRemoveContactParams is the params the main process sends back to the renderer after updating a contact.
type MainProcessToRendererRemoveContactParams struct {
	Error        bool
	ErrorMessage string
}

// newRemoveContactCall is the constructor for the RemoveContact local procedure call.
// It should only receive the repos that are needed. In this case the contact repo.
// Param contactStorer storer.ContactStorer is the contact repo needed to remove a contact record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newRemoveContactCall(contactStorer storer.ContactStorer, rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		RemoveContactCallID,
		func(params []byte, call func([]byte)) {
			mainProcessRemoveContact(params, call, contactStorer)
		},
		rendererReceiveAndDispatchRemoveContact,
		rendererSendPayload,
	)
}

// mainProcessRemoveContact is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererRemoveContactParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param contactStorer is the contact repo.
// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Remove the contact from the repo. Call back any errors or not found.
// 3. Call the renderer back with no errors.
func mainProcessRemoveContact(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessRemoveContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		log.Println("mainProcessRemoveContact error is ", err.Error())
		message := fmt.Sprintf("mainProcessRemoveContact: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererRemoveContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2. Remove the contact from the repo.
	err := contactStorer.RemoveContact(rxparams.ID)
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessRemoveContact: contactStorer.RemoveContact(rxparams.ID): error is %s\n", err.Error())
		txparams := &MainProcessToRendererRemoveContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the contact record. No errors.
	txparams := &MainProcessToRendererRemoveContactParams{}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

// rendererReceiveAndDispatchRemoveContact is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a MainProcessToRendererRemoveContactParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *MainProcessToRendererRemoveContactParams.
// 2. Dispatch the *MainProcessToRendererRemoveContactParams.
func rendererReceiveAndDispatchRemoveContact(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererRemoveContactParams.
	rxparams := &MainProcessToRendererRemoveContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainRemoveContact defined about.
		rxparams = &MainProcessToRendererRemoveContactParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *MainProcessToRendererRemoveContactParams to the renderer panel callers that want to handle the RemoveContact call backs.
	dispatch(rxparams)
}

/*

	For renderer code see "github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/panels/RemoveButton/RemoveContactConfirmPanel/caller.go"

*/
