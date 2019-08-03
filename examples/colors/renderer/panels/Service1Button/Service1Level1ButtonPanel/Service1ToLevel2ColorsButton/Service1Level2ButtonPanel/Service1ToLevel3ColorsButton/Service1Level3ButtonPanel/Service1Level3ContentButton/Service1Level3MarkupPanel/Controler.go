package service1level3markuppanel

import (
	"github.com/pkg/errors"
)

/*

	Panel name: Service1Level3MarkupPanel

*/

// panelControler controls user input.
type panelControler struct {
	group      *panelGroup
	presenter  *panelPresenter
	caller     *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelControler members.
	// example:

	// import "syscall/js"

	addCustomerName   js.Value
	addCustomerSubmit js.Value

	*/
}

// defineControlsSetHandlers defines controler members and sets their handlers.
// Returns the error.
func (controler *panelControler) defineControlsSetHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controler *panelControler) defineControlsSetHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set their handlers.
	// example:

	// import "syscall/js"

	var cb js.Func

	// Define the customer name input field.
	if controler.addCustomerName = notJS.GetElementByID("addCustomerName"); controler.addCustomerName == null {
		err = errors.New("unable to find #addCustomerName")
		return
	}

	// Define the submit button and set it's handler.
	if controler.addCustomerSubmit = notJS.GetElementByID("addCustomerSubmit"); controler.addCustomerSubmit == null {
		err = errors.New("unable to find #addCustomerSubmit")
		return
	}
	// see render/viewtools/callback.go
	// use the event call back func and set propagations.
	cb = tools.RegisterEventCallBack(controler.handleSubmit, true, true, true)
	notJS.SetOnClick(controler.addCustomerSubmit, cb)

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.
// example:

// import "github.com/josephbudd/kickwasm/examples/colors/domain/store/record"

func (controler *panelControler) handleSubmit(event js.Value) (nilReturn interface{}) {
	name := strings.TrimSpace(notJS.GetValue(controler.addCustomerName))
	if len(name) == 0 {
		tools.Error("Customer Name is required.")
		return
	}
	r := &record.Customer{
		Name: name,
	}
	controler.caller.AddCustomer(r)
	return
}

*/

// initialCalls runs the first code that the controler needs to run.
func (controler *panelControler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.
	// example:

	controler.customerSelectWidget.start()

	*/

}
