package templates

// SpawnTabPanelController is the genereric renderer panel controller template.
const SpawnTabPanelController = `package {{call .PackageNameCase .PanelName}}

import (
	"github.com/pkg/errors"

	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

/*

	Panel name: {{.PanelName}}

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

	// example:

	// my spawn template has a name input field and a submit button.
	// <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">
	// <button id="addCustomerSubmit{{.SpawnID}}">Close</button>

	import "syscall/js"

	addCustomerName   js.Value
	addCustomerSubmit js.Value

	*/
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

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

// example:

// import "{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"

func (controller *panelController) handleSubmit(e viewtools.Event) (nilReturn interface{}) {
	// See renderer/viewtools/event.go.
	// The viewtools.Event funcs.
	//   e.PreventDefaultBehavior()
	//   e.StopCurrentPhasePropagation()
	//   e.StopAllPhasePropagation()
	//   target := e.Target
	//   event := e.Event
	name := strings.TrimSpace(notJS.GetValue(controller.addCustomerName))
	if len(name) == 0 {
		tools.Error("Customer Name is required.")
		return
	}
	r := &record.CustomerRecord{
		Name: name,
	}
	controller.caller.AddCustomer(r)
	return
}

*/


func (controller *panelController) UnSpawning() {

	/* NOTE TO DEVELOPER. Step 4 of 5.

	// This func is called when this tab and it's panels are in the process of unspawning.
	// So if you have some cleaning up to do then do it now.
	//
	// For example if you have a widget that needs to be unspawned
	//   because maybe it has a go routine running that needs to be stopped
	//   then do it here.

	// example:

	controller.myWidget.UnSpawn()

	*/
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 5 of 5.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	// example:

	controller.customerSelectWidget.start()

	*/
}
`
