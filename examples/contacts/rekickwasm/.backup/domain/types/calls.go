package types

// CallID is the unique id for a map[CallID]caller.Renderer or a map[CallID]caller.MainProcessor
type CallID uint64

// Payload is a the information transported between the main process and the renderer.
type Payload struct {
	Procedure CallID
	Params    string
}

