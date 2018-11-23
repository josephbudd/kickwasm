package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newGetContactsPageStatesCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.GetContactsPageStatesCallID,
		rendererReceiveAndDispatchGetContactsPageStates,
		rendererSendPayload,
	)
}

// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererGetContactsPageStatesParams.
// 2. Dispatch the *types.MainProcessToRendererGetContactsPageStatesParams.
func rendererReceiveAndDispatchGetContactsPageStates(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageStatesParams.
	rxparams := &types.MainProcessToRendererGetContactsPageStatesParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happend during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContactsPageStates.
		rxparams = &types.MainProcessToRendererGetContactsPageStatesParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererGetContactsPageStatesParams to the renderer panel callers that want to handle the GetContactsPageStates call backs.
	dispatch(rxparams)
}
