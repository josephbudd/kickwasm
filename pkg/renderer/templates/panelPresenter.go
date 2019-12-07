package templates

// PanelPresenter is the genereric renderer panel presenter template.
const PanelPresenter = `// +build js, wasm

package {{call .PackageNameCase .PanelName}}

import ({{ if .IsTabSiblingPanel }}
	"strings"

	"{{.ApplicationGitPath}}{{.ImportRendererMarkup}}"
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
	messenger      *panelMessenger
	tabPanelHeader *markup.Element
	tabButton      *markup.Element{{else}}group      *panelGroup
	controller *panelController
	messenger  *panelMessenger{{ end }}

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
{{ if .IsTabSiblingPanel }}// Tab button label.

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
{{ end }}
/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	presenter.editCustomerName.SetValue(record.Name)
}

*/
`
