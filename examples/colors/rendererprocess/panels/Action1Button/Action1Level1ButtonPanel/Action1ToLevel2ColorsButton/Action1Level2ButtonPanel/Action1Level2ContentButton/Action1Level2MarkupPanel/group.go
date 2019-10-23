// +build js, wasm

package action1level2markuppanel

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
	action1Level2MarkupPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

	if group.action1Level2MarkupPanel = notJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton-Action1Level2MarkupPanel"); group.action1Level2MarkupPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton-Action1Level2MarkupPanel")
		return
	}

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showAction1Level2MarkupPanel shows the panel you named Action1Level2MarkupPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton-Action1Level2MarkupPanel.
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
func (group *panelGroup) showAction1Level2MarkupPanel(force bool) {
	tools.ShowPanelInButtonGroup(group.action1Level2MarkupPanel, force)
}