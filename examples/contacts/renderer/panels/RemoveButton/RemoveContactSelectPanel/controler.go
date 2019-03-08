package removecontactselectpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/widgets"
)

/*

	Panel name: RemoveContactSelectPanel

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

	contactRemoveSelect *widgets.ContactFVList
}

// defineControlsSetHandlers defines controler members and sets their handlers.
// Returns the error.
func (panelControler *Controler) defineControlsSetHandlers() (err error) {

	defer func() {
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
	if div = notJS.GetElementByID("contactRemoveSelect"); div == js.Null() {
		err = errors.New(`unable to find #contactRemoveSelect`)
		return
	}

	panelControler.contactRemoveSelect = widgets.NewContactFVList(
		// div
		div,
		// onSizeFunc
		// Called when there are records in the db.
		func() {
			panelControler.panelGroup.showRemoveContactSelectPanel(false)
		},
		// onNoSizeFunc
		// Called when there are no records in the db.
		func() {
			panelControler.panelGroup.showRemoveContactNotReadyPanel(false)
		},
		// hideFunc
		tools.ElementHide,
		// showFunc
		tools.ElementShow,
		// isShownFunc
		tools.ElementIsShown,
		// notjs
		notJS,
		// tools
		tools,
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

	panelControler.contactRemoveSelect.Start()

}
