package EditContactSelectPanel

import (
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/widgets"
)

/*

	Panel name: EditContactSelectPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactSelectPanel

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

	contactEditSelect *widgets.ContactFVList
}

// defineControlsSetHandlers defines controler members and sets their handlers.
func (panelControler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notjs := panelControler.notJS
	panelControler.contactEditSelect = widgets.NewContactFVList(
		// div
		notjs.GetElementByID("contactEditSelect"),
		// onSizeFunc
		// Called when there are records in the db.
		func() {
			panelControler.panel.showEditContactSelectPanel(false)
		},
		// onNoSizeFunc
		// Called when there are no records in the db.
		func() {
			panelControler.panel.showEditContactNotReadyPanel(false)
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

*/

// initialCalls runs the first code that the controler needs to run.
func (panelControler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

	panelControler.contactEditSelect.Start()

}
