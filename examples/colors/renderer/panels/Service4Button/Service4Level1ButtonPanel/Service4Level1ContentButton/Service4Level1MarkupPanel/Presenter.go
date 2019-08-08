package service4level1markuppanel

import (
	"github.com/pkg/errors"
)

/*

	Panel name: Service4Level1MarkupPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group      *panelGroup
	controller *panelController
	caller     *panelCaller

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.
	// example:

	editCustomerName js.Value

	*/
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
	if presenter.editCustomerName = notJS.GetElementByID("editCustomerName"); presenter.editCustomerName == null {
		err = errors.New("unable to find #editCustomerName")
		return
	}

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.
// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	notJS.SetValue(presenter.editCustomerName, record.Name)
}

*/
