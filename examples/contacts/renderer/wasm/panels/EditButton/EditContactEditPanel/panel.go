package EditContactEditPanel

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/transports/calls"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/states"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: EditContactEditPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactEditPanel

*/

// Panel has a controler, presenter and caller.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/wasm/viewtools
	notjs     *kicknotjs.NotJS

	editContactNotReadyPanel js.Value

	editContactSelectPanel js.Value

	editContactEditPanel js.Value
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notjs *kicknotjs.NotJS, connection *calls.Calls) *Panel {
	panel := &Panel{
		tools: tools,
	}

	panel.editContactNotReadyPanel = notjs.GetElementByID("tabsMasterView-home-pad-EditButton-EditContactNotReadyPanel")

	panel.editContactSelectPanel = notjs.GetElementByID("tabsMasterView-home-pad-EditButton-EditContactSelectPanel")

	panel.editContactEditPanel = notjs.GetElementByID("tabsMasterView-home-pad-EditButton-EditContactEditPanel")
	// initialize controler, presenter, caller.
	controler := &Controler{
		panel:  panel,
		quitCh: quitCh,
		tools:  tools,
		notjs:  notjs,
	}
	controler.defineControlsSetHandlers()
	presenter := &Presenter{
		panel: panel,
		tools: tools,
		notjs: notjs,
	}
	presenter.defineMembers()
	caller := &Caller{
		panel:      panel,
		quitCh:     quitCh,
		connection: connection,
		tools:      tools,
		notjs:      notjs,
		states:     states.NewStates(),
	}
	caller.addMainProcessCallBacks()
	// finish controler, presenter, caller.
	controler.presenter = presenter
	controler.caller = caller
	presenter.controler = controler
	presenter.caller = caller
	caller.controler = controler
	caller.presenter = presenter
	// finish panel
	panel.controler = controler
	panel.presenter = presenter
	panel.caller = caller
	return panel
}

/*
	Show panel funcs.

	Calls these from the controler, presenter and caller.
*/

// showEditContactNotReadyPanel shows the panel you named EditContactNotReadyPanel while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #tabsMasterView-home-pad-EditButton-EditContactNotReadyPanel, which you named EditContactNotReadyPanel, becomes visible.
/* Your note for this panel is:
tell the user that he/she has not added any contacts yet.
*/
func (panel *Panel) showEditContactNotReadyPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.editContactNotReadyPanel, force)
}

// showEditContactSelectPanel shows the panel you named EditContactSelectPanel while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #tabsMasterView-home-pad-EditButton-EditContactSelectPanel, which you named EditContactSelectPanel, becomes visible.
/* Your note for this panel is:
A mapvlist allowing the user to select a contact to edit.
*/
func (panel *Panel) showEditContactSelectPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.editContactSelectPanel, force)
}

// showEditContactEditPanel shows the panel you named EditContactEditPanel while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #tabsMasterView-home-pad-EditButton-EditContactEditPanel, which you named EditContactEditPanel, becomes visible.
/* Your note for this panel is:
edit the form
*/
func (panel *Panel) showEditContactEditPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.editContactEditPanel, force)
}

// InitialCalls runs the first code that the panel needs to run.
func (p *Panel) InitialCalls() {
	p.controler.initialCalls()
	p.caller.initialCalls()
}
