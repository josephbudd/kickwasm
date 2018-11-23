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

func newUpdateContactCall(contactStorer storer.ContactStorer) *calling.MainProcess {
	return calling.NewMainProcess(
		callids.UpdateContactCallID,
		func(params []byte, call func([]byte)) {
			mainProcessUpdateContact(params, call, contactStorer)
		},
	)
}

// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Update the contact from the storer. Call back any errors.
// 3. Call the renderer back with the contact record.
func mainProcessUpdateContact(params []byte, callBackToRenderer func([]byte), contactStorer storer.ContactStorer) {
	rxparams := &types.RendererToMainProcessUpdateContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessUpdateContact: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererUpdateContactParams{
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
		log.Println(message)
		txparams := &types.MainProcessToRendererUpdateContactParams{
			Error:        true,
			ErrorMessage: message,
			State:        rxparams.State,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the contact record.
	txparams := &types.MainProcessToRendererUpdateContactParams{
		Record: rxparams.Record,
		State:  rxparams.State,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
