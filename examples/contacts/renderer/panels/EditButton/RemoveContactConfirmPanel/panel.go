package RemoveContactConfirmPanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/interfaces/panelHelper"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactConfirmPanel
	Panel id:   tabsMasterView-home-pad-EditButton-RemoveContactConfirmPanel

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools

	removeContactNotReadyPanel js.Value

	removeContactSelectPanel js.Value

	removeContactConfirmPanel js.Value
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notJS *notjs.NotJS, connection map[types.CallID]caller.Renderer, helper panelHelper.Helper) *Panel {
	panel := &Panel{
		tools: tools,
	}

	panel.removeContactNotReadyPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-RemoveContactNotReadyPanel")

	panel.removeContactSelectPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-RemoveContactSelectPanel")

	panel.removeContactConfirmPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-RemoveContactConfirmPanel")
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

// showRemoveContactNotReadyPanel shows the panel you named RemoveContactNotReadyPanel while hiding any other panels in it's group.
// That panel's id is tabsMasterView-home-pad-EditButton-RemoveContactNotReadyPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
there are no contacts
*/
func (panel *Panel) showRemoveContactNotReadyPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.removeContactNotReadyPanel, force)
}

// showRemoveContactSelectPanel shows the panel you named RemoveContactSelectPanel while hiding any other panels in it's group.
// That panel's id is tabsMasterView-home-pad-EditButton-RemoveContactSelectPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
A mapvlist allowing the user to select a contact to remove.
*/
func (panel *Panel) showRemoveContactSelectPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.removeContactSelectPanel, force)
}

// showRemoveContactConfirmPanel shows the panel you named RemoveContactConfirmPanel while hiding any other panels in it's group.
// This panel's id is tabsMasterView-home-pad-EditButton-RemoveContactConfirmPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
form for confirmation of the record removal
*/
func (panel *Panel) showRemoveContactConfirmPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.removeContactConfirmPanel, force)
}

// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controler.initialCalls()
	panel.caller.initialCalls()
}
