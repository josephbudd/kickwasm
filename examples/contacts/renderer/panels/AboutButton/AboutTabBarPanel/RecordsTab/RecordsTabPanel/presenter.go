package recordstabpanel

import (
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RecordsTabPanel

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

	*/

	recordsCountP js.Value
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

	*/

	null := js.Null()

	if panelPresenter.recordsCountP = panelPresenter.notJS.GetElementByID("recordsCountP"); panelPresenter.recordsCountP == null {
		err = errors.New("unable to find #recordsCountP")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

// DisplayRecordCount displays the contact record count.
func (panelPresenter *Presenter) DisplayRecordCount(count uint64) {
	var message string
	switch count {
	case 0:
		message = "You don't have any contact records yet."
	case 1:
		message = "You only have 1 contact record."
	default:
		message = fmt.Sprintf("You have a total of %d records now.", count)
	}
	panelPresenter.notJS.SetInnerText(panelPresenter.recordsCountP, message)
}
