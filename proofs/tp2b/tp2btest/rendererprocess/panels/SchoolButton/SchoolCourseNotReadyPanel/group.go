// +build js, wasm

package schoolcoursenotreadypanel

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/api/markup"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/framework/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// panelGroup is a group of 2 panels.
// It also has show panel funcs for each panel in this panel group.
type panelGroup struct {
	schoolCourseNotReadyPanel js.Value
	schoolPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(group *panelGroup) defineMembers(): %w", err)
		}
	}()

    var panel *markup.Element
 if panel = document.ElementByID("mainMasterView-home-pad-SchoolButton-SchoolCourseNotReadyPanel"); panel == nil {
	err = fmt.Errorf("unable to find #mainMasterView-home-pad-SchoolButton-SchoolCourseNotReadyPanel")
		return
    }
    group.schoolCourseNotReadyPanel = panel.JSValue()
 if panel = document.ElementByID("mainMasterView-home-pad-SchoolButton-SchoolPanel"); panel == nil {
	err = fmt.Errorf("unable to find #mainMasterView-home-pad-SchoolButton-SchoolPanel")
		return
    }
    group.schoolPanel = panel.JSValue()

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showSchoolCourseNotReadyPanel shows the panel you named SchoolCourseNotReadyPanel while hiding any other panels in this panel group.
// This panel's id is mainMasterView-home-pad-SchoolButton-SchoolCourseNotReadyPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
Displayed when there are on courses.
*/
func (group *panelGroup) showSchoolCourseNotReadyPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.schoolCourseNotReadyPanel, force)
}

// showSchoolPanel shows the panel you named SchoolPanel while hiding any other panels in this panel group.
// That panel's id is mainMasterView-home-pad-SchoolButton-SchoolPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:

*/
func (group *panelGroup) showSchoolPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.schoolPanel, force)
}