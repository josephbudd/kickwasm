package types

// RendererToMainProcessGetLicenseParams are the GetLicensePage function parameters that the renderer sends to the main process.
type RendererToMainProcessGetLicenseParams struct {
	Liscense []string
}

// MainProcessToRendererGetLicenseParams are the GetLicensePage function parameters that the main process sends to the renderer.
type MainProcessToRendererGetLicenseParams struct {
	License      []string
	Error        bool
	ErrorMessage string
}
