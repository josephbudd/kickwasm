package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/fix"
)

// Rotate 2 panels in a tab.
// DO NOT EVER MOVE ProveButtonPanel OUT OF ProveButton

const (
	kickwasmDotYAML = `title: Rotate 2 panels in a tab.
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
- name: FirstButton
  label: 1
  panels:
  - name: FirstButtonPanel
    tabs:
    - name: FirstButtonPanelFirstTab
      label: FirstButtonPanel FirstTab
      panels:
      - name: FirstButtonPanelFirstTabPanel
        note: FirstButtonPanelFirstTabPanel.
        markup: <p>FirstButtonPanel FirstTab Panel</p>
      - name: FirstButtonPanelSecondTabPanel
        note: FirstButtonPanelSecondTabPanel.
        markup: <p>FirstButtonPanel SecondTab Panel</p>
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
	rekickwasmDotYAML = `title: Rotate 2 panels in a tab.
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
- name: FirstButton
  label: 1
  panels:
  - name: FirstButtonPanel
    tabs:
    - name: FirstButtonPanelFirstTab
      label: FirstButtonPanel FirstTab
      panels:
      - name: FirstButtonPanelSecondTabPanel
        note: FirstButtonPanelSecondTabPanel.
        markup: <p>FirstButtonPanel SecondTab Panel</p>
      - name: FirstButtonPanelFirstTabPanel
        note: FirstButtonPanelFirstTabPanel.
        markup: <p>FirstButtonPanel FirstTab Panel</p>
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

	checkFirstTabOrder(&msg)

	return
}

func checkFirstTabOrder(msg *[]string) {
	bnpn := proofs.ButtonNamePanelNames()
	var pNames []string
	var found bool
	if pNames, found = bnpn["FirstButton"]; !found {
		*msg = append(*msg, "FirstButton not found")
		return
	}
	// Have the first button so check it's panels.
	if len(pNames) != 1 {
		*msg = append(*msg, fmt.Sprintf("The \"FirstButton\" has %%d panels", len(pNames)))
		return
	}
	if pNames[0] != "FirstButtonPanel" {
		*msg = append(*msg, "The \"FirstButton\"'s panel is not \"FirstButtonPanel\"")
		return
	}
	// Have the panel so check it's tab order
	pntn := proofs.PanelNameTabNames()
	var tNames []string
	if tNames, found = pntn["FirstButtonPanel"]; !found {
		*msg = append(*msg, "The \"FirstButtonPanel\" is not in proof.PanelNameTabNames()")
		return
	}
	if len(tNames) != 1 {
		*msg = append(*msg, fmt.Sprintf("\"FirstButtonPanel\" has %%d tabs", len(tNames)))
		return
	}
	if tNames[0] != "FirstButtonPanelFirstTab" {
		*msg = append(*msg, "The \"FirstButtonPanel\"'s first tab is not \"FirstButtonPanelFirstTab\"")
		return
	}
	tnpn := proofs.TabNamePanelNames()
	if pNames, found = tnpn["FirstButtonPanelFirstTab"]; !found {
		*msg = append(*msg, "The \"FirstButtonPanelFirstTab\" is not in proof.TabNamePanelNames()")
		return
	}
	if pNames[0] != "FirstButtonPanelSecondTabPanel" {
		*msg = append(*msg, "The \"FirstButtonPanelFirstTab\"'s first panel is not \"FirstButtonPanelSecondTabPanel\"")
		return
	}
	if pNames[1] != "FirstButtonPanelFirstTabPanel" {
		*msg = append(*msg, "The \"FirstButtonPanelFirstTab\"'s first panel is not \"FirstButtonPanelFirstTabPanel\"")
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
	err = fix.Refactor(appName, "Rotate 2 panels in a tab.", sourceCodeFolderPath, kickwasmDotYAML, rekickwasmDotYAML, proveDotGo, false, testing)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
