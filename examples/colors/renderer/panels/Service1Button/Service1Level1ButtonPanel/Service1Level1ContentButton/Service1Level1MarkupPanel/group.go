package service1level1markuppanel

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

	service1Level1MarkupPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

	if group.service1Level1MarkupPanel = notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-Service1Level1ContentButton-Service1Level1MarkupPanel"); group.service1Level1MarkupPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-Service1Level1ContentButton-Service1Level1MarkupPanel")
		return
	}

	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/

// showService1Level1MarkupPanel shows the panel you named Service1Level1MarkupPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-Service1Level1ContentButton-Service1Level1MarkupPanel.
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
func (group *panelGroup) showService1Level1MarkupPanel(force bool) {
	tools.ShowPanelInButtonGroup(group.service1Level1MarkupPanel, force)
}