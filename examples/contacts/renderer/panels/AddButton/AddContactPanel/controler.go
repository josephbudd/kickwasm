package addcontactpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: AddContactPanel

*/

// Controler is a HelloPanel Controler.
type Controler struct {
	panelGroup *PanelGroup
	presenter  *Presenter
	caller     *Caller
	quitCh     chan struct{}    // send an empty struct to start the quit process.
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

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
// Returns the error.
func (panelControler *Controler) defineControlsSetHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelControler *Controler) defineControlsSetHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.

	*/

	notjs := panelControler.notJS
	tools := panelControler.tools
	null := js.Null()

	if panelControler.contactAddName = notjs.GetElementByID("contactAddName"); panelControler.contactAddName == null {
		err = errors.New("unable to find #contactAddName")
		return
	}
	if panelControler.contactAddAddress1 = notjs.GetElementByID("contactAddAddress1"); panelControler.contactAddAddress1 == null {
		err = errors.New("unable to find #contactAddAddress1")
		return
	}
	if panelControler.contactAddAddress2 = notjs.GetElementByID("contactAddAddress2"); panelControler.contactAddAddress2 == null {
		err = errors.New("unable to find #contactAddAddress2")
		return
	}
	if panelControler.contactAddCity = notjs.GetElementByID("contactAddCity"); panelControler.contactAddCity == null {
		err = errors.New("unable to find #contactAddCity")
		return
	}
	if panelControler.contactAddState = notjs.GetElementByID("contactAddState"); panelControler.contactAddState == null {
		err = errors.New("unable to find #contactAddState")
		return
	}
	if panelControler.contactAddZip = notjs.GetElementByID("contactAddZip"); panelControler.contactAddZip == null {
		err = errors.New("unable to find #contactAddZip")
		return
	}
	if panelControler.contactAddPhone = notjs.GetElementByID("contactAddPhone"); panelControler.contactAddPhone == null {
		err = errors.New("unable to find #contactAddPhone")
		return
	}
	if panelControler.contactAddEmail = notjs.GetElementByID("contactAddEmail"); panelControler.contactAddEmail == null {
		err = errors.New("unable to find #contactAddEmail")
		return
	}
	if panelControler.contactAddSocial = notjs.GetElementByID("contactAddSocial"); panelControler.contactAddSocial == null {
		err = errors.New("unable to find #contactAddSocial")
		return
	}

	if panelControler.contactAddSubmit = notjs.GetElementByID("contactAddSubmit"); panelControler.contactAddSubmit == null {
		err = errors.New("unable to find #contactAddSubmit")
		return
	}
	cb := tools.RegisterEventCallBack(panelControler.handleSubmit, true, true, true)
	notjs.SetOnClick(panelControler.contactAddSubmit, cb)

	if panelControler.contactAddCancel = notjs.GetElementByID("contactAddCancel"); panelControler.contactAddCancel == null {
		err = errors.New("unable to find #contactAddCancel")
		return
	}
	cb = tools.RegisterEventCallBack(panelControler.handleCancel, true, true, true)
	notjs.SetOnClick(panelControler.contactAddCancel, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (panelControler *Controler) handleSubmit(event js.Value) interface{} {
	record := panelControler.getRecord()
	tools := panelControler.tools
	if len(record.Name) == 0 {
		tools.Error("Name is required.")
		return nil
	}
	if len(record.Address1) == 0 {
		tools.Error("Address1 is required.")
		return nil
	}
	if len(record.City) == 0 {
		tools.Error("City is required.")
		return nil
	}
	if len(record.State) == 0 {
		tools.Error("State is required.")
		return nil
	}
	if len(record.Zip) == 0 {
		tools.Error("Zip is required.")
		return nil
	}
	if len(record.Email) == 0 && len(record.Phone) == 0 {
		tools.Error("Either Email or Phone is required.")
		return nil
	}
	panelControler.caller.updateContact(record)
	return nil
}

func (panelControler *Controler) handleCancel(event js.Value) interface{} {
	panelControler.presenter.clearForm()
	panelControler.tools.Back()
	return nil
}

func (panelControler *Controler) getRecord() *types.ContactRecord {
	notjs := panelControler.notJS
	return &types.ContactRecord{
		Name:     notjs.GetValue(panelControler.contactAddName),
		Address1: notjs.GetValue(panelControler.contactAddAddress1),
		Address2: notjs.GetValue(panelControler.contactAddAddress2),
		City:     notjs.GetValue(panelControler.contactAddCity),
		State:    notjs.GetValue(panelControler.contactAddState),
		Zip:      notjs.GetValue(panelControler.contactAddZip),
		Phone:    notjs.GetValue(panelControler.contactAddPhone),
		Email:    notjs.GetValue(panelControler.contactAddEmail),
		Social:   notjs.GetValue(panelControler.contactAddSocial),
	}
}

// initialCalls runs the first code that the controler needs to run.
func (panelControler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
