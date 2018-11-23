package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newUpdateContactCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.UpdateContactCallID,
		rendererReceiveAndDispatchUpdateContact,
		rendererSendPayload,
	)
}

// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererUpdateContactParams.
// 2. Dispatch the *types.MainProcessToRendererUpdateContactParams.
func rendererReceiveAndDispatchUpdateContact(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererUpdateContactParams.
	rxparams := &types.MainProcessToRendererUpdateContactParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveUpdateContact.
		rxparams = &types.MainProcessToRendererUpdateContactParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererUpdateContactParams to the renderer panel callers that want to handle the UpdateContact call backs.
	dispatch(rxparams)
}
