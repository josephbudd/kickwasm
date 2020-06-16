// +build js, wasm

package provebuttonpanel

import (
	"strings"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/domain/data/loglevels"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/domain/lpc/message"
)

/*

	Panel name: ProveButtonPanel

*/

// panelMessenger communicates with the main process via an asynchrounous connection.
type panelMessenger struct {
	group      *panelGroup
	presenter  *panelPresenter
	controller *panelController
}

func (messenger *panelMessenger) LogFail(err error) {
	msg := &message.LogRendererToMainProcess{
		Level:   loglevels.LogLevelError,
		Message: "Failed tp2b test.\n"+strings.ReplaceAll(err.Error(), "\n", "<br/>"),
	}
	sendCh <- msg
}

func (messenger *panelMessenger) logPass() {
	msg := &message.LogRendererToMainProcess{
		Level:   loglevels.LogLevelInfo,
		Message: "Move panel from a tab to a button. Passed test.",
	}
	sendCh <- msg
}

func (messenger *panelMessenger) rxLog(msg *message.LogMainProcessToRenderer) {
	rendererProcessCtxCancel()
}

// dispatchMessages dispatches LPC messages from the main process.
// It stops when it receives on the eoj channel.
func (messenger *panelMessenger) dispatchMessages() {
	go func() {
		for {
			select {
			case <-rendererProcessCtx.Done():
				return
			case msg := <-receiveCh:
				// A message sent from the main process to the renderer.
				switch msg := msg.(type) {

				case *message.LogMainProcessToRenderer:
					messenger.rxLog(msg)
				}
			}
		}
	}()

	return
}

// initialSends sends the first messages to the main process.
func (messenger *panelMessenger) initialSends() {
}
