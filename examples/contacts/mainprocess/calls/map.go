package calls

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

// TODO: Add your calls.
// Example:
//      callids.AddConactCallID: newAddContactCall(contactStorer)

// GetCallMap returns a map of each mainprocess call.
func GetCallMap(contactStore storer.ContactStorer) map[types.CallID]caller.MainProcesser {
	return map[types.CallID]caller.MainProcesser{
		callids.LogCallID:                                  newLogCall(),
		callids.UpdateContactCallID:                        newUpdateContactCall(contactStore),
		callids.RemoveContactCallID:                        newRemoveContactCall(contactStore),
		callids.GetContactCallID:                           newGetContactCall(contactStore),
		callids.GetContactsPageCitiesMatchStateCallID:      newGetContactsPageCitiesMatchStateCall(contactStore),
		callids.GetContactsPageRecordsMatchStateCityCallID: newGetContactsPageRecordsMatchStateCityCall(contactStore),
		callids.GetContactsPageStatesCallID:                newGetContactsPageStatesCall(contactStore),
	}
}
