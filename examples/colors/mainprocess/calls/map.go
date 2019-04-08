package calls

import (
	"github.com/josephbudd/kickwasm/examples/colors/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
)

// TODO: Add your calls.
// Example:
//      callids.AddConactCallID: newAddContactCall(contactStorer)

// GetCallMap returns a map of each mainprocess call.
func GetCallMap(colorStore storer.ColorStorer) map[types.CallID]caller.MainProcesser {
	return map[types.CallID]caller.MainProcesser{
		callids.LogCallID: newLogCall(),
	}
}
