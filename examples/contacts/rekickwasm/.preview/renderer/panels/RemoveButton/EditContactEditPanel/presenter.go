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

// Presenter writes to the panel
type Presenter struct {
	panel     *Panel
	controler *Controler
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools
	notJS     *notjs.NotJS

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
func (panelPresenter *Presenter) defineMembers() {

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.

	*/

	notjs := panelPresenter.notJS
	panelPresenter.contactEditName = notjs.GetElementByID("contactEditName")
	panelPresenter.contactEditAddress1 = notjs.GetElementByID("contactEditAddress1")
	panelPresenter.contactEditAddress2 = notjs.GetElementByID("contactEditAddress2")
	panelPresenter.contactEditCity = notjs.GetElementByID("contactEditCity")
	panelPresenter.contactEditState = notjs.GetElementByID("contactEditState")
	panelPresenter.contactEditZip = notjs.GetElementByID("contactEditZip")
	panelPresenter.contactEditPhone = notjs.GetElementByID("contactEditPhone")
	panelPresenter.contactEditEmail = notjs.GetElementByID("contactEditEmail")
	panelPresenter.contactEditSocial = notjs.GetElementByID("contactEditSocial")
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

func (panelPresenter *Presenter) fillForm(record *types.ContactRecord) {
	notjs := panelPresenter.notJS
	notjs.SetValue(panelPresenter.contactEditName, record.Name)
	notjs.SetValue(panelPresenter.contactEditAddress1, record.Address1)
	notjs.SetValue(panelPresenter.contactEditAddress2, record.Address2)
	notjs.SetValue(panelPresenter.contactEditCity, record.City)
	notjs.SetValue(panelPresenter.contactEditState, record.State)
	notjs.SetValue(panelPresenter.contactEditZip, record.Zip)
	notjs.SetValue(panelPresenter.contactEditPhone, record.Phone)
	notjs.SetValue(panelPresenter.contactEditEmail, record.Email)
	notjs.SetValue(panelPresenter.contactEditSocial, record.Social)
}
