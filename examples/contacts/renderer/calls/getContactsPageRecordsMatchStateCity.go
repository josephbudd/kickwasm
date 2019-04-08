package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newGetContactsPageRecordsMatchStateCityCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.GetContactsPageRecordsMatchStateCityCallID,
		rendererReceiveAndDispatchGetContactsPageRecordsMatchStateCity,
		rendererSendPayload,
	)
}

// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams.
// 2. Dispatch the *types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams.
func rendererReceiveAndDispatchGetContactsPageRecordsMatchStateCity(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetContactsPageRecordsMatchStateCityParams.
	rxparams := &types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happen during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetContactsPageRecordsMatchStateCity.
		rxparams = &types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams to the renderer panel callers that want to handle the GetContactsPageRecordsMatchStateCity call backs.
	dispatch(rxparams)
}
