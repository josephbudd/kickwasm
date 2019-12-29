// +build js, wasm

package helloworldtemplatepanel

import (
	"context"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/framework/lpc"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/paneling"
)

/*

	Panel name: HelloWorldTemplatePanel

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
