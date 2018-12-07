package CreditTabPanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/interfaces/panelHelper"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: CreditTabPanel
	Panel id:   tabsMasterView_home_pad_AboutButton_AboutTabBarPanel_tab_bar-CreditTabPanel-inner-CreditTabPanel

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see /renderer/viewtools

	creditTabPanel js.Value
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notJS *notjs.NotJS, connection map[types.CallID]caller.Renderer, helper panelHelper.Helper) *Panel {
	panel := &Panel{
		tools: tools,
	}

	panel.creditTabPanel = notJS.GetElementByID("tabsMasterView_home_pad_AboutButton_AboutTabBarPanel_tab_bar-CreditTabPanel-inner-CreditTabPanel")
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

// showCreditTabPanel shows the panel you named CreditTabPanel while hiding any other panels in it's group.
// This panel will become visible only when this group of panels becomes visible.
/* Your note for this panel is:
static text that display the credits.
*/
func (panel *Panel) showCreditTabPanel() {
	panel.tools.ShowPanelInTabGroup(panel.creditTabPanel)
}


// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controler.initialCalls()
	panel.caller.initialCalls()
}
