package templates

// SpawnPanelPresenter is the genereric renderer panel presenter template.
const SpawnPanelPresenter = `package {{call .PackageNameCase .PanelName}}

import (
	"syscall/js"

	"github.com/pkg/errors"
)

/*

	Panel name: {{.PanelName}}

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	uniqueID       uint64
	group          *panelGroup
	controler      *panelControler
	caller         *panelCaller
	tabButton      js.Value
	tabPanelHeader js.Value

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.
	// example:

	addCustomerName js.Value

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
	// my spawn template has a name input field and a submit button.
	// <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">

	var id string

	// Define the customer name input field.
	// Build it's id using the uniqueID.
	id = tools.FixSpawnID("addCustomerName{{.SpawnID}}", controler.uniqueID)
	if presenter.addCustomerName = notJS.GetElementByID(id); presenter.addCustomerName == null {
		err = errors.New("unable to find #" + id)
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
	notJS.SetInnerText(presenter.tabPanelHeader, heading)
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.
// example:

// displayCustomer displays the customer in the add customer form.
// This short example only uses the customer name field in the form.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	notJS.SetValue(presenter.addCustomerName, record.Name)
}

*/
`
