package AddContactPanel

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: AddContactPanel
	Panel id:   tabsMasterView-home-pad-AddButton-AddContactPanel

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

	contactAddName     js.Value
	contactAddAddress1 js.Value
	contactAddAddress2 js.Value
	contactAddCity     js.Value
	contactAddState    js.Value
	contactAddZip      js.Value
	contactAddPhone    js.Value
	contactAddEmail    js.Value
	contactAddSocial   js.Value
	contactAddSubmit   js.Value
	contactAddCancel   js.Value
}

// defineControlsSetHandlers defines controler members and sets their handlers.
func (panelControler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notjs := panelControler.notJS
	panelControler.contactAddName = notjs.GetElementByID("contactAddName")
	panelControler.contactAddAddress1 = notjs.GetElementByID("contactAddAddress1")
	panelControler.contactAddAddress2 = notjs.GetElementByID("contactAddAddress2")
	panelControler.contactAddCity = notjs.GetElementByID("contactAddCity")
	panelControler.contactAddState = notjs.GetElementByID("contactAddState")
	panelControler.contactAddZip = notjs.GetElementByID("contactAddZip")
	panelControler.contactAddPhone = notjs.GetElementByID("contactAddPhone")
	panelControler.contactAddEmail = notjs.GetElementByID("contactAddEmail")
	panelControler.contactAddSocial = notjs.GetElementByID("contactAddSocial")

	panelControler.contactAddSubmit = notjs.GetElementByID("contactAddSubmit")
	panelControler.contactAddCancel = notjs.GetElementByID("contactAddCancel")

	cb := notjs.RegisterCallBack(panelControler.handleSubmit)
	notjs.SetOnClick(panelControler.contactAddSubmit, cb)
	cb = notjs.RegisterCallBack(panelControler.handleCancel)
	notjs.SetOnClick(panelControler.contactAddCancel, cb)
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (panelControler *Controler) handleSubmit(args []js.Value) {
	record := panelControler.getRecord()
	tools := panelControler.tools
	if len(record.Name) == 0 {
		tools.Alert(fmt.Sprintf("len(record.Name) is %d, %q", len(record.Name), record.Name))
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

func (panelControler *Controler) handleCancel(args []js.Value) {
	panelControler.presenter.clearForm()
	panelControler.tools.Back()
}

func (panelControler *Controler) getRecord() *types.ContactRecord {
	notjs := panelControler.notJS
	return &types.ContactRecord{
		Name:     notjs.GetValue(panelControler.contactAddName),
		Address1: notjs.GetValue(panelControler.contactAddAddress1),
		Address2: notjs.GetValue(panelControler.contactAddAddress2),
		City:     notjs.GetValue(panelControler.contactAddCity),
		State:    notjs.GetValue(panelControler.contactAddState),
		Zip:      notjs.GetValue(panelControler.contactAddZip),
		Phone:    notjs.GetValue(panelControler.contactAddPhone),
		Email:    notjs.GetValue(panelControler.contactAddEmail),
		Social:   notjs.GetValue(panelControler.contactAddSocial),
	}
}

// initialCalls runs the first code that the controler needs to run.
func (panelControler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
