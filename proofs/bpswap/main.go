package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/fix"
)

// Swap buttons between 2 panels.
// DO NOT EVER MOVE ProveButtonPanel OUT OF ProveButton

const (
	kickwasmDotYAML = `title: Swap buttons between 2 panels.
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
      - name: FirstButtonPanelFirstButtonPanel
        note: FirstButtonPanelFirstButtonPanel.
        markup: <p>FirstButtonPanel FirstButton Panel</p>
    - name: FirstButtonPanelSecondButton
      label: FirstButtonPanel SecondButton
      panels:
      - name: FirstButtonPanelSecondButtonPanel
        note: FirstButtonPanelSecondButtonPanel.
        markup: <p>FirstButtonPanel SecondButton Panel</p>
- name: SecondButton
  label: 2
  panels:
  - name: SecondButtonPanel
    buttons:
    - name: SecondButtonPanelFirstButton
      label: SecondButtonPanel FirstButton
      panels:
      - name: SecondButtonPanelFirstButtonPanel
        note: SecondButtonPanelFirstButtonPanel.
        markup: <p>SecondButtonPanel FirstButton Panel</p>
    - name: SecondButtonPanelSecondButton
      label: SecondButtonPanel SecondButton
      panels:
      - name: SecondButtonPanelSecondButtonPanel
        note: SecondButtonPanelSecondButtonPanel.
        markup: <p>SecondButtonPanel SecondButton Panel</p>
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
	rekickwasmDotYAML = `title: Swap buttons between 2 panels.
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
- name: FirstButton
  label: 1
  panels:
  - name: FirstButtonPanel
    buttons:
    - name: SecondButtonPanelFirstButton
      label: SecondButtonPanel FirstButton
      panels:
      - name: SecondButtonPanelFirstButtonPanel
        note: SecondButtonPanelFirstButtonPanel.
        markup: <p>SecondButtonPanel FirstButton Panel</p>
    - name: FirstButtonPanelSecondButton
      label: FirstButtonPanel SecondButton
      panels:
      - name: FirstButtonPanelSecondButtonPanel
        note: FirstButtonPanelSecondButtonPanel.
        markup: <p>FirstButtonPanel SecondButton Panel</p>
- name: SecondButton
  label: 2
  panels:
  - name: SecondButtonPanel
    buttons:
    - name: FirstButtonPanelFirstButton
      label: FirstButtonPanel FirstButton
      panels:
      - name: FirstButtonPanelFirstButtonPanel
        note: FirstButtonPanelFirstButtonPanel.
        markup: <p>FirstButtonPanel FirstButton Panel</p>
    - name: SecondButtonPanelSecondButton
      label: SecondButtonPanel SecondButton
      panels:
      - name: SecondButtonPanelSecondButtonPanel
        note: SecondButtonPanelSecondButtonPanel.
        markup: <p>SecondButtonPanel SecondButton Panel</p>
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

	checkFirstPanelButtonOrder(&msg)
	checkSecondPanelButtonOrder(&msg)

	return
}

func checkFirstPanelButtonOrder(msg *[]string) {
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
	// Have the panel so check it's button order
	pnbn := proofs.PanelNameButtonNames()
	var bNames []string
	if bNames, found = pnbn["FirstButtonPanel"]; !found {
		*msg = append(*msg, "The panel \"FirstButtonPanel\" is not in proof.PanelNameButtonNames()")
		return
	}
	if len(bNames) != 2 {
		*msg = append(*msg, fmt.Sprintf("\"FirstButtonPanel\" has %%d buttons", len(bNames)))
		return
	}
	if bNames[0] != "SecondButtonPanelFirstButton" {
		*msg = append(*msg, "The \"FirstButtonPanel\"'s first button is not \"SecondButtonPanelFirstButton\"")
		return
	}
	if bNames[1] != "FirstButtonPanelSecondButton" {
		*msg = append(*msg, "The \"FirstButtonPanel\"'s second button is not \"FirstButtonPanelSecondButton\"")
		return
	}
}

func checkSecondPanelButtonOrder(msg *[]string) {
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
	// Have the panel so check it's button order
	pnbn := proofs.PanelNameButtonNames()
	var bNames []string
	if bNames, found = pnbn["SecondButtonPanel"]; !found {
		*msg = append(*msg, "The panel \"SecondButtonPanel\" is not in proof.PanelNameButtonNames()")
		return
	}
	if len(bNames) != 2 {
		*msg = append(*msg, fmt.Sprintf("\"SecondButtonPanel\" has %%d buttons", len(bNames)))
		return
	}
	if bNames[0] != "FirstButtonPanelFirstButton" {
		*msg = append(*msg, "The \"SecondButtonPanel\"'s first button is not \"FirstButtonPanelFirstButton\"")
		return
	}
	if bNames[1] != "SecondButtonPanelSecondButton" {
		*msg = append(*msg, "The \"SecondButtonPanel\"'s second button is not \"SecondButtonPanelSecondButton\"")
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
	err = fix.Refactor(appName, "Swap buttons between 2 panels.", sourceCodeFolderPath, kickwasmDotYAML, rekickwasmDotYAML, proveDotGo, false, testing)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}