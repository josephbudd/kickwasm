package abouttabpanel

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

	aboutTabPanel js.Value
}

func (panelGroup *PanelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelGroup *PanelGroup) defineMembers()")
		}
	}()

	notJS := panelGroup.notJS
	null := js.Null()

	if panelGroup.aboutTabPanel = notJS.GetElementByID("tabsMasterView_home_pad_AboutButton_AboutTabBarPanel_tab_bar-AboutTabPanel-inner-AboutTabPanel"); panelGroup.aboutTabPanel == null {
		err = errors.New("unable to find #tabsMasterView_home_pad_AboutButton_AboutTabBarPanel_tab_bar-AboutTabPanel-inner-AboutTabPanel")
		return
	}


	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showAboutTabPanel shows the panel you named AboutTabPanel while hiding any other panels in this panel group.
// This panel will become visible only when this group of panels becomes visible.
/* Your note for this panel is:
dynamic text that display the author and version.
*/
func (panelGroup *PanelGroup) showAboutTabPanel() {
	panelGroup.tools.ShowPanelInTabGroup(panelGroup.aboutTabPanel)
}


