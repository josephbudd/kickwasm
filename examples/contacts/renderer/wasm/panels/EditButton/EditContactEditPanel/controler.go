package EditContactEditPanel

import (
	//"syscall/js"

	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
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
	tools     *viewtools.Tools // see /renderer/wasm/viewtools
	notjs     *kicknotjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Controler members.

	*/

	record *records.ContactRecord

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
	contactEditCancel   js.Value
}

// defineControlsSetHandlers defines controler members and sets their handlers.
func (controler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notjs := controler.notjs
	controler.contactEditName = notjs.GetElementByID("contactEditName")
	controler.contactEditAddress1 = notjs.GetElementByID("contactEditAddress1")
	controler.contactEditAddress2 = notjs.GetElementByID("contactEditAddress2")
	controler.contactEditCity = notjs.GetElementByID("contactEditCity")
	controler.contactEditState = notjs.GetElementByID("contactEditState")
	controler.contactEditZip = notjs.GetElementByID("contactEditZip")
	controler.contactEditPhone = notjs.GetElementByID("contactEditPhone")
	controler.contactEditEmail = notjs.GetElementByID("contactEditEmail")
	controler.contactEditSocial = notjs.GetElementByID("contactEditSocial")
	controler.contactEditSubmit = notjs.GetElementByID("contactEditSubmit")
	controler.contactEditCancel = notjs.GetElementByID("contactEditCancel")

	cb := notjs.RegisterCallBack(controler.handleSubmit)
	notjs.SetOnClick(controler.contactEditSubmit, cb)
	cb = notjs.RegisterCallBack(controler.handleCancel)
	notjs.SetOnClick(controler.contactEditCancel, cb)
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controler *Controler) handleGetContact(record *records.ContactRecord) {
	controler.presenter.fillForm(record)
	controler.panel.showEditContactEditPanel(false)
}

func (controler *Controler) handleCancel(args []js.Value) {
	controler.panel.showEditContactSelectPanel(false)
}

func (controler *Controler) handleSubmit(args []js.Value) {
	tools := controler.tools
	record := controler.getForm()
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
	controler.caller.updateContact(record)
}

func (controler *Controler) getForm() *records.ContactRecord {
	notjs := controler.notjs
	return &records.ContactRecord{
		ID:       controler.record.ID,
		Name:     notjs.GetValue(controler.contactEditName),
		Address1: notjs.GetValue(controler.contactEditAddress1),
		Address2: notjs.GetValue(controler.contactEditAddress2),
		City:     notjs.GetValue(controler.contactEditCity),
		State:    notjs.GetValue(controler.contactEditState),
		Phone:    notjs.GetValue(controler.contactEditPhone),
		Email:    notjs.GetValue(controler.contactEditEmail),
		Social:   notjs.GetValue(controler.contactEditSocial),
	}
}

// initialCalls runs the first code that the controler needs to run.
func (controler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.

	*/

}
