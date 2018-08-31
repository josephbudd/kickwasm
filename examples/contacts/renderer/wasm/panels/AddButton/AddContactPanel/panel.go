package AddContactPanel

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/transports/calls"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: AddContactPanel
	Panel id:   tabsMasterView-home-pad-AddButton-AddContactPanel

*/

// A panel has a controler, presenter and caller.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/wasm/viewtools
	notjs     *kicknotjs.NotJS

	addContactPanel js.Value
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notjs *kicknotjs.NotJS, connection *calls.Calls) *Panel {
	panel := &Panel{
		tools: tools,
	}

	panel.addContactPanel = notjs.GetElementByID("tabsMasterView-home-pad-AddButton-AddContactPanel")
	// initialize controler, presenter, caller.
	controler := &Controler{
		panel:  panel,
		quitCh: quitCh,
		tools:  tools,
		notjs:  notjs,
	}
	controler.defineControlsSetHandlers()
	presenter := &Presenter{
		panel:   panel,
		tools:   tools,
		notjs:   notjs,
	}
	presenter.defineMembers()
	caller := &Caller{
		panel:      panel,
		quitCh:     quitCh,
		connection: connection,
		tools:      tools,
		notjs:      notjs,
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

// showAddContactPanel shows the panel you named AddContactPanel while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #tabsMasterView-home-pad-AddButton-AddContactPanel, which you named AddContactPanel, becomes visible.
/* Your note for this panel is:
Adding a contact.
Allow the user to enter contact info into a form and submit or cancel.

*/
func (panel *Panel) showAddContactPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.addContactPanel, force)
}

// InitialCalls runs the first code that the panel needs to run.
func (p *Panel) InitialCalls() {
	p.controler.initialCalls()
	p.caller.initialCalls()
}
