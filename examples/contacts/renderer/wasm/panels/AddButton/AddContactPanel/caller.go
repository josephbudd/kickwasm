package AddContactPanel

import (
	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/transports/calls"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: AddContactPanel
	Panel id:   tabsMasterView-home-pad-AddButton-AddContactPanel

*/

// Caller communicates with the main process via an asynchrounous connection.
type Caller struct {
	panel      *Panel
	presenter  *Presenter
	controler  *Controler
	quitCh     chan struct{} // send an empty struct to start the quit process.
	connection *calls.Calls
	tools      *viewtools.Tools // see /renderer/wasm/viewtools
	notjs      *kicknotjs.NotJS
}

// setMainProcessCallBacks tells the main process what funcs to call back to.
func (caller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 1 of 3.

	// Tell the main processs to call back to your funcs.
	// example:

	caller.connection.AddCustomer.AddCallBack(caller.addCustomerCB)

	*/

	caller.connection.UpdateContact.AddCallBack(caller.updateContactCB)

}

// UpdateContact

func (caller *Caller) updateContact(record *records.ContactRecord) {
	params := &calls.RendererToMainProcessUpdateContactParams{
		Record: record,
	}
	caller.connection.UpdateContact.CallMainProcess(params)
}

func (caller *Caller) updateContactCB(params interface{}) {
	switch params := params.(type) {
	case *calls.MainProcessToRendererUpdateContactParams:
		if params.Error {
			caller.tools.Error(params.ErrorMessage)
			return
		}
		// No errors so show the contact record.
		caller.tools.Success("Contact updated.")
		caller.panel.showAddContactPanel(false)
	default:
		// default should only happen during development.
		// It means that the mainprocess func "mainProcessReceiveUpdateContact" passed the wrong type of param to callBackToRenderer.
		caller.tools.Error("Wrong param type send from mainProcessReceiveUpdateContact")
		caller.panel.showAddContactPanel(false)
	}
}

// initialCalls makes the first calls to the main process.
func (caller *Caller) initialCalls() {}
