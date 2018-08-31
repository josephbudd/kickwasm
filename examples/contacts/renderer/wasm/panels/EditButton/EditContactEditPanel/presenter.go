package EditContactEditPanel

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: EditContactEditPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactEditPanel

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

	contactEditName     js.Value
	contactEditAddress1 js.Value
	contactEditAddress2 js.Value
	contactEditCity     js.Value
	contactEditState    js.Value
	contactEditZip      js.Value
	contactEditPhone    js.Value
	contactEditEmail    js.Value
	contactEditSocial   js.Value
}

// defineMembers defines the Presenter members by their html elements.
func (presenter *Presenter) defineMembers() {

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.

	*/

	notjs := presenter.notjs
	presenter.contactEditName = notjs.GetElementByID("contactEditName")
	presenter.contactEditAddress1 = notjs.GetElementByID("contactEditAddress1")
	presenter.contactEditAddress2 = notjs.GetElementByID("contactEditAddress2")
	presenter.contactEditCity = notjs.GetElementByID("contactEditCity")
	presenter.contactEditState = notjs.GetElementByID("contactEditState")
	presenter.contactEditZip = notjs.GetElementByID("contactEditZip")
	presenter.contactEditPhone = notjs.GetElementByID("contactEditPhone")
	presenter.contactEditEmail = notjs.GetElementByID("contactEditEmail")
	presenter.contactEditSocial = notjs.GetElementByID("contactEditSocial")
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

func (presenter *Presenter) fillForm(record *records.ContactRecord) {
	notjs := presenter.notjs
	notjs.SetValue(presenter.contactEditName, record.Name)
	notjs.SetValue(presenter.contactEditAddress1, record.Address1)
	notjs.SetValue(presenter.contactEditAddress2, record.Address2)
	notjs.SetValue(presenter.contactEditCity, record.City)
	notjs.SetValue(presenter.contactEditState, record.State)
	notjs.SetValue(presenter.contactEditZip, record.Zip)
	notjs.SetValue(presenter.contactEditPhone, record.Phone)
	notjs.SetValue(presenter.contactEditEmail, record.Email)
	notjs.SetValue(presenter.contactEditSocial, record.Social)
}
