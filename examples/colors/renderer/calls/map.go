package calls

import (
	"github.com/josephbudd/kickwasm/examples/colors/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
)

// TODO: Add your calls.
// Example:
//      callids.AddCustomerCallID: newAddCustomerCall(rendererSendPayload, customerStorer)

// GetCallMap returns a render call map.
func GetCallMap(rendererSendPayload func(payload []byte) error) map[types.CallID]caller.Renderer {
	return map[types.CallID]caller.Renderer{
		callids.LogCallID: newLogCall(rendererSendPayload),
	}
}
