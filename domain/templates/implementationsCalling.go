package templates

// ImplementationsMainProcessGo is the template for domain/implementations/calling/mainprocess.go.
const ImplementationsMainProcessGo = `package calling

import (
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// MainProcess is a procedure call from the main process to the renderer and the renderer to the main process.
// MainProcess implements caller.MainProcess with func Process
type MainProcess struct {
	ID                 types.CallID
	mainprocessReceive func(params []byte, call func(params []byte))
}

// NewMainProcess constructs a new call.
func NewMainProcess(
	id types.CallID,
	mainprocessReceive func(params []byte, callback func(params []byte)),
) *MainProcess {
	return &MainProcess{
		ID:                 id,
		mainprocessReceive: mainprocessReceive,
	}
}

// Process passes the params to the main process.
func (call *MainProcess) Process(params []byte, callback func(params []byte)) {
	call.mainprocessReceive(params, callback)
}
`

// ImplementationsRendererGo is the template for domain/implementations/calling/renderer.go.
const ImplementationsRendererGo = `package calling

import (
	"encoding/json"

	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// Renderer is a procedure call from the main process to the renderer and the renderer to the main process.
// Renderer implements caller.Renderer with funcs CallMainProcess, AddCallBack, Dispatch.
type Renderer struct {
	ID                      types.CallID
	rendererReceiveDispatch func(params []byte, dispatch func(interface{}))
	rendererSendPayload     func(params []byte) error
	rendererListeners       []func(params interface{})
}

// NewRenderer constructs a new call.
func NewRenderer(
	id types.CallID,
	rendererReceiveDispatch func(params []byte, dispatch func(interface{})),
	rendererSendPayload func(payload []byte) error,
) *Renderer {
	return &Renderer{
		ID:                      id,
		rendererReceiveDispatch: rendererReceiveDispatch,
		rendererSendPayload:     rendererSendPayload,
		rendererListeners:       make([]func(interface{}), 0, 10),
	}
}

// CallMainProcess calls the main process from the renderer
func (call *Renderer) CallMainProcess(params interface{}) {
	paramsbb, _ := json.Marshal(params)
	payload := &types.Payload{
		Procedure: call.ID,
		Params:    string(paramsbb),
	}
	payloadbb, _ := json.Marshal(payload)
	call.rendererSendPayload(payloadbb)
}

// AddCallBack add the call back param f to this call's call back from the main process.
func (call *Renderer) AddCallBack(f func(interface{})) {
	call.rendererListeners = append(call.rendererListeners, f)
}

// Dispatch passes the params to the renderer call back funcs.
func (call *Renderer) Dispatch(params []byte) {
	call.rendererReceiveDispatch(
		params,
		func(params interface{}) {
			for _, f := range call.rendererListeners {
				f(params)
			}
		},
	)
}
`
