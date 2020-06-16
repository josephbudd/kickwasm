// +build js, wasm

package viewtools

import (
	"fmt"
	"path/filepath"
	"strings"
	"syscall/js"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/api/event"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/framework/callback"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/framework/spawnpack"
)

// FixSpawnID fixes a spawn html template element id using the spawns unique id.
// It replaces "{{.SpawnID}}" with the spawn's unique id.
func FixSpawnID(id string, uniqueID uint64) (fixedID string) {
	fixedID = strings.ReplaceAll(id, spawnIDReplacePattern, fmt.Sprint(uniqueID))
	return
}
 
// SpawnTab adds an html tab button and panels to the html document.
// It returns their uniqueID, the tab button id, panels names mapped to their ids.
func SpawnTab(tabBarID, tabName, tabLabel, tabPanelHeadingText string, userContentPanelPaths []string) (tabButton, tabPanelHeader js.Value, uniqueID uint64, panelNameID map[string]string, err error) {
	var classList js.Value
	uniqueID = callback.NewEventHandlerID()
	// Find the tab bar.
	var tabBar js.Value
	if tabBar = getElementByID(document, tabBarID); tabBar.IsNull() {
		err = fmt.Errorf("Unable to find tab bar #" + tabBarID)
		return
	}
	// get the number of tabs in this tab bar.
	nTabButtons := tabBar.Get("children").Length()
	// Create the button and add it to the tab bar.
	tabButton = document.Call("createElement", "BUTTON")
	tabButtonID := BuildSpawnTabButtonID(tabBarID, tabName, uniqueID)
	tabButton.Set("id", tabButtonID)
	classList = tabButton.Get("classList")
	classList.Call("add", TabClassName)

	label := document.Call("createTextNode", tabLabel)
	tabButton.Call("appendChild", label)
	tabBar.Call("appendChild", tabButton)
	var tabPanelVisiblilityClass string
	if nTabButtons == 0 {
		// if this is the first tab then bring it up front.
		tabButton.Call("focus")
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
	underTabBarDivID := buildSpawnUnderTabBarID(tabBarID)
	var underTabBarDiv js.Value
	if underTabBarDiv = getElementByID(document, underTabBarDivID); underTabBarDiv.IsNull() {
		err = fmt.Errorf("Unable to find under tab bar #" + underTabBarDivID)
		return
	}
	// Create the tab panel
	tabPanel := document.Call("createElement", "DIV")
	tabPanelID := buildSpawnTabButtonPanelID(tabButtonID)
	tabPanel.Set("id", tabPanelID)
	classList = tabPanel.Get("classList")
	classList.Call("add", TabPanelClassName)
	classList.Call("add", PanelWithHeadingClassName)
	classList.Call("add", tabPanelVisiblilityClass)
	tabPanelHeader = document.Call("createElement", "H3")
	tabPanelHeadingID := buildSpawnTabButtonPanelHeadingID(tabButtonID)
	tabPanelHeader.Set("id", tabPanelHeadingID)
	classList = tabPanelHeader.Get("classList")
	classList.Call("add", PanelHeadingClassName)
	heading := document.Call("createTextNode", tabPanelHeadingText)
	tabPanelHeader.Call("appendChild", heading)
	tabPanel.Call("appendChild", tabPanelHeader)
	// Inside the tab panel is the inner panel
	//   which is the wrapper of group of user content (markup) panels.
	// Create the group
	group := document.Call("createElement", "DIV")
	innerPanelID := buildSpawnTabButtonInnerPanelID(tabButtonID)
	group.Set("id", innerPanelID)
	classList = group.Get("classList")
	classList.Call("add", TabPanelGroupClassName)
	tabPanel.Call("appendChild", group)
	// Create each user content & it's markup panel inside the group panel
	l := len(userContentPanelPaths)
	panelNameID = make(map[string]string, l)
	uniqueIDString := fmt.Sprint(uniqueID)
	userContentPanels := make([]js.Value, len(userContentPanelPaths))
	for i, path := range userContentPanelPaths {
		var markupbb []byte
		var found bool
		if markupbb, found = spawnpack.Contents(path); !found {
			err = fmt.Errorf("Unable to find the spawn template at " + path)
			return
		}
		// get the panel name.
		base := filepath.Base(path)
		l := len(base) - len(filepath.Ext(base))
		panelName := base[:l]
		// user content panel wraps the markup panel
		userContentPanel := document.Call("createElement", "DIV")
		userContentPanelID := buildSpawnTabButtonInnerMarkupPanelID(tabButtonID, panelName)
		userContentPanels[i] = userContentPanel
		panelNameID[panelName] = userContentPanelID
		userContentPanel.Set("id", userContentPanelID)
		classList = userContentPanel.Get("classList")
		classList.Call("add", UserContentClassName)
		classList.Call("add", SliderPanelInnerSiblingClassName)
		if i == 0 {
			// The first markup panel is the defauilt and is visible.
			classList = userContentPanel.Get("classList")
			classList.Call("add", SeenClassName)
		} else {
			// The other markup panels are hidden.
			classList = userContentPanel.Get("classList")
			classList.Call("add", UnSeenClassName)
		}
		var hvscroll bool
		if hvscroll, found = panelNameHVScroll[panelName]; !found {
			err = fmt.Errorf("Unable to find panel name %q in panelNameHVScroll", panelName)
			return
		}
		if hvscroll {
			classList = userContentPanel.Get("classList")
			classList.Call("add", HVScrollClassName)
		} else {
			classList = userContentPanel.Get("classList")
			classList.Call("add", VScrollClassName)
		}
		group.Call("appendChild", userContentPanel)

		// markup panel inside the user content panel.
		markupPanel := document.Call("createElement", "DIV")
		classList = markupPanel.Get("classList")
		classList.Call("add", SeenClassName)
		markup := strings.ReplaceAll(string(markupbb), spawnIDReplacePattern, uniqueIDString)
		markupPanel.Set("innerHTML", markup)
		userContentPanel.Call("appendChild", markupPanel)
	}
	// now that the tab panel is fully constructed
	//   add it to the under tab bar div.
	underTabBarDiv.Call("appendChild", tabPanel)

	// add the tab bar button panel group
	buttonPanelsMap[tabButtonID] = userContentPanels
	// tab button onclick handler
	setTabBarSpawnButtonOnClick(tabButton, uniqueID)

	// Re format the tabs in the tab bar.
	reformatTabBarTabs(tabBar)
	return
}

// UnSpawnTab removes a tab button and panels from the html document.
// Returns the error.
func UnSpawnTab(tabButton js.Value) (err error ) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("UnSpawnTab: %w", err)
		}
	}()

	tabButtonID := tabButton.Get("id").String()

	// Step 1:
	// If there is a tab to the left if there is one
	//   then force click that tab bringing it up front with it's panel.
	tabBar := tabButton.Get("parentNode")
	siblings := tabBar.Get("children")
	l := siblings.Length()
	var i int
	for i = 0; i < l; i++ {
		sibling := siblings.Index(i)
		if sibling.Equal(tabButton) {
			break
		}
	}
	if i > 0 {
		// The tab being closed is not the first tab.
		// So click on the tab before this tab.
		i--
	} else {
		// The tab being closed is the first tab.
		// So click on the second tab.
		i++
	}
	ForceTabButtonClick(siblings.Index(i))

	// Step 2: The tab button.
	// Remove this tab from the tab bar in the DOM.
	// cleanup data.
	tabBarID := tabBar.Get("id").String()
	removeSpawnTabBarButtonFromLastPanelLevels(tabBarID, tabButtonID)
	// remove the button the DOM.
	tabBar.Call("removeChild", tabButton)

	// Step 3: The tab panels.
	// Remove the tab panels from the DOM.
	panels := buttonPanelsMap[tabButtonID]
	underTabBarDiv := panels[0].Get("parentNode")
	for _, p := range panels {
		underTabBarDiv.Call("removeChild", p)
	}
	// cleanup data.
	// remove the tab bar button panel group.
	delete(buttonPanelsMap, tabButtonID)

	// Reformat the tabs in the tab bar.
	reformatTabBarTabs(tabBar)

	return
}

// BuildSpawnTabButtonID forms an id for a spawned tab bar button.
func BuildSpawnTabButtonID(tabBarID, tabName string, uniqueID uint64) (id string) {
	id = fmt.Sprintf("%s-%s(%d)", tabBarID, tabName, uniqueID)
	return
}

// BuildSpawnTabButtonMarkupPanelID forms an id for a spawned tab bar button's markup panel.
// This exists mainly for spawn panel groups.
func BuildSpawnTabButtonMarkupPanelID(tabBarID, tabName, panelName string, uniqueID uint64) (id string) {
	buttonID := BuildSpawnTabButtonID(tabBarID, tabName, uniqueID)
	id = buildSpawnTabButtonInnerMarkupPanelID(buttonID, panelName)
	return
}

// buildSpawnUnderTabBarID forms an id for a spawned tab bar button.
func buildSpawnUnderTabBarID(tabBarID string) (id string) {
	id = tabBarID + "-under-tab-bar"
	return
}

// buildSpawnTabButtonPanelID forms an id for a spawned tab bar button.
func buildSpawnTabButtonPanelID(buttonID string) (id string) {
	id = buttonID + "Panel"
	return
}

// buildSpawnTabButtonPanelHeadingID forms an id for a spawned tab bar button button's inner panel.
func buildSpawnTabButtonPanelHeadingID(buttonID string) (id string) {
	id = buildSpawnTabButtonPanelID(buttonID) + "-panel-heading"
	return
}

// buildSpawnTabButtonInnerPanelID forms an id for a spawned tab bar button button's inner panel.
func buildSpawnTabButtonInnerPanelID(buttonID string) (id string) {
	id = buildSpawnTabButtonPanelID(buttonID) + "-inner"
	return
}

// buildSpawnTabButtonInnerMarkupPanelID forms an id for a spawned tab bar button's markup panel.
func buildSpawnTabButtonInnerMarkupPanelID(buttonID, panelName string) (id string) {
	id = buildSpawnTabButtonInnerPanelID(buttonID) + "-" + panelName
	return
}

func addTabBarButton(tabBar, newButton js.Value) {
	tabBar.Call("appendChild", newButton)
}

func setTabBarSpawnButtonOnClick(tabBarButton js.Value, uniqueID uint64) {
	f := func(e event.Event) (nilReturn interface{}) {
		if e.JSTarget.Get("tagName").String() != "BUTTON" {
			// The user slid the mouse over the button and raised the mouse button.
			// Button is not known so ignore.
			return
		}
		handleTabButtonOnClick(e.JSTarget)
		return
	}
	callback.AddEventHandler(f, tabBarButton, "click", false, uniqueID)
}

func removeSpawnTabBarButtonFromLastPanelLevels(tabBarID, panelID string) {
	if tabberTabBarLastPanel[tabBarID] == panelID {
		tabberTabBarLastPanel[tabBarID] = ""
	}
}
