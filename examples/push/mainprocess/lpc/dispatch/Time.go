package dispatch

import (
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
// Param rxmessage *message.TimeRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.TimeMainProcessToRenderer message back to the renderer.
// Param eojing lpc.EOJer ( End Of Job ) is an interface for your go routine to receive a stop signal.
//   It signals go routines that they must stop because the main process is ending.
//   So only use it inside a go routine if you have one.
//   In your go routine
//     1. Get a channel to listen to with eojing.NewEOJ().
//     2. Before your go routine returns, release that channel with eojing.Release().
// Param stores is a struct the contains each of your stores.
func handleTime(rxmessage *message.TimeRendererToMainProcess, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
	return
}
