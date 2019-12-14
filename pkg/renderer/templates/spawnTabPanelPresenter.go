package templates

// SpawnPanelPresenter is the genereric renderer panel presenter template.
const SpawnPanelPresenter = `// +build js, wasm

package {{call .PackageNameCase .PanelName}}

import (
	"fmt"
	"strings"

	"{{.ApplicationGitPath}}{{.ImportRendererDOM}}"
	"{{.ApplicationGitPath}}{{.ImportRendererMarkup}}"
)

/*

	Panel name: {{.PanelName}}

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	uniqueID       uint64
	document       *dom.DOM
	group          *panelGroup
	controller     *panelController
	messenger      *panelMessenger
	tabButton      *markup.Element
	tabPanelHeader *markup.Element

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

	// example:

	addCustomerName *markup.Element

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

	import "{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"

	// my spawn template has a name input field and a submit button.
	// <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">

	var id string

	// Define the customer name input field.
	// Build it's id using the uniqueID.
	id = display.SpawnID("addCustomerName{{.SpawnID}}", controller.uniqueID)
	if presenter.addCustomerName = presenter.document.ElementByID(id); presenter.addCustomerName == nil {
		err = fmt.Errorf("unable to find #" + id)
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

import "{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"

// displayCustomer displays the customer in the add customer form.
// This short example only uses the customer name field in the form.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	presenter.addCustomerName.SetValue(record.Name)
}

*/
`
