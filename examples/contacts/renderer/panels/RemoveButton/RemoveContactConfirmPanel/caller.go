package RemoveContactConfirmPanel

import (
	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/states"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactConfirmPanel
	Panel id:   tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel

*/

// Caller communicates with the main process via an asynchrounous connection.
type Caller struct {
	panel      *Panel
	presenter  *Presenter
	controler  *Controler
	quitCh     chan struct{} // send an empty struct to start the quit process.
	connection types.RendererCallMap
	tools      *viewtools.Tools // see /renderer/viewtools
	notjs      *kicknotjs.NotJS
	// my added members
	serviceStates       *states.States
	removeContactCaller caller.Renderer
}

// setMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 1 of 3.

	// Tell the main processs to call back to your funcs.

	*/

	getContact := panelCaller.connection[calling.GetContactCallID]
	getContact.AddCallBack(panelCaller.getContactCB)

	panelCaller.removeContactCaller = panelCaller.connection[calling.RemoveContactCallID]
	panelCaller.removeContactCaller.AddCallBack(panelCaller.removeContactCB)

}

/* NOTE TO DEVELOPER. Step 2 of 3.

// Define calls to the main process and their and call backs.

*/

// Get Contact

func (panelCaller *Caller) getContactCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererGetContactParams:
		if params.State&panelCaller.serviceStates.Remove == panelCaller.serviceStates.Remove {
			if params.Error {
				return
			}
			// no error so let the remove confirm panel handle the call back.
			panelCaller.controler.handleGetContact(params.Record)
		}
	}
}

// Remove Contact

func (panelCaller *Caller) removeContact(id uint64) {
	params := &calling.RendererToMainProcessRemoveContactParams{
		ID: id,
	}
	panelCaller.removeContactCaller.CallMainProcess(params)
}

func (panelCaller *Caller) removeContactCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererRemoveContactParams:
		if params.Error {
			panelCaller.tools.Error(params.ErrorMessage)
			return
		}
		// the select panelCaller needs to handle this.
	}
}

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 3 of 3.

	// Make any initial calls to the main process that must be made when the app starts.

	*/

}
