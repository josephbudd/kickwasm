package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newRemoveContactCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.RemoveContactCallID,
		rendererReceiveAndDispatchRemoveContact,
		rendererSendPayload,
	)
}

// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererRemoveContactParams.
// 2. Dispatch the *types.MainProcessToRendererRemoveContactParams.
func rendererReceiveAndDispatchRemoveContact(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererRemoveContactParams.
	rxparams := &types.MainProcessToRendererRemoveContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveRemoveContact.
		rxparams = &types.MainProcessToRendererRemoveContactParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererRemoveContactParams to the renderer panel callers that want to handle the RemoveContact call backs.
	dispatch(rxparams)
}
