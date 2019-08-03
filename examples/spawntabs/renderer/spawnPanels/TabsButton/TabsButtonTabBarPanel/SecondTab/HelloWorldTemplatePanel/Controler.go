package helloworldtemplatepanel

import (
	"syscall/js"

	"github.com/pkg/errors"
)

/*

	Panel name: HelloWorldTemplatePanel

*/

// panelControler controls user input.
type panelControler struct {
	uniqueID  uint64
	panel     *spawnedPanel
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller
	unspawn   func() error

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelControler members.

	*/

	// closeSpawnButton{{.SpawnID}}
	closeSpawnButton js.Value
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

	// Define the panelControler members by their html elements.
	// Set their handlers.

	*/

	var id string

	// Define the self close button and set it's handler.
	id = tools.FixSpawnID("closeSpawnButton{{.SpawnID}}", controler.uniqueID)
	if controler.closeSpawnButton = notJS.GetElementByID(id); controler.closeSpawnButton == null {
		err = errors.New("unable to find #" + id)
		return
	}
	// see render/viewtools/callback.go
	// use the event call back func for spawned panels and set propagations.
	cb := tools.RegisterSpawnEventCallBack(controler.handleClick, true, true, true, controler.uniqueID)
	notJS.SetOnClick(controler.closeSpawnButton, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controler *panelControler) handleClick(event js.Value) (nilReturn interface{}) {
	if err := controler.unspawn(); err != nil {
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
