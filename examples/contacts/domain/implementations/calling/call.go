package calling

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
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

