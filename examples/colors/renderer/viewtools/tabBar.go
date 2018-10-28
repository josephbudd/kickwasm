package viewtools

import (
	"strings"
	"syscall/js"
)

/*
	WARNING:

	DO NOT EDIT THIS FILE.

*/

// ForceTabButtonClick implements the behavior of a tab button being clicked by the user.
func (tools *Tools) ForceTabButtonClick(button js.Value) {
	tools.handleTabButtonOnClick(button)
}

func (tools *Tools) initializeTabBar() {
	notjs := tools.notjs
	tools.tabberLastPanelID = ""
	tools.tabberLastPanelLevels = make(map[string]string)

	cb := tools.notjs.RegisterCallBack(
		func(args []js.Value) {
			target := notjs.GetEventTarget(args[0])
			tools.handleTabButtonOnClick(target)
		},
	)
	for id := range tools.tabberLastPanelLevels {
		tabbar := notjs.GetElementByID(id)
		tools.setTabBarOnClicks(tabbar, cb)
	}
}

func (tools *Tools) setTabBarOnClicks(tabbar js.Value, cb js.Callback) {
	notjs := tools.notjs
	children := notjs.ChildrenSlice(tabbar)
	for _, ch := range children {
		if notjs.TagName(ch) == "BUTTON" {
			ch.Set("onclick", cb)
		}
	}
}

func (tools *Tools) handleTabButtonOnClick(button js.Value) {
	tools.setTabButtonFocus(button)
	nextpanelid := tools.notjs.ID(button) + "Panel"
	if nextpanelid != tools.tabberLastPanelID {
		// clear this level
		parts := strings.Split(nextpanelid, "-")
		nextpanellevel := parts[0]
		tools.IDHide(tools.tabberLastPanelLevels[nextpanellevel])
		// show the next panel
		tools.IDShow(nextpanelid)
		// remember next panel. it is now the last panel.
		tools.tabberLastPanelID = nextpanelid
		tools.tabberLastPanelLevels[nextpanellevel] = nextpanelid
	}
	tools.SizeApp()
}

func (tools *Tools) setTabButtonFocus(tabinfocus js.Value) {
	// focus the tab now in focus
	notjs := tools.notjs
	notjs.ClassListReplaceClass(tabinfocus, UnSelectedTabClassName, SelectedTabClassName)
	p := notjs.ParentNode(tabinfocus)
	children := notjs.ChildrenSlice(p)
	for _, ch := range children {
		if ch != tabinfocus && notjs.TagName(ch) == "BUTTON" && notjs.ClassListContains(ch, SelectedTabClassName) {
			notjs.ClassListReplaceClass(ch, SelectedTabClassName, UnSelectedTabClassName)
		}
	}
}
