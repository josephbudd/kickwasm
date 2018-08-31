package RemoveContactConfirmPanel

import (
	//"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: RemoveContactConfirmPanel
	Panel id:   tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel

*/

// Controler is a HelloPanel Controler.
type Controler struct {
	panel     *Panel
	presenter *Presenter
	caller    *Caller
	quitCh    chan struct{}    // send an empty struct to start the quit process.
	tools     *viewtools.Tools // see /renderer/wasm/viewtools
	notjs     *kicknotjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Controler members.
	// example:

	addCustomerName   js.Value
	addCustomerSubmit js.Value

	*/
}

// defineControlsSetHandlers defines controler members and sets their handlers.
func (controler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.
	// example:

	// Define controler members.
	notjs := controler.notjs
	controler.addCustomerName := notjs.GetElementByID("addCustomerName")
	controler.addCustomerSubmit := notjs.GetElementByID("addCustomerSubmit")

	// Set handlers.
	cb := notjs.RegisterCallBack(controler.handleSubmit)
	notjs.SetOnClick(controler.addCustomerSubmit, cb)

	*/
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.
// example:

func (controler *Controler) handleSubmit([]js.Value) {
	name := strings.TrimSpace(controler.notjs.GetValue(controler.addCustomerName))
	if len(name) == 0 {
		controler.tools.Error("Customer Name is required.")
		return
	}
	record := &records.Customer{
		Name: name,
	}
	controler.caller.AddCustomer(record)
}

*/

// initialCalls runs the first code that the controler needs to run.
func (controler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.
	// example:

	controler.customerSelectWidget.start()

	*/

}
