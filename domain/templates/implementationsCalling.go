package templates

// ImplementationsCallingMapGo is the template for domain/implementations/calling/callmap.go
const ImplementationsCallingMapGo = `{{$Dot := .}}package calling

import (
	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesCallers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesStorers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

var nextid types.CallID

// CallConstructor creates a new Call.
type CallConstructor func(rendererSendPayload func(payload []byte) error) *Call

// TODO: Add your calls.
// Example:
//      AddConactCallID: newAddContactCall(rendererSendPayload, contactStorer)
func makeCallMap(rendererSendPayload func(payload []byte) error, {{range $i, $repo := .Repos}}{{if ne $i 0}}, {{end}}{{call $Dot.LowerCamelCase $repo}}Storer storer.{{$repo}}Storer{{end}}) map[types.CallID]interface{} {
	return map[types.CallID]interface{}{
		LogCallID:      newLogCall(rendererSendPayload),{{if .AddAbout}}
		GetAboutCallID: newGetAboutCall(rendererSendPayload),{{end}}
	}
}

func nextCallID() types.CallID {
	id := nextid
	nextid++
	return id
}

// GetMainProcessCallsMap returns an id call map needed for the main process.
func GetMainProcessCallsMap({{range $i, $repo := .Repos}}{{if ne $i 0}}, {{end}}{{call $Dot.LowerCamelCase $repo}}Storer storer.{{$repo}}Storer{{end}}) map[types.CallID]caller.MainProcesser {
	cmap := makeCallMap(nil, {{range $i, $repo := .Repos}}{{if ne $i 0}}, {{end}}{{call $Dot.LowerCamelCase $repo}}Storer{{end}})
	mpmap := make(map[types.CallID]caller.MainProcesser)
	for k, v := range cmap {
		mpmap[k] = v.(caller.MainProcesser)
	}
	return mpmap
}

// GetRendererCallMap returns an id call map needed for the renderer.
// Param rendererSendPayload is the renderer's func that gets the payload sent to the main process.
func GetRendererCallMap(rendererSendPayload func(payload []byte) error) map[types.CallID]caller.Renderer {
	cmap := makeCallMap(rendererSendPayload, {{range $i, $repo := .Repos}}{{if ne $i 0}}, {{end}}nil{{end}})
	rmap := make(map[types.CallID]caller.Renderer)
	for k, v := range cmap {
		rmap[k] = v.(caller.Renderer)
	}
	return rmap
}

`

// ImplementationsCallGo is the template for domain/implementations/calling/call.go.
const ImplementationsCallGo = `package calling

import (
	"encoding/json"

	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// Call is a procedure call from the main process to the renderer and the renderer to the main process.
// Call implements caller.MainProcess and caller.Renderer
type Call struct {
	ID                      types.CallID
	mainprocessReceive      func(params []byte, call func(params []byte))
	rendererReceiveDispatch func(params []byte, dispatch func(interface{}))
	rendererSendPayload     func(params []byte) error
	rendererListeners       []func(params interface{})
}

// newCall constructs a new call.
func newCall(
	id types.CallID,
	mainprocessReceive func(params []byte, callback func(params []byte)),
	rendererReceiveDispatch func(params []byte, dispatch func(interface{})),
	rendererSendPayload func(payload []byte) error,
) *Call {
	return &Call{
		ID:                      id,
		mainprocessReceive:      mainprocessReceive,
		rendererReceiveDispatch: rendererReceiveDispatch,
		rendererSendPayload:     rendererSendPayload,
		rendererListeners:       make([]func(interface{}), 0, 10),
	}
}

// CallMainProcess calls the main process from the renderer
func (call *Call) CallMainProcess(params interface{}) {
	paramsbb, _ := json.Marshal(params)
	payload := &types.Payload{
		Procedure: call.ID,
		Params:    string(paramsbb),
	}
	payloadbb, _ := json.Marshal(payload)
	call.rendererSendPayload(payloadbb)
}

// AddCallBack add the call back param f to this call's call back from the main process.
func (call *Call) AddCallBack(f func(interface{})) {
	call.rendererListeners = append(call.rendererListeners, f)
}

// MainProcessReceive passes the params to the main process.
func (call *Call) MainProcessReceive(params []byte, callback func(params []byte)) {
	call.mainprocessReceive(params, callback)
}

// RendererReceiveAndDispatch passes the params to the renderer.
func (call *Call) RendererReceiveAndDispatch(params []byte) {
	call.rendererReceiveDispatch(params, call.dispatch)
}

// dispatch dispatches the data to the listeners
func (call *Call) dispatch(params interface{}) {
	for _, f := range call.rendererListeners {
		f(params)
	}
}

`

// CallingLogGo is the domain/implementations/calling/log.go template.
const CallingLogGo = `package calling

import (
	"encoding/json"
	"fmt"
	"log"
)

// The Log call id.
var LogCallID = nextCallID()

// Log types are the type message that is logged.
const (
	LogTypeNil int = iota
	LogTypeInfo
	LogTypeWarning
	LogTypeError
	LogTypeFatal
)

// The call Log must have 2 params.
// 1. RendererToMainProcessLogParams
// 2. MainProcessToRendererLogParams

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

// The Log call must have a constructor of type CallConstructor.

func newLogCall(rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		LogCallID,
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

	import 	"{{.ApplicationGitPath}}{{.ImportDomainImplementationsCalling}}"

	func (caller *Caller) setCallBacks() {
		logger := caller.connections[calling.LogCallID]
		logger.AddCallBack(caller.LogCB)
	}

	// Log a fatal message.
	func (caller *Caller) LogFatal(message string) {
		params := &calling.RendererToMainProcessLogParams{
			Type: calling.LogTypeFatal,
			Message: message,
		}
		logger := caller.connections[calling.LogCallID]
		logger.CallMainProcess(params)
	}

	// LogCB Log call back from the main process.
	func (caller *Caller) LogCB(params interface{}) {
		switch params := params.(type) {
		case *calling.MainProcessToRendererLogParams:
			if params.Error {
				caller.tools.Error(params.ErrorMessage)
				return
			}
		}
	}

*/

`

// CallingExampleGoTxt is the template for domain/implementations/calling/example.go.
const CallingExampleGoTxt = `

/*
	Below is a GetCustomer example file.
	It demonstrates how you need to define your own local procedure calls.

	I would probably name it mainprocess/calling/getCustomer.go

	In my application I would also have similar files for
	  * GetCustomers
	  * UpdateCustomer
	  * RemoveCustomer

	A total of 5 things must be done.
	1. Define the call id.
	2. Define the params that the renderer sends to the mainprocess.
	   * In this case: RendererToMainProcessGetCustomerParams
	3. Define the params that the mainprocess sends to the renderer.
	   * In this case: MainProcessToRendererGetCustomerParams
	4. Define the constructor.
	   * In this case: newGetCustomerCall
	   * The constructor only needs 2 funcs defined for it.
	     * In this case...
	     1. "mainProcessReceiveGetCustomer" which is the complete main process job.
		 2. "rendererReceiveAndDispatchGetCustomer" which is a simple renderer setup for the dispath process.
	5. Add this Call to the map constructor in map.go func makeCallMap.
		ex: GetCustomerID: newGetCustomerCall,
*/


package calling

import (
	"encoding/json"
	"fmt"
	"log"

	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesStorers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// ID
var GetCustomerID = nextCallID()

// RendererToMainProcessGetCustomerParams is the GetCustomer function parameters that the renderer sends to the main process.
type RendererToMainProcessGetCustomerParams struct {
	ID    uint64
}

// MainProcessToRendererGetCustomerParams is the GetCustomer function parameters that the main process sends to the renderer.
type MainProcessToRendererGetCustomerParams struct {
	Error        bool
	ErrorMessage string
	Record       *types.CustomerRecord
}

// newGetCustomerCall is the constructor for the GetCustomer local procedure call.
// It should only receive the repos that are needed. In this case the customer repo.
// Param customerStorer storer.CustomerStorer is the customer repo needed to get a customer record.
// Param rendererSendPayload: is a kickasm generated renderer func that sends data to the main process.
func newGetCustomerCall(customerStorer storer.CustomerStorer, rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		GetCustomerID,
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetCustomer(params, call, customerStorer)
		},
		rendererReceiveAndDispatchGetCustomer,
		rendererSendPayload,
	)
}

// mainProcessReceiveGetCustomer is a main process func.
// This is how the main process receives a call from the renderer.
// Param params is a []byte of a MainProcessToRendererGetCustomerParams
// Param callBackToRenderer is a func that calls back to the renderer.
// Param customerStorer is the customer repo.
// The func is simple:
// 1. Unmarshall the params. Call back any errors.
// 2. Get the customer from the repo. Call back any errors or not found.
// 3. Call the renderer back with the customer record.
func mainProcessReceiveGetCustomer(params []byte, callBackToRenderer func(params []byte), customerStorer storer.CustomerStorer) {
	// 1. Unmarshall the params.
	rxparams := &RendererToMainProcessGetCustomerParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// Calling back the error.
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
	customer, err := customerStorer.GetCustomer(rxparams.ID)
	if err != nil {
		// Calling back the error.
		message := fmt.Sprintf("mainProcessGetCustomer: customerStorer.GetCustomer(rxparams.ID): error is %s\n", err.Error())
		txparams := &MainProcessToRendererGetCustomerParams{
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
	// 2. Dispatch the *MainProcessToRendererGetCustomerParams to the renderer panel caller that want to handle the GetCustomer call backs.
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
		getCustomer := caller.connections[calling.GetCustomerID]
		if err := getCustomer.CallMainProcess(params); err != nil {
			caller.tools.Error(err.Error())
		}
	}

	// getCustomerCB handles a call back from the main process.
	// This func is simple:
	// Use switch params.(type) to get the *calls.MainProcessToRendererGetCustomerParams.
	// 1. Process the params.
	func (caller *Caller) getCustomerCB(params interface{}) {
		switch params := params.(type) {
		case *calling.MainProcessToRendererGetCustomerParams:
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

// CallingGetAboutGo is the domain/calling/about.go template.
const CallingGetAboutGo = `package calling

import (
	"encoding/json"
	"log"

	"{{.ApplicationGitPath}}{{.ImportMainProcessServicesAbout}}"
)

// The GetAbout call id.
var GetAboutCallID = nextCallID()

// MainProcessToRendererGetAboutParams are the GetAbout function parameters that the main process sends to the renderer.
type MainProcessToRendererGetAboutParams struct {
	Error        bool
	ErrorMessage string
	Version      string
	Releases     map[string]map[string][]string
	Contributors map[string]string
	Credits      []about.Credit
	Licenses     []about.License
}

func newGetAboutCall(rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		GetAboutCallID,
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetAbout(params, call)
		},
		rendererReceiveAndDispatchGetAbout,
		rendererSendPayload,
	)
}

func rendererCallMainProcessGetAbout(params interface{}, call func([]byte)) error {
	call([]byte{})
	return nil
}

func mainProcessReceiveGetAbout(params []byte, callBack func(params []byte)) {
	log.Println("mainProcessReceiveGetAbout")
	txparams := &MainProcessToRendererGetAboutParams{
		Version:      about.String(),
		Releases:     about.GetReleases(),
		Contributors: about.GetContributors(),
		Credits:      about.GetCredits(),
		Licenses:     about.GetLicenses(),
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBack(txparamsbb)
}

func rendererReceiveAndDispatchGetAbout(params []byte, dispatch func(interface{})) {
	rxparams := &MainProcessToRendererGetAboutParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		rxparams = &MainProcessToRendererGetAboutParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	dispatch(rxparams)
}

/*

	For renderer code see {{.ApplicationGitPath}}{{.OutputRendererPanels}}/AboutButton/AboutPanel/caller.go

*/

`
