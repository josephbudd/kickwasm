package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newGetAboutCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.GetAboutCallID,
		rendererReceiveAndDispatchGetAbout,
		rendererSendPayload,
	)
}

// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererGetAboutParams.
// 2. Dispatch the *types.MainProcessToRendererGetAboutParams.
func rendererReceiveAndDispatchGetAbout(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetAboutParams.
	rxparams := &types.MainProcessToRendererGetAboutParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happen during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetAbout.
		rxparams = &types.MainProcessToRendererGetAboutParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererGetAboutParams to the renderer panel callers that want to handle the GetAbout call backs.
	dispatch(rxparams)
}
