package prove

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/framework/proofs"
)

// Pass will not return an error if rekickwasm worked.
func Pass() (err error) {

	msg := make([]string, 0, 10)
	defer func() {
		if len(msg) > 0 {
			err = fmt.Errorf(strings.Join(msg, "\n"))
		}
	}()

	checkButton(&msg)

	return
}

func checkButton(msg *[]string) {
	bnpn := proofs.ButtonNamePanelNames()
	var pNames []string
	var found bool
	var pntn map[string][]string
	var tNames []string
	var tnpn map[string][]string

	// First Button.
	if pNames, found = bnpn["SchoolButton"]; !found {
		*msg = append(*msg, "SchoolButton not found")
		return
	}
	// Have the first button so check it's panels.
	if len(pNames) != 2 {
		*msg = append(*msg, fmt.Sprintf("The \"SchoolButton\" has %d panels", len(pNames)))
		return
	}
	if pNames[0] != "SchoolCourseNotReadyPanel" {
		*msg = append(*msg, "The \"SchoolButton\"'s 1st panel is not \"SchoolCourseNotReadyPanel\"")
		return
	} 
	if pNames[1] != "SchoolPanel" {
		*msg = append(*msg, "The \"SchoolButton\"'s 2nd panel is not \"SchoolPanel\"")
		return
	}
	// Have the panel so check it's tab order
	pntn = proofs.PanelNameTabNames()
	if tNames, found = pntn["SchoolPanel"]; !found {
		*msg = append(*msg, "The tab \"SchoolPanel\" is not in proof.PanelNameTabNames()")
		return
	}
	if len(tNames) != 3 {
		*msg = append(*msg, fmt.Sprintf("\"SchoolPanel\" has %d tabs", len(tNames)))
		return
	}
	// First tab.
	if tNames[0] != "SchoolCourseTab" {
		*msg = append(*msg, "The \"SchoolPanel\"'s first tab is not \"SchoolCourseTab\"")
		return
	}
	tnpn = proofs.TabNamePanelNames()
	if pNames, found = tnpn["SchoolCourseTab"]; !found {
		*msg = append(*msg, "The tab \"SchoolCourseTab\" is not in proof.TabNamePanelNames()")
		return
	}
	if len(pNames) != 2 {
		*msg = append(*msg, fmt.Sprintf("\"SchoolCourseTab\" has %d tabs", len(pNames)))
		return
	}
	if pNames[0] != "SchoolCourseSelectPanel" {
		*msg = append(*msg, "The \"SchoolCourseTab\"'s first panel is not \"SchoolCourseSelectPanel\"")
		return
	}
	if pNames[1] != "SchoolCourseStatsPanel" {
		*msg = append(*msg, "The \"SchoolCourseTab\"'s first panel is not \"SchoolCourseStatsPanel\"")
		return
	}
	// Second tab.
	if tNames[1] != "SchoolCourseCopyTab" {
		*msg = append(*msg, "The \"SchoolPanel\"'s second tab is not \"SchoolCourseCopyTab\"")
		return
	}
	if pNames, found = tnpn["SchoolCourseCopyTab"]; !found {
		*msg = append(*msg, "The tab \"SchoolCourseCopyTab\" is not in proof.TabNamePanelNames()")
		return
	}
	if len(pNames) != 4 {
		*msg = append(*msg, fmt.Sprintf("\"SchoolCourseCopyTab\" has %d tabs", len(pNames)))
		return
	}
	if pNames[0] != "SchoolCourseCopyCompletedPanel" {
		*msg = append(*msg, "The \"SchoolCourseCopyTab\"'s first panel is not \"SchoolCourseCopyCompletedPanel\"")
		return
	}
	if pNames[1] != "SchoolCourseCopyChoosePanel" {
		*msg = append(*msg, "The \"SchoolCourseCopyTab\"'s first panel is not \"SchoolCourseCopyChoosePanel\"")
		return
	}
	if pNames[2] != "SchoolCourseCopyPracticePanel" {
		*msg = append(*msg, "The \"SchoolCourseCopyTab\"'s first panel is not \"SchoolCourseCopyPracticePanel\"")
		return
	}
	if pNames[3] != "SchoolCourseCopyTestPanel" {
		*msg = append(*msg, "The \"SchoolCourseCopyTab\"'s first panel is not \"SchoolCourseCopyTestPanel\"")
		return
	}
}
