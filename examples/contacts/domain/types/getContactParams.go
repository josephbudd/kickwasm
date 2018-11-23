package types

// RendererToMainProcessGetContactParams are the GetContactsPage function parameters that the renderer sends to the main process.
type RendererToMainProcessGetContactParams struct {
	ID    uint64
	State uint64
}

// MainProcessToRendererGetContactParams are the GetContactsPage function parameters that the main process sends to the renderer.
type MainProcessToRendererGetContactParams struct {
	Record       *ContactRecord
	State        uint64
	Error        bool
	ErrorMessage string
}
