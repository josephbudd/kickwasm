// +build js, wasm

package action1level4markuppanel

import (
	"context"

	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/api/dom"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/framework/lpc"
)

/*

	Panel name: Action1Level4MarkupPanel

*/

var (
	// rendererProcessCtx is the renderer process's context.
	rendererProcessCtx context.Context

	// rendererProcessCtxCancel is the renderer process's context cancel func.
	// Calling it will stop the entire renderer process.
	// To gracefully stop the entire renderer process use either of the api funcs
	//   application.GracefullyClose(cancelFunc context.CancelFunc)
	//   or application.NewGracefullyCloseHandler(cancelFunc context.CancelFunc) (handler func(e event.Event) (nilReturn interface{})).
	rendererProcessCtxCancel context.CancelFunc

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// The document object module.
	document *dom.DOM
)
