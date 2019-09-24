package createpanel

import (
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"

	secondtab "github.com/josephbudd/kickwasm/examples/spawntabs/renderer/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/viewtools"
)

/*

	Panel name: CreatePanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller
	eventCh   chan viewtools.Event

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController fields.

	*/

	newHelloWorldButton js.Value
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define each controller in the GUI by it's html element.
	// Handle each controller's events.

	// example:

	// Define the customer name text input GUI controller.
	if controller.addCustomerName = notJS.GetElementByID("addCustomerName"); controller.addCustomerName == null {
		err = errors.New("unable to find #addCustomerName")
		return
	}

	// Define the submit button GUI controller.
	if controller.addCustomerSubmit = notJS.GetElementByID("addCustomerSubmit"); controller.addCustomerSubmit == null {
		err = errors.New("unable to find #addCustomerSubmit")
		return
	}
	// Handle the submit button's onclick event.
	tools.AddEventHandler(controller.handleSubmit, controller.addCustomerSubmit, "click", false)

	*/

	// Define the submit button and set it's handler.
	if controller.newHelloWorldButton = notJS.GetElementByID("newHelloWorldButton"); controller.newHelloWorldButton == null {
		err = errors.New("unable to find #newHelloWorldButton")
		return
	}
	// Handle the button's onclick event.
	tools.AddEventHandler(controller.handleClick, controller.newHelloWorldButton, "click", false)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controller *panelController) handleClick(event viewtools.Event) (nilReturn interface{}) {
	spawnCount++
	n := spawnCount
	tabLabel := fmt.Sprintf("Tab %d", n)
	panelHeading := fmt.Sprintf("Panel Heading %d", n)
	if _, err := secondtab.Spawn(tabLabel, panelHeading, nil); err != nil {
		tools.Error(err.Error())
	}
	return
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/
}
