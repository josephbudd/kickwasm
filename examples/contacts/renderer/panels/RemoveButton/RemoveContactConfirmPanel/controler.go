package RemoveContactConfirmPanel

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
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
	tools     *viewtools.Tools // see /renderer/viewtools
	notjs     *kicknotjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Controler members.

	*/

	record              *types.ContactRecord
	contactRemoveSubmit js.Value
	contactRemoveCancel js.Value
}

// defineControlsSetHandlers defines panelControler members and sets their handlers.
func (panelControler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notjs := panelControler.notjs
	panelControler.contactRemoveSubmit = notjs.GetElementByID("contactRemoveSubmit")
	panelControler.contactRemoveCancel = notjs.GetElementByID("contactRemoveCancel")

	cb := notjs.RegisterEventCallBack(false, false, false, panelControler.handleSubmit)
	notjs.SetOnClick(panelControler.contactRemoveSubmit, cb)
	cb = notjs.RegisterEventCallBack(false, false, false, panelControler.handleCancel)
	notjs.SetOnClick(panelControler.contactRemoveCancel, cb)
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (panelControler *Controler) handleSubmit(event js.Value) {
	panelControler.caller.removeContact(panelControler.record.ID)
}

func (panelControler *Controler) handleCancel(event js.Value) {
	panelControler.panel.showRemoveContactSelectPanel(false)
}

func (panelControler *Controler) handleGetContact(record *types.ContactRecord) {
	panelControler.record = record
	panelControler.presenter.displayRecord(record)
	panelControler.panel.showRemoveContactConfirmPanel(false)
}

// initialCalls runs the first code that the panelControler needs to run.
func (panelControler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.
	// example:

	panelControler.customerSelectWidget.start()

	*/

}
