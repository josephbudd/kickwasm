package RemoveContactConfirmPanel

import (
	"github.com/josephbudd/kicknotjs"

	//"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/transports/calls"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
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
	connection *calls.Calls
	tools      *viewtools.Tools // see /renderer/wasm/viewtools
	notjs      *kicknotjs.NotJS
}

// setMainProcessCallBacks tells the main process what funcs to call back to.
func (caller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 1 of 3.

	// Tell the main processs to call back to your funcs.
	// example:

	caller.connection.AddCustomer.AddCallBack(caller.addCustomerCB)

	*/

}

/* NOTE TO DEVELOPER. Step 2 of 3.

// Define calls to the main process and their and call backs.
// example:

// Add Customer.

func (caller *Caller) addCustomer(record *records.CustomerRecord) {
	params := &RendererToMainProcessAddCustomerParams{
		Record: record,
	}
	caller.connection.AddCustomer.CallMainProcess(params)
}

func (caller *Caller) addCustomerCB(params interface{}) {
	switch params := params.(type) {
	case *calls.MainProcessToRendererAddCustomerParams:
		if params.Error {
			caller.tools.Error(params.ErrorMessage)
			return
		}
		// no errors
		caller.Tools.Success("Customer Added.")
	}
}

*/

// initialCalls makes the first calls to the main process.
func (caller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 3 of 3.

	// Make any initial calls to the main process that must be made when the app starts.
	// example:

	params := calls.RendererToMainProcessLogParams{
		Type: calls.LogTypeInfo,
		Message: "Started",
	}
	caller.connection.Log.CallMainProcess(params)

	*/

}
