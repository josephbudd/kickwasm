package AddContactPanel

import (
	//"syscall/js"

	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
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
	tools     *viewtools.Tools // see /renderer/wasm/viewtools
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
func (presenter *Presenter) defineMembers() {

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.

	*/

	notjs := presenter.notjs
	presenter.contactAddName = notjs.GetElementByID("contactAddName")
	presenter.contactAddAddress1 = notjs.GetElementByID("contactAddAddress1")
	presenter.contactAddAddress2 = notjs.GetElementByID("contactAddAddress2")
	presenter.contactAddCity = notjs.GetElementByID("contactAddCity")
	presenter.contactAddState = notjs.GetElementByID("contactAddState")
	presenter.contactAddZip = notjs.GetElementByID("contactAddZip")
	presenter.contactAddPhone = notjs.GetElementByID("contactAddPhone")
	presenter.contactAddEmail = notjs.GetElementByID("contactAddEmail")
	presenter.contactAddSocial = notjs.GetElementByID("contactAddSocial")
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

func (presenter *Presenter) showRecord(record *records.ContactRecord) {
	notjs := presenter.notjs
	notjs.SetValue(presenter.contactAddName, record.Name)
	notjs.SetValue(presenter.contactAddAddress1, record.Address1)
	notjs.SetValue(presenter.contactAddAddress2, record.Address2)
	notjs.SetValue(presenter.contactAddCity, record.City)
	notjs.SetValue(presenter.contactAddState, record.State)
	notjs.SetValue(presenter.contactAddZip, record.Zip)
	notjs.SetValue(presenter.contactAddPhone, record.Phone)
	notjs.SetValue(presenter.contactAddEmail, record.Email)
	notjs.SetValue(presenter.contactAddSocial, record.Social)
}
