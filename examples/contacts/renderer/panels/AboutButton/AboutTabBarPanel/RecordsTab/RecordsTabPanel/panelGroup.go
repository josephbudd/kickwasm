package RecordsTabPanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
	"github.com/pkg/errors"
)

// PanelGroup is a group of 1 panel.
// It also has a show panel func for each panel in this panel group.
type PanelGroup struct {
	tools *viewtools.Tools
	notJS *notjs.NotJS

	recordsTabPanel js.Value
}

func (panelGroup *PanelGroup) defineMembers() (err error) {
	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelGroup *PanelGroup) defineMembers()")
		}
	}()

	notJS := panelGroup.notJS
	null := js.Null()

	if panelGroup.recordsTabPanel = notJS.GetElementByID("tabsMasterView_home_pad_AboutButton_AboutTabBarPanel_tab_bar-RecordsTabPanel-inner-RecordsTabPanel"); panelGroup.recordsTabPanel == null {
		err = errors.New("unable to find #tabsMasterView_home_pad_AboutButton_AboutTabBarPanel_tab_bar-RecordsTabPanel-inner-RecordsTabPanel")
		return
	}

	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showRecordsTabPanel shows the panel you named RecordsTabPanel while hiding any other panels in this panel group.
// This panel will become visible only when this group of panels becomes visible.
/* Your note for this panel is:
display text about the number of records in a paragraph
*/
func (panelGroup *PanelGroup) showRecordsTabPanel() {
	panelGroup.tools.ShowPanelInTabGroup(panelGroup.recordsTabPanel)
}
