package EditContactEditPanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: EditContactEditPanel

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
func (panelPresenter *Presenter) defineMembers() (err error) {
	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelPresenter *Presenter) defineMembers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.

	*/

	notjs := panelPresenter.notJS
	null := js.Null()

	if panelPresenter.contactEditName = notjs.GetElementByID("contactEditName"); panelPresenter.contactEditName == null {
		err = errors.New(`unable to find #contactEditName`)
		return
	}
	if panelPresenter.contactEditAddress1 = notjs.GetElementByID("contactEditAddress1"); panelPresenter.contactEditAddress1 == null {
		err = errors.New(`unable to find #contactEditAddress1`)
		return
	}
	if panelPresenter.contactEditAddress2 = notjs.GetElementByID("contactEditAddress2"); panelPresenter.contactEditAddress2 == null {
		err = errors.New(`unable to find #contactEditAddress2`)
		return
	}
	if panelPresenter.contactEditCity = notjs.GetElementByID("contactEditCity"); panelPresenter.contactEditCity == null {
		err = errors.New(`unable to find #contactEditCity`)
		return
	}
	if panelPresenter.contactEditState = notjs.GetElementByID("contactEditState"); panelPresenter.contactEditState == null {
		err = errors.New(`unable to find #contactEditState`)
		return
	}
	if panelPresenter.contactEditZip = notjs.GetElementByID("contactEditZip"); panelPresenter.contactEditZip == null {
		err = errors.New(`unable to find #contactEditZip`)
		return
	}
	if panelPresenter.contactEditPhone = notjs.GetElementByID("contactEditPhone"); panelPresenter.contactEditPhone == null {
		err = errors.New(`unable to find #contactEditPhone`)
		return
	}
	if panelPresenter.contactEditEmail = notjs.GetElementByID("contactEditEmail"); panelPresenter.contactEditEmail == null {
		err = errors.New(`unable to find #contactEditEmail`)
		return
	}
	if panelPresenter.contactEditSocial = notjs.GetElementByID("contactEditSocial"); panelPresenter.contactEditSocial == null {
		err = errors.New(`unable to find #contactEditSocial`)
		return
	}

	return
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
