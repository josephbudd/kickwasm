// +build js, wasm

package pushpanel

import (
	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/markup"
	"github.com/pkg/errors"
)

/*

	Panel name: PushPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group      *panelGroup
	controller *panelController
	messenger  *panelMessenger

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

	// example:

	editCustomerName *markup.Element

	*/

	timeSpan *markup.Element
}

// defineMembers defines the panelPresenter members by their html elements.
// Returns the error.
func (presenter *panelPresenter) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(presenter *panelPresenter) defineMembers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your panelPresenter members.

	// example:

	// Define the edit form's customer name input field.
	if presenter.editCustomerName = document.ElementByID("editCustomerName"); presenter.editCustomerName == nil {
		err = errors.New("unable to find #editCustomerName")
		return
	}

	*/

	// Define timeSpan output field.
	if presenter.timeSpan = document.ElementByID("timeSpan"); presenter.timeSpan == nil {
		err = errors.New("unable to find #timeSpan")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	presenter.editCustomerName.SetValue(record.Name)
}

*/

// displayTimeSpan displays time.
func (presenter *panelPresenter) displayTimeSpan(s string) {
	presenter.timeSpan.SetInnerText(s)
}
