package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/services/about"
)

func newGetAboutCall() *calling.MainProcess {
	return calling.NewMainProcess(
		callids.GetAboutCallID,
		func(params []byte, call func([]byte)) {
			mainProcessGetAbout(params, call)
		},
	)
}

// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Get the about information from the about service.
// 3. Call the renderer back with
//     * the author
//     * the version
func mainProcessGetAbout(params []byte, callBackToRenderer func(params []byte)) {
	// 1. Unmarshall the params.
	rxparams := &types.RendererToMainProcessGetAboutParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetAbout: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetAboutParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2. Get the about information from the about service.
	author, version := about.GetAuthorVersion()
	// 3. Call the renderer back with the correct record.
	txparams := &types.MainProcessToRendererGetAboutParams{
		Author:  author,
		Version: version,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
