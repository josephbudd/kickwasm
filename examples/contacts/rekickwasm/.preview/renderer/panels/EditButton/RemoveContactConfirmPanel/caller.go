package RemoveContactConfirmPanel

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
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
	connection map[types.CallID]caller.Renderer
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Caller members.

	*/

	// my added members
	state               uint64
	removeContactCaller caller.Renderer
}

// addMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define your added Caller members.
	// Tell the main processs to call back to your funcs.

	*/

	getContact := panelCaller.connection[callids.GetContactCallID]
	getContact.AddCallBack(panelCaller.getContactCB)

	panelCaller.removeContactCaller = panelCaller.connection[callids.RemoveContactCallID]
	panelCaller.removeContactCaller.AddCallBack(panelCaller.removeContactCB)

}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Define calls to the main process and their and call backs.

*/

// Get Contact

func (panelCaller *Caller) getContactCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererGetContactParams:
		if params.State&panelCaller.state == panelCaller.state {
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
	params := &types.RendererToMainProcessRemoveContactParams{
		ID: id,
	}
	panelCaller.removeContactCaller.CallMainProcess(params)
}

func (panelCaller *Caller) removeContactCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererRemoveContactParams:
		if params.Error {
			panelCaller.tools.Error(params.ErrorMessage)
			return
		}
		panelCaller.tools.Success("Contact removed.")
	}
}

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make any initial calls to the main process that must be made when the app starts.

	*/

}
