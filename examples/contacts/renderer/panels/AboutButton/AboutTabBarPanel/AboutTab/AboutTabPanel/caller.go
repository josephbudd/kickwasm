package abouttabpanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: AboutTabPanel

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

	getAboutConnection caller.Renderer
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

	// Get author, version
	var found bool

	// Define the connection.
	if panelCaller.getAboutConnection, found = panelCaller.connection[callids.GetAboutCallID]; !found {
		err = errors.New("unable to find panelCaller.connection[callids.GetAboutCallID]")
		return
	}
	// Have the connection call back to my call back handler.
	panelCaller.getAboutConnection.AddCallBack(panelCaller.getAboutCB)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// 3.1: Define your funcs which call to the main process.
// 3.2: Define your funcs which the main process calls back to.

*/

// Get About.

func (panelCaller *Caller) getAbout() {
	params := &types.RendererToMainProcessGetAboutParams{}
	panelCaller.getAboutConnection.CallMainProcess(params)
}

func (panelCaller *Caller) getAboutCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererGetAboutParams:
		if params.Error {
			panelCaller.tools.Error(params.ErrorMessage)
			return
		}
		panelCaller.presenter.displayAbout(params.Author, params.Version)
	}
}

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	//4: Make any initial calls to the main process that must be made when the app starts.

	*/

	panelCaller.getAbout()
}
