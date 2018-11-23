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

func newGetContactCall(contactStorer storer.ContactStorer) *calling.MainProcess {
	return calling.NewMainProcess(
		callids.GetContactCallID,
		func(params []byte, call func([]byte)) {
			mainProcessGetContact(params, call, contactStorer)
		},
	)
}

// The func is simple:
// 1.  Unmarshall the params. Call back any errors.
// 2.a Get the record from the store. Call back any errors.
// 3.  Call the renderer back with
//     * the same State
//     * the requested record.
func mainProcessGetContact(params []byte, callBackToRenderer func(params []byte), contactStorer storer.ContactStorer) {
	// 1. Unmarshall the params.
	rxparams := &types.RendererToMainProcessGetContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContact: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2.a Get the record from the store.
	r, err := contactStorer.GetContact(rxparams.ID)
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetContact: contactStorer.GetContact(rxparams.ID): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetContactParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the correct record.
	txparams := &types.MainProcessToRendererGetContactParams{
		State:  rxparams.State,
		Record: r,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
