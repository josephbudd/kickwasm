package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/fix"
)

// Swap tabs between 2 panels.
// DO NOT EVER MOVE ProveButtonPanel OUT OF ProveButton

const (
	kickwasmDotYAML = `title: Swap tabs between 2 panels.
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
    - name: FirstButtonPanelSecondTab
      label: FirstButtonPanel SecondTab
      panels:
      - name: FirstButtonPanelSecondTabPanel
        note: FirstButtonPanelSecondTabPanel.
        markup: <p>FirstButtonPanel SecondTab Panel</p>
- name: SecondButton
  label: 1
  panels:
  - name: SecondButtonPanel
    tabs:
    - name: SecondButtonPanelFirstTab
      label: SecondButtonPanel FirstTab
      panels:
      - name: SecondButtonPanelFirstTabPanel
        note: SecondButtonPanelFirstTabPanel.
        markup: <p>SecondButtonPanel FirstTab Panel</p>
    - name: SecondButtonPanelSecondTab
      label: SecondButtonPanel SecondTab
      panels:
      - name: SecondButtonPanelSecondTabPanel
        note: SecondButtonPanelSecondTabPanel.
        markup: <p>SecondButtonPanel SecondTab Panel</p>
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
	rekickwasmDotYAML = `title: Swap tabs between 2 panels.
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
    - name: SecondButtonPanelSecondTab
      label: SecondButtonPanel SecondTab
      panels:
      - name: SecondButtonPanelSecondTabPanel
        note: SecondButtonPanelSecondTabPanel.
        markup: <p>SecondButtonPanel SecondTab Panel</p>
- name: SecondButton
  label: 1
  panels:
  - name: SecondButtonPanel
    tabs:
    - name: SecondButtonPanelFirstTab
      label: SecondButtonPanel FirstTab
      panels:
      - name: SecondButtonPanelFirstTabPanel
        note: SecondButtonPanelFirstTabPanel.
        markup: <p>SecondButtonPanel FirstTab Panel</p>
    - name: FirstButtonPanelSecondTab
      label: FirstButtonPanel SecondTab
      panels:
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

	proveDotGo = `package prove

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/rendererprocess/framework/proofs"
	"github.com/pkg/errors"
)

// Pass will not return an error if rekickwasm worked.
func Pass() (err error) {

	msg := make([]string, 0, 10)
	defer func() {
		if len(msg) > 0 {
			err = errors.New(strings.Join(msg, "\n"))
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
		*msg = append(*msg, "The tab \"FirstButtonPanel\" is not in proof.PanelNameTabNames()")
		return
	}
	if len(tNames) != 2 {
		*msg = append(*msg, fmt.Sprintf("\"FirstButtonPanel\" has %%d tabs", len(tNames)))
		return
	}
	if tNames[0] != "FirstButtonPanelFirstTab" {
		*msg = append(*msg, "The \"FirstButtonPanel\"'s first tab is not \"FirstButtonPanelFirstTab\"")
		return
	}
	if tNames[1] != "SecondButtonPanelSecondTab" {
		*msg = append(*msg, "The \"FirstButtonPanel\"'s second tab is not \"SecondButtonPanelSecondTab\"")
		return
	}
}

func checkSecondTabOrder(msg *[]string) {
	bnpn := proofs.ButtonNamePanelNames()
	var pNames []string
	var found bool
	if pNames, found = bnpn["SecondButton"]; !found {
		*msg = append(*msg, "SecondButton not found")
		return
	}
	// Have the first button so check it's panels.
	if len(pNames) != 1 {
		*msg = append(*msg, fmt.Sprintf("The \"SecondButton\" has %%d panels", len(pNames)))
		return
	}
	if pNames[0] != "SecondButtonPanel" {
		*msg = append(*msg, "The \"SecondButton\"'s panel is not \"SecondButtonPanel\"")
		return
	}
	// Have the panel so check it's tab order
	pntn := proofs.PanelNameTabNames()
	var tNames []string
	if tNames, found = pntn["SecondButtonPanel"]; !found {
		*msg = append(*msg, "The tab \"SecondButtonPanel\" is not in proof.PanelNameTabNames()")
		return
	}
	if len(tNames) != 2 {
		*msg = append(*msg, fmt.Sprintf("\"SecondButtonPanel\" has %%d tabs", len(tNames)))
		return
	}
	if tNames[0] != "SecondButtonPanelFirstTab" {
		*msg = append(*msg, "The \"SecondButtonPanel\"'s first tab is not \"SecondButtonPanelFirstTab\"")
		return
	}
	if tNames[1] != "FirstButtonPanelSecondTab" {
		*msg = append(*msg, "The \"SecondButtonPanel\"'s second tab is not \"FirstButtonPanelSecondTab\"")
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
	err = fix.Refactor(appName, "Swap tabs between 2 panels.", sourceCodeFolderPath, kickwasmDotYAML, rekickwasmDotYAML, proveDotGo, false, testing)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
