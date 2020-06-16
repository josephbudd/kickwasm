package dispatch

import (
	"context"
	"fmt"
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
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxMessage *message.LogRendererToMainProcess is the params received from the renderer.
// Param sending is the channel to use to send a *message.LogMainProcessToRenderer to the renderer.
// Builds an error for loglevels.LogLevelError and loglevels.LogLevelFatal.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleLog(ctx context.Context, rxMessage *message.LogRendererToMainProcess, sending lpc.Sending, errChan chan error) {

	var err error
	defer func() {
		if err != nil {
			errChan <- err
		}
	}()
	
	var msg string
	switch rxMessage.Level {
	case loglevels.LogLevelInfo:
		msg = "colors: Info: " + rxMessage.Message
	case loglevels.LogLevelWarning:
		msg = "colors: Warning: " + rxMessage.Message
	case loglevels.LogLevelError:
		msg = "colors: Error: " + rxMessage.Message
		err = fmt.Errorf(msg)
	case loglevels.LogLevelFatal:
		msg = "colors: Fatal: " + rxMessage.Message
		err = fmt.Errorf(msg)
	default:
		msg = fmt.Sprintf("colors: %d: %s", rxMessage.Level, rxMessage.Message)
	}
	// Log the message from the renderer.
	log.Println(msg)
	// Send an update back to the renderer.
	txMessage := &message.LogMainProcessToRenderer{
		Level:   rxMessage.Level,
		Message: rxMessage.Message,
		Fatal:   rxMessage.Level == loglevels.LogLevelFatal,
	}
	sending <- txMessage
	return
}
