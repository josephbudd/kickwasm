package templates

// CallsGo is the mainprocess/calls/calls.go template.
const CallsGo = `{{$Dot := .}}package calls

import (
	"{{.ApplicationGitPath}}{{.ImportMainProcessBehaviorRepoi}}"
)

/*

	TODO:

	1. Complete the definition of the Calls struct.
	   example:
		   AddCustomer *LPC

	2. In func newCalls, complete the construction of &Calls{...}
	   example:
           AddCustomer: newAddCustomerLPC(customerRepo, rendererSendPayload),

	3. In func newCallsMap, complete the construction of callsMap.
	   example:
       callsMap[callsStruct.AddCustomer.ID] = callsStruct.AddCustomer

*/

// Calls is the calls between the main process and the renderer.
// TODO: you need to add your procedure names to this struct.
// example: GetCustomer *LPC
type Calls struct {
	Log      *LPC
{{if .AddAbout}}	GetAbout *LPC{{end}}
}

// newCalls constructs a new Calls
// TODO: You need to complete the inline construction of &Calls
// example:
//   GetCustomer: newGetCustomerLPC(rendererSendPayload),
func newCalls({{range $i, $repo := .Repos}}{{if ne $i 0}}, {{end}}{{call $Dot.LowerCamelCase $repo}}Repo repoi.{{call $Dot.CamelCase $repo}}RepoI{{end}}, rendererSendPayload func(params []byte) error) *Calls {
	lpcID = 0
	return &Calls{
		Log:      newLogLPC(rendererSendPayload),
{{if .AddAbout}}		GetAbout: newGetAboutLPC(rendererSendPayload),{{end}}
	}
}

// newCallsMap constructs new Calls for the renderer.
// TODO: You need to complete the construction of callsMap.
// example:
//   callsMap[callsStruct.GetCustomer.ID] = callsStruct.GetCustomer
func newCallsMap(callsStruct *Calls) map[int]*LPC {
	callsMap := make(map[int]*LPC)
	// build the map from Calls
	callsMap[callsStruct.Log.ID] = callsStruct.Log
{{if .AddAbout}}	callsMap[callsStruct.GetAbout.ID] = callsStruct.GetAbout{{end}}

	return callsMap
}

// NewCallsAndMap constructs a new Calls and a new map of LPC ids matched to their LPC
func NewCallsAndMap({{range $i, $repo := .Repos}}{{if ne $i 0}}, {{end}}{{call $Dot.LowerCamelCase $repo}}Repo repoi.{{$repo}}RepoI{{end}}, rendererSendPayload func(params []byte) error) (callsStruct *Calls, callsMap map[int]*LPC) {
	// make the calls
	callsStruct = newCalls({{range $i, $repo := .Repos}}{{if ne $i 0}}, {{end}}{{call $Dot.LowerCamelCase $repo}}Repo{{end}}, rendererSendPayload)
	// make the map
	callsMap = newCallsMap(callsStruct)
	// return both
	return
}
`

// CallsLogGo is the mainprocess/calls/log.go template.
const CallsLogGo = `{{$Dot := .}}package calls

import (
	"encoding/json"
	"fmt"
	"log"
)

// Log types are the type message that is logged.
const (
	LogTypeNil int = iota
	LogTypeInfo
	LogTypeWarning
	LogTypeError
	LogTypeFatal
)

// RendererToMainProcessLogParams are the Log function parameters that the renderer sends to the main process.
type RendererToMainProcessLogParams struct {
	Type    int
	Message string
}

// MainProcessToRendererLogParams are the Log function parameters that the main process sends to the renderer.
type MainProcessToRendererLogParams struct {
	Error        bool
	ErrorMessage string
	Type         int
}

func newLogLPC(rendererSendPayload func(payload []byte) error) *LPC {
	return newLPC(
		// main process receive
		func(params []byte, call func([]byte)) {
			mainProcessReceiveLog(params, call)
		},
		// renderer receive dispatch
		func(params []byte, dispatch func(interface{})) {
			rxparams := &MainProcessToRendererLogParams{}
			if err := json.Unmarshal(params, rxparams); err != nil {
				rxparams = &MainProcessToRendererLogParams{
					Error:        true,
					ErrorMessage: err.Error(),
				}
			}
			dispatch(rxparams)
		},
		rendererSendPayload,
	)
}

func mainProcessReceiveLog(params []byte, callBackToRenderer func(params []byte)) {
	log.Println("mainProcessReceiveLog")
	rxparams := &RendererToMainProcessLogParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		log.Println("mainProcessReceiveLog error is ", err.Error())
		message := fmt.Sprintf("mainProcessLog: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererLogParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
	}
	switch rxparams.Type {
	case LogTypeInfo:
		log.Println("Renderer Log: Info: ", rxparams.Message)
	case LogTypeWarning:
		log.Println("Renderer Log: Warning: ", rxparams.Message)
	case LogTypeError:
		log.Println("Renderer Log: Error: ", rxparams.Message)
	case LogTypeFatal:
		log.Println("Renderer Log: Fatal: ", rxparams.Message)
	default:
		message := "mainProcessReceiveLog: Error unknown rxparams.Type"
		log.Println(message)
		txparams := &MainProcessToRendererLogParams{
			Type:         rxparams.Type,
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return

	}
	txparams := &MainProcessToRendererLogParams{
		Type: rxparams.Type,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

func rendererReceiveAndDispatchLog(params []byte, dispatch func(interface{})) {
	rxparams := &MainProcessToRendererLogParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		rxparams = &MainProcessToRendererLogParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	dispatch(rxparams)
}

/*

	Here is an example of a panel's caller calling the mainprocess' "Log" procedure.

	import 	"{{.ApplicationGitPath}}/mainprocess/calls"

	func (caller *Caller) setCallBacks() {
		caller.connection.Log.AddCallBack(caller.LogCB)
	}

	// Log a fatal message.
	func (caller *Caller) LogFatal(message string) {
		params := &calls.RendererToMainProcessLogParams{
			Type: calls.LogTypeFatal,
			Message: message,
		}
		caller.connection.Log.CallMainProcess(params)
	}

	// LogCB Log call back from the main process.
	func (caller *Caller) LogCB(params interface{}) {
		switch params.(type) {
		case *calls.MainProcessToRendererLogParams:
			if params.Error {
				caller.tools.Error(params.ErrorMessage)
				return
			}
		}
	}

*/
`

// CallsLPCGo is the mainprocess/calls/lpc.go file.
const CallsLPCGo = `package calls

import (
	"encoding/json"
)

// Payload is the payload passed between the main process sends and the renderer.
type Payload struct {
	Procedure int
	Params    string
}

var (
	lpcID int
)

func nextLPCID() int {
	id := lpcID
	lpcID++
	return id
}

// LPC is a local procedure call.
type LPC struct {
	ID                      int
	mainprocessReceive      func(params []byte, call func(params []byte))
	rendererReceiveDispatch func(params []byte, dispatch func(interface{}))
	rendererSendPayload     func(params []byte) error
	rendererListeners       []func(params interface{})
}

func newLPC(
	mainprocessReceive func(params []byte, callback func(params []byte)),
	rendererReceiveDispatch func(params []byte, dispatch func(interface{})),
	rendererSendPayload func(payload []byte) error,
) *LPC {
	return &LPC{
		ID:                      nextLPCID(),
		mainprocessReceive:      mainprocessReceive,
		rendererReceiveDispatch: rendererReceiveDispatch,
		rendererSendPayload:     rendererSendPayload,
		rendererListeners:       make([]func(interface{}), 0, 10),
	}
}

// CallMainProcess calls the main process from the renderer
func (lpc *LPC) CallMainProcess(params interface{}) {
	paramsbb, _ := json.Marshal(params)
	payload := &Payload{
		Procedure: lpc.ID,
		Params:    string(paramsbb),
	}
	payloadbb, _ := json.Marshal(payload)
	lpc.rendererSendPayload(payloadbb)
}

// MainProcessReceive passes the params to the main process.
func (lpc *LPC) MainProcessReceive(params []byte, callback func(params []byte)) {
	lpc.mainprocessReceive(params, callback)
}

// RendererReceiveAndDispatch passes the params to the renderer.
func (lpc *LPC) RendererReceiveAndDispatch(params []byte) {
	lpc.rendererReceiveDispatch(params, lpc.dispatch)
}

// dispatch dispatches the data to the listeners
func (lpc *LPC) dispatch(params interface{}) {
	for _, f := range lpc.rendererListeners {
		f(params)
	}
}

// AddCallBack add the call back param f to this lpc's call back from the main process.
func (lpc *LPC) AddCallBack(f func(interface{})) {
	lpc.rendererListeners = append(lpc.rendererListeners, f)
}
`

// CallsExampleGo is the mainprocess/calls/example.go template.
const CallsExampleGo = `

/*
	Below is a GetCustomer example file.
	It demonstrates how you need to define your own local procedure calls.

	I would probably name it mainprocess/calls/getCustomer.go

	In my application I would also have similar files for
	  * GetCustomers
	  * UpdateCustomer
	  * RemoveCustomer

	A total of 5 things must be done.
	1. Define the params that the renderer sends to the mainprocess.
	   * In this case: RendererToMainProcessGetCustomerParams
	2. Define the params that the mainprocess sends to the renderer.
	   * In this case: MainProcessToRendererGetCustomerParams
	3. Define the constructor.
	   * In this case: newGetCustomerLPC
	   * The constructor only needs 2 funcs defined for it.
	     * In this case...
	     1. "mainProcessReceiveGetCustomer" which is the complete main process job.
	     2. "rendererReceiveAndDispatchGetCustomer" which is a simple renderer setup for the dispath process.
*/


package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"{{.ApplicationGitPath}}{{.ImportMainProcessBehaviorRepoi}}"
)

// RendererToMainProcessGetCustomerParams is the GetCustomer function parameters that the renderer sends to the main process.
type RendererToMainProcessGetCustomerParams struct {
	ID    uint64
}

// MainProcessToRendererGetCustomerParams is the GetCustomer function parameters that the main process sends to the renderer.
type MainProcessToRendererGetCustomerParams struct {
	Error        bool
	ErrorMessage string
	Record       *records.CustomerRecord
}

// newGetCustomerLPC is the constructor for the GetCustomer local procedure call.
// It should only receive the repos that are needed. In this case the customer repo.
// Param customerRepo repoi.CustomerRepoI is the customer repo needed to get a customer record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetCustomerLPC(customerRepo repoi.CustomerRepoI, rendererSendPayload func(payload []byte) error) *LPC {
	return newLPC(
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetCustomer(params, call, customerRepo)
		},
		rendererReceiveAndDispatchGetCustomer,
		rendererSendPayload,
	)
}

// mainProcessReceiveGetCustomer is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererGetCustomerParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param customerRepo is the customer repo.
// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Get the customer from the repo. Call back any errors or not found.
// 3. Call the renderer back with the customer record.
func mainProcessReceiveGetCustomer(params []byte, callBackToRenderer func(params []byte), customerRepo repoi.CustomerRepoI) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessGetCustomerParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Call back the error.
		log.Println("mainProcessReceiveGetCustomer error is ", err.Error())
		message := fmt.Sprintf("mainProcessGetCustomer: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetCustomerParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 2. Get the customer from the repo.
	customer, err := customerRepo.GetCustomer(rxparams.ID)
	if err != nil {
		// Call back the error.
		message := fmt.Sprintf("mainProcessGetCustomer: customerRepo.GetCustomer(rxparams.ID): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetCustomerParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	if customer == nil {
		// Call back "record not found".
		message := "mainProcessGetCustomer: customerRepo.GetCustomer(rxparams.ID): error is Record Not Found"
		txparams := &MainProcessToRendererGetCustomerParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return
	}
	// 3. Call the renderer back with the customer record.
	txparams := &MainProcessToRendererGetCustomerParams{
		Record:        customer,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

// rendererReceiveAndDispatchGetCustomer is a renderer func.
// It receives and dispatches the params sent by the main process.
// Param params is a []byte of a MainProcessToRendererGetCustomerParams
// Param dispatch is a func that dispatches params to the renderer call backs.
// This func is simple.
// 1. Unmarshall params into a *MainProcessToRendererGetCustomerParams.
// 2. Dispatch the *MainProcessToRendererGetCustomerParams.
func rendererReceiveAndDispatchGetCustomer(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetCustomerParams.
	rxparams := &MainProcessToRendererGetCustomerParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetCustomer defined about.
		rxparams = &MainProcessToRendererGetCustomerParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *MainProcessToRendererGetCustomerParams to the renderer panel callers that want to handle the GetCustomer call backs.
	dispatch(rxparams)
}

/*

	So here is some renderer code.
	This is some code for a panel's caller file.

	import 	"{{.ApplicationGitPath}}/mainprocess/calls"

	func (caller *Caller) setCallBacks() {
		caller.connection.GetCustomer.AddCallBack(caller.getCustomerCB)
	}

	// getCustomer calls the main process GetCustomer procedure.
	func (caller *Caller) getCustomer(id uint64) {
		params := &calls.RendererToMainProcessGetCustomerParams{
			ID: id,
		}
		if err := caller.connection.GetCustomer.CallMainProcess(params); err != nil {
			caller.tools.Error(err.Error())
		}
	}

	// getCustomerCB handles a call back from the main process.
	// This func is simple:
	// Use switch params.(type) to get the *calls.MainProcessToRendererGetCustomerParams.
	// 1. Process the params.
	func (caller *Caller) getCustomerCB(params interface{}) {
		switch params.(type) {
		case *calls.MainProcessToRendererGetCustomerParams:
			if params.Error {
				caller.tools.Error(params.ErrorMessage)
				return
			}
			// No errors so show the customer record.
			caller.presenter.showCustomer(params.Record)
		default:
			// default should only happen during development.
			// It means that the mainprocess func "mainProcessReceiveGetCustomer" passed the wrong type of param to callBackToRenderer.
			caller.tools.Error("Wrong param type send from mainProcessReceiveGetCustomer")
		}
	}

*/

`
