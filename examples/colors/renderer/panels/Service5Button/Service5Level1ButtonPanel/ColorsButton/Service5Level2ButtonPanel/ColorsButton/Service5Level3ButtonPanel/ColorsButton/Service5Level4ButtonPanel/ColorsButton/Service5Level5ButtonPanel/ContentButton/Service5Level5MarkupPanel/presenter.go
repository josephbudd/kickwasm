package Service5Level5MarkupPanel

import (
	//"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/colors/renderer/viewtools"
)

/*

	Panel name: Service5Level5MarkupPanel
	Panel id:   tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ColorsButton-Service5Level5ButtonPanel-ContentButton-Service5Level5MarkupPanel

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
