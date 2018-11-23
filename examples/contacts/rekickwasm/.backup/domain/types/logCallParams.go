package types

// RendererToMainProcessLogCallParams are the Log function parameters that the renderer sends to the main process.
type RendererToMainProcessLogCallParams struct {
	Level   uint64
	Message string
}

// MainProcessToRendererLogCallParams are the Log function parameters that the main process sends to the renderer.
type MainProcessToRendererLogCallParams struct {
	Error        bool
	ErrorMessage string
	Level        uint64
}

