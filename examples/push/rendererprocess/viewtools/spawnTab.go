// +build js, wasm

package viewtools

import (
	"fmt"
	"path/filepath"
	"strings"
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/spawnpack"
)

// FixSpawnID fixes a spawn html template element id using the spawns unique id.
// It replaces "{{.SpawnID}}" with the spawn's unique id.
func (tools *Tools) FixSpawnID(id string, uniqueID uint64) (fixedID string) {
	fixedID = strings.ReplaceAll(id, spawnIDReplacePattern, fmt.Sprint(uniqueID))
	return
}
 
// SpawnTab adds an html tab button and panels to the html document.
// It returns their uniqueID, the tab button id, panels names mapped to their ids.
func (tools *Tools) SpawnTab(tabBarID, tabName, tabLabel, tabPanelHeadingText string, userContentPanelPaths []string) (tabButton, tabPanelHeader js.Value, uniqueID uint64, panelNameID map[string]string, err error) {
	uniqueID = tools.newSpawnID()
	notJS := tools.NotJS
	null := js.Null()
	// Find the tab bar.
	var tabBar js.Value
	if tabBar = notJS.GetElementByID(tabBarID); tabBar == null {
		err = errors.New("Unable to find tab bar #" + tabBarID)
		return
	}
	// get the number of tabs in this tab bar.
	_, nTabButtons := notJS.Children(tabBar)
	// Create the button and add it to the tab bar.
	tabButton = notJS.CreateElementBUTTON()
	tabButtonID := tools.BuildSpawnTabButtonID(tabBarID, tabName, uniqueID)
	notJS.SetID(tabButton, tabButtonID)
	notJS.ClassListAddClass(tabButton, TabClassName)
	notJS.ClassListAddClass(tabButton, UnSelectedTabClassName)
	label := notJS.CreateTextNode(tabLabel)
	notJS.AppendChild(tabButton, label)
	notJS.AppendChild(tabBar, tabButton)
	var tabPanelVisiblilityClass string
	if nTabButtons == 0 {
		// if this is the first tab then bring it up front.
		notJS.Focus(tabButton)
		// this will only be visible if it is the first and only tab button.
		tabPanelVisiblilityClass = SeenClassName
	} else {
		// not the first tab
		// * so leave it in back.
		// * the tab panel is hidden.
		tabPanelVisiblilityClass = UnSeenClassName
	}
	// Panels
	// Find the under tab bar div.
	// The tab's panel is inside the under tab bar div.
	underTabBarDivID := tools.buildSpawnUnderTabBarID(tabBarID)
	var underTabBarDiv js.Value
	if underTabBarDiv = notJS.GetElementByID(underTabBarDivID); underTabBarDiv == null {
		err = errors.New("Unable to find under tab bar #" + underTabBarDivID)
		return
	}
	// Create the tab panel
	tabPanel := notJS.CreateElementDIV()
	tabPanelID := tools.buildSpawnTabButtonPanelID(tabButtonID)
	notJS.SetID(tabPanel, tabPanelID)
	notJS.ClassListAddClass(tabPanel, TabPanelClassName)
	notJS.ClassListAddClass(tabPanel, PanelWithHeadingClassName)
	notJS.ClassListAddClass(tabPanel, tabPanelVisiblilityClass)
	tabPanelHeader = notJS.CreateElementH3()
	tabPanelHeadingID := tools.buildSpawnTabButtonPanelHeadingID(tabButtonID)
	notJS.SetID(tabPanelHeader, tabPanelHeadingID)
	notJS.ClassListAddClass(tabPanelHeader, PanelHeadingClassName)
	heading := notJS.CreateTextNode(tabPanelHeadingText)
	notJS.AppendChild(tabPanelHeader, heading)
	notJS.AppendChild(tabPanel, tabPanelHeader)
	// Inside the tab panel is the inner panel
	//   which is the wrapper of group of user content (markup) panels.
	// Create the group
	group := notJS.CreateElementDIV()
	innerPanelID := tools.buildSpawnTabButtonInnerPanelID(tabButtonID)
	notJS.SetID(group, innerPanelID)
	notJS.ClassListAddClass(group, TabPanelGroupClassName)
	notJS.ClassListAddClass(group, UserContentClassName)
	notJS.AppendChild(tabPanel, group)
	// Create each user content & it's markup panel inside the group panel
	l := len(userContentPanelPaths)
	panelNameID = make(map[string]string, l)
	uniqueIDString := fmt.Sprint(uniqueID)
	userContentPanels := make([]js.Value, len(userContentPanelPaths))
	var firstMarkupPanelID string
	for i, path := range userContentPanelPaths {
		var markupbb []byte
		var found bool
		if markupbb, found = spawnpack.Contents(path); !found {
			err = errors.New("Unable to find the spawn template at " + path)
			return
		}
		// get the panel name.
		base := filepath.Base(path)
		l := len(base) - len(filepath.Ext(base))
		panelName := base[:l]
		// user content panel wraps the markup panel
		userContentPanel := notJS.CreateElementDIV()
		userContentPanelID := tools.buildSpawnTabButtonInnerMarkupPanelID(tabButtonID, panelName)
		userContentPanels[i] = userContentPanel
		panelNameID[panelName] = userContentPanelID
		notJS.SetID(userContentPanel, userContentPanelID)
		notJS.ClassListAddClass(userContentPanel, UserContentClassName)
		notJS.ClassListAddClass(userContentPanel, SliderPanelInnerSiblingClassName)
		if i == 0 {
			// The first markup panel is the defauilt and is visible.
			notJS.ClassListAddClass(userContentPanel, SeenClassName)
		} else {
			// The other markup panels are hidden.
			notJS.ClassListAddClass(userContentPanel, UnSeenClassName)
		}
		var hvscroll bool
		if hvscroll, found = tools.panelNameHVScroll[panelName]; !found {
			emsg := fmt.Sprintf("Unable to find panel name %q in tools.panelNameHVScroll", panelName)
			err = errors.New(emsg)
			return
		}
		if hvscroll {
			notJS.ClassListAddClass(userContentPanel, HVScrollClassName)
		} else {
			notJS.ClassListAddClass(userContentPanel, VScrollClassName)
		}
		notJS.AppendChild(group, userContentPanel)
		// markup panel inside the user content panel.
		markupPanel := notJS.CreateElementDIV()
		notJS.ClassListAddClass(markupPanel, SeenClassName)
		// if i == 0 {
		// 	// The first markup panel is the defauilt and is visible.
		// 	notJS.ClassListAddClass(markupPanel, SeenClassName)
		// } else {
		// 	// The other markup panels are hidden.
		// 	notJS.ClassListAddClass(markupPanel, UnSeenClassName)
		// }
		markup := strings.ReplaceAll(string(markupbb), spawnIDReplacePattern, uniqueIDString)
		notJS.SetInnerHTML(markupPanel, markup)
		notJS.AppendChild(userContentPanel, markupPanel)
	}
	// now that the tab panel is fully constructed
	//   add it to the under tab bar div.
	notJS.AppendChild(underTabBarDiv, tabPanel)
	// add the tab bar button panel group
	tools.buttonPanelsMap[tabButtonID] = userContentPanels
	// If this is the first button added to the tab bar
	//   then it is the default panel.
	// Set it as the last panel.
	if _, length := notJS.Children(tabBar); length == 1 {
		tools.addNewSpawnTabBarToLastPanelLevels(tabBarID, firstMarkupPanelID)
	}
	// tab button onclick handler
	tools.setTabBarSpawnButtonOnClick(tabButton, uniqueID)
	return
}

// UnSpawnTab removes a tab button and panels from the html document.
// Returns the error.
func (tools *Tools) UnSpawnTab(tabButton js.Value) (err error ) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "tools.UnSpawnTab")
		}
	}()

	notJS := tools.NotJS
	tabButtonID := notJS.ID(tabButton)

	// Step 1:
	// If there is a tab to the left if there is one
	//   then force click that tab bringing it up front with it's panel.
	tabBar := notJS.ParentNode(tabButton)
	siblings := notJS.ChildrenSlice(tabBar)
	var sibling js.Value
	var i int
	for i, sibling = range siblings {
		if sibling == tabButton {
			break
		}
	}
	if i > 0 {
		tools.ForceTabButtonClick(siblings[i-1])
	}

	// Step 2: The tab button.
	// Remove this tab from the tab bar in the DOM.
	// cleanup data.
	tabBarID := notJS.ID(tabBar)
	tools.removeSpawnTabBarButtonFromLastPanelLevels(tabBarID, tabButtonID)
	// remove the button the DOM.
	notJS.RemoveChild(tabBar, tabButton)

	// Step 3: The tab panels.
	// Remove the tab panels from the DOM.
	panels := tools.buttonPanelsMap[tabButtonID]
	underTabBarDiv := notJS.ParentNode(panels[0])
	for _, p := range panels {
		notJS.RemoveChild(underTabBarDiv, p)
	}
	// cleanup data.
	// remove the tab bar button panel group.
	delete(tools.buttonPanelsMap, tabButtonID)

	return
}

// BuildSpawnTabButtonID forms an id for a spawned tab bar button.
func (tools *Tools) BuildSpawnTabButtonID(tabBarID, tabName string, uniqueID uint64) (id string) {
	id = fmt.Sprintf("%s-%s(%d)", tabBarID, tabName, uniqueID)
	return
}

// BuildSpawnTabButtonMarkupPanelID forms an id for a spawned tab bar button's markup panel.
// This exists mainly for spawn panel groups.
func (tools *Tools) BuildSpawnTabButtonMarkupPanelID(tabBarID, tabName, panelName string, uniqueID uint64) (id string) {
	buttonID := tools.BuildSpawnTabButtonID(tabBarID, tabName, uniqueID)
	id = tools.buildSpawnTabButtonInnerMarkupPanelID(buttonID, panelName)
	return
}

func (tools *Tools) newSpawnID() (id uint64) {
	// spawn ids begin with 1.
	tools.spawnID++
	id = tools.spawnID
	return
}

// buildSpawnUnderTabBarID forms an id for a spawned tab bar button.
func (tools *Tools) buildSpawnUnderTabBarID(tabBarID string) (id string) {
	id = tabBarID + "-under-tab-bar"
	return
}

// buildSpawnTabButtonPanelID forms an id for a spawned tab bar button.
func (tools *Tools) buildSpawnTabButtonPanelID(buttonID string) (id string) {
	id = buttonID + "Panel"
	return
}

// buildSpawnTabButtonPanelHeadingID forms an id for a spawned tab bar button button's inner panel.
func (tools *Tools) buildSpawnTabButtonPanelHeadingID(buttonID string) (id string) {
	id = tools.buildSpawnTabButtonPanelID(buttonID) + "-panel-heading"
	return
}

// buildSpawnTabButtonInnerPanelID forms an id for a spawned tab bar button button's inner panel.
func (tools *Tools) buildSpawnTabButtonInnerPanelID(buttonID string) (id string) {
	id = tools.buildSpawnTabButtonPanelID(buttonID) + "-inner"
	return
}

// buildSpawnTabButtonInnerMarkupPanelID forms an id for a spawned tab bar button's markup panel.
func (tools *Tools) buildSpawnTabButtonInnerMarkupPanelID(buttonID, panelName string) (id string) {
	id = tools.buildSpawnTabButtonInnerPanelID(buttonID) + "-" + panelName
	return
}

func (tools *Tools) addTabBarButton(tabBar, newButton js.Value) {
	notJS := tools.NotJS
	notJS.AppendChild(tabBar, newButton)
}

func (tools *Tools) setTabBarSpawnButtonOnClick(tabBarButton js.Value, uniqueID uint64) {
	notJS := tools.NotJS
	cb := tools.RegisterSpawnEventCallBack(
		func(event js.Value) interface{} {
			target := notJS.GetEventTarget(event)
			tools.handleTabButtonOnClick(target)
			return nil
		},
		true, true, true,
		uniqueID,
	)
	tabBarButton.Set("onclick", cb)
}

func (tools *Tools) addNewSpawnTabBarToLastPanelLevels(newTabBarID, firstAndOnlyPanelID string) {
	tools.tabberTabBarLastPanel[newTabBarID] = firstAndOnlyPanelID
}

func (tools *Tools) removeSpawnTabBarButtonFromLastPanelLevels(tabBarID, panelID string) {
	if tools.tabberTabBarLastPanel[tabBarID] == panelID {
		tools.tabberTabBarLastPanel[tabBarID] = ""
	}
}
