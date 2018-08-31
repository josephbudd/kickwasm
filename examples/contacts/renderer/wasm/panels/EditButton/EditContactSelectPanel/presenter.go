package EditContactSelectPanel

import (
	//"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: EditContactSelectPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactSelectPanel

*/

// Presenter writes to the panel
type Presenter struct {
	panel     *Panel
	controler *Controler
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/wasm/viewtools
	notjs     *kicknotjs.NotJS

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your Presenter members here.
	// example:

	customerName js.Value

	*/
}

// defineMembers defines the Presenter members by their html elements.
func (presenter *Presenter) defineMembers() {

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.
	// example:

	presenter.customerName = presenter.notjs.GetElementByID("customerName")

	*/
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.
// example:

// displayCustomer displays the customer in the panel.
func (presenter *Presenter) displayCustomer(record *records.CustomerRecord) {
	presenter.notjs.SetInnerText(presenter.customerName, record.Name)
}

*/
