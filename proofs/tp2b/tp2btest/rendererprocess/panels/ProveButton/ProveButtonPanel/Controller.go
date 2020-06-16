// +build js, wasm

package provebuttonpanel

import (
	"fmt"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/prove"
)
		
/*

	Panel name: ProveButtonPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	messenger *panelMessenger
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(controller *panelController) defineControlsHandlers(): %!w(string=Move panel from a tab to a button.)", err)
		}
	}()

	return
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	if err := prove.Pass(); err != nil {
		controller.messenger.LogFail(err)
	} else {
		controller.messenger.logPass()
	}
}
