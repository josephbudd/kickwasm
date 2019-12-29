// +build js, wasm

package createpanel

import (
	"context"

	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/dom"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework/lpc"
)

/*

	Panel name: CreatePanel

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

	// spawnCount is the number of spawns.
	spawnCount uint
)
