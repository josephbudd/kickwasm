package types

// RendererToMainProcessRemoveContactParams are the RemoveContact function parameters that the renderer sends to the main process.
type RendererToMainProcessRemoveContactParams struct {
	ID uint64
}

// MainProcessToRendererRemoveContactParams are the RemoveContact function parameters that the main process sends to the renderer.
type MainProcessToRendererRemoveContactParams struct {
	Error        bool
	ErrorMessage string
}
