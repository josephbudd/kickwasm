package EditContactEditPanel

import (
	//"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: EditContactEditPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactEditPanel

*/

// Presenter writes to the panel
type Presenter struct {
	panel     *Panel
	controler *Controler
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools
	notjs     *kicknotjs.NotJS

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your Presenter members here.
	// example:

	customerName js.Value

	*/
}

// defineMembers defines the Presenter members by their html elements.
func (panelPresenter *Presenter) defineMembers() {

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.
	// example:

	panelPresenter.customerName = panelPresenter.notjs.GetElementByID("customerName")

	*/
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.
// example:

// displayCustomer displays the customer in the panel.
func (panelPresenter *Presenter) displayCustomer(record *records.CustomerRecord) {
	panelPresenter.notjs.SetInnerText(panelPresenter.customerName, record.Name)
}

*/
