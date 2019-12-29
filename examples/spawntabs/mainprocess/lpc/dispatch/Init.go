package dispatch

import (
	"context"

	"github.com/josephbudd/kickwasm/examples/spawntabs/domain/lpc/message"
	"github.com/josephbudd/kickwasm/examples/spawntabs/domain/store"
	"github.com/josephbudd/kickwasm/examples/spawntabs/mainprocess/lpc"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.
	Kicklpc will not edit this file.

*/

// handleInit is the *message.InitRendererToMainProcess handler.
//   The InitRendererToMainProcess message signals that
//   * the renderer process is up and running,
//   * the main process may push messages to the renderer process.
//   The message is sent from rendererprocess/Main.go which you can edit.
// handleInit's response back to the renderer is the *message.InitMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.InitRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.InitMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleInit(ctx context.Context, rxmessage *message.InitRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	return
}
