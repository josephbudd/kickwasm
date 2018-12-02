package EditContactEditPanel

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
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
	connection map[types.CallID]caller.Renderer
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Caller members.

	*/

	// my added members
	state uint64
	// my added callers
	// interfaces which can be mocked in tests.
	updateContactCaller caller.Renderer
}

// addMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define your added Caller members.
	// Tell the main processs to call back to your funcs.

	*/

	panelCaller.updateContactCaller = panelCaller.connection[callids.UpdateContactCallID]
	panelCaller.updateContactCaller.AddCallBack(panelCaller.updateContactCB)

	panelCaller.connection[callids.GetContactCallID].AddCallBack(panelCaller.getContactCB)

}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Define calls to the main process and their and call backs.

*/

// UpdateContact

func (panelCaller *Caller) updateContact(record *types.ContactRecord) {
	params := &types.RendererToMainProcessUpdateContactParams{
		Record: record,
		State:  panelCaller.state,
	}
	panelCaller.updateContactCaller.CallMainProcess(params)
}

func (panelCaller *Caller) updateContactCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererUpdateContactParams:
		if params.State&panelCaller.state == panelCaller.state {
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
			// No errors so show the contact record.
			panelCaller.tools.Success("Contact updated.")
			panelCaller.panel.showEditContactSelectPanel(false)
		}
	default:
		// default should only happen during development.
		// It means that the mainprocess func "mainProcessReceiveUpdateContact" passed the wrong type of param to callBackToRenderer.
		panelCaller.tools.Error("Wrong param type send from mainProcessReceiveUpdateContact")
	}
}

// GetContact

func (panelCaller *Caller) getContactCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererGetContactParams:
		if params.State&panelCaller.state == panelCaller.state {
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
			}
			// no error so let the edit panel handle the call back.
			panelCaller.controler.handleGetContact(params.Record)
		}
	}
}

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make any initial calls to the main process that must be made when the app starts.

	*/

}
