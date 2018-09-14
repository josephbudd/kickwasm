package calling

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

var nextid types.CallID

// CallConstructor creates a new Call.
type CallConstructor func(rendererSendPayload func(payload []byte) error) *Call

// TODO: Add your calls.
// Example:
//      AddConactCallID: newAddContactCall(rendererSendPayload, contactStorer)
func makeCallMap(rendererSendPayload func(payload []byte) error, contactStorer storer.ContactStorer) map[types.CallID]interface{} {
	return map[types.CallID]interface{}{
		LogCallID:      newLogCall(rendererSendPayload),
		GetAboutCallID: newGetAboutCall(rendererSendPayload),
	}
}

func nextCallID() types.CallID {
	id := nextid
	nextid++
	return id
}

// GetMainProcessCallsMap returns an id call map needed for the main process.
func GetMainProcessCallsMap(contactStorer storer.ContactStorer) map[types.CallID]caller.MainProcesser {
	cmap := makeCallMap(nil, contactStorer)
	mpmap := make(map[types.CallID]caller.MainProcesser)
	for k, v := range cmap {
		mpmap[k] = v.(caller.MainProcesser)
	}
	return mpmap
}

// GetRendererCallMap returns an id call map needed for the renderer.
// Param rendererSendPayload is the renderer's func that gets the payload sent to the main process.
func GetRendererCallMap(rendererSendPayload func(payload []byte) error) map[types.CallID]caller.Renderer {
	cmap := makeCallMap(rendererSendPayload, nil)
	rmap := make(map[types.CallID]caller.Renderer)
	for k, v := range cmap {
		rmap[k] = v.(caller.Renderer)
	}
	return rmap
}

