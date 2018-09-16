package calling

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

// UpdateContactCallID is the UpdateContact call id.
var UpdateContactCallID = nextCallID()

//RendererToMainProcessUpdateContactParams is the param the renderer sends to the main process to update a contact.
type RendererToMainProcessUpdateContactParams struct {
	Record *types.ContactRecord
	State  uint64
}

// MainProcessToRendererUpdateContactParams is the params the main process sends back to the renderer after updating a contact.
type MainProcessToRendererUpdateContactParams struct {
	Record       *types.ContactRecord
	State        uint64
	Error        bool
	ErrorMessage string
}

// newUpdateContactCall is the constructor for the UpdateContact Call.
// It should only receive the repos that are needed. In this case the contact repo.
// Param contactStorer storer.ContactStorer is the contact repo needed to update a contact record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newUpdateContactCall(contactStorer storer.ContactStorer, rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		UpdateContactCallID,
		func(params []byte, call func([]byte)) {
			mainProcessReceiveUpdateContact(params, call, contactStorer)
		},
		rendererReceiveAndDispatchUpdateContact,
		rendererSendPayload,
	)
}

// mainProcessReceiveUpdateContact is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererUpdateContactParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param contactStorer is the contact repo.
// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Update the contact from the repo. Call back any errors or not found.
// 3. Call the renderer back with the contact record.
func mainProcessReceiveUpdateContact(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
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
	err := contactStorer.UpdateContact(rxparams.Record)
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessUpdateContact: contactStorer.UpdateContact(rxparams.ID): error is %s\n", err.Error())
		txparams := &MainProcessToRendererUpdateContactParams{
			Error:        true,
			ErrorMessage: message,
			State:        rxparams.State,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the contact record.
	txparams := &MainProcessToRendererUpdateContactParams{
		Record: rxparams.Record,
		State:  rxparams.State,
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

	For renderer code see "github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/panels/EditButton/EditContactEditPanel/caller.go"

*/
