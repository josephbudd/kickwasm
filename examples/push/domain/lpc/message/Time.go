package message

// TimeRendererToMainProcess is the Time message that the renderer sends to the main process.
type TimeRendererToMainProcess struct {
}

// TimeMainProcessToRenderer is the Time message that the main process sends to the renderer.
type TimeMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string

	Time string
}
