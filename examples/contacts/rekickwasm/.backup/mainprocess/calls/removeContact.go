package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newRemoveContactCall(contactStorer storer.ContactStorer) *calling.MainProcess {
	return calling.NewMainProcess(
		callids.RemoveContactCallID,
		func(params []byte, call func([]byte)) {
			mainProcessRemoveContact(params, call, contactStorer)
		},
	)
}

// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Remove the contact from the repo. Call back any errors.
// 3. Call the renderer back with no errors.
func mainProcessRemoveContact(params []byte, callBackToRenderer func([]byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &types.RendererToMainProcessRemoveContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessRemoveContact: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererRemoveContactParams{
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
		log.Println(message)
		txparams := &types.MainProcessToRendererRemoveContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the contact record. No errors.
	txparams := &types.MainProcessToRendererRemoveContactParams{}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
