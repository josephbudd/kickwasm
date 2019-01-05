package Service2Level1MarkupPanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/colors/site/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/site/viewtools"
)

// PanelGroup is a group of 1 panel.
// It also has a show panel func for each panel in this panel group.
type PanelGroup struct {
	tools *viewtools.Tools
	notJS *notjs.NotJS

	service2Level1MarkupPanel js.Value
}

func (panelGroup *PanelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelGroup *PanelGroup) defineMembers()")
		}
	}()

	notJS := panelGroup.notJS
	null := js.Null()

	if panelGroup.service2Level1MarkupPanel = notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ContentButton-Service2Level1MarkupPanel"); panelGroup.service2Level1MarkupPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ContentButton-Service2Level1MarkupPanel")
		return
	}


	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showService2Level1MarkupPanel shows the panel you named Service2Level1MarkupPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ContentButton-Service2Level1MarkupPanel.
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
func (panelGroup *PanelGroup) showService2Level1MarkupPanel(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.service2Level1MarkupPanel, force)
}

