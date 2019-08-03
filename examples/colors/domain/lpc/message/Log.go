package message

// LogRendererToMainProcess is the Log message that the renderer sends to the main process.
type LogRendererToMainProcess struct {
	Level   uint64
	Message string
}

// LogMainProcessToRenderer is the Log message that the main process sends to the renderer.
type LogMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string
	Level        uint64
	Message      string
}
