package EditContactEditPanel

import (
	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/transports/calls"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/states"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: EditContactEditPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactEditPanel

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

	states *states.States
}

// setMainProcessCallBacks tells the main process what funcs to call back to.
func (caller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 1 of 3.

	// Tell the main processs to call back to your funcs.

	*/

	caller.connection.UpdateContact.AddCallBack(caller.updateContactCB)
	caller.connection.GetContact.AddCallBack(caller.getContactCB)
}

/* NOTE TO DEVELOPER. Step 2 of 3.

// Define calls to the main process and their and call backs.

*/

// Get Contact.

func (caller *Caller) getContactCB(params interface{}) {
	switch params := params.(type) {
	case *calls.MainProcessToRendererGetContactParams:
		// can not access Panel.Caller
		if params.State&caller.states.Edit == caller.states.Edit {
			if params.Error {
				// let the select panel caller handle this.
				return
			}
			// no errors so let the user edit this record.
			caller.controler.handleGetContact(params.Record)
		}
	}
}

// Update Contact.

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
		// no errors
		caller.tools.Success("Contact Added.")
	}
}

// initialCalls makes the first calls to the main process.
func (caller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 3 of 3.

	// Make any initial calls to the main process that must be made when the app starts.

	*/

}
