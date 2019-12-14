package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/fix"
)

// Add an Initial Button with tabs.
// DO NOT EVER MOVE ProveButtonPanel OUT OF ProveButton
//
// The refactor rotates the two initial button positions.

const (
	kickwasmDotYAML = `title: Add an Initial Button with tabs.
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
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
	rekickwasmDotYAML = `title: Add an Initial Button with tabs.
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
- name: FirstButton
  label: 1
  panels:
  - name: FirstButtonPanel
    tabs:
    - name: FirstButtonPanelTab
      label: First
      panels:
      - name: FirstButtonPanelTabPanel
        note: The FirstButtonPanelTabPanel
        markup: <p>The FirstButtonPanelTabPanel</p>
- name: SecondButton
  label: 2
  panels:
  - name: SecondButtonPanel
    tabs:
    - name: SecondButtonPanelTab
      label: Second
      panels:
      - name: SecondButtonPanelTabPanel
        note: The SecondButtonPanelTabPanel
        markup: <p>The SecondButtonPanelTabPanel</p>
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

	checkInitButtonOrder(&msg)
	checkInitButtonPanels(&msg)
	checkTabPanels(&msg)

	return
}

func checkInitButtonOrder(msg *[]string) {
	bb := proofs.HomeButtonsNames()
	var errs bool
	if bb[0] != "FirstButton" {
		*msg = append(*msg, "bb[0] != \"FirstButton\"")
		errs = true
	}
	if bb[1] != "SecondButton" {
		*msg = append(*msg, "bb[1] != \"SecondButton\"")
		errs = true
	}
	if bb[2] != "ProveButton" {
		*msg = append(*msg, "bb[2] != \"ProveButton\"")
		errs = true
	}
	if errs {
		*msg = append(*msg, fmt.Sprintf("bb is %%#v", bb))
	}
}

func checkInitButtonPanels(msg *[]string) {
	bpn := proofs.ButtonNamePanelNames()
	var bad bool
	var found bool
	if _, found = bpn["ProveButton"]; !found {
		*msg = append(*msg, "ProveButton not found")
		bad = true
	}
	if _, found = bpn["FirstButton"]; !found {
		*msg = append(*msg, "FirstButton not found")
		bad = true
	}
	if _, found = bpn["SecondButton"]; !found {
		*msg = append(*msg, "SecondButton not found")
		bad = true
	}
	if bad {
		return
	}
	var pNames []string
	pNames = bpn["ProveButton"]
	if len(pNames) == 0 {
		*msg = append(*msg, "ProveButton has no panels")
	} else {
		if pNames[0] != "ProveButtonPanel" {
			*msg = append(*msg, "ProveButton does not have ProveButtonPanel")
		}
	}
	pNames = bpn["FirstButton"]
	if len(pNames) == 0 {
		*msg = append(*msg, "FirstButton has no panels")
	} else {
		if pNames[0] != "FirstButtonPanel" {
			*msg = append(*msg, "FirstButton does not have FirstButtonPanel")
		}
	}
	pNames = bpn["SecondButton"]
	if len(pNames) == 0 {
		*msg = append(*msg, "SecondButton has no panels")
	} else {
		if pNames[0] != "SecondButtonPanel" {
			*msg = append(*msg, "SecondButton does not have SecondButtonPanel")
		}
	}
}

func checkTabPanels(msg *[]string) {
	tnpn := proofs.TabNamePanelNames()
	if _, found := tnpn["FirstButtonPanelTab"]; !found {
		*msg = append(*msg, "FirstButtonPanelTab not found.")
		return
	}
	pNames := tnpn["FirstButtonPanelTab"]
	if len(pNames) == 0 {
		*msg = append(*msg, "FirstButtonPanelTab has no panels.")
		return
	}
	if pNames[0] != "FirstButtonPanelTabPanel" {
		*msg = append(*msg, "FirstButtonPanelTab does not have FirstButtonPanelTabPanel")
	}

	if _, found := tnpn["SecondButtonPanelTab"]; !found {
		*msg = append(*msg, "SecondButtonPanelTab not found.")
		return
	}
	pNames = tnpn["SecondButtonPanelTab"]
	if len(pNames) == 0 {
		*msg = append(*msg, "SecondButtonPanelTab has no panels.")
		return
	}
	if pNames[0] != "SecondButtonPanelTabPanel" {
		*msg = append(*msg, "SecondButtonPanelTab does not have SecondButtonPanelTabPanel")
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
	err = fix.Refactor(appName, "Add an Initial Button with tabs.", sourceCodeFolderPath, kickwasmDotYAML, rekickwasmDotYAML, proveDotGo, false, testing)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
