package dispatch

import (
	"github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
	"github.com/josephbudd/kickwasm/examples/push/domain/store"
	"github.com/josephbudd/kickwasm/examples/push/mainprocess/lpc"
	"github.com/josephbudd/kickwasm/examples/push/mainprocess/services/timing"
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
//   The message is sent from renderer/Main.go which you can edit.
// handleInit's response back to the renderer is the *message.InitMainProcessToRenderer.
// Param rxmessage *message.InitRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.InitMainProcessToRenderer message back to the renderer.
// Param eojing lpc.EOJer ( End Of Job ) is an interface for your go routine to receive a stop signal.
//   It signals go routines that they must stop because the main process is ending.
//   So only use it inside a go routine if you have one.
//   In your go routine
//     1. Get a channel to listen to with eojing.NewEOJ().
//     2. Before your go routine returns, release that channel with eojing.Release().
// Param stores is a struct the contains each of your stores.
func handleInit(rxmessage *message.InitRendererToMainProcess, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
	timing.Do(sending, eojing, stores)
	return
}
