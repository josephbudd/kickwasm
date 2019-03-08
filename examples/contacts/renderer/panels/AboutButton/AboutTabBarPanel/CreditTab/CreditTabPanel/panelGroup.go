package credittabpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

// PanelGroup is a group of 1 panel.
// It also has a show panel func for each panel in this panel group.
type PanelGroup struct {
	tools *viewtools.Tools
	notJS *notjs.NotJS

	creditTabPanel js.Value
}

func (panelGroup *PanelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelGroup *PanelGroup) defineMembers()")
		}
	}()

	notJS := panelGroup.notJS
	null := js.Null()

	if panelGroup.creditTabPanel = notJS.GetElementByID("tabsMasterView_home_pad_AboutButton_AboutTabBarPanel_tab_bar-CreditTabPanel-inner-CreditTabPanel"); panelGroup.creditTabPanel == null {
		err = errors.New("unable to find #tabsMasterView_home_pad_AboutButton_AboutTabBarPanel_tab_bar-CreditTabPanel-inner-CreditTabPanel")
		return
	}


	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showCreditTabPanel shows the panel you named CreditTabPanel while hiding any other panels in this panel group.
// This panel will become visible only when this group of panels becomes visible.
/* Your note for this panel is:
static text that display the credits.
*/
func (panelGroup *PanelGroup) showCreditTabPanel() {
	panelGroup.tools.ShowPanelInTabGroup(panelGroup.creditTabPanel)
}


