package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/fix"
)

// Move a panel from a tab to a button.
// DO NOT EVER MOVE ProveButtonPanel OUT OF ProveButton

const (
	kickwasmDotYAML = `title: Linux CW Trainer

importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
  
buttons:
- name: SchoolButton
  label: School
  heading: School
  panels:
  - name: SchoolPanel
    tabs:
    - name: SchoolCourseTab
      label: School
      panels:
      - name: SchoolCourseNotReadyPanel
        note: Displayed when there are on courses.
        markup: <p>You haven't created any courses yet.</p>
      - name: SchoolCourseSelectPanel
        note: |
          A vlist for selecting the current school course.
          Displayed when there is no current school course selected.
        markup: <div id="schoolCourseSelectWrapper"></div>
      - name: SchoolCourseStatsPanel
        note: Displays the stats of the current course.
        markup: |
          <p>TBD</p>
    - name: SchoolCourseCopyTab
      label: Copy
      panels:
      - name: SchoolCourseCopyCompletedPanel
        note: |
          Displayed when the user has successfully copied for this lesson.
          Maybe it should display the results.
        markup: <p>TBD</p>
      - name: SchoolCourseCopyChoosePanel
        note: |
          Displayed when the user has successfully copied for this lesson.
          Maybe it should display the results.
        markup: |
          <h4>Copy</h4>
          <p>
          <button id="schoolCourseCopyChoosePractice">Practice</button>
          <button id="schoolCourseCopyChooseTest">Test</button>
          </p>
      - name: SchoolCourseCopyPracticePanel
        note: A page to let the user copy without recording the results.
        markup: |
          <p>
            <button id="schoolCourseCopyPracticeStart">Start</button>
            <button id="schoolCourseCopyPracticeCheck" class="unseen">Check</button>
          </p>
          <textarea id="schoolCourseCopyPracticeCopy" class="resize-me-width"></textarea>
          <div id="schoolCourseCopyPracticeText" class="resize-me-width"></div>
      - name: SchoolCourseCopyTestPanel
        note: A page to let the user copy with results recorded.
        markup: |
          <p>
            <button id="schoolCourseCopyTestStart">Start</button>
            <button id="schoolCourseCopyTestCheck" class="unseen">Check</button>
          </p>
          <textarea id="schoolCourseCopyTestCopy" class="resize-me-width"></textarea>
          <div id="schoolCourseCopyTestText" class="resize-me-width"></div>
    - name: SchoolCourseKeyTab
      label: Key
      panels:
      - name: SchoolCourseKeyCompletedPanel
        note: |
          Displayed when the user has successfully copied for this lesson.
          Maybe it should display the results.
        markup: <p>TBD</p>
      - name: SchoolCourseKeyChoosePanel
        note: |
          Displayed when the user has successfully copied for this lesson.
          Maybe it should display the results.
        markup: |
          <h4>Key</h4>
          <p>
            <button id="schoolCourseKeyChoosePractice">Practice</button>
            <button id="schoolCourseKeyChooseTest">Test</button>
          </p>
      - name: SchoolCourseKeyPracticePanel
        note: A page to let the user copy without recording the results.
        markup: |
          <h3 id="schoolCourseKeyPracticeH"></h3>
          <p>
          <label for="schoolCourseKeyPracticeMetronomeOn">Use the metronome.</label>
          <input type="checkbox" id="schoolCourseKeyPracticeMetronomeOn" checked />
          </p>
          <p id="schoolCourseKeyPracticeP">
            <button id="schoolCourseKeyPracticeStart">Start</button>
            <button id="schoolCourseKeyPracticeCheck" class="unseen">Check</button>
          </p>
          <div id="schoolCourseKeyPracticeKey" class="user-not-keying resize-me-width"></div>
          <div id="schoolCourseKeyPracticeCopy" class="resize-me-width"></div>
      - name: SchoolCourseKeyTestPanel
        note: A page to let the user copy with results recorded.
        markup: |
          <h3 id="schoolCourseKeyTestH"></h3>
          <p id="schoolCourseKeyTestP">
            <button id="schoolCourseKeyTestStart">Start</button>
            <button id="schoolCourseKeyTestCheck" class="unseen">Check</button>
          </p>
          <div id="schoolCourseKeyTestKey" class="user-not-keying resize-me-width"></div>
          <div id="schoolCourseKeyTestCopy" class="resize-me-width"></div>
- name: ProveButton
  label: Prove
  panels:
  - name: ProveButtonPanel
    note: |
      The panel for the prove button.
      Run the tests from this panel.
      Do not move the ProveButtonPanel from the ProveButton.
    markup: |
      <p>Prove It!</p>
`
	rekickwasmDotYAML = `title: Linux CW Trainer

importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
  
buttons:
- name: SchoolButton
  label: School
  heading: School
  panels:
  - name: SchoolCourseNotReadyPanel
    note: Displayed when there are on courses.
    markup: <p>You haven't created any courses yet.</p>
  - name: SchoolPanel
    tabs:
    - name: SchoolCourseTab
      label: School
      panels:
      - name: SchoolCourseSelectPanel
        note: |
          A vlist for selecting the current school course.
          Displayed when there is no current school course selected.
        markup: <div id="schoolCourseSelectWrapper"></div>
      - name: SchoolCourseStatsPanel
        note: Displays the stats of the current course.
        markup: |
          <p>TBD</p>
    - name: SchoolCourseCopyTab
      label: Copy
      panels:
      - name: SchoolCourseCopyCompletedPanel
        note: |
          Displayed when the user has successfully copied for this lesson.
          Maybe it should display the results.
        markup: <p>TBD</p>
      - name: SchoolCourseCopyChoosePanel
        note: |
          Displayed when the user has successfully copied for this lesson.
          Maybe it should display the results.
        markup: |
          <h4>Copy</h4>
          <p>
            <button id="schoolCourseCopyChoosePractice">Practice</button>
            <button id="schoolCourseCopyChooseTest">Test</button>
          </p>
      - name: SchoolCourseCopyPracticePanel
        note: A page to let the user copy without recording the results.
        markup: |
          <p>
            <button id="schoolCourseCopyPracticeStart">Start</button>
            <button id="schoolCourseCopyPracticeCheck" class="unseen">Check</button>
          </p>
          <textarea id="schoolCourseCopyPracticeCopy" class="resize-me-width"></textarea>
          <div id="schoolCourseCopyPracticeText" class="resize-me-width"></div>
      - name: SchoolCourseCopyTestPanel
        note: A page to let the user copy with results recorded.
        markup: |
          <p>
            <button id="schoolCourseCopyTestStart">Start</button>
            <button id="schoolCourseCopyTestCheck" class="unseen">Check</button>
          </p>
          <textarea id="schoolCourseCopyTestCopy" class="resize-me-width"></textarea>
          <div id="schoolCourseCopyTestText" class="resize-me-width"></div>
    - name: SchoolCourseKeyTab
      label: Key
      panels:
      - name: SchoolCourseKeyCompletedPanel
        note: |
          Displayed when the user has successfully copied for this lesson.
          Maybe it should display the results.
        markup: <p>TBD</p>
      - name: SchoolCourseKeyChoosePanel
        note: |
          Displayed when the user has successfully copied for this lesson.
          Maybe it should display the results.
        markup: |
          <h4>Key</h4>
          <p>
            <button id="schoolCourseKeyChoosePractice">Practice</button>
            <button id="schoolCourseKeyChooseTest">Test</button>
          </p>
      - name: SchoolCourseKeyPracticePanel
        note: A page to let the user copy without recording the results.
        markup: |
          <h3 id="schoolCourseKeyPracticeH"></h3>
          <p>
            <label for="schoolCourseKeyPracticeMetronomeOn">Use the metronome.</label>
            <input type="checkbox" id="schoolCourseKeyPracticeMetronomeOn" checked />
          </p>
          <p id="schoolCourseKeyPracticeP">
            <button id="schoolCourseKeyPracticeStart">Start</button>
            <button id="schoolCourseKeyPracticeCheck" class="unseen">Check</button>
          </p>
          <div id="schoolCourseKeyPracticeKey" class="user-not-keying resize-me-width"></div>
          <div id="schoolCourseKeyPracticeCopy" class="resize-me-width"></div>
      - name: SchoolCourseKeyTestPanel
        note: A page to let the user copy with results recorded.
        markup: |
          <h3 id="schoolCourseKeyTestH"></h3>
          <p id="schoolCourseKeyTestP">
            <button id="schoolCourseKeyTestStart">Start</button>
            <button id="schoolCourseKeyTestCheck" class="unseen">Check</button>
          </p>
          <div id="schoolCourseKeyTestKey" class="user-not-keying resize-me-width"></div>
          <div id="schoolCourseKeyTestCopy" class="resize-me-width"></div>
- name: ProveButton
  label: Prove
  panels:
  - name: ProveButtonPanel
    note: |
      The panel for the prove button.
      Run the tests from this panel.
      Do not move the ProveButtonPanel from the ProveButton.
    markup: |
      <p>Prove It!</p>
`

	proveDotGo = `package prove

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/rendererprocess/framework/proofs"
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
		*msg = append(*msg, fmt.Sprintf("The \"SchoolButton\" has %%d panels", len(pNames)))
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
		*msg = append(*msg, fmt.Sprintf("\"SchoolPanel\" has %%d tabs", len(tNames)))
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
		*msg = append(*msg, fmt.Sprintf("\"SchoolCourseTab\" has %%d tabs", len(pNames)))
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
		*msg = append(*msg, fmt.Sprintf("\"SchoolCourseCopyTab\" has %%d tabs", len(pNames)))
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
`
)

func main() {
	var testing = false
	var wd string
	var err error
	if wd, err = os.Getwd(); err != nil {
		return
	}
	appName := filepath.Base(wd)
	sourceCodeFolderPath := filepath.Join(wd, appName+"test")
	err = fix.Refactor(appName, "Move panel from a tab to a button.", sourceCodeFolderPath, kickwasmDotYAML, rekickwasmDotYAML, proveDotGo, false, testing)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
