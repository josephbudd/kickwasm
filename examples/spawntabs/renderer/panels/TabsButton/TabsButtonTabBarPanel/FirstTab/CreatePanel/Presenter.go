package createpanel

import (
	"strings"
	"syscall/js"


	"github.com/pkg/errors"
)

/*

	Panel name: CreatePanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group          *panelGroup
	controller     *panelController
	caller         *panelCaller
	tabPanelHeader js.Value
	tabButton      js.Value

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
// Tab button label.

func (presenter *panelPresenter) getTabLabel() (label string) {
	label = notJS.GetInnerText(presenter.tabButton)
	return
}

func (presenter *panelPresenter) setTabLabel(label string) {
	notJS.SetInnerText(presenter.tabButton, label)
}

// Tab panel heading.

func (presenter *panelPresenter) getTabPanelHeading() (heading string) {
	heading = notJS.GetInnerText(presenter.tabPanelHeader)
	return
}

func (presenter *panelPresenter) setTabPanelHeading(heading string) {
	heading = strings.TrimSpace(heading)
	if len(heading) == 0 {
		tools.ElementHide(presenter.tabPanelHeader)
	} else {
		tools.ElementShow(presenter.tabPanelHeader)
	}
	notJS.SetInnerText(presenter.tabPanelHeader, heading)
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	notJS.SetValue(presenter.editCustomerName, record.Name)
}

*/
