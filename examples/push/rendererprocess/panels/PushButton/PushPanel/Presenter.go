// +build js, wasm

package pushpanel

import (
	"syscall/js"

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

	*/

	timeSpan js.Value
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

	*/

	// Define timeSpan output field.
	if presenter.timeSpan = notJS.GetElementByID("timeSpan"); presenter.timeSpan == null {
		err = errors.New("unable to find #timeSpan")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

*/

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayTimeSpan(s string) {
	notJS.SetInnerText(presenter.timeSpan, s)
}
