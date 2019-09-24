package helloworldtemplatepanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/viewtools"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/widgets"
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

	widget *widgets.Button
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

	*/

	// The button widget handles it's own events not this controller.
	controller.widget = widgets.NewButton(tools, notJS)
	id := tools.FixSpawnID("widgetWrapper{{.SpawnID}}", controller.uniqueID)
	widgetWrapper := notJS.GetElementByID(id)
	controller.widget.Spawn(widgetWrapper, "Close", controller.handleClick)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

*/

// handleClick unspawns.
// First it unspawns the widget.
// Then it unspawns the tab.
func (controller *panelController) handleClick(event viewtools.Event) (nilInterface interface{}) {
	// Unspawn this panel's tab and all of it's panels.
	if err := controller.unspawn(); err != nil {
		tools.Error(err.Error())
	}
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

	// Unspawn the widget.
	controller.widget.UnSpawn()
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 5 of 5.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/
}
