package EditContactSelectPanel

import (
	"github.com/josephbudd/kicknotjs"

	//"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/transports/calls"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/states"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: EditContactSelectPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactSelectPanel

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

	caller.connection.GetContactsPage.AddCallBack(caller.getContactsPageCB)
	caller.connection.UpdateContact.AddCallBack(caller.updateContactCB)
	caller.connection.GetContact.AddCallBack(caller.getContactCB)

}

// GetState returns this caller's state which is the controler's select id.
func (caller *Caller) GetState() uint64 {
	return caller.controler.contactEditSelectID
}

/* NOTE TO DEVELOPER. Step 2 of 3.

// Define calls to the main process and their and call backs.

*/

// UpdateContact

func (caller *Caller) updateContactCB(params interface{}) {
	switch params := params.(type) {
	case *calls.MainProcessToRendererUpdateContactParams:
		if params.Error {
			return
		}
		// the contacts repo has been changed to restart the contact selector.
		caller.controler.contactEditSelect.Start()
	}
}

// GetContact

func (caller *Caller) getContact(id uint64) {
	params := &calls.RendererToMainProcessGetContactParams{
		ID:    id,
		State: caller.controler.contactEditSelectID,
	}
	caller.connection.GetContact.CallMainProcess(params)
}

func (caller *Caller) getContactCB(params interface{}) {
	switch params := params.(type) {
	case *calls.MainProcessToRendererGetContactParams:
		if params.State&caller.states.Edit == caller.states.Edit {
			if params.Error {
				caller.tools.Error(params.ErrorMessage)
			}
			// no error so let the edit panel handle the call back.
		}
	}
}

// GetContactsPage

func (caller *Caller) getContactsPage(sortedIndex, pageSize, state uint64) {
	params := &calls.RendererToMainProcessGetContactsPageParams{
		SortedIndex: sortedIndex,
		PageSize:    pageSize,
		State:       state & caller.states.Edit,
	}
	caller.connection.GetContactsPage.CallMainProcess(params)
}

func (caller *Caller) getContactsPageCB(params interface{}) {
	switch params := params.(type) {
	case *calls.MainProcessToRendererGetContactsPageParams:
		if params.State&caller.states.Edit == caller.states.Edit {
			if params.Error {
				caller.tools.Error(params.ErrorMessage)
				return
			}
			// ok
			caller.controler.handleContactsPage(params.Records, params.SortedIndex, params.RecordCount, params.State)
		}
	}
}

// initialCalls makes the first calls to the main process.
func (caller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 3 of 3.

	// Make any initial calls to the main process that must be made when the app starts.

	*/

}
