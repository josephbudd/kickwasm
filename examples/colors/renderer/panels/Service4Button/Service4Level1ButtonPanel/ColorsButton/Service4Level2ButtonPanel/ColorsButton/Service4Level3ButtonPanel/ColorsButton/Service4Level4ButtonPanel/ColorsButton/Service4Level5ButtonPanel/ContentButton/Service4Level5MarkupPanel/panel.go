package Service4Level5MarkupPanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/interfaces/panelHelper"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/viewtools"
)

/*

	Panel name: Service4Level5MarkupPanel
	Panel id:   tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton-Service4Level5ButtonPanel-ContentButton-Service4Level5MarkupPanel

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools

	service4Level5MarkupPanel js.Value
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notJS *notjs.NotJS, connection map[types.CallID]caller.Renderer, helper panelHelper.Helper) *Panel {
	panel := &Panel{
		tools: tools,
	}

	panel.service4Level5MarkupPanel = notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton-Service4Level5ButtonPanel-ContentButton-Service4Level5MarkupPanel")
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

// showService4Level5MarkupPanel shows the panel you named Service4Level5MarkupPanel while hiding any other panels in it's group.
// This panel's id is tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton-Service4Level5ButtonPanel-ContentButton-Service4Level5MarkupPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
This is the only content.
Brought to you in the first service color.

*/
func (panel *Panel) showService4Level5MarkupPanel(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.service4Level5MarkupPanel, force)
}

// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controler.initialCalls()
	panel.caller.initialCalls()
}
