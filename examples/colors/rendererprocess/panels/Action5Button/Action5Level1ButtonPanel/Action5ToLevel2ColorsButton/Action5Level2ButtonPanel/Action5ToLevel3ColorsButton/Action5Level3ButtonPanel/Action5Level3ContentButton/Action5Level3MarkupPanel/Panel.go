// +build js, wasm

package action5level3markuppanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/lpc"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/paneling"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/viewtools"
)

/*

	Panel name: Action5Level3MarkupPanel

*/

// Panel has a controller, presenter and messenger.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controller *panelController
	presenter  *panelPresenter
	messenger  *panelMessenger
}

// NewPanel constructs a new panel.
func NewPanel(quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending, vtools *viewtools.Tools, njs *notjs.NotJS, help *paneling.Help) (panel *Panel, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "Action5Level3MarkupPanel")
		}
	}()

	quitCh = quitChan
	eojCh = eojChan
	receiveCh = receiveChan
	sendCh = sendChan
	tools = vtools
	notJS = njs

	group := &panelGroup{}
	controller := &panelController{
		group:   group,
		eventCh: make(chan viewtools.Event, 1024),
	}
	presenter := &panelPresenter{
		group: group,
	}
	messenger := &panelMessenger{
		group: group,
	}

	/* NOTE TO DEVELOPER. Step 1 of 1.

	// Set any controller, presenter or messenger members that you added.
	// Use your custom help funcs if needed.

	// example:

	messenger.state = help.GetStateAdd()

	*/

	controller.presenter = presenter
	controller.messenger = messenger
	presenter.controller = controller
	presenter.messenger = messenger
	messenger.controller = controller
	messenger.presenter = presenter

	// completions
	if err = group.defineMembers(); err != nil {
		return
	}
	if err = controller.defineControlsHandlers(); err != nil {
		return
	}
	if err = presenter.defineMembers(); err != nil {
		return
	}

	// No errors so define the panel.
	panel = &Panel{
		controller: controller,
		presenter:  presenter,
		messenger:  messenger,
	}
	return
}

// StartDispatchers starts the event and message dispatchers.
func (panel *Panel) StartDispatchers() {
	panel.messenger.dispatchMessages()
}

// InitialJobs runs the first code that the panel needs to run.
func (panel *Panel) InitialJobs() {
	panel.controller.initialCalls()
	panel.messenger.initialSends()
}
