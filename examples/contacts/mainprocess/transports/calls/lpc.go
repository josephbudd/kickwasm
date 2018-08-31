package calls

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
