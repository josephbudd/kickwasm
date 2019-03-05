package Service2Level2MarkupPanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/colors/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/viewtools"
)

/*

	Panel name: Service2Level2MarkupPanel

*/

// Controler is a HelloPanel Controler.
type Controler struct {
	panelGroup *PanelGroup
	presenter *Presenter
	caller    *Caller
	quitCh    chan struct{}    // send an empty struct to start the quit process.
	tools     *viewtools.Tools // see /renderer/viewtools
	notJS     *notjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Controler members.
	// example:

	// import "syscall/js"

	addCustomerName   js.Value
	addCustomerSubmit js.Value

	*/
}

// defineControlsSetHandlers defines controler members and sets their handlers.
// Returns the error.
func (panelControler *Controler) defineControlsSetHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelControler *Controler) defineControlsSetHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set their handlers.
	// example:

	// import "syscall/js"

	notJS := panelControler.notJS
	tools := panelControler.tools
	null := js.Null()

	// Define the customer name input field.
	if panelControler.customerName = notJS.GetElementByID("customerName"); panelControler.customerName == null {
		err = errors.New("unable to find #customerName")
		return
	}

	// Define the submit button and set it's handler.
	if panelControler.addCustomerSubmit = notJS.GetElementByID("addCustomerSubmit"); panelControler.addCustomerSubmit == null {
		err = errors.New("unable to find #addCustomerSubmit")
		return
	}
	cb := notJS.RegisterCallBack(panelControler.handleSubmit)
	notJS.SetOnClick(panelControler.addCustomerSubmit, cb)

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.
// example:

// import "github.com/josephbudd/kickwasm/examples/colors/domain/types"

func (panelControler *Controler) handleSubmit(this js.Value, args []js.Value) interface{} {
	name := strings.TrimSpace(panelControler.notJS.GetValue(panelControler.addCustomerName))
	if len(name) == 0 {
		panelControler.tools.Error("Customer Name is required.")
		return nil
	}
	record := &types.CustomerRecord{
		Name: name,
	}
	panelControler.caller.AddCustomer(record)
	return nil
}

*/

// initialCalls runs the first code that the controler needs to run.
func (panelControler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.
	// example:

	panelControler.customerSelectWidget.start()

	*/

}

