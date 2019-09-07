package viewtools

import (
	"syscall/js"
)

// ShowPanelInButtonGroup shows a panel in a button pad button group and hides the other panels in the group.
func (tools *Tools) ShowPanelInButtonGroup(panel js.Value, force bool) {
	if force && tools.hereIsVisible() {
		// show this panel.
		// the app is resized by HideShow.
		tools.HideShow(tools.here, panel)
	} else {
		// not forcing so just set the panel to be visible when the user makes it visible.
		_, isVisible := tools.ShowInGroup(panel, ToBeSeenClassName, ToBeUnSeenClassName)
		if isVisible {
			// this panel is visible anyway so resize the app.
			tools.SizeApp()
		}
	}
}

// ShowPanelInTabGroup shows a panel in a tab button group and hides the other panels in the group.
func (tools *Tools) ShowPanelInTabGroup(panel js.Value) {
	_, isVisible := tools.ShowInGroup(panel, SeenClassName, UnSeenClassName)
	if isVisible {
		// this tab panel is visible anyway so resize the app.
		tools.SizeApp()
	}
}

// ShowInGroup only works with panel groups decended from the slider collection.  ( panels shown with the back button on the left side. )
// It shows one panel while hiding the other panels in a panel group.
// It does so by adding and removing classes to panel classLists.
// Param target is the panel to be shown.
// Param showClass is the class name for showing target.
// Param hideClass is the class name for hiding the other panels in target's group.
// Returns 2 params
// 1. if param target has an ancestor which is the slider collections. ( panels shown with the back button on the left side. )
// 2. if the param target becomes visible.
func (tools *Tools) ShowInGroup(target js.Value, showClass, hideClass string) (isSliderSub, isVisible bool) {
	notJS := tools.NotJS
	isSliderSub = notJS.ParentNode(target) == tools.tabsMasterviewHomeSliderCollection
	// tab sibling panels are in sliders but they are special.
	isTabSibling := notJS.ClassListContains(target, SliderPanelInnerSiblingClassName)
	if !(isSliderSub || isTabSibling) {
		// not in the slider collection
		isSliderSub = (isSliderSub || isTabSibling)
		return
	}
	targetInGroup := false
	var divs []js.Value
	for _, divs = range tools.buttonPanelsMap {
		for _, div := range divs {
			if target == div {
				// target is in this group
				targetInGroup = true
				break
			}
		}
		if targetInGroup {
			// target is in this group
			break
		}
	}
	if !targetInGroup {
		// target not in group so not a slider sub.
		isSliderSub = false
		return
	}
	// yes target is a slider div
	tools.setInGroup(divs, target, showClass, hideClass)
	// check for visibility
	for _, div := range divs {
		if notJS.ClassListContains(div, SeenClassName) {
			isVisible = true
			break
		}
	}
	// only really visible if slider is visible
	isVisible = notJS.ClassListContains(tools.tabsMasterviewHomeSlider, SeenClassName) && isVisible
	if isVisible {
		// set tools.here
		if isTabSibling {
			// tools.here is never a tab panel sibling.
			// tools.here must be a slider panel.
			// Find the correct ancestor.
			sliderPanel := notJS.ParentNode(target)
			for {
				if notJS.ClassListContains(sliderPanel, SliderPanelClassName) {
					break
				}
				sliderPanel = notJS.ParentNode(sliderPanel)
			}
			tools.here = sliderPanel
		} else {
			// target is the new here.
			tools.here = target
		}
		// here is now this slider sub panel.
		// here is never a tab panel sibling.
		if showClass == ToBeSeenClassName {
			// set target for class seen
			// set non targets for class unseen
			tools.setInGroup(divs, target, SeenClassName, UnSeenClassName)
		}
	}
	isSliderSub = isSliderSub || isTabSibling
	return
}

//HideShow hides one panel and shows another panel.
// both panels must have the parentNode == SliderPresenter.sliderCollection.
func (tools *Tools) HideShow(hideDiv, showDiv js.Value) {
	// hide the hide div
	notJS := tools.NotJS
	isSliderH := tools.hideInGroup(hideDiv, SeenClassName, UnSeenClassName)
	// show the show div
	isSliderS, isVisibleS := tools.ShowInGroup(showDiv, SeenClassName, UnSeenClassName)
	if isSliderS {
		// reset the back button's color class.
		backColorLevel := showDiv.Call("getAttribute", "backColorLevel").String()
		firstClass := notJS.ClassListGetClassAt(tools.tabsMasterviewHomeSliderBack, 0)
		notJS.ClassListReplaceClass(tools.tabsMasterviewHomeSliderBack, firstClass, backColorLevel)
	}
	if isSliderH && isSliderS {
		// the slider was visible for the hideDiv and so it still is for the showDiv
		tools.SizeApp()
		return
	}
	// hideDiv and showDiv are not both sliders
	if !isVisibleS {
		// showDiv, the div to show is not visible
		if isSliderH {
			// hideDiv is in the slider collection
			tools.hideSlider()
		} else {
			// hideDiv is not in the slider collection, its a master div or home or some little thing in a panel
			tools.ElementHide(hideDiv)
		}
		if isSliderS {
			// showDiv is in the slider collection
			tools.showSlider()
		} else {
			// showDiv is not in the slider collection
			tools.ElementShow(showDiv)
		}
	}
	tools.SizeApp()
}

// toBeShownInGroup returns if the class is set to become visible when it's panel group is made visible.
// Returns 2 params
// 1. if param target has an ancestor which is the slider collections. ( panels shown with the back button on the left side. )
// 2. if the param target becomes visible.
func (tools *Tools) toBeShownInGroup(target js.Value) (bool, bool) {
	return tools.ShowInGroup(target, ToBeSeenClassName, ToBeUnSeenClassName)
}

// Returns is the target is a slider sub panel, a child of the slider collection div.
func (tools *Tools) toBeHiddenInGroup(target js.Value) bool {
	return tools.hideInGroup(target, ToBeSeenClassName, ToBeUnSeenClassName)
}

// setInGroup works with a group of panels from tools.buttonPanelsMap.
// It sets target's to setClass and removes unSetClass.
// It sets the other panel's to unSetClass and removes setClass.
func (tools *Tools) setInGroup(group []js.Value, target js.Value, setClass, unSetClass string) {
	notJS := tools.NotJS
	for _, panel := range group {
		if panel != target {
			notJS.ClassListReplaceClass(panel, setClass, unSetClass)
		}
	}
	notJS.ClassListReplaceClass(target, unSetClass, setClass)
}

// hideInGroup hides target in a group.
// Returns is the target is a slider sub panel, a child of the slider collection div.
func (tools *Tools) hideInGroup(target js.Value, showClass, hideClass string) (isSliderSub bool) {
	notJS := tools.NotJS
	parentNode := notJS.ParentNode(target)
	isSliderSub = parentNode == tools.tabsMasterviewHomeSliderCollection
	if !isSliderSub {
		// not in the slider collection.
		return
	}
	notJS.ClassListReplaceClass(target, showClass, hideClass)
	return
}

func (tools *Tools) initializeGroups() {
	// build the buttonPanelsMap
	var buttonid string
	var panel js.Value
	// Tabs TabsButton button.
	buttonid = "tabsMasterView-home-pad-TabsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-TabsButton-TabsButtonTabBarPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-TabsButton-TabsButtonTabBarPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Tabs FirstTab button.
	buttonid = "tabsMasterView_home_pad_TabsButton_TabsButtonTabBarPanel_tab_bar-FirstTab"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView_home_pad_TabsButton_TabsButtonTabBarPanel_tab_bar-FirstTabPanel-inner-CreatePanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView_home_pad_TabsButton_TabsButtonTabBarPanel_tab_bar-FirstTabPanel-inner-CreatePanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
}
