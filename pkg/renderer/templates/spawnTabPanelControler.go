package templates

// SpawnTabPanelControler is the genereric renderer panel controler template.
const SpawnTabPanelControler = `package {{call .PackageNameCase .PanelName}}

import (
	"github.com/pkg/errors"
)

/*

	Panel name: {{.PanelName}}

*/

// panelControler controls user input.
type panelControler struct {
	uniqueID  uint64
	panel     *spawnedPanel
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller
	unspawn    func() error

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelControler members.
	// example:

	// my spawn template has a name input field and a submit button.
	// <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">
	// <button id="addCustomerSubmit{{.SpawnID}}">Close</button>

	// import "syscall/js"

	addCustomerName   js.Value
	addCustomerSubmit js.Value

	*/
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
	// example:

	// import "fmt"
	// import "syscall/js"

	var id string
	var cb js.Func

	// Define the customer name input field.
	id = tools.FixSpawnID("addCustomerName{{.SpawnID}}", controler.uniqueID)
	if controler.addCustomerName = notJS.GetElementByID(id); controler.addCustomerName == null {
		err = errors.New("unable to find #" + id)
		return
	}

	// Define the submit button and set it's handler.
	id = tools.FixSpawnID("addCustomerSubmit{{.SpawnID}}", controler.uniqueID)
	if controler.addCustomerSubmit = notJS.GetElementByID(id); controler.addCustomerSubmit == null {
		err = errors.New("unable to find #" + id)
		return
	}
	// see render/viewtools/callback.go
	// use the event call back func for spawned panels and set propagations.
	cb = tools.RegisterSpawnEventCallBack(controler.handleSubmit, true, true, true, controler.uniqueID)
	notJS.SetOnClick(controler.addCustomerSubmit, cb)

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.
// example:

// import "{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"

func (controler *panelControler) handleSubmit(event js.Value) (nilReturn interface{}) {
	name := strings.TrimSpace(notJS.GetValue(controler.addCustomerName))
	if len(name) == 0 {
		tools.Error("Customer Name is required.")
		return
	}
	r := &record.CustomerRecord{
		Name: name,
	}
	controler.caller.AddCustomer(r)
	return
}

*/

// initialCalls runs the first code that the controler needs to run.
func (controler *panelControler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.
	// example:

	controler.customerSelectWidget.start()

	*/

}
`
