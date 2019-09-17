package message

// InitRendererToMainProcess is the Init message that the renderer sends to the main process.
// InitRendererToMainProcess signals that
// * the renderer process is up and running,
// * the main process may push messages to the renderer process.
type InitRendererToMainProcess struct {
}

// InitMainProcessToRenderer is the Init message that the main process sends to the renderer.
type InitMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string
}
