package types

// RendererToMainProcessGetAboutParams are the GetAboutPage function parameters that the renderer sends to the main process.
type RendererToMainProcessGetAboutParams struct{}

// MainProcessToRendererGetAboutParams are the GetAboutPage function parameters that the main process sends to the renderer.
type MainProcessToRendererGetAboutParams struct {
	Author       string
	Version      []string
	Error        bool
	ErrorMessage string
}
