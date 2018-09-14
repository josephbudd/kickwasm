package RemoveContactConfirmPanel

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactConfirmPanel
	Panel id:   tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools
	notjs     *kicknotjs.NotJS

	removeContactNotReadyPanel js.Value

	removeContactSelectPanel js.Value

	removeContactConfirmPanel js.Value
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notjs *kicknotjs.NotJS, connection types.RendererCallMap) *Panel {
	panel := &Panel{
		tools: tools,
	}

	panel.removeContactNotReadyPanel = notjs.GetElementByID("tabsMasterView-home-pad-RemoveButton-RemoveContactNotReadyPanel")

	panel.removeContactSelectPanel = notjs.GetElementByID("tabsMasterView-home-pad-RemoveButton-RemoveContactSelectPanel")

	panel.removeContactConfirmPanel = notjs.GetElementByID("tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel")
	// initialize controler, presenter, caller.
	controler := &Controler{
		panel:  panel,
		quitCh: quitCh,
		tools:  tools,
		notjs:  notjs,
	}
	presenter := &Presenter{
		panel:   panel,
		tools:   tools,
		notjs:   notjs,
	}
	caller := &Caller{
		panel:      panel,
		quitCh:     quitCh,
		connection: connection,
		tools:      tools,
		notjs:      notjs,
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

	Calls these from the controler, presenter and caller.
*/

// showRemoveContactNotReadyPanel shows the panel you named RemoveContactNotReadyPanel while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #tabsMasterView-home-pad-RemoveButton-RemoveContactNotReadyPanel, which you named RemoveContactNotReadyPanel, becomes visible.
/* Your note for this panel is:
there are no contacts
*/
func (panel *Panel) showRemoveContactNotReadyPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.removeContactNotReadyPanel, force)
}

// showRemoveContactSelectPanel shows the panel you named RemoveContactSelectPanel while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #tabsMasterView-home-pad-RemoveButton-RemoveContactSelectPanel, which you named RemoveContactSelectPanel, becomes visible.
/* Your note for this panel is:
A mapvlist allowing the user to select a contact to remove.
*/
func (panel *Panel) showRemoveContactSelectPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.removeContactSelectPanel, force)
}

// showRemoveContactConfirmPanel shows the panel you named RemoveContactConfirmPanel while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #tabsMasterView-home-pad-RemoveButton-RemoveContactConfirmPanel, which you named RemoveContactConfirmPanel, becomes visible.
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
