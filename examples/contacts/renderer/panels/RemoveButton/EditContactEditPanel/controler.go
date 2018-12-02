package EditContactEditPanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: EditContactEditPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactEditPanel

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
func (panelControler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notjs := panelControler.notJS
	panelControler.contactEditName = notjs.GetElementByID("contactEditName")
	panelControler.contactEditAddress1 = notjs.GetElementByID("contactEditAddress1")
	panelControler.contactEditAddress2 = notjs.GetElementByID("contactEditAddress2")
	panelControler.contactEditCity = notjs.GetElementByID("contactEditCity")
	panelControler.contactEditState = notjs.GetElementByID("contactEditState")
	panelControler.contactEditZip = notjs.GetElementByID("contactEditZip")
	panelControler.contactEditPhone = notjs.GetElementByID("contactEditPhone")
	panelControler.contactEditEmail = notjs.GetElementByID("contactEditEmail")
	panelControler.contactEditSocial = notjs.GetElementByID("contactEditSocial")
	panelControler.contactEditSubmit = notjs.GetElementByID("contactEditSubmit")
	panelControler.contactEditReset = notjs.GetElementByID("contactEditReset")
	panelControler.contactEditCancel = notjs.GetElementByID("contactEditCancel")

	cb := notjs.RegisterCallBack(panelControler.handleSubmit)
	notjs.SetOnClick(panelControler.contactEditSubmit, cb)
	cb = notjs.RegisterCallBack(panelControler.handleReset)
	notjs.SetOnClick(panelControler.contactEditReset, cb)
	cb = notjs.RegisterCallBack(panelControler.handleCancel)
	notjs.SetOnClick(panelControler.contactEditCancel, cb)
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (panelControler *Controler) handleGetContact(record *types.ContactRecord) {
	panelControler.record = record
	panelControler.presenter.fillForm(record)
	panelControler.panel.showEditContactEditPanel(false)
}

func (panelControler *Controler) handleCancel(args []js.Value) {
	panelControler.panel.showEditContactSelectPanel(false)
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
	if len(record.Phone) == 0 {
		tools.Error("Phone is required.")
		return
	}
	if len(record.Email) == 0 {
		tools.Error("Email is required.")
		return
	}
	if len(record.Social) == 0 {
		tools.Error("Social is required.")
		return
	}
	panelControler.caller.updateContact(record)
}

func (panelControler *Controler) getForm() *types.ContactRecord {
	notjs := panelControler.notJS
	return &types.ContactRecord{
		ID:       panelControler.record.ID,
		Name:     notjs.GetValue(panelControler.contactEditName),
		Address1: notjs.GetValue(panelControler.contactEditAddress1),
		Address2: notjs.GetValue(panelControler.contactEditAddress2),
		City:     notjs.GetValue(panelControler.contactEditCity),
		State:    notjs.GetValue(panelControler.contactEditState),
		Zip:      notjs.GetValue(panelControler.contactEditZip),
		Phone:    notjs.GetValue(panelControler.contactEditPhone),
		Email:    notjs.GetValue(panelControler.contactEditEmail),
		Social:   notjs.GetValue(panelControler.contactEditSocial),
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
