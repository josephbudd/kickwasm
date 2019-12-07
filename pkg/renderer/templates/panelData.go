package templates

// PanelData is the genereric renderer panel template.
const PanelData = `{{$Dot := .}}// +build js, wasm

package {{call .PackageNameCase .PanelName}}

import (
	"{{.ApplicationGitPath}}{{.ImportRendererDOM}}"
	"{{.ApplicationGitPath}}{{.ImportRendererLPC}}"
)

/*

	Panel name: {{.PanelName}}

*/

var (
	// quitCh will close the application
	quitCh chan struct{}

	// eojCh will signal go routines to stop and return because the application is ending.
	eojCh chan struct{}

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// The document object module.
	document *dom.DOM
)
`
