package dispatch

import (
	"context"

	"github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
	"github.com/josephbudd/kickwasm/examples/push/domain/store"
	"github.com/josephbudd/kickwasm/examples/push/mainprocess/lpc"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.
	Kicklpc will not edit this file.

*/

// handleTime is the *message.TimeRendererToMainProcess handler.
// It's response back to the renderer is the *message.TimeMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.TimeRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.TimeMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleTime(ctx context.Context, rxmessage *message.TimeRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	return
}
