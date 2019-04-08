package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newGetContactsPageCitiesMatchStateCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.GetContactsPageCitiesMatchStateCallID,
		rendererReceiveAndDispatchGetContactsPageCitiesMatchState,
		rendererSendPayload,
	)
}

// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererGetContactsPageCitiesMatchStateParams.
// 2. Dispatch the *types.MainProcessToRendererGetContactsPageCitiesMatchStateParams.
func rendererReceiveAndDispatchGetContactsPageCitiesMatchState(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageCitiesMatchStateParams.
	rxparams := &types.MainProcessToRendererGetContactsPageCitiesMatchStateParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happen during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContactsPageCitiesMatchState.
		rxparams = &types.MainProcessToRendererGetContactsPageCitiesMatchStateParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererGetContactsPageCitiesMatchStateParams to the renderer panel callers that want to handle the GetContactsPageCitiesMatchState call backs.
	dispatch(rxparams)
}
