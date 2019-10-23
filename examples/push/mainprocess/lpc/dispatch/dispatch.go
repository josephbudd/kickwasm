package dispatch

import (
	"log"

	"github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
	"github.com/josephbudd/kickwasm/examples/push/domain/store"
	"github.com/josephbudd/kickwasm/examples/push/mainprocess/lpc"
)

/*
	DO NOT EDIT THIS FILE.

	USE THE TOOL kicklpc TO ADD OR REMOVE LPC Messages.

	kicklpc will edit this file for you.

*/

// Do dispatches local process communications messages received from the renderer.
// They are dispatched to the main process handlers here in package dispatch.
// You are required to code the functionality into those handlers.
func Do(cargo interface{}, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
	switch cargo := cargo.(type) {
	case *message.LogRendererToMainProcess:
		// Log does not need a lpc.EOJer because it does not have a go routine.
		handleLog(cargo, sending)
	case *message.InitRendererToMainProcess:
		// Init signals that
		// * the renderer process is up and running,
		// * the main process may push messages to the renderer process.
		handleInit(cargo, sending, eojing, stores)
	case *message.TimeRendererToMainProcess:
		handleTime(cargo, sending, eojing, stores)
	default:
		log.Println("dispatch Do: unknown cargo type.")
	}
}