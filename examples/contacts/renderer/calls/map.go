package calls

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

// TODO: Add your calls.
// Example:
//      callids.AddCustomerCallID: newAddCustomerCall(rendererSendPayload, customerStorer)

// GetCallMap returns a render call map.
func GetCallMap(rendererSendPayload func(payload []byte) error) map[types.CallID]caller.Renderer {
	return map[types.CallID]caller.Renderer{
		callids.LogCallID:                                  newLogCall(rendererSendPayload),
		callids.UpdateContactCallID:                        newUpdateContactCall(rendererSendPayload),
		callids.RemoveContactCallID:                        newRemoveContactCall(rendererSendPayload),
		callids.GetContactCallID:                           newGetContactCall(rendererSendPayload),
		callids.GetContactsPageCitiesMatchStateCallID:      newGetContactsPageCitiesMatchStateCall(rendererSendPayload),
		callids.GetContactsPageRecordsMatchStateCityCallID: newGetContactsPageRecordsMatchStateCityCall(rendererSendPayload),
		callids.GetContactsPageStatesCallID:                newGetContactsPageStatesCall(rendererSendPayload),
		callids.GetAboutCallID:                             newGetAboutCall(rendererSendPayload),
		callids.GetLicenseCallID:                           newGetLicenseCall(rendererSendPayload),
	}
}
