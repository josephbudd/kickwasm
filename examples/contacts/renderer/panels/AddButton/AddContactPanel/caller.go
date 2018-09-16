package AddContactPanel

import (
	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/states"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
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
	connection types.RendererCallMap
	tools      *viewtools.Tools // see /renderer/viewtools
	notjs      *kicknotjs.NotJS
	// my added members
	serviceStates *states.States
	// my added callers
	// interfaces which can be mocked in tests.
	updateContactCaller caller.Renderer
}

// setMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 1 of 3.

	// Tell the main processs to call back to your funcs.

	*/

	panelCaller.updateContactCaller = panelCaller.connection[calling.UpdateContactCallID]
	panelCaller.updateContactCaller.AddCallBack(panelCaller.updateContactCB)

}

/* NOTE TO DEVELOPER. Step 2 of 3.

// Define calls to the main process and their and call backs.

*/

// UpdateContact

func (panelCaller *Caller) updateContact(record *types.ContactRecord) {
	params := &calling.RendererToMainProcessUpdateContactParams{
		Record: record,
		State:  panelCaller.serviceStates.Add,
	}
	panelCaller.updateContactCaller.CallMainProcess(params)
}

func (panelCaller *Caller) updateContactCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererUpdateContactParams:
		if params.State&panelCaller.serviceStates.Add == panelCaller.serviceStates.Add {
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
			// No errors so show the contact record.
			panelCaller.tools.Success("Contact added.")
			panelCaller.presenter.clearForm()
			panelCaller.panel.showAddContactPanel(false)
		}
	default:
		// default should only happen during development.
		// It means that the mainprocess func "mainProcessReceiveUpdateContact" passed the wrong type of param to callBackToRenderer.
		panelCaller.tools.Error("Wrong param type send from mainProcessReceiveUpdateContact")
		panelCaller.panel.showAddContactPanel(false)
	}
}

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 3 of 3.

	// Make any initial calls to the main process that must be made when the app starts.

	*/

}
