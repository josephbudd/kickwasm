package RemoveContactNotReadyPanel

import (
	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	//"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactNotReadyPanel
	Panel id:   tabsMasterView-home-pad-RemoveButton-RemoveContactNotReadyPanel

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
}

// setMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 1 of 3.

	// Tell the main processs to call back to your funcs.
	// example:

	addCustomerCall := panelCaller.connection[calling.AddCustomerCallId]
	addCustomerCall.AddCallBack(panelCaller.addCustomerCB)

	*/

}

/* NOTE TO DEVELOPER. Step 2 of 3.

// Define calls to the main process and their and call backs.
// example:

// Add Customer.

func (panelCaller *Caller) addCustomer(record *types.CustomerRecord) {
	params := &calls.RendererToMainProcessAddCustomerParams{
		Record: record,
	}
	addCustomerCall := panelCaller.connection[calling.AddCustomerCallId]
	addCustomerCall.CallMainProcess(params)
}

func (panelCaller *Caller) addCustomerCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererAddCustomerParams:
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

	/* NOTE TO DEVELOPER. Step 3 of 3.

	// Make any initial calls to the main process that must be made when the app starts.
	// example:

	params := calls.RendererToMainProcessLogParams{
		Type: calls.LogTypeInfo,
		Message: "Started",
	}
	logCall := panelCaller.connection[calling.LogCallID]
	logCall.CallMainProcess(params)

	*/

}
