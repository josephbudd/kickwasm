// +build js, wasm

package action3level3markuppanel

import (
	"syscall/js"

	"github.com/pkg/errors"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// panelGroup is a group of 1 panel.
// It also has a show panel func for each panel in this panel group.
type panelGroup struct {
	action3Level3MarkupPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

	if group.action3Level3MarkupPanel = notJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton-Action3Level3MarkupPanel"); group.action3Level3MarkupPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton-Action3Level3MarkupPanel")
		return
	}

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showAction3Level3MarkupPanel shows the panel you named Action3Level3MarkupPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton-Action3Level3MarkupPanel.
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
func (group *panelGroup) showAction3Level3MarkupPanel(force bool) {
	tools.ShowPanelInButtonGroup(group.action3Level3MarkupPanel, force)
}
