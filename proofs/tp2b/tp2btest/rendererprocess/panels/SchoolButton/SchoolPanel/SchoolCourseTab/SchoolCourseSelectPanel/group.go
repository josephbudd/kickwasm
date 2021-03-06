// +build js, wasm

package schoolcourseselectpanel

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
	schoolCourseSelectPanel js.Value
	schoolCourseStatsPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(group *panelGroup) defineMembers(): %w", err)
		}
	}()

    var panel *markup.Element
 if panel = document.ElementByID("mainMasterView_home_pad_SchoolButton_SchoolPanel_tab_bar-SchoolCourseTabPanel-inner-SchoolCourseSelectPanel"); panel == nil {
	err = fmt.Errorf("unable to find #mainMasterView_home_pad_SchoolButton_SchoolPanel_tab_bar-SchoolCourseTabPanel-inner-SchoolCourseSelectPanel")
		return
    }
    group.schoolCourseSelectPanel = panel.JSValue()
 if panel = document.ElementByID("mainMasterView_home_pad_SchoolButton_SchoolPanel_tab_bar-SchoolCourseTabPanel-inner-SchoolCourseStatsPanel"); panel == nil {
	err = fmt.Errorf("unable to find #mainMasterView_home_pad_SchoolButton_SchoolPanel_tab_bar-SchoolCourseTabPanel-inner-SchoolCourseStatsPanel")
		return
    }
    group.schoolCourseStatsPanel = panel.JSValue()

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showSchoolCourseSelectPanel shows the panel you named SchoolCourseSelectPanel while hiding any other panels in this panel group.
// This panel will become visible only when this group of panels becomes visible.
/* Your note for this panel is:
A vlist for selecting the current school course.
Displayed when there is no current school course selected.

*/
func (group *panelGroup) showSchoolCourseSelectPanel() {
	viewtools.ShowPanelInTabGroup(group.schoolCourseSelectPanel)
}


// showSchoolCourseStatsPanel shows the panel you named SchoolCourseStatsPanel while hiding any other panels in this panel group.
// That panel will become visible only when this group of panels becomes visible.
/* Your note for that panel is:
Displays the stats of the current course.
*/
func (group *panelGroup) showSchoolCourseStatsPanel() {
	viewtools.ShowPanelInTabGroup(group.schoolCourseStatsPanel)
}

