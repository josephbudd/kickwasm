package templates

// CallsMapGo is the template for renderer/calls/callmap.go
const CallsMapGo = `package calls

import (
	"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesCallers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// TODO: Add your calls.
// Example:
//      callids.AddCustomerCallID: newAddCustomerCall(rendererSendPayload, customerStorer)

// GetCallMap returns a render call map.
func GetCallMap(rendererSendPayload func(payload []byte) error) map[types.CallID]caller.Renderer {
	return map[types.CallID]caller.Renderer{
		callids.LogCallID:           newLogCall(rendererSendPayload),{{if .AddAbout}}
		callids.GetAboutCallID:      newGetAboutCall(rendererSendPayload),{{end}}
	}
}

`

// CallsLogGo is the renderer/calls/log.go template.
const CallsLogGo = `package calls

import (
	"encoding/json"

	"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
	"{{.ApplicationGitPath}}{{.ImportDomainImplementationsCalling}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

func newLogCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.LogCallID,
		rendererReceiveAndDispatchLog,
		rendererSendPayload,
	)
}

func rendererReceiveAndDispatchLog(params []byte, dispatch func(interface{})) {
	rxparams := &types.MainProcessToRendererLogCallParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		rxparams = &types.MainProcessToRendererLogCallParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	dispatch(rxparams)
}

/*

	Here is an example of a panel's caller calling the mainprocess' "Log" procedure.

	import (
		"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
		"{{.ApplicationGitPath}}{{.ImportDomainImplementationsCalling}}"
		"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
	)

	func (panelCaller *Caller) setCallBacks() {
		logger := panelCaller.connections[callids.LogCallID]
		logger.AddCallBack(panelCaller.LogCB)
	}

	// Log a fatal message.
	func (panelCaller *Caller) LogFatal(message string) {
		params := &types.RendererToMainProcessLogParams{
			Type: types.LogTypeFatal,
			Message: message,
		}
		logger := panelCaller.connections[callids.LogCallID]
		logger.CallMainProcess(params)
	}

	// LogCB Log call back from the main process.
	func (panelCaller *Caller) LogCB(params interface{}) {
		switch params := params.(type) {
		case *types.MainProcessToRendererLogCallParams:
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
		}
	}

*/

`

// CallsExampleGoTxt is the template for renderer/calls/example.go.
const CallsExampleGoTxt = `package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
	"{{.ApplicationGitPath}}{{.ImportDomainImplementationsCalling}}"
	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesCallers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesStorers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

/* The 6 Step Process.
	Step 1 must be completed outside of this package.
	Steps 2 & 3 must be completed together outside of this package.
	Steps 4 & 5 must be completed together inside this package.

	STEP 1: See domain/types.exampleGo.txt for what must be done in domain/types.

	STEPS 2 & 3: See mainprocess/calls/exampleGo.txt for what must be done in mainprocess/calls.

	STEPS 4 & 5:

		4. Define the constructor.
		* In this case: newGetCustomerCall
		* The constructor only needs 1 func defined for it. That func will receive and dispatch the params sent by the main process to your renderer panel caller funcs which will process the params.
			* In this case...
			1. "rendererReceiveAndDispatchGetCustomer" which will receive and dispatch the params.
		5. Add this Call to func makeCallMap in renderer/calls/map.go.
			ex: GetCustomerCallID: newGetCustomerCall(rendererSendPayload, contactStorer),

*/

// STEP 4.0 Define the constructor.
//
// newGetCustomerCall is the constructor for the GetCustomer Call.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetCustomerCall(rendererSendPayload func(payload []byte) error) *Call {
	return calling.NewRenderer(
		callids.GetCustomerID,
		rendererReceiveAndDispatchGetCustomer,
		rendererSendPayload,
	)
}

// STEP 4.1. Define the rendererReceiveAndDispatchGetCustomer.
//
// rendererReceiveAndDispatchGetCustomer is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a types.MainProcessToRendererGetCustomerParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererGetCustomerParams.
// 2. Dispatch the *types.MainProcessToRendererGetCustomerParams.
func rendererReceiveAndDispatchGetCustomer(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *types.MainProcessToRendererGetCustomerParams.
	rxparams := &types.MainProcessToRendererGetCustomerParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error should only happend during the development stage.
		// It means a conflict with the txparams in func processGetCustomer defined about.
		//   See mainprocess/calls/exampleGo.txt
		rxparams = &types.MainProcessToRendererGetCustomerParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererGetCustomerParams to the renderer panel caller that want to handle the GetCustomer call backs.
	dispatch(rxparams)
}

/*

	So here is some renderer code.
	This is some code for a panel's caller file.

	"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"

	func (panelCaller *Caller) setCallBacks() {
		getCustomerCall := panelCaller.connection[callids.GetCustomerCallID]
		getCustomerCall.AddCallBack(panelCaller.getCustomerCB)
	}

	// getCustomer calls the main process GetCustomer procedure.
	func (panelCaller *Caller) getCustomer(id uint64) {
		params := &types.RendererToMainProcessGetCustomerParams{
			ID: id,
		}
		getCustomer := panelCaller.connections[callids.GetCustomerID]
		if err := getCustomer.CallMainProcess(params); err != nil {
			panelCaller.tools.Error(err.Error())
		}
	}

	// getCustomerCB handles a call back from the main process.
	// This func is simple:
	// Use switch params.(type) to get the *calls.MainProcessToRendererGetCustomerParams.
	// 1. Process the params.
	func (panelCaller *Caller) getCustomerCB(params interface{}) {
		switch params := params.(type) {
		case *types.MainProcessToRendererGetCustomerParams:
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
			// No errors so show the customer record.
			panelCaller.presenter.showCustomer(params.Record)
		default:
			// default should only happen during development.
			// It means that in renderer/calls/exampleGo.txt, func rendererReceiveAndDispatchGetCustomer has unmarshalled rxparams to something other than *calls.MainProcessToRendererGetCustomerParams.
			panelCaller.tools.Error("Wrong param type send from processGetCustomer")
		}
	}

*/

`

// CallsGetAboutGo is the domain/calling/about.go template.
const CallsGetAboutGo = `package calls

import (
	"encoding/json"

	"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
	"{{.ApplicationGitPath}}{{.ImportDomainImplementationsCalling}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

func newGetAboutCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.GetAboutCallID,
		rendererReceiveAndDispatchGetAbout,
		rendererSendPayload,
	)
}

func rendererReceiveAndDispatchGetAbout(params []byte, dispatch func(interface{})) {
	rxparams := &types.MainProcessToRendererGetAboutParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		rxparams = &types.MainProcessToRendererGetAboutParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	dispatch(rxparams)
}

/*

	For renderer code see /renderer/panels/AboutButton/AboutPanel/caller.go

*/

`
