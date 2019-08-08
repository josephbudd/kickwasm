package createpanel

import (
	"fmt"
	"syscall/js"

	secondtab "github.com/josephbudd/kickwasm/examples/spawntabs/renderer/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab"
	"github.com/pkg/errors"
)

/*

	Panel name: CreatePanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController members.

	*/

	newHelloWorldButton js.Value
}

// defineControlsSetHandlers defines controller members and sets their handlers.
// Returns the error.
func (controller *panelController) defineControlsSetHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsSetHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controller members by their html elements.
	// Set their handlers.

	*/

	// Define the submit button and set it's handler.
	if controller.newHelloWorldButton = notJS.GetElementByID("newHelloWorldButton"); controller.newHelloWorldButton == null {
		err = errors.New("unable to find #newHelloWorldButton")
		return
	}
	// see render/viewtools/callback.go
	// use the event call back func and set propagations.
	cb := tools.RegisterEventCallBack(controller.handleClick, true, true, true)
	notJS.SetOnClick(controller.newHelloWorldButton, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controller *panelController) handleClick(event js.Value) (nilReturn interface{}) {
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
