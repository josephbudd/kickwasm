package AddContactPanel

import (
	//"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: AddContactPanel
	Panel id:   tabsMasterView-home-pad-AddButton-AddContactPanel

*/

// Controler is a HelloPanel Controler.
type Controler struct {
	panel     *Panel
	presenter *Presenter
	caller    *Caller
	quitCh    chan struct{}    // send an empty struct to start the quit process.
	tools     *viewtools.Tools // see /renderer/viewtools
	notjs     *kicknotjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Controler members.
	// example:

	addCustomerName   js.Value
	addCustomerSubmit js.Value

	*/
}

// defineControlsSetHandlers defines controler members and sets their handlers.
func (panelControler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.
	// example:

	// Define controler members.
	notjs := panelControler.notjs
	panelControler.addCustomerName := notjs.GetElementByID("addCustomerName")
	panelControler.addCustomerSubmit := notjs.GetElementByID("addCustomerSubmit")

	// Set handlers.
	cb := notjs.RegisterCallBack(panelControler.handleSubmit)
	notjs.SetOnClick(panelControler.addCustomerSubmit, cb)

	*/
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.
// example:

func (panelControler *Controler) handleSubmit([]js.Value) {
	name := strings.TrimSpace(panelControler.notjs.GetValue(panelControler.addCustomerName))
	if len(name) == 0 {
		panelControler.tools.Error("Customer Name is required.")
		return
	}
	record := &records.Customer{
		Name: name,
	}
	panelControler.caller.AddCustomer(record)
}

*/

// initialCalls runs the first code that the controler needs to run.
func (panelControler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.
	// example:

	panelControler.customerSelectWidget.start()

	*/

}
