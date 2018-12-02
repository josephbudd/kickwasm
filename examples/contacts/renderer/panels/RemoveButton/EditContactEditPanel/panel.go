package EditContactEditPanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/interfaces/panelHelper"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: EditContactEditPanel
	Panel id:   tabsMasterView-home-pad-RemoveButton-EditContactEditPanel

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools

	editContactNotReadyPanel js.Value

	editContactSelectPanel js.Value

	editContactEditPanel js.Value
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notJS *notjs.NotJS, connection map[types.CallID]caller.Renderer, helper panelHelper.Helper) *Panel {
	panel := &Panel{
		tools: tools,
	}

	panel.editContactNotReadyPanel = notJS.GetElementByID("tabsMasterView-home-pad-RemoveButton-EditContactNotReadyPanel")

	panel.editContactSelectPanel = notJS.GetElementByID("tabsMasterView-home-pad-RemoveButton-EditContactSelectPanel")

	panel.editContactEditPanel = notJS.GetElementByID("tabsMasterView-home-pad-RemoveButton-EditContactEditPanel")
	// initialize controler, presenter, caller.
	controler := &Controler{
		panel:  panel,
		quitCh: quitCh,
		tools:  tools,
		notJS:  notJS,
	}
	presenter := &Presenter{
		panel:   panel,
		tools:   tools,
		notJS:   notJS,
	}
	caller := &Caller{
		panel:      panel,
		quitCh:     quitCh,
		connection: connection,
		tools:      tools,
		notJS:      notJS,
	}
	// settings
	panel.controler = controler
	panel.presenter = presenter
	panel.caller = caller
	controler.presenter = presenter
	controler.caller = caller
	presenter.controler = controler
	presenter.caller = caller
	caller.controler = controler
	caller.presenter = presenter
	// completions
	controler.defineControlsSetHandlers()
	presenter.defineMembers()
	caller.addMainProcessCallBacks()
	return panel
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showEditContactNotReadyPanel shows the panel you named EditContactNotReadyPanel while hiding any other panels in it's group.
// That panel's id is tabsMasterView-home-pad-RemoveButton-EditContactNotReadyPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
tell the user that he/she has not added any contacts yet.
*/
func (panel *Panel) showEditContactNotReadyPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.editContactNotReadyPanel, force)
}

// showEditContactSelectPanel shows the panel you named EditContactSelectPanel while hiding any other panels in it's group.
// That panel's id is tabsMasterView-home-pad-RemoveButton-EditContactSelectPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
A mapvlist allowing the user to select a contact to edit.
*/
func (panel *Panel) showEditContactSelectPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.editContactSelectPanel, force)
}

// showEditContactEditPanel shows the panel you named EditContactEditPanel while hiding any other panels in it's group.
// This panel's id is tabsMasterView-home-pad-RemoveButton-EditContactEditPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
edit the form
*/
func (panel *Panel) showEditContactEditPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.editContactEditPanel, force)
}

// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controler.initialCalls()
	panel.caller.initialCalls()
}
