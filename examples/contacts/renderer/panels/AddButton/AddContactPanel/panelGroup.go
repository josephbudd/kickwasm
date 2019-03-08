package addcontactpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

// PanelGroup is a group of 1 panel.
// It also has a show panel func for each panel in this panel group.
type PanelGroup struct {
	tools *viewtools.Tools
	notJS *notjs.NotJS

	addContactPanel js.Value
}

func (panelGroup *PanelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelGroup *PanelGroup) defineMembers()")
		}
	}()

	notJS := panelGroup.notJS
	null := js.Null()

	if panelGroup.addContactPanel = notJS.GetElementByID("tabsMasterView-home-pad-AddButton-AddContactPanel"); panelGroup.addContactPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-AddButton-AddContactPanel")
		return
	}


	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showAddContactPanel shows the panel you named AddContactPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-AddButton-AddContactPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
Adding a contact.
Allow the user to enter contact info into a form and submit or cancel.

*/
func (panelGroup *PanelGroup) showAddContactPanel(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.addContactPanel, force)
}

