// +build js, wasm

package schoolcoursenotreadypanel

import (
	"context"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/api/dom"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/framework/lpc"
)

/*

	Panel name: SchoolCourseNotReadyPanel

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
