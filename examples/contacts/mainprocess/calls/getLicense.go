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

func newGetLicenseCall() *calling.MainProcess {
	return calling.NewMainProcess(
		callids.GetLicenseCallID,
		func(params []byte, call func([]byte)) {
			mainProcessGetLicense(params, call)
		},
	)
}

// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Get the license information from the about service.
// 3. Call the renderer back with
//     * the license
func mainProcessGetLicense(params []byte, callBackToRenderer func(params []byte)) {
	// 1. Unmarshall the params.
	rxparams := &types.RendererToMainProcessGetLicenseParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetLicense: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		log.Println(message)
		txparams := &types.MainProcessToRendererGetLicenseParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2. Get the license information from the about service.
	license := about.GetLicense()
	// 3. Call the renderer back with the correct record.
	txparams := &types.MainProcessToRendererGetLicenseParams{
		License: license,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
