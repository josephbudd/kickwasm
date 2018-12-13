package EditContactSelectPanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/widgets"
	"github.com/pkg/errors"
)

/*

	Panel name: EditContactSelectPanel

*/

// Controler is a HelloPanel Controler.
type Controler struct {
	panelGroup *PanelGroup
	presenter  *Presenter
	caller     *Caller
	quitCh     chan struct{}    // send an empty struct to start the quit process.
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Controler members.

	*/

	contactEditSelect *widgets.ContactFVList
}

// defineControlsSetHandlers defines controler members and sets their handlers.
func (panelControler *Controler) defineControlsSetHandlers() (err error) {
	defer func() {
		// close and check for the error
		if err != nil {
			err = errors.WithMessage(err, "(panelControler *Controler) defineControlsSetHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notJS := panelControler.notJS
	tools := panelControler.tools
	var div js.Value
	if div = notJS.GetElementByID("contactEditSelect"); div == js.Null() {
		err = errors.New(`unable to find #contactEditSelect`)
		return
	}
	panelControler.contactEditSelect = widgets.NewContactFVList(
		// div
		div,
		// onSizeFunc
		// Called when there are records in the db.
		func() {
			panelControler.panelGroup.showEditContactSelectPanel(false)
		},
		// onNoSizeFunc
		// Called when there are no records in the db.
		func() {
			panelControler.panelGroup.showEditContactNotReadyPanel(false)
		},
		// hideFunc
		tools.ElementHide,
		// showFunc
		tools.ElementShow,
		// isShownFunc
		tools.ElementIsShown,
		// notjs
		notJS,
		// ContactGetter
		panelControler.caller,
	)

	return
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
