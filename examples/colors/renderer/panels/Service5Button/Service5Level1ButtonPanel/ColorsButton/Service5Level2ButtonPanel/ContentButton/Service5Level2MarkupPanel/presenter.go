package service5level2markuppanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/colors/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/viewtools"
)

/*

	Panel name: Service5Level2MarkupPanel

*/

// Presenter writes to the panel
type Presenter struct {
	panelGroup *PanelGroup
	controler  *Controler
	caller     *Caller
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your Presenter members here.
	// example:

	// import "syscall/js"

	customerName js.Value

	*/
}

// defineMembers defines the Presenter members by their html elements.
// Returns the error.
func (panelPresenter *Presenter) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelPresenter *Presenter) defineMembers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.
	// example:

	// import "syscall/js"

	notJS := panelPresenter.notJS
	tools := panelPresenter.tools
	null := js.Null()

	// Define the customer name input field.
	if panelPresenter.customerName = notJS.GetElementByID("customerName"); panelPresenter.customerName == null {
		err = errors.New("unable to find #customerName")
		return
	}

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.
// example:

// displayCustomer displays the customer in the panel.
func (panelPresenter *Presenter) displayCustomer(record *types.CustomerRecord) {
	panelPresenter.notJS.SetInnerText(panelPresenter.customerName, record.Name)
}

*/
