package removecontactconfirmpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactConfirmPanel

*/

// Presenter writes to the panel
type Presenter struct {
	panelGroup *PanelGroup
	controler  *Controler
	caller     *Caller
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

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
// Returns the error.
func (panelPresenter *Presenter) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelPresenter *Presenter) defineMembers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.

	*/

	notJS := panelPresenter.notJS
	null := js.Null()

	if panelPresenter.contactRemoveName = notJS.GetElementByID("contactRemoveName"); panelPresenter.contactRemoveName == null {
		err = errors.New(`unable to find #contactRemoveName`)
		return
	}
	if panelPresenter.contactRemoveAddress1 = notJS.GetElementByID("contactRemoveAddress1"); panelPresenter.contactRemoveAddress1 == null {
		err = errors.New(`unable to find #contactRemoveAddress1`)
		return
	}
	if panelPresenter.contactRemoveAddress2 = notJS.GetElementByID("contactRemoveAddress2"); panelPresenter.contactRemoveAddress2 == null {
		err = errors.New(`unable to find #contactRemoveAddress2`)
		return
	}
	if panelPresenter.contactRemoveCity = notJS.GetElementByID("contactRemoveCity"); panelPresenter.contactRemoveCity == null {
		err = errors.New(`unable to find #contactRemoveCity`)
		return
	}
	if panelPresenter.contactRemoveState = notJS.GetElementByID("contactRemoveState"); panelPresenter.contactRemoveState == null {
		err = errors.New(`unable to find #contactRemoveState`)
		return
	}
	if panelPresenter.contactRemoveZip = notJS.GetElementByID("contactRemoveZip"); panelPresenter.contactRemoveZip == null {
		err = errors.New(`unable to find #contactRemoveZip`)
		return
	}
	if panelPresenter.contactRemovePhone = notJS.GetElementByID("contactRemovePhone"); panelPresenter.contactRemovePhone == null {
		err = errors.New(`unable to find #contactRemovePhone`)
		return
	}
	if panelPresenter.contactRemoveEmail = notJS.GetElementByID("contactRemoveEmail"); panelPresenter.contactRemoveEmail == null {
		err = errors.New(`unable to find #contactRemoveEmail`)
		return
	}
	if panelPresenter.contactRemoveSocial = notJS.GetElementByID("contactRemoveSocial"); panelPresenter.contactRemoveSocial == null {
		err = errors.New(`unable to find #contactRemoveSocial`)
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

func (panelPresenter *Presenter) displayRecord(record *types.ContactRecord) {
	notJS := panelPresenter.notJS
	notJS.SetInnerText(panelPresenter.contactRemoveName, record.Name)
	notJS.SetInnerText(panelPresenter.contactRemoveAddress1, record.Address1)
	notJS.SetInnerText(panelPresenter.contactRemoveAddress2, record.Address2)
	notJS.SetInnerText(panelPresenter.contactRemoveCity, record.City)
	notJS.SetInnerText(panelPresenter.contactRemoveState, record.State)
	notJS.SetInnerText(panelPresenter.contactRemoveZip, record.Zip)
	notJS.SetInnerText(panelPresenter.contactRemovePhone, record.Phone)
	notJS.SetInnerText(panelPresenter.contactRemoveEmail, record.Email)
	notJS.SetInnerText(panelPresenter.contactRemoveSocial, record.Social)
}
