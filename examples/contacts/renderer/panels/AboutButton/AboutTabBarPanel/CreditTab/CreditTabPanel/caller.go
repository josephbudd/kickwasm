package credittabpanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: CreditTabPanel

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

	addCustomerConnection caller.Renderer

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

	// 2.1: Define each one of your Caller connection members as a conection to the main process.
	// 2.2: Tell the caller connection to the main processs to add a call back to each of your call back funcs.

	// example:

	import "github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"

	var found bool

	// Add customer.
	// Define the connection.
	if panelCaller.addCustomerConnection, found = panelCaller.connection[callids.AddCustomerCallId]; !found {
		err = errors.New("unable to find panelCaller.connection[callids.AddCustomerCallId]")
		return
	}
	// Have the connection call back to my call back handler.
	panelCaller.addCustomerConnection.AddCallBack(panelCaller.addCustomerCB)

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// 3.1: Define your funcs which call to the main process.
// 3.2: Define your funcs which the main process calls back to.

// example:

// Add Customer.

func (panelCaller *Caller) addCustomer(record *types.CustomerRecord) {
	params := &types.RendererToMainProcessAddCustomerParams{
		Record: record,
	}
	panelCaller.addCustomerConnection.CallMainProcess(params)
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

	import "github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	import "github.com/josephbudd/kickwasm/examples/contacts/domain/data/loglevels"

	params := types.RendererToMainProcessLogParams{
		Level:   loglevels.LogLevelInfo,
		Message: "Started",
	}
	logConnection := panelCaller.connection[callids.LogCallID]
	logConnection.CallMainProcess(params)

	*/

}

