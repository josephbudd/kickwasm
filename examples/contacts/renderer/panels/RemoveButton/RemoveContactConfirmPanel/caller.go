package RemoveContactConfirmPanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactConfirmPanel

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
	state                   uint64
	removeContactConnection caller.Renderer
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
	var con caller.Renderer

	if panelCaller.removeContactConnection, found = panelCaller.connection[callids.RemoveContactCallID]; !found {
		err = errors.New(`unable to find panelCaller.connection[callids.RemoveContactCallID]`)
		return
	}
	panelCaller.removeContactConnection.AddCallBack(panelCaller.removeContactCB)

	if con, found = panelCaller.connection[callids.GetContactCallID]; !found {
		err = errors.New(`unable to find panelCaller.connection[callids.GetContactCallID]`)
		return
	}
	con.AddCallBack(panelCaller.getContactCB)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// 3.1: Define your funcs which call to the main process.
// 3.2: Define your funcs which the main process calls back to.

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
	panelCaller.removeContactConnection.CallMainProcess(params)
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

	//4: Make any initial calls to the main process that must be made when the app starts.

	*/

}
