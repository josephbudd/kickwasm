package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/fix"
)

const (
	kickwasmDotYAML = `title: %[1]s test
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
- name: FirstButton
  label: 1
  panels:
  - name: FirstButtonPanel
    note: The panel for the first button.
    markup: <p>1</p>
- name: ProveButton
  label: 2
  panels:
  - name: ProveButtonPanel
    note: |
      The panel for the second button.
      Run the tests from this panel.
    markup: |
      <p>2</p>
`
	rekickwasmDotYAML = `title: %[1]s test
importPath: github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest
buttons:
- name: ProveButton
  label: 2
  panels:
  - name: ProveButtonPanel
    note: |
      The panel for the second button.
      Run the tests from this panel.
    markup: |
      <p>2</p>
- name: FirstButton
  label: 1
  panels:
  - name: FirstButtonPanel
    note: The panel for the first button.
    markup: <p>1</p>
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

	homeButtonOrder(&msg)
	homeButtonPanels(&msg)

	return
}

func homeButtonOrder(msg *[]string) {
	bb := proofs.HomeButtonsNames()
	var errs bool
	if bb[0] != "FirstButton" {
		*msg = append(*msg, "bb[0] != \"FirstButton\"")
		errs = true
	}
	if bb[1] != "ProveButton" {
		*msg = append(*msg, "bb[1] != \"ProveButton\"")
		errs = true
	}
	if errs {
		*msg = append(*msg, fmt.Sprintf("bb is %%#v", bb))
	}
}

func homeButtonPanels(msg *[]string) {
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
	if bad {
		return
	}
	pNames2 := bpn["ProveButton"]
	if len(pNames2) == 0 {
		*msg = append(*msg, "ProveButton has no panels")
	} else {
		if pNames2[0] != "ProveButtonPanel" {
			*msg = append(*msg, "ProveButton does not have ProveButtonPanel")
		}
	}
	pNames1 := bpn["FirstButton"]
	if len(pNames1) == 0 {
		*msg = append(*msg, "FirstButton has no panels")
	} else {
		if pNames1[0] != "FirstButtonPanel" {
			*msg = append(*msg, "FirstButton does not have FirstButtonPanel")
		}
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
	err = fix.Refactor(appName, sourceCodeFolderPath, kickwasmDotYAML, rekickwasmDotYAML, proveDotGo, true, testing)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
