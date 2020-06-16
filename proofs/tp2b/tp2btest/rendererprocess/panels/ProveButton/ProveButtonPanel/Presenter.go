// +build js, wasm

package provebuttonpanel

import (
	"fmt"
)

/*

	Panel name: ProveButtonPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group      *panelGroup
	controller *panelController
	messenger  *panelMessenger

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

	// example:

	import "github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/api/markup"

	editCustomerName *markup.Element

	*/
}

// defineMembers defines the panelPresenter members by their html elements.
// Returns the error.
func (presenter *panelPresenter) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(presenter *panelPresenter) defineMembers(): %w", err)
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your panelPresenter members.

	// example:

	// Define the edit form's customer name input field.
	if presenter.editCustomerName = document.ElementByID("editCustomerName"); presenter.editCustomerName == nil {
		err = fmt.Errorf("unable to find #editCustomerName")
		return
	}

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(r *record.CustomerRecord) {
	presenter.editCustomerName.SetValue(r.Name)
}

*/