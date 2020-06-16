package templates

// SpawnTabPanelData is the genereric renderer spawn data.go template.
const SpawnTabPanelData = `{{$Dot := .}}// +build js, wasm

package {{call .PackageNameCase .PanelName}}

import (
	"context"

	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkLPC}}"
	"{{.ApplicationGitPath}}{{.ImportRendererPaneling}}"
)

/*

	Panel name: {{.PanelName}}

*/

var (
	// rendererProcessCtxCancel is the renderer process's context cancel func.
	// Calling it will stop the entire renderer process.
	// To gracefully stop the entire renderer process use the api funcs
	//   application.GracefullyClose(rendererProcessCtxCancel)
	//   or application.NewGracefullyCloseHandler(rendererProcessCtxCancel)
	rendererProcessCtxCancel context.CancelFunc

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// Your customized paneling.Help for initializing your panels.
	help *paneling.Help
)
`
