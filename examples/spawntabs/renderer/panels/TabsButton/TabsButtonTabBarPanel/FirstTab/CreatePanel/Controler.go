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

// panelControler controls user input.
type panelControler struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelControler members.

	*/

	newHelloWorldButton js.Value
}

// defineControlsSetHandlers defines controler members and sets their handlers.
// Returns the error.
func (controler *panelControler) defineControlsSetHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controler *panelControler) defineControlsSetHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set their handlers.

	*/

	// Define the submit button and set it's handler.
	if controler.newHelloWorldButton = notJS.GetElementByID("newHelloWorldButton"); controler.newHelloWorldButton == null {
		err = errors.New("unable to find #newHelloWorldButton")
		return
	}
	// see render/viewtools/callback.go
	// use the event call back func and set propagations.
	cb := tools.RegisterEventCallBack(controler.handleClick, true, true, true)
	notJS.SetOnClick(controler.newHelloWorldButton, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controler *panelControler) handleClick(event js.Value) (nilReturn interface{}) {
	spawnCount++
	n := spawnCount
	tabLabel := fmt.Sprintf("Tab %d", n)
	panelHeading := fmt.Sprintf("Panel Heading %d", n)
	if _, err := secondtab.Spawn(tabLabel, panelHeading, nil); err != nil {
		tools.Error(err.Error())
	}
	return
}

// initialCalls runs the first code that the controler needs to run.
func (controler *panelControler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
