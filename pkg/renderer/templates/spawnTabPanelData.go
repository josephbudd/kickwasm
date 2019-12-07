package templates

// SpawnTabPanelData is the genereric renderer spawn data.go template.
const SpawnTabPanelData = `{{$Dot := .}}// +build js, wasm

package {{call .PackageNameCase .PanelName}}

import (
	"{{.ApplicationGitPath}}{{.ImportRendererLPC}}"
	"{{.ApplicationGitPath}}{{.ImportRendererPaneling}}"
)

/*

	Panel name: {{.PanelName}}

*/

var (
	// quitCh will close the application
	quitCh chan struct{}

	// eojCh will close each panel messenger's message dispatcher go routine.
	eojCh chan struct{}

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// Your customized paneling.Help for initializing your panels.
	help *paneling.Help
)
`
