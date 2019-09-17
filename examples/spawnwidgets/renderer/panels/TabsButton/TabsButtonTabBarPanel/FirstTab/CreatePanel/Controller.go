package createpanel

import (
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"

	secondtab "github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/viewtools"
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

	/* NOTE TO DEVELOPER. Step 1 of 5.

	// Declare your panelController members.

	*/

	newHelloWorldButton js.Value
}

// defineControlsReceiveEvents defines controller members and starts receiving their events.
// Returns the error.
func (controller *panelController) defineControlsReceiveEvents() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsReceiveEvents()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 5.

	// Define the controller members by their html elements.
	// Receive their events.

	*/

	// Define the submit button and set it's handler.
	if controller.newHelloWorldButton = notJS.GetElementByID("newHelloWorldButton"); controller.newHelloWorldButton == null {
		err = errors.New("unable to find #newHelloWorldButton")
		return
	}
	// Receive the submit button's onclick event.
	controller.receiveEvent(controller.newHelloWorldButton, "onclick", false, false, false)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

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

// dispatchEvents dispatches events from the controls.
// It stops when it receives on the eoj channel.
func (controller *panelController) dispatchEvents() {
	go func() {
		var event viewtools.Event
		for {
			select {
			case <-eojCh:
				return
			case event = <-controller.eventCh:
				// An event that this controller is receiving from one of its members.
				switch event.Target {

				/* NOTE TO DEVELOPER. Step 4 of 5.

				// 4.1.a: Add a case for each controller member
				//          that you are receiving events for.
				// 4.1.b: In that case statement, pass the event to your event handler.

				*/

				case controller.newHelloWorldButton:
					controller.handleClick(event.Event)
				}
			}
		}
	}()

	return
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 5 of 5.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}

// receiveEvent gets this controller listening for element's event.
// Param elements if the controler's element.
// Param event is the event ex: "onclick".
// Param preventDefault indicates if the default behavior of the event must be prevented.
// Param stopPropagation indicates if the event's propogation must be stopped.
// Param stopImmediatePropagation indicates if the event's immediate propogation must be stopped.
func (controller *panelController) receiveEvent(element js.Value, event string, preventDefault, stopPropagation, stopImmediatePropagation bool) {
	tools.SendEvent(controller.eventCh, element, event, preventDefault, stopPropagation, stopImmediatePropagation)
}
