package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newGetContactCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.GetContactCallID,
		rendererReceiveAndDispatchGetContact,
		rendererSendPayload,
	)
}

// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererGetContactParams.
// 2. Dispatch the *types.MainProcessToRendererGetContactParams.
func rendererReceiveAndDispatchGetContact(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactParams.
	rxparams := &types.MainProcessToRendererGetContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happen during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContact.
		rxparams = &types.MainProcessToRendererGetContactParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererGetContactParams to the renderer panel callers that want to handle the GetContact call backs.
	dispatch(rxparams)
}
