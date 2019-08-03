package dispatch

import (
	"log"

	"github.com/josephbudd/kickwasm/examples/colors/domain/data/loglevels"
	"github.com/josephbudd/kickwasm/examples/colors/domain/lpc/message"
	"github.com/josephbudd/kickwasm/examples/colors/mainprocess/lpc"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.
	Kicklpc will not edit this file.

*/

// handleLog logs a renderer message to the application log.
// Param rxMessage *message.LogRendererToMainProcess is the params received from the renderer.
// Param sending is the channel to use to send a *message.LogMainProcessToRenderer to the renderer.
func handleLog(rxMessage *message.LogRendererToMainProcess, sending lpc.Sending) {
	var msg string
	var errorMsg = ""
	var isErr bool
	switch rxMessage.Level {
	case loglevels.LogLevelInfo:
		msg = "Renderer Log: Info: " + rxMessage.Message
	case loglevels.LogLevelWarning:
		msg = "Renderer Log: Warning: " + rxMessage.Message
	case loglevels.LogLevelError:
		msg = "Renderer Log: Error: " + rxMessage.Message
	case loglevels.LogLevelFatal:
		msg = "Renderer Log: Fatal: " + rxMessage.Message
	default:
		errorMsg = "Unknown Level"
		isErr = true
		msg = "Renderer Log: ???: " + rxMessage.Message
	}
	// Log the message from the renderer.
	log.Println(msg)
	// Send an update back to the renderer.
	txMessage := &message.LogMainProcessToRenderer{
		Message:      msg,
		ErrorMessage: errorMsg,
		Error:        isErr,
	}
	sending <- txMessage
	return
}
