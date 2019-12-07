// +build js, wasm

package createpanel

import (
	"strings"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/markup"

	"github.com/pkg/errors"
)

/*

	Panel name: CreatePanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group          *panelGroup
	controller     *panelController
	messenger      *panelMessenger
	tabPanelHeader *markup.Element
	tabButton      *markup.Element

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

	// example:

	editCustomerName *markup.Element

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
	if presenter.editCustomerName = document.ElementByID("editCustomerName"); presenter.editCustomerName == nil {
		err = errors.New("unable to find #editCustomerName")
		return
	}

	*/

	return
}
// Tab button label.

func (presenter *panelPresenter) getTabLabel() (label string) {
	label = presenter.tabButton.InnerText()
	return
}

func (presenter *panelPresenter) setTabLabel(label string) {
	presenter.tabButton.SetInnerText(label)
}

// Tab panel heading.

func (presenter *panelPresenter) getTabPanelHeading() (heading string) {
	heading = presenter.tabPanelHeader.InnerText()
	return
}

func (presenter *panelPresenter) setTabPanelHeading(heading string) {
	heading = strings.TrimSpace(heading)
	if len(heading) == 0 {
		presenter.tabPanelHeader.Hide()
	} else {
		presenter.tabPanelHeader.Show()
	}
	presenter.tabPanelHeader.SetInnerText(heading)
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	presenter.editCustomerName.SetValue(record.Name)
}

*/
