package Service1Level2MarkupPanel

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/viewtools"
)

/*

	Panel name: Service1Level2MarkupPanel
	Panel id:   tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ContentButton-Service1Level2MarkupPanel

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools
	notjs     *kicknotjs.NotJS

	service1Level2MarkupPanel js.Value
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notjs *kicknotjs.NotJS, connection map[types.CallID]caller.Renderer) *Panel {
	panel := &Panel{
		tools: tools,
	}

	panel.service1Level2MarkupPanel = notjs.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ContentButton-Service1Level2MarkupPanel")
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

// showService1Level2MarkupPanel shows the panel you named Service1Level2MarkupPanel while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ContentButton-Service1Level2MarkupPanel, which you named Service1Level2MarkupPanel, becomes visible.
/* Your note for this panel is:
This is the only content.
Brought to you in the first service color.

*/
func (panel *Panel) showService1Level2MarkupPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.service1Level2MarkupPanel, force)
}

// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controler.initialCalls()
	panel.caller.initialCalls()
}
