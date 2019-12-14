package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/fix"
)

// Swap panels between 2 buttons.
// DO NOT EVER MOVE ProveButtonPanel OUT OF ProveButton

const (
	kickwasmDotYAML = `title: Swap panels between 2 buttons.
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
- name: FirstButton
  label: 1
  panels:
  - name: FirstButtonPanel
    buttons:
    - name: FirstButtonPanelFirstButton
      label: FirstButtonPanel FirstButton
      panels:
      - name: FirstButtonPanelFirstButtonFirstPanel
        note: FirstButtonPanelFirstButtonFirstPanel.
        markup: <p>FirstButtonPanel FirstButton FirstPanel</p>
      - name: FirstButtonPanelFirstButtonSecondPanel
        note: FirstButtonPanelFirstButtonSecondPanel.
        markup: <p>FirstButtonPanel FirstButton SecondPanel</p>
- name: SecondButton
  label: 2
  panels:
  - name: SecondButtonPanel
    buttons:
    - name: SecondButtonPanelFirstButton
      label: SecondButtonPanel FirstButton
      panels:
      - name: SecondButtonPanelFirstButtonFirstPanel
        note: SecondButtonPanelFirstButtonFirstPanel.
        markup: <p>SecondButtonPanel FirstButton FirstPanel</p>
      - name: SecondButtonPanelFirstButtonSecondPanel
        note: SecondButtonPanelFirstButtonSecondPanel.
        markup: <p>SecondButtonPanel FirstButton SecondPanel</p>
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
	rekickwasmDotYAML = `title: Swap panels between 2 buttons.
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
- name: FirstButton
  label: 1
  panels:
  - name: FirstButtonPanel
    buttons:
    - name: FirstButtonPanelFirstButton
      label: FirstButtonPanel FirstButton
      panels:
      - name: SecondButtonPanelFirstButtonFirstPanel
        note: SecondButtonPanelFirstButtonFirstPanel.
        markup: <p>SecondButtonPanel FirstButton FirstPanel</p>
      - name: FirstButtonPanelFirstButtonFirstPanel
        note: FirstButtonPanelFirstButtonFirstPanel.
        markup: <p>FirstButtonPanel FirstButton FirstPanel</p>
- name: SecondButton
  label: 2
  panels:
  - name: SecondButtonPanel
    buttons:
    - name: SecondButtonPanelFirstButton
      label: SecondButtonPanel FirstButton
      panels:
      - name: FirstButtonPanelFirstButtonSecondPanel
        note: FirstButtonPanelFirstButtonSecondPanel.
        markup: <p>FirstButtonPanel FirstButton SecondPanel</p>
      - name: SecondButtonPanelFirstButtonSecondPanel
        note: SecondButtonPanelFirstButtonSecondPanel.
        markup: <p>SecondButtonPanel FirstButton SecondPanel</p>
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

	checkFirstButton(&msg)
	checkSecondButton(&msg)

	return
}

func checkFirstButton(msg *[]string) {
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
	// Have the panel so check it's button
	pnbn := proofs.PanelNameButtonNames()
	var bNames []string
	if bNames, found = pnbn["FirstButtonPanel"]; !found {
		*msg = append(*msg, "The panel \"FirstButtonPanel\" is not in proof.PanelNameButtonNames()")
		return
	}
	if len(bNames) != 1 {
		*msg = append(*msg, fmt.Sprintf("\"FirstButtonPanel\" has %%d buttons", len(bNames)))
		return
	}
	if bNames[0] != "FirstButtonPanelFirstButton" {
		*msg = append(*msg, "The \"FirstButtonPanel\"'s first button is not \"FirstButtonPanelFirstButton\"")
		return
	}
	// Get this button's panels.
	pNames = bnpn["FirstButtonPanelFirstButton"]
	if pNames[0] != "SecondButtonPanelFirstButtonFirstPanel" {
		*msg = append(*msg, "The \"FirstButtonPanel\"'s first panel is not \"SecondButtonPanelFirstButtonFirstPanel\"")
		return
	}
	if pNames[1] != "FirstButtonPanelFirstButtonFirstPanel" {
		*msg = append(*msg, "The \"FirstButtonPanel\"'s second panel is not \"FirstButtonPanelFirstButtonFirstPanel\"")
		return
	}
}

func checkSecondButton(msg *[]string) {
/*
- name: SecondButton
  label: 2
  panels:
  - name: SecondButtonPanel
    buttons:
    - name: SecondButtonPanelFirstButton
      label: SecondButtonPanel FirstButton
      panels:
      - name: FirstButtonPanelFirstButtonSecondPanel
        note: FirstButtonPanelFirstButtonSecondPanel.
        markup: <p>FirstButtonPanel FirstButton SecondPanel</p>
      - name: SecondButtonPanelFirstButtonSecondPanel
        note: SecondButtonPanelFirstButtonSecondPanel.
        markup: <p>SecondButtonPanel FirstButton SecondPanel</p>
*/
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
	// Have the panel so check it's button
	pnbn := proofs.PanelNameButtonNames()
	var bNames []string
	if bNames, found = pnbn["SecondButtonPanel"]; !found {
		*msg = append(*msg, "The panel \"SecondButtonPanel\" is not in proof.PanelNameButtonNames()")
		return
	}
	if len(bNames) != 1 {
		*msg = append(*msg, fmt.Sprintf("\"SecondButtonPanel\" has %%d buttons", len(bNames)))
		return
	}
	if bNames[0] != "SecondButtonPanelFirstButton" {
		*msg = append(*msg, "The \"SecondButtonPanel\"'s first button is not \"SecondButtonPanelFirstButton\"")
		return
	}
	// Get this button's panels.
	pNames = bnpn["SecondButtonPanelFirstButton"]
	if pNames[0] != "FirstButtonPanelFirstButtonSecondPanel" {
		*msg = append(*msg, "The \"SecondButtonPanel\"'s first panel is not \"FirstButtonPanelFirstButtonSecondPanel\"")
		return
	}
	if pNames[1] != "SecondButtonPanelFirstButtonSecondPanel" {
		*msg = append(*msg, "The \"SecondButtonPanel\"'s second panel is not \"SecondButtonPanelFirstButtonSecondPanel\"")
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
	err = fix.Refactor(appName, "Swap panels between 2 buttons.", sourceCodeFolderPath, kickwasmDotYAML, rekickwasmDotYAML, proveDotGo, false, testing)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
