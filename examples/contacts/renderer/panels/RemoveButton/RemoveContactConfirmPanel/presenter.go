package RemoveContactConfirmPanel

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactConfirmPanel
	Panel id:   tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel

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

	contactRemoveName     js.Value
	contactRemoveAddress1 js.Value
	contactRemoveAddress2 js.Value
	contactRemoveCity     js.Value
	contactRemoveState    js.Value
	contactRemoveZip      js.Value
	contactRemovePhone    js.Value
	contactRemoveEmail    js.Value
	contactRemoveSocial   js.Value
}

// defineMembers defines the Presenter members by their html elements.
func (panelPresenter *Presenter) defineMembers() {

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.

	*/

	notjs := panelPresenter.notjs
	panelPresenter.contactRemoveName = notjs.GetElementByID("contactRemoveName")
	panelPresenter.contactRemoveAddress1 = notjs.GetElementByID("contactRemoveAddress1")
	panelPresenter.contactRemoveAddress2 = notjs.GetElementByID("contactRemoveAddress2")
	panelPresenter.contactRemoveCity = notjs.GetElementByID("contactRemoveCity")
	panelPresenter.contactRemoveState = notjs.GetElementByID("contactRemoveState")
	panelPresenter.contactRemoveZip = notjs.GetElementByID("contactRemoveZip")
	panelPresenter.contactRemovePhone = notjs.GetElementByID("contactRemovePhone")
	panelPresenter.contactRemoveEmail = notjs.GetElementByID("contactRemoveEmail")
	panelPresenter.contactRemoveSocial = notjs.GetElementByID("contactRemoveSocial")
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

func (panelPresenter *Presenter) displayRecord(record *types.ContactRecord) {
	notjs := panelPresenter.notjs
	notjs.SetInnerText(panelPresenter.contactRemoveName, record.Name)
	notjs.SetInnerText(panelPresenter.contactRemoveAddress1, record.Address1)
	notjs.SetInnerText(panelPresenter.contactRemoveAddress2, record.Address2)
	notjs.SetInnerText(panelPresenter.contactRemoveCity, record.City)
	notjs.SetInnerText(panelPresenter.contactRemoveState, record.State)
	notjs.SetInnerText(panelPresenter.contactRemoveZip, record.Zip)
	notjs.SetInnerText(panelPresenter.contactRemovePhone, record.Phone)
	notjs.SetInnerText(panelPresenter.contactRemoveEmail, record.Email)
	notjs.SetInnerText(panelPresenter.contactRemoveSocial, record.Social)
}
