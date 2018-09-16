package AddContactPanel

import (
	"syscall/js"

	"github.com/josephbudd/kick/examples/fvlist/mainprocess/ports/records"
	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: AddContactPanel
	Panel id:   tabsMasterView-home-pad-AddButton-AddContactPanel

*/

// Presenter writes to the panel
type Presenter struct {
	panel     *Panel
	controler *Controler
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools
	notjs     *kicknotjs.NotJS

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your Presenter members here.

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
}

// defineMembers defines the Presenter members by their html elements.
func (panelPresenter *Presenter) defineMembers() {

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.

	*/

	notjs := panelPresenter.notjs
	panelPresenter.contactAddName = notjs.GetElementByID("contactAddName")
	panelPresenter.contactAddAddress1 = notjs.GetElementByID("contactAddAddress1")
	panelPresenter.contactAddAddress2 = notjs.GetElementByID("contactAddAddress2")
	panelPresenter.contactAddCity = notjs.GetElementByID("contactAddCity")
	panelPresenter.contactAddState = notjs.GetElementByID("contactAddState")
	panelPresenter.contactAddZip = notjs.GetElementByID("contactAddZip")
	panelPresenter.contactAddPhone = notjs.GetElementByID("contactAddPhone")
	panelPresenter.contactAddEmail = notjs.GetElementByID("contactAddEmail")
	panelPresenter.contactAddSocial = notjs.GetElementByID("contactAddSocial")
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

func (panelPresenter *Presenter) showRecord(record *records.ContactRecord) {
	notjs := panelPresenter.notjs
	notjs.SetValue(panelPresenter.contactAddName, record.Name)
	notjs.SetValue(panelPresenter.contactAddAddress1, record.Address1)
	notjs.SetValue(panelPresenter.contactAddAddress2, record.Address2)
	notjs.SetValue(panelPresenter.contactAddCity, record.City)
	notjs.SetValue(panelPresenter.contactAddState, record.State)
	notjs.SetValue(panelPresenter.contactAddZip, record.Zip)
	notjs.SetValue(panelPresenter.contactAddPhone, record.Phone)
	notjs.SetValue(panelPresenter.contactAddEmail, record.Email)
	notjs.SetValue(panelPresenter.contactAddSocial, record.Social)
}

func (panelPresenter *Presenter) clearForm() {
	notjs := panelPresenter.notjs
	notjs.ClearValue(panelPresenter.contactAddName)
	notjs.ClearValue(panelPresenter.contactAddAddress1)
	notjs.ClearValue(panelPresenter.contactAddAddress2)
	notjs.ClearValue(panelPresenter.contactAddCity)
	notjs.ClearValue(panelPresenter.contactAddState)
	notjs.ClearValue(panelPresenter.contactAddZip)
	notjs.ClearValue(panelPresenter.contactAddPhone)
	notjs.ClearValue(panelPresenter.contactAddEmail)
	notjs.ClearValue(panelPresenter.contactAddSocial)
}
