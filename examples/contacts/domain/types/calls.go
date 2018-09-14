package types

import "github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"

// CallID is the unique id for a RendererCallMap or a MainProcessCallsMap
type CallID uint64

// Payload is a the information transported between the main process and the renderer.
type Payload struct {
	Procedure CallID
	Params    string
}

// RendererCallMap is each call id mapped to its Renderer interface.
type RendererCallMap map[CallID]caller.Renderer

// MainProcessCallsMap is each call id mapped to its MainProcessor interfaces.
type MainProcessCallsMap map[CallID]caller.MainProcesser

