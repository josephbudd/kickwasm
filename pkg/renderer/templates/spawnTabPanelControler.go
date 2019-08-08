package templates

// SpawnTabPanelController is the genereric renderer panel controller template.
const SpawnTabPanelController = `package {{call .PackageNameCase .PanelName}}

import (
	"github.com/pkg/errors"
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
	unspawn   func() error

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController members.
	// example:

	// my spawn template has a name input field and a submit button.
	// <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">
	// <button id="addCustomerSubmit{{.SpawnID}}">Close</button>

	// import "syscall/js"

	addCustomerName   js.Value
	addCustomerSubmit js.Value

	*/
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
	// example:

	// import "fmt"
	// import "syscall/js"

	var id string
	var cb js.Func

	// Define the customer name input field.
	id = tools.FixSpawnID("addCustomerName{{.SpawnID}}", controller.uniqueID)
	if controller.addCustomerName = notJS.GetElementByID(id); controller.addCustomerName == null {
		err = errors.New("unable to find #" + id)
		return
	}

	// Define the submit button and set it's handler.
	id = tools.FixSpawnID("addCustomerSubmit{{.SpawnID}}", controller.uniqueID)
	if controller.addCustomerSubmit = notJS.GetElementByID(id); controller.addCustomerSubmit == null {
		err = errors.New("unable to find #" + id)
		return
	}
	// see render/viewtools/callback.go
	// use the event call back func for spawned panels and set propagations.
	cb = tools.RegisterSpawnEventCallBack(controller.handleSubmit, true, true, true, controller.uniqueID)
	notJS.SetOnClick(controller.addCustomerSubmit, cb)

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.
// example:

// import "{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"

func (controller *panelController) handleSubmit(event js.Value) (nilReturn interface{}) {
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

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.
	// example:

	controller.customerSelectWidget.start()

	*/

}
`
