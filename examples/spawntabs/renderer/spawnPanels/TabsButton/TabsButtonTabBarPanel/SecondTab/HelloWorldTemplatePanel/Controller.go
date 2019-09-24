package helloworldtemplatepanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/viewtools"
)

/*

	Panel name: HelloWorldTemplatePanel

*/

// panelController controls user input.
type panelController struct {
	uniqueID  uint64
	panel     *spawnedPanel
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller
	eventCh   chan viewtools.Event
	unspawn   func() error

	/* NOTE TO DEVELOPER. Step 1 of 5.

	// Declare your panelController members.

	*/

	//closeSpawnButton{{.SpawnID}}
	closeSpawnButton js.Value
	//input{{.SpawnID}}
	input js.Value
	// setButton{{.SpawnID}}
	setButton js.Value
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 5.

	// Define each controller in the GUI by it's html element.
	// Handle each controller's events.

	// example:

	var id string

	// Define the customer name input field.
	id = tools.FixSpawnID("addCustomerName{{.SpawnID}}", controller.uniqueID)
	if controller.addCustomerName = notJS.GetElementByID(id); controller.addCustomerName == null {
		err = errors.New("unable to find #" + id)
		return
	}

	// Define the submit button.
	id = tools.FixSpawnID("addCustomerSubmit{{.SpawnID}}", controller.uniqueID)
	if controller.addCustomerSubmit = notJS.GetElementByID(id); controller.addCustomerSubmit == null {
		err = errors.New("unable to find #" + id)
		return
	}
	// Handle the submit button's onclick event.
	tools.AddSpawnEventHandler(controller.handleSubmit, controller.addCustomerSubmit, "click", false, controller.uniqueID)

	*/

	var id string

	// Define the self close button and set it's handler.
	id = tools.FixSpawnID("closeSpawnButton{{.SpawnID}}", controller.uniqueID)
	if controller.closeSpawnButton = notJS.GetElementByID(id); controller.closeSpawnButton == null {
		err = errors.New("unable to find #" + id)
		return
	}
	// Handle the submit button's onclick event.
	tools.AddSpawnEventHandler(controller.handleClick, controller.closeSpawnButton, "click", false, controller.uniqueID)

	// Define the label input.
	id = tools.FixSpawnID("input{{.SpawnID}}", controller.uniqueID)
	if controller.input = notJS.GetElementByID(id); controller.input == null {
		err = errors.New("unable to find #" + id)
		return
	}

	// Define the label set button and set it's handler.
	id = tools.FixSpawnID("setButton{{.SpawnID}}", controller.uniqueID)
	if controller.setButton = notJS.GetElementByID(id); controller.setButton == null {
		err = errors.New("unable to find #" + id)
		return
	}
	// Handle the submit button's onclick event.
	tools.AddSpawnEventHandler(controller.handleSetClick, controller.setButton, "click", false, controller.uniqueID)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

*/

func (controller *panelController) handleClick(event viewtools.Event) (nilReturn interface{}) {
	if err := controller.unspawn(); err != nil {
		tools.Error(err.Error())
	}
	return
}

func (controller *panelController) handleSetClick(event viewtools.Event) (nilReturn interface{}) {
	var text string
	if text = notJS.GetValue(controller.input); len(text) == 0 {
		tools.Error("Enter some text for the tab label.")
		return
	}
	controller.presenter.setTabLabel(text)
	return
}

func (controller *panelController) UnSpawning() {

	/* NOTE TO DEVELOPER. Step 4 of 5.

	// This func is called when this tab and it's panels are in the process of unspawning.
	// So if you have some cleaning up to do then do it now.
	//
	// For example if you have a widget that needs to be unspawned
	//   because maybe it has a go routine running that needs to be stopped
	//   then do it here.

	*/
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 5 of 5.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/
}
