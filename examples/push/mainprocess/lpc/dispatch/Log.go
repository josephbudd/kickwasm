package dispatch

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/push/domain/data/loglevels"
	"github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
	"github.com/josephbudd/kickwasm/examples/push/mainprocess/lpc"
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
		msg = "push: Info: " + rxMessage.Message
	case loglevels.LogLevelWarning:
		msg = "push: Warning: " + rxMessage.Message
	case loglevels.LogLevelError:
		msg = "push: Error: " + rxMessage.Message
		err = errors.New(msg)
	case loglevels.LogLevelFatal:
		msg = "push: Fatal: " + rxMessage.Message
		err = errors.New(msg)
	default:
		msg = fmt.Sprintf("push: %d: %s", rxMessage.Level, rxMessage.Message)
	}
	// Log the message from the renderer.
	log.Println(msg)
	// Send an update back to the renderer.
	// In this case no errors.
	txMessage := &message.LogMainProcessToRenderer{
		Level:   rxMessage.Level,
		Message: rxMessage.Message,
		Fatal:   rxMessage.Level == loglevels.LogLevelFatal,
	}
	sending <- txMessage
	return
}
