package RemoveContactSelectPanel

import (
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/widgets"
)

/*

	Panel name: RemoveContactSelectPanel
	Panel id:   tabsMasterView-home-pad-RemoveButton-RemoveContactSelectPanel

*/

// Controler is a HelloPanel Controler.
type Controler struct {
	panel     *Panel
	presenter *Presenter
	caller    *Caller
	quitCh    chan struct{}    // send an empty struct to start the quit process.
	tools     *viewtools.Tools // see /renderer/viewtools
	notJS     *notjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Controler members.

	*/

	contactRemoveSelect *widgets.ContactFVList
}

// defineControlsSetHandlers defines controler members and sets their handlers.
func (panelControler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notjs := panelControler.notJS
	panelControler.contactRemoveSelect = widgets.NewContactFVList(
		// div
		notjs.GetElementByID("contactRemoveSelect"),
		// onSizeFunc
		// Called when there are records in the db.
		func() {
			panelControler.panel.showRemoveContactSelectPanel(false)
		},
		// onNoSizeFunc
		// Called when there are no records in the db.
		func() {
			panelControler.panel.showRemoveContactNotReadyPanel(false)
		},
		// hideFunc
		panelControler.tools.ElementHide,
		// showFunc
		panelControler.tools.ElementShow,
		// isShownFunc
		panelControler.tools.ElementIsShown,
		// notjs
		panelControler.notJS,
		// ContactGetter
		panelControler.caller,
	)
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.
// example:

func (panelControler *Controler) handleSubmit([]js.Value) {
	name := strings.TrimSpace(panelControler.notJS.GetValue(panelControler.addCustomerName))
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

	*/

	panelControler.contactRemoveSelect.Start()

}
