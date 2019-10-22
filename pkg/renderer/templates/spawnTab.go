package templates

// SpawnTabPrepare is the tab's prepare.go template.
const SpawnTabPrepare = `{{$Dot := .}}// +build js, wasm

package {{call .PackageNameCase .TabName}}

import ({{ range .PrepareImports }}
	{{.}}{{end}}
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// Prepare initializes this package in preparation for spawning.
func Prepare(quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending, vtools *viewtools.Tools, njs *notjs.NotJS, help *paneling.Help) {
	tools = vtools
{{ range .PanelNames }}
	{{ call $Dot.PackageNameCase . }}.Prepare(quitChan, eojChan, receiveChan, sendChan, vtools, njs, help){{end}}
}
`

// SpawnTabSpawn is the tab's spawn.go template.
const SpawnTabSpawn = `{{$Dot := .}}// +build js, wasm

package {{call .PackageNameCase .TabName}}

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/pkg/errors"
{{ range .SpawnImports }}
	{{.}}{{end}}
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

const (
	tabBarID = "{{.TabBarID}}"
	tabName  = "{{.TabName}}"
)

var (
	markupTemplatePaths = {{.MarkupTemplatePaths}}
	tools  *viewtools.Tools
)

// Tab represents a tab that will spawn and unspawn an html tab bar tab.
type Tab struct {
	uniqueID      uint64
	hTMLButton    js.Value
	prepareToUnSpawns []func()
}

// Spawn creates the DOM elements and go code for a tab.
// This is called by the tab bar.
// Param tabLabel is the label in the tab button. The button's innerText.
// Param panelHeading is the heading for each panel.
// Param panelData is an empty interface passed to each panel's spawn func.
// Returns the tab unspawn func and the error.
func Spawn(tabLabel, panelHeading string, panelData interface{}) (unspawn func() error, err error) {

	defer func() {
		if err != nil {
			message := fmt.Sprintf("%s.Spawn()", "{{call .PackageNameCase .TabName}}")
			err = errors.WithMessage(err, message)
		}
	}()

	// Spawn the DOM elements.
	var tabButton js.Value
	var tabPanelHeader js.Value
	var uniqueID uint64
	var panelNameID map[string]string
	if tabButton, tabPanelHeader, uniqueID, panelNameID, err = tools.SpawnTab(tabBarID, tabName, tabLabel, panelHeading, markupTemplatePaths); err != nil {
		return
	}
	// Define the tab.
	tab := &Tab{
		hTMLButton:    tabButton,
		uniqueID:      uniqueID,
		prepareToUnSpawns: make([]func(), 0, 20),
	}
	unspawn = tab.unSpawn
	// Build the go code.
	var f func()
{{ range .PanelNames }}
	if f, err = {{ call $Dot.PackageNameCase . }}.BuildPanel(uniqueID, tabButton, tabPanelHeader, panelNameID, panelData, unspawn); err != nil {
		return
	}
	tab.prepareToUnSpawns = append(tab.prepareToUnSpawns, f){{end}}
	tools.IncSpawnedPanels(len(tab.prepareToUnSpawns))
	return
}

// unSpawn totally removes the DOM elements and go code of a spawned tab.
// Returns the error.
func (tab *Tab) unSpawn() (err error) {

	defer func() {
		if err != nil {
			message := fmt.Sprintf("{{call .PackageNameCase .TabName}}.unSpawn(): uniqueID is %d", tab.uniqueID)
			err = errors.WithMessage(err, message)
		}
	}()

	tools.DecSpawnedPanels(len(tab.prepareToUnSpawns))

	messages := make([]string, 0, 2)
	// Remove the tab and panels from the DOM.
	if err = tools.UnSpawnTab(tab.hTMLButton); err != nil {
		messages = append(messages, err.Error())
	}
	// Unregister each panel controller's javascript call backs.
	if err = tools.UnRegisterCallBacks(tab.uniqueID); err != nil {
		messages = append(messages, err.Error())
	}
	// Stop each panel messenger's message dispatcher.
	for _, prepareToUnSpawn := range tab.prepareToUnSpawns {
		prepareToUnSpawn()
	}

	// construct a new error from the accumulated errors.
	if len(messages) > 0 {
		err = errors.New(strings.Join(messages, "\n"))
	}
	return
}

func (tab *Tab) Tab() js.Value {
	return tab.hTMLButton
}

func (tab *Tab) UniqueID() uint64 {
	return tab.uniqueID
}
`
