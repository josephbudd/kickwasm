package types

// RendererToMainProcessUpdateContactParams are the UpdateContact function parameters that the renderer sends to the main process.
type RendererToMainProcessUpdateContactParams struct {
	Record *ContactRecord
	State  uint64
}

// MainProcessToRendererUpdateContactParams are the UpdateContact function parameters that the main process sends to the renderer.
type MainProcessToRendererUpdateContactParams struct {
	Error        bool
	ErrorMessage string
	Record       *ContactRecord
	State        uint64
}
