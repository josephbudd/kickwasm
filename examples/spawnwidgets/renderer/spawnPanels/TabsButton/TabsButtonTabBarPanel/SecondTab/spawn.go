package secondtab

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/pkg/errors"

	helloworldtemplatepanel "github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab/HelloWorldTemplatePanel"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

const (
	tabBarID = "tabsMasterView_home_pad_TabsButton_TabsButtonTabBarPanel_tab_bar"
	tabName  = "SecondTab"
)

var (
	markupTemplatePaths = []string{"spawnTemplates/TabsButton/TabsButtonTabBarPanel/SecondTab/HelloWorldTemplatePanel.tmpl"}
	tools  *viewtools.Tools
)

// Tab represents a tab that will spawn and unspawn an html tab bar tab.
type Tab struct {
	uniqueID      uint64
	hTMLButton    js.Value
	stopListeners []func()
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
			message := fmt.Sprintf("%s.Spawn()", "secondtab")
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
		stopListeners: make([]func(), 0, 20),
	}
	unspawn = tab.unSpawn
	// Build the go code.
	var f func()

	if f, err = helloworldtemplatepanel.BuildPanel(uniqueID, tabButton, tabPanelHeader, panelNameID, panelData, unspawn); err != nil {
		return
	}
	tab.stopListeners = append(tab.stopListeners, f)
	tools.IncSpawnedPanels(len(tab.stopListeners))
	return
}

// unSpawn totally removes the DOM elements and go code of a spawned tab.
// Returns the error.
func (tab *Tab) unSpawn() (err error) {

	defer func() {
		if err != nil {
			message := fmt.Sprintf("secondtab.unSpawn(): uniqueID is %d", tab.uniqueID)
			err = errors.WithMessage(err, message)
		}
	}()

	tools.DecSpawnedPanels(len(tab.stopListeners))

	messages := make([]string, 0, 2)
	// Remove the tab and panels from the DOM.
	if err = tools.UnSpawnTab(tab.hTMLButton); err != nil {
		messages = append(messages, err.Error())
	}
	// Unregister each panel controller's javascript call backs.
	if err = tools.UnRegisterCallBacks(tab.uniqueID); err != nil {
		messages = append(messages, err.Error())
	}
	// Stop each panel's caller's message listener.
	for _, stopListener := range tab.stopListeners {
		stopListener()
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
