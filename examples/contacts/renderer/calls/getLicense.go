package calls

import (
	"encoding/json"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newGetLicenseCall(rendererSendPayload func(payload []byte) error) *calling.Renderer {
	return calling.NewRenderer(
		callids.GetLicenseCallID,
		rendererReceiveAndDispatchGetLicense,
		rendererSendPayload,
	)
}

// This func is simple.
// 1. Unmarshall params into a *types.MainProcessToRendererGetLicenseParams.
// 2. Dispatch the *types.MainProcessToRendererGetLicenseParams.
func rendererReceiveAndDispatchGetLicense(params []byte, dispatch func(interface{})) {
	// 1. Unmarshall params into a *MainProcessToRendererGetLicenseParams.
	rxparams := &types.MainProcessToRendererGetLicenseParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		// This error will only happen during the development stage.
		// It means a conflict with the txparams in func mainProcessReceiveGetLicense.
		rxparams = &types.MainProcessToRendererGetLicenseParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	// 2. Dispatch the *types.MainProcessToRendererGetLicenseParams to the renderer panel callers that want to handle the GetLicense call backs.
	dispatch(rxparams)
}
