package AddContactPanel

import (
	//"syscall/js"

	"fmt"
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
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
	tools     *viewtools.Tools // see /renderer/wasm/viewtools
	notjs     *kicknotjs.NotJS

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
func (controler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notjs := controler.notjs
	controler.contactAddName = notjs.GetElementByID("contactAddName")
	controler.contactAddAddress1 = notjs.GetElementByID("contactAddAddress1")
	controler.contactAddAddress2 = notjs.GetElementByID("contactAddAddress2")
	controler.contactAddCity = notjs.GetElementByID("contactAddCity")
	controler.contactAddState = notjs.GetElementByID("contactAddState")
	controler.contactAddZip = notjs.GetElementByID("contactAddZip")
	controler.contactAddPhone = notjs.GetElementByID("contactAddPhone")
	controler.contactAddEmail = notjs.GetElementByID("contactAddEmail")
	controler.contactAddSocial = notjs.GetElementByID("contactAddSocial")

	controler.contactAddSubmit = notjs.GetElementByID("contactAddSubmit")
	controler.contactAddCancel = notjs.GetElementByID("contactAddCancel")

	cb := notjs.RegisterCallBack(controler.handleSubmit)
	notjs.SetOnClick(controler.contactAddSubmit, cb)
	cb = notjs.RegisterCallBack(controler.handleCancel)
	notjs.SetOnClick(controler.contactAddCancel, cb)
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controler *Controler) handleSubmit(args []js.Value) {
	record := controler.getRecord()
	tools := controler.tools
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
	controler.caller.updateContact(record)
}

func (controler *Controler) handleCancel(args []js.Value) {
	controler.tools.Back()
}

func (controler *Controler) getRecord() *records.ContactRecord {
	notjs := controler.notjs
	return &records.ContactRecord{
		Name:     notjs.GetValue(controler.contactAddName),
		Address1: notjs.GetValue(controler.contactAddAddress1),
		Address2: notjs.GetValue(controler.contactAddAddress2),
		City:     notjs.GetValue(controler.contactAddCity),
		State:    notjs.GetValue(controler.contactAddState),
		Zip:      notjs.GetValue(controler.contactAddZip),
		Phone:    notjs.GetValue(controler.contactAddPhone),
		Email:    notjs.GetValue(controler.contactAddEmail),
		Social:   notjs.GetValue(controler.contactAddSocial),
	}
}

// initialCalls runs the first code that the controler needs to run.
func (controler *Controler) initialCalls() {}
