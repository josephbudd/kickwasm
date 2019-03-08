package removecontactconfirmpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactConfirmPanel

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

	record              *types.ContactRecord
	contactRemoveSubmit js.Value
	contactRemoveCancel js.Value
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
	null := js.Null()

	// submit button
	if panelControler.contactRemoveSubmit = notJS.GetElementByID("contactRemoveSubmit"); panelControler.contactRemoveSubmit == null {
		err = errors.New(`unable to find #contactRemoveSubmit`)
		return
	}
	cb := tools.RegisterEventCallBack(panelControler.handleSubmit, true, true, true)
	notJS.SetOnClick(panelControler.contactRemoveSubmit, cb)
	// cancel button
	if panelControler.contactRemoveCancel = notJS.GetElementByID("contactRemoveCancel"); panelControler.contactRemoveCancel == null {
		err = errors.New(`unable to find #contactRemoveCancel`)
		return
	}
	cb = tools.RegisterEventCallBack(panelControler.handleCancel, true, true, true)
	notJS.SetOnClick(panelControler.contactRemoveCancel, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (panelControler *Controler) handleSubmit(event js.Value) interface{} {
	panelControler.caller.removeContact(panelControler.record.ID)
	return nil
}

func (panelControler *Controler) handleCancel(event js.Value) interface{} {
	panelControler.panelGroup.showRemoveContactSelectPanel(false)
	return nil
}

func (panelControler *Controler) handleGetContact(record *types.ContactRecord) {
	panelControler.record = record
	panelControler.presenter.displayRecord(record)
	panelControler.panelGroup.showRemoveContactConfirmPanel(false)
}

// initialCalls runs the first code that the controler needs to run.
func (panelControler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
