package templates

// PanelPresenter is the genereric renderer panel presenter template.
const PanelPresenter = `// +build js, wasm

package {{call .PackageNameCase .PanelName}}

import (
	"fmt"{{ if .IsTabSiblingPanel }}
	"strings"

	"{{.ApplicationGitPath}}{{.ImportRendererAPIMarkup}}"
{{ end }}
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

	// example:{{ if not .IsTabSiblingPanel }}

	import "{{.ApplicationGitPath}}{{.ImportRendererAPIMarkup}}"{{ end }}

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
{{ if .IsTabSiblingPanel }}
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
{{ end }}
/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(r *record.CustomerRecord) {
	presenter.editCustomerName.SetValue(r.Name)
}

*/
`
