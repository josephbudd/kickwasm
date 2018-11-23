package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
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
		"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
		"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
		"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
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

