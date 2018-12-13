package EditContactEditPanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: EditContactEditPanel

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

	record *types.ContactRecord

	contactEditName     js.Value
	contactEditAddress1 js.Value
	contactEditAddress2 js.Value
	contactEditCity     js.Value
	contactEditState    js.Value
	contactEditZip      js.Value
	contactEditPhone    js.Value
	contactEditEmail    js.Value
	contactEditSocial   js.Value
	contactEditSubmit   js.Value
	contactEditReset    js.Value
	contactEditCancel   js.Value
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
	null := js.Null()

	// Define the name input.
	if panelControler.contactEditName = notJS.GetElementByID("contactEditName"); panelControler.contactEditName == null {
		err = errors.New(`unable to find #contactEditName`)
		return
	}
	// Define the address 1 input.
	if panelControler.contactEditAddress1 = notJS.GetElementByID("contactEditAddress1"); panelControler.contactEditAddress1 == null {
		err = errors.New(`unable to find #contactEditAddress1`)
		return
	}
	// Define the address 2 input.
	if panelControler.contactEditAddress2 = notJS.GetElementByID("contactEditAddress2"); panelControler.contactEditAddress2 == null {
		err = errors.New(`unable to find #contactEditAddress2`)
		return
	}
	// Define the city input.
	if panelControler.contactEditCity = notJS.GetElementByID("contactEditCity"); panelControler.contactEditCity == null {
		err = errors.New(`unable to find #contactEditCity`)
		return
	}
	// Define the state input.
	if panelControler.contactEditState = notJS.GetElementByID("contactEditState"); panelControler.contactEditState == null {
		err = errors.New(`unable to find #contactEditState`)
		return
	}
	// Define the zip input.
	if panelControler.contactEditZip = notJS.GetElementByID("contactEditZip"); panelControler.contactEditZip == null {
		err = errors.New(`unable to find #contactEditZip`)
		return
	}
	// Define the phone input.
	if panelControler.contactEditPhone = notJS.GetElementByID("contactEditPhone"); panelControler.contactEditPhone == null {
		err = errors.New(`unable to find #contactEditPhone`)
		return
	}
	// Define the email input.
	if panelControler.contactEditEmail = notJS.GetElementByID("contactEditEmail"); panelControler.contactEditEmail == null {
		err = errors.New(`unable to find #contactEditEmail`)
		return
	}
	// Define the social input.
	if panelControler.contactEditSocial = notJS.GetElementByID("contactEditSocial"); panelControler.contactEditSocial == null {
		err = errors.New(`unable to find #contactEditSocial`)
		return
	}
	// Define the submit button and set it's handler.
	if panelControler.contactEditSubmit = notJS.GetElementByID("contactEditSubmit"); panelControler.contactEditSubmit == null {
		err = errors.New(`unable to find #contactEditSubmit`)
		return
	}
	cb := notJS.RegisterCallBack(panelControler.handleSubmit)
	notJS.SetOnClick(panelControler.contactEditSubmit, cb)
	// Define the reset button and set it's handler.
	if panelControler.contactEditReset = notJS.GetElementByID("contactEditReset"); panelControler.contactEditReset == null {
		err = errors.New(`unable to find #contactEditReset`)
		return
	}
	cb = notJS.RegisterCallBack(panelControler.handleReset)
	notJS.SetOnClick(panelControler.contactEditReset, cb)
	// Define the cancel button and set it's handler.
	if panelControler.contactEditCancel = notJS.GetElementByID("contactEditCancel"); panelControler.contactEditCancel == null {
		err = errors.New(`unable to find #contactEditCancel`)
		return
	}
	cb = notJS.RegisterCallBack(panelControler.handleCancel)
	notJS.SetOnClick(panelControler.contactEditCancel, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (panelControler *Controler) handleGetContact(record *types.ContactRecord) {
	panelControler.record = record
	panelControler.presenter.fillForm(record)
	panelControler.panelGroup.showEditContactEditPanel(false)
}

func (panelControler *Controler) handleCancel(args []js.Value) {
	panelControler.panelGroup.showEditContactSelectPanel(false)
}

func (panelControler *Controler) handleSubmit(args []js.Value) {
	tools := panelControler.tools
	record := panelControler.getForm()
	if len(record.Name) == 0 {
		tools.Error("Name is required.")
		return
	}
	if len(record.Address1) == 0 {
		tools.Error("Address1 is required.")
		return
	}
	if len(record.City) == 0 {
		tools.Error("City is required.")
		return
	}
	if len(record.State) == 0 {
		tools.Error("State is required.")
		return
	}
	if len(record.Zip) == 0 {
		tools.Error("Zip is required.")
		return
	}
	if len(record.Email) == 0 && len(record.Phone) == 0 {
		tools.Error("Either Email or Phone is required.")
		return
	}
	panelControler.caller.updateContact(record)
}

func (panelControler *Controler) getForm() *types.ContactRecord {
	notJS := panelControler.notJS
	return &types.ContactRecord{
		ID:       panelControler.record.ID,
		Name:     notJS.GetValue(panelControler.contactEditName),
		Address1: notJS.GetValue(panelControler.contactEditAddress1),
		Address2: notJS.GetValue(panelControler.contactEditAddress2),
		City:     notJS.GetValue(panelControler.contactEditCity),
		State:    notJS.GetValue(panelControler.contactEditState),
		Zip:      notJS.GetValue(panelControler.contactEditZip),
		Phone:    notJS.GetValue(panelControler.contactEditPhone),
		Email:    notJS.GetValue(panelControler.contactEditEmail),
		Social:   notJS.GetValue(panelControler.contactEditSocial),
	}
}

func (panelControler *Controler) handleReset(args []js.Value) {
	panelControler.presenter.fillForm(panelControler.record)
}

// initialCalls runs the first code that the controler needs to run.
func (panelControler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
