package removecontactconfirmpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

// PanelGroup is a group of 3 panels.
// It also has show panel funcs for each panel in this panel group.
type PanelGroup struct {
	tools *viewtools.Tools
	notJS *notjs.NotJS

	removeContactNotReadyPanel js.Value
	removeContactSelectPanel js.Value
	removeContactConfirmPanel js.Value
}

func (panelGroup *PanelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelGroup *PanelGroup) defineMembers()")
		}
	}()

	notJS := panelGroup.notJS
	null := js.Null()

	if panelGroup.removeContactNotReadyPanel = notJS.GetElementByID("tabsMasterView-home-pad-RemoveButton-RemoveContactNotReadyPanel"); panelGroup.removeContactNotReadyPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-RemoveButton-RemoveContactNotReadyPanel")
		return
	}
	if panelGroup.removeContactSelectPanel = notJS.GetElementByID("tabsMasterView-home-pad-RemoveButton-RemoveContactSelectPanel"); panelGroup.removeContactSelectPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-RemoveButton-RemoveContactSelectPanel")
		return
	}
	if panelGroup.removeContactConfirmPanel = notJS.GetElementByID("tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel"); panelGroup.removeContactConfirmPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel")
		return
	}


	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showRemoveContactNotReadyPanel shows the panel you named RemoveContactNotReadyPanel while hiding any other panels in this panel group.
// That panel's id is tabsMasterView-home-pad-RemoveButton-RemoveContactNotReadyPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
there are no contacts
*/
func (panelGroup *PanelGroup) showRemoveContactNotReadyPanel(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.removeContactNotReadyPanel, force)
}

// showRemoveContactSelectPanel shows the panel you named RemoveContactSelectPanel while hiding any other panels in this panel group.
// That panel's id is tabsMasterView-home-pad-RemoveButton-RemoveContactSelectPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
A mapvlist allowing the user to select a contact to remove.
*/
func (panelGroup *PanelGroup) showRemoveContactSelectPanel(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.removeContactSelectPanel, force)
}

// showRemoveContactConfirmPanel shows the panel you named RemoveContactConfirmPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
form for confirmation of the record removal
*/
func (panelGroup *PanelGroup) showRemoveContactConfirmPanel(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.removeContactConfirmPanel, force)
}

