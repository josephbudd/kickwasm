package templates

// PanelPresenter is the genereric renderer panel presenter template.
const PanelPresenter = `package {{call .PackageNameCase .PanelName}}

import ({{ if .IsTabSiblingPanel }}
	"strings"
	"syscall/js"

{{ end }}
	"github.com/pkg/errors"
)

/*

	Panel name: {{.PanelName}}

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	{{ if .IsTabSiblingPanel }}group          *panelGroup
	controller     *panelController
	caller         *panelCaller
	tabPanelHeader js.Value{{else}}group      *panelGroup
	controller *panelController
	caller     *panelCaller{{ end }}

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
{{ if .IsTabSiblingPanel }}// Tab panel heading.

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
{{ end }}
/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.
// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	notJS.SetValue(presenter.editCustomerName, record.Name)
}

*/
`
