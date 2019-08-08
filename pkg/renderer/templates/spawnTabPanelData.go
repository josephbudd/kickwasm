package templates

// SpawnTabPanelData is the genereric renderer spawn data.go template.
const SpawnTabPanelData = `{{$Dot := .}}package {{call .PackageNameCase .PanelName}}

import (
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportRendererLPC}}"
	"{{.ApplicationGitPath}}{{.ImportRendererNotJS}}"
	"{{.ApplicationGitPath}}{{.ImportRendererPaneling}}"
	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

/*

	Panel name: {{.PanelName}}

*/

var (
	// quitCh will close the application
	quitCh chan struct{}

	// eojCh will close each caller's go routine.
	eojCh chan struct{}

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// The framework's renderer API.
	tools *viewtools.Tools

	// Some javascipt like dom functions written in go.
	notJS *notjs.NotJS

	// Your customized paneling.Help for initializing your panels.
	help *paneling.Help

	// The javascript null value.
	null = js.Null()
)
`
