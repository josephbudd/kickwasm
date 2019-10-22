// +build js, wasm

package action4level5markuppanel

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
	action4Level5MarkupPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

	if group.action4Level5MarkupPanel = notJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton-Action4Level5MarkupPanel"); group.action4Level5MarkupPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton-Action4Level5MarkupPanel")
		return
	}

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showAction4Level5MarkupPanel shows the panel you named Action4Level5MarkupPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton-Action4Level5MarkupPanel.
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
func (group *panelGroup) showAction4Level5MarkupPanel(force bool) {
	tools.ShowPanelInButtonGroup(group.action4Level5MarkupPanel, force)
}
