package templates

// CallsMapGo is the template for mainprocess/calls/callmap.go
const CallsMapGo = `{{$Dot := .}}package calls

import (
	"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesCallers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesStorers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// TODO: Add your calls.
// Example:
//      callids.AddConactCallID: newAddContactCall(contactStorer)

// GetCallMap returns a map of each mainprocess call.
func GetCallMap({{range $i, $store := .Stores}}{{if ne $i 0}}, {{end}}{{call $Dot.LowerCamelCase $store}}Store storer.{{$store}}Storer{{end}}) map[types.CallID]caller.MainProcesser {
	return map[types.CallID]caller.MainProcesser{
		callids.LogCallID: newLogCall(),
	}
}
`

// CallsLogGo is the mainprocess/calls/log.go template.
const CallsLogGo = `package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
	"{{.ApplicationGitPath}}{{.ImportDomainDataLogLevels}}"
	"{{.ApplicationGitPath}}{{.ImportDomainImplementationsCalling}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

func newLogCall() *calling.MainProcess {
	return calling.NewMainProcess(
		callids.LogCallID,
		// main process receive
		func(params []byte, call func([]byte)) {
			processLog(params, call)
		},
	)
}

func processLog(params []byte, callBackToRenderer func(params []byte)) {
	rxparams := &types.RendererToMainProcessLogCallParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		log.Println("processLog error is ", err.Error())
		message := fmt.Sprintf("mainProcessLog: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &types.MainProcessToRendererLogCallParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
	}
	switch rxparams.Level {
	case loglevels.LogLevelInfo:
		log.Println("Renderer Log: Info: ", rxparams.Message)
	case loglevels.LogLevelWarning:
		log.Println("Renderer Log: Warning: ", rxparams.Message)
	case loglevels.LogLevelError:
		log.Println("Renderer Log: Error: ", rxparams.Message)
	case loglevels.LogLevelFatal:
		log.Println("Renderer Log: Fatal: ", rxparams.Message)
	default:
		message := "processLog: Error unknown rxparams.Level"
		log.Println(message)
		txparams := &types.MainProcessToRendererLogCallParams{
			Level:        rxparams.Level,
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return

	}
	txparams := &types.MainProcessToRendererLogCallParams{
		Level: rxparams.Level,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

/*

	Here is an example of a panel's caller calling the mainprocess' "Log" procedure.

	import (
		"{{.ApplicationGitPath}}{{.ImportDomainDataCallIDs}}"
		"{{.ApplicationGitPath}}{{.ImportDomainImplementationsCalling}}"
	)

	func (panelCaller *Caller) setCallBacks() {
		logger := panelCaller.connections[callerids.LogCallID]
		logger.AddCallBack(panelCaller.LogCB)
	}

	// Log a fatal message.
	func (panelCaller *Caller) LogFatal(message string) {
		params := &types.RendererToMainProcessLogParams{
			Level:   loglevels.LogLevelFatal,
			Message: message,
		}
		logger := panelCaller.connections[callerids.LogCallID]
		logger.CallMainProcess(params)
	}

	// LogCB Log call back from the main process.
	func (panelCaller *Caller) LogCB(params interface{}) {
		switch params := params.(type) {
		case *calling.MainProcessToRendererLogCallParams:
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
		}
	}

*/
`

// CallsExampleGoTxt is the template for mainprocess/calls/example.go.
const CallsExampleGoTxt = `package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/myapp/domain/data/callids"
	"github.com/myapp/domain/implementations/calling"
	"github.com/myapp/domain/interfaces/storer"
	"github.com/myapp/mainprocess/services/keyService"
)

/* The 6 Step Process.
	Step 1 must be completed outside of this package.
	Step 2 must be completed outside of this package.
	Steps 3 & 4 must be completed together inside this package.
	Steps 5 & 6 must be completed together outside of this package.

	STEP 1: See domain/implementations/data/callids/exampleGo.txt for what you must first defined in the domain/data/callids package.

	STEP 2: Define the call parameter types in domain/types.

		1. The param which is passed from the renderer to the main process.
		2. The param which is passed from the main process to the renderer.
	
		The folder **domain/types/** contains files which define the call params. You will create the 2 params for each of your calls.
		
		Below is the file **domain/types/updateContactCallParams.go** from the example/contacts program. I always give the MainProcessToRenderer param an **Error** and **ErrorMessage** so that I know if there was any error. In this case the params also contain a **State** which indictates adding or editing a contact record.

		==================================================================================================================================
		package types
		
		// RendererToMainProcessUpdateContactParams are the UpdateContact function parameters that the renderer sends to the main process.
		type RendererToMainProcessUpdateContactParams struct {
			Record       *ContactRecord
			State        uint64
		}
		
		// MainProcessToRendererUpdateContactParams are the UpdateContact function parameters that the main process sends to the renderer.
		type MainProcessToRendererUpdateContactParams struct {
			Error        bool
			ErrorMessage string
			Record       *ContactRecord
			State        uint64
		}
		===================================================================================================================================
	
	STEPS 3 & 4:
		3. Define the constructor.
		* In this case: newGetCustomerCall
		* The constructor needs the customer storer.
		* The constructor only needs 1 other func defined for it. That func will receive and process the params sent by the renderer.
			* In this case...
			1 "processGetCustomer" which is the complete main process job.

		4. Add this Call to func makeCallMap in map.go.
		ex: data.GetCustomerCallID: newGetCustomerCall(rendererSendPayload, contactStorer),

	STEPS 5 & 6: See renderer/calls/exampleGo.txt for what must be done in the renderer/calls package.

*/

// STEP 3.1: Define the constructor.
//
// newGetCustomerCall is the constructor for the GetCustomer Call.
// Param customerStorer is the key code test results storer.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetCustomerCall(customerStorer storer.KeyCodeStorer) *calling.MainProcess {
	return calling.NewMainProcess(
		data.GetCustomerCallID,
		func(params []byte, call func([]byte)) {
			processGetCustomer(params, call, customerStorer)
		},
	)
}

// STEP 3.2: Define the processGetCustomer.
// 
// processGetCustomer is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererGetCustomerParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param customerStorer is the customer store.
// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Get the customer from the store. Call back any errors or not found.
// 3. Call the renderer back with the customer record.
func processGetCustomer(params []byte, callBackToRenderer func(params []byte), customerStorer storer.CustomerStorer) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessGetCustomerParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Calling back the error.
		log.Println("processGetCustomer error is ", err.Error())
		message := fmt.Sprintf("mainProcessGetCustomer: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &types.MainProcessToRendererGetCustomerParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2. Get the customer from the store.
	customer, err := customerStorer.GetCustomer(rxparams.ID)
	if err != nil {
		// Calling back the error.
		message := fmt.Sprintf("mainProcessGetCustomer: customerStorer.GetCustomer(rxparams.ID): error is %s\n", err.Error())
		txparams := &types.MainProcessToRendererGetCustomerParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	if customer == nil {
		// Calling back "record not found".
		message := "mainProcessGetCustomer: customerStorer.GetCustomer(rxparams.ID): error is Record Not Found"
		txparams := &types.MainProcessToRendererGetCustomerParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the customer record.
	txparams := &types.MainProcessToRendererGetCustomerParams{
		Record: customer,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}
`
