package AddContactPanel

import (
	"syscall/js"

	"github.com/josephbudd/kick/examples/fvlist/mainprocess/ports/records"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
	"github.com/pkg/errors"
)

/*

	Panel name: AddContactPanel

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

	if panelPresenter.contactAddName = notjs.GetElementByID("contactAddName"); panelPresenter.contactAddName == null {
		err = errors.New("unable to find #contactAddName")
		return
	}
	if panelPresenter.contactAddAddress1 = notjs.GetElementByID("contactAddAddress1"); panelPresenter.contactAddAddress1 == null {
		err = errors.New("unable to find #contactAddAddress1")
		return
	}
	if panelPresenter.contactAddAddress2 = notjs.GetElementByID("contactAddAddress2"); panelPresenter.contactAddAddress2 == null {
		err = errors.New("unable to find #contactAddAddress2")
		return
	}
	if panelPresenter.contactAddCity = notjs.GetElementByID("contactAddCity"); panelPresenter.contactAddCity == null {
		err = errors.New("unable to find #contactAddCity")
		return
	}
	if panelPresenter.contactAddState = notjs.GetElementByID("contactAddState"); panelPresenter.contactAddState == null {
		err = errors.New("unable to find #contactAddState")
		return
	}
	if panelPresenter.contactAddZip = notjs.GetElementByID("contactAddZip"); panelPresenter.contactAddZip == null {
		err = errors.New("unable to find #contactAddZip")
		return
	}
	if panelPresenter.contactAddPhone = notjs.GetElementByID("contactAddPhone"); panelPresenter.contactAddPhone == null {
		err = errors.New("unable to find #contactAddPhone")
		return
	}
	if panelPresenter.contactAddEmail = notjs.GetElementByID("contactAddEmail"); panelPresenter.contactAddEmail == null {
		err = errors.New("unable to find #contactAddEmail")
		return
	}
	if panelPresenter.contactAddSocial = notjs.GetElementByID("contactAddSocial"); panelPresenter.contactAddSocial == null {
		err = errors.New("unable to find #contactAddSocial")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

func (panelPresenter *Presenter) showRecord(record *records.ContactRecord) {
	notjs := panelPresenter.notJS
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
	notjs := panelPresenter.notJS
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
