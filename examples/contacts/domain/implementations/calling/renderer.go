package calling

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
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
