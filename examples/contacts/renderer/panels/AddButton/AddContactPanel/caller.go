package addcontactpanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: AddContactPanel

*/

// Caller communicates with the main process via an asynchrounous connection.
type Caller struct {
	panelGroup *PanelGroup
	presenter  *Presenter
	controler  *Controler
	quitCh     chan struct{} // send an empty struct to start the quit process.
	connection map[types.CallID]caller.Renderer
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// 1: Declare your Caller members.

	*/

	// my added members
	state uint64
	// my added callers
	// interfaces which can be mocked in tests.
	updateContactConnection caller.Renderer
}

// addMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelCaller *Caller) addMainProcessCallBacks()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// 2.1: Define each one of your Caller connection members as a conection to the main process.
	// 2.2: Tell the caller connection to the main processs to add a call back to each of your call back funcs.

	*/

	var found bool

	if panelCaller.updateContactConnection, found = panelCaller.connection[callids.UpdateContactCallID]; !found {
		err = errors.New("unable to find panelCaller.connection[callids.UpdateContactCallID]")
		return
	}
	panelCaller.updateContactConnection.AddCallBack(panelCaller.updateContactCB)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// 3.1: Define your funcs which call to the main process.
// 3.2: Define your funcs which the main process calls back to.

*/

// UpdateContact

func (panelCaller *Caller) updateContact(record *types.ContactRecord) {
	params := &types.RendererToMainProcessUpdateContactParams{
		Record: record,
		State:  panelCaller.state,
	}
	panelCaller.updateContactConnection.CallMainProcess(params)
}

func (panelCaller *Caller) updateContactCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererUpdateContactParams:
		if params.State&panelCaller.state == panelCaller.state {
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
			// No errors.
			// Update the user and clear the form.
			panelCaller.tools.Success("Contact added.")
			panelCaller.presenter.clearForm()
			panelCaller.panelGroup.showAddContactPanel(false)
		}
	default:
		// default should only happen during development.
		// It means that I screwed up in the mainprocess func "mainProcessReceiveUpdateContact". It means that I passed the wrong type of param to callBackToRenderer.
		panelCaller.tools.Error("Wrong param type send from mainProcessReceiveUpdateContact")
		panelCaller.panelGroup.showAddContactPanel(false)
	}
}

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	//4: Make any initial calls to the main process that must be made when the app starts.

	*/

}
