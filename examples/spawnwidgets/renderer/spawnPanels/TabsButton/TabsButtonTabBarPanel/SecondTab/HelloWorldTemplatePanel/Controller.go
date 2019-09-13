package helloworldtemplatepanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/viewtools"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/widgets"
)

/*

	Panel name: HelloWorldTemplatePanel

*/

// panelController controls user input.
type panelController struct {
	uniqueID     uint64
	panel        *spawnedPanel
	group        *panelGroup
	presenter    *panelPresenter
	caller       *panelCaller
	eventCh      chan viewtools.Event
	unSpawningCh chan struct{}
	unspawn      func() error

	/* NOTE TO DEVELOPER. Step 1 of 5.

	// Declare your panelController members.

	*/

	widget *widgets.Button
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

	controller.widget = widgets.NewButton(tools, notJS, eojCh)
	id := tools.FixSpawnID("widgetWrapper{{.SpawnID}}", controller.uniqueID)
	widgetWrapper := notJS.GetElementByID(id)
	controller.widget.Spawn(widgetWrapper, "Close", controller.handleClick)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

*/

func (controller *panelController) handleClick(event js.Value) {
	controller.widget.UnSpawn()
	if err := controller.unspawn(); err != nil {
		tools.Error(err.Error())
	}
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
			case <-controller.unSpawningCh:
				return
			case event = <-controller.eventCh:
				// An event that this controller is receiving from one of its members.
				switch event.Target {

				/* NOTE TO DEVELOPER. Step 4 of 5.

				// 4.1.a: Add a case for each controller member
				//          that you are receiving events for.
				// 4.1.b: In that case statement, pass the event to your event handler.

				*/
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

func (controller *panelController) receiveEvent(element js.Value, event string, preventDefault, stopPropagation, stopImmediatePropagation bool) {
	tools.SendSpawnEvent(controller.eventCh, element, event, preventDefault, stopPropagation, stopImmediatePropagation, controller.uniqueID)
}
