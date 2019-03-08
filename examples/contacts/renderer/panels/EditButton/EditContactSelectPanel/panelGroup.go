package editcontactselectpanel

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

	editContactNotReadyPanel js.Value
	editContactSelectPanel js.Value
	editContactEditPanel js.Value
}

func (panelGroup *PanelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelGroup *PanelGroup) defineMembers()")
		}
	}()

	notJS := panelGroup.notJS
	null := js.Null()

	if panelGroup.editContactNotReadyPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-EditContactNotReadyPanel"); panelGroup.editContactNotReadyPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-EditButton-EditContactNotReadyPanel")
		return
	}
	if panelGroup.editContactSelectPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-EditContactSelectPanel"); panelGroup.editContactSelectPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-EditButton-EditContactSelectPanel")
		return
	}
	if panelGroup.editContactEditPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-EditContactEditPanel"); panelGroup.editContactEditPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-EditButton-EditContactEditPanel")
		return
	}


	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showEditContactNotReadyPanel shows the panel you named EditContactNotReadyPanel while hiding any other panels in this panel group.
// That panel's id is tabsMasterView-home-pad-EditButton-EditContactNotReadyPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
tell the user that he/she has not added any contacts yet.
*/
func (panelGroup *PanelGroup) showEditContactNotReadyPanel(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.editContactNotReadyPanel, force)
}

// showEditContactSelectPanel shows the panel you named EditContactSelectPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-EditButton-EditContactSelectPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
A mapvlist allowing the user to select a contact to edit.
*/
func (panelGroup *PanelGroup) showEditContactSelectPanel(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.editContactSelectPanel, force)
}

// showEditContactEditPanel shows the panel you named EditContactEditPanel while hiding any other panels in this panel group.
// That panel's id is tabsMasterView-home-pad-EditButton-EditContactEditPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
edit the form
*/
func (panelGroup *PanelGroup) showEditContactEditPanel(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.editContactEditPanel, force)
}

