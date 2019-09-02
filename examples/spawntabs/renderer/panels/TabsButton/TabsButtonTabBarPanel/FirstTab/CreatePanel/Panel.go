package createpanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/lpc"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/paneling"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/viewtools"
)

/*

	Panel name: CreatePanel

*/

// Panel has a controller, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controller *panelController
	presenter  *panelPresenter
	caller     *panelCaller
}

// NewPanel constructs a new panel.
func NewPanel(quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending, vtools *viewtools.Tools, njs *notjs.NotJS, help *paneling.Help) (panel *Panel, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "CreatePanel")
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
		group:          group,
		tabPanelHeader: notJS.GetElementByID("tabsMasterView_home_pad_TabsButton_TabsButtonTabBarPanel_tab_bar-FirstTabPanel-H3"),
	}
	caller := &panelCaller{
		group: group,
	}

	/* NOTE TO DEVELOPER. Step 1 of 1.

	// Set any controller, presenter or caller members that you added.
	// Use your custom help funcs if needed.
	// example:

	caller.state = help.GetStateAdd()

	*/

	controller.presenter = presenter
	controller.caller = caller
	presenter.controller = controller
	presenter.caller = caller
	caller.controller = controller
	caller.presenter = presenter

	// completions
	if err = group.defineMembers(); err != nil {
		return
	}
	if err = controller.defineControlsReceiveEvents(); err != nil {
		return
	}
	if err = presenter.defineMembers(); err != nil {
		return
	}

	// No errors so define the panel.
	panel = &Panel{
		controller: controller,
		presenter:  presenter,
		caller:     caller,
	}
	return
}

// StartDispatchers starts the event and message dispatchers.
func (panel *Panel) StartDispatchers() {
	panel.controller.dispatchEvents()
	panel.caller.dispatchMessages()
}

// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controller.initialCalls()
	panel.caller.initialCalls()
}
