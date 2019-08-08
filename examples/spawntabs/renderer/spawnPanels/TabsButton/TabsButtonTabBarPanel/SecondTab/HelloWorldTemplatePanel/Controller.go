package helloworldtemplatepanel

import (
	"syscall/js"

	"github.com/pkg/errors"
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
	unspawn   func() error

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController members.

	*/

	// closeSpawnButton{{.SpawnID}}
	closeSpawnButton js.Value
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

	// Define the panelController members by their html elements.
	// Set their handlers.

	*/

	var id string

	// Define the self close button and set it's handler.
	id = tools.FixSpawnID("closeSpawnButton{{.SpawnID}}", controller.uniqueID)
	if controller.closeSpawnButton = notJS.GetElementByID(id); controller.closeSpawnButton == null {
		err = errors.New("unable to find #" + id)
		return
	}
	// see render/viewtools/callback.go
	// use the event call back func for spawned panels and set propagations.
	cb := tools.RegisterSpawnEventCallBack(controller.handleClick, true, true, true, controller.uniqueID)
	notJS.SetOnClick(controller.closeSpawnButton, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controller *panelController) handleClick(event js.Value) (nilReturn interface{}) {
	if err := controller.unspawn(); err != nil {
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
