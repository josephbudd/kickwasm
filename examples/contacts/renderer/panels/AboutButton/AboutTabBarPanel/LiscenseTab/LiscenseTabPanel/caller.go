package LiscenseTabPanel

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
	"github.com/pkg/errors"
)

/*

	Panel name: LiscenseTabPanel

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

	// example:

	addCustomerCall caller.Renderer

	*/
}

// addMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() (err error) {
	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelCaller *Caller) addMainProcessCallBacks()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// 2.1: Define each one of your added Caller members.
	// 2.2: Tell the main processs to add a call back to each of your call back funcs.

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// 3.1: Define your funcs which call to the main process.
// 3.2: Define your funcs which the main process calls back to.

// example:

import "github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"


// Add Customer.

func (panelCaller *Caller) addCustomer(record *types.CustomerRecord) {
	params := &calls.RendererToMainProcessAddCustomerParams{
		Record: record,
	}
	panelCaller.addCustomerCall.CallMainProcess(params)
}

func (panelCaller *Caller) addCustomerCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererAddCustomerParams:
		if params.Error {
			panelCaller.tools.Error(params.ErrorMessage)
			return
		}
		// no errors
		panelCaller.tools.Success("Customer Added.")
	}
}

*/

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	//4: Make any initial calls to the main process that must be made when the app starts.

	// example:

	params := types.RendererToMainProcessLogParams{
		Level:   loglevels.LogLevelInfo,
		Message: "Started",
	}
	logCall := panelCaller.connection[callids.LogCallID]
	logCall.CallMainProcess(params)

	*/

}