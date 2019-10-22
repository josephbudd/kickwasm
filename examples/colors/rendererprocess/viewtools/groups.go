// +build js, wasm

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
	// Action1Button Action1Button button.
	buttonid = "tabsMasterView-home-pad-Action1Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level1ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1Level1ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1Level1ContentButton-Action1Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1Level1ContentButton-Action1Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1ToLevel2ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level2ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton-Action1Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton-Action1Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1ToLevel3ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level3ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1Level3ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1Level3ContentButton-Action1Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1Level3ContentButton-Action1Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1ToLevel4ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level4ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1Level4ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1Level4ContentButton-Action1Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1Level4ContentButton-Action1Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1ToLevel5ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level5ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel-Action1Level5ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel-Action1Level5ContentButton-Action1Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel-Action1Level5ContentButton-Action1Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Button button.
	buttonid = "tabsMasterView-home-pad-Action2Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level1ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2Level1ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2Level1ContentButton-Action2Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2Level1ContentButton-Action2Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2ToLevel2ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level2ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2Level2ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2Level2ContentButton-Action2Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2Level2ContentButton-Action2Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2ToLevel3ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level3ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2Level3ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2Level3ContentButton-Action2Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2Level3ContentButton-Action2Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2ToLevel4ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level4ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2Level4ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2Level4ContentButton-Action2Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2Level4ContentButton-Action2Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2ToLevel5ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level5ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel-Action2Level5ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel-Action2Level5ContentButton-Action2Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel-Action2Level5ContentButton-Action2Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Button button.
	buttonid = "tabsMasterView-home-pad-Action3Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level1ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3Level1ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3Level1ContentButton-Action3Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3Level1ContentButton-Action3Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3ToLevel2ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level2ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3Level2ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3Level2ContentButton-Action3Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3Level2ContentButton-Action3Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3ToLevel3ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level3ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton-Action3Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton-Action3Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3ToLevel4ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level4ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3Level4ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3Level4ContentButton-Action3Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3Level4ContentButton-Action3Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3ToLevel5ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level5ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel-Action3Level5ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel-Action3Level5ContentButton-Action3Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel-Action3Level5ContentButton-Action3Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Button button.
	buttonid = "tabsMasterView-home-pad-Action4Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level1ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4Level1ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4Level1ContentButton-Action4Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4Level1ContentButton-Action4Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4ToLevel2ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level2ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4Level2ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4Level2ContentButton-Action4Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4Level2ContentButton-Action4Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4ToLevel3ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level3ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4Level3ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4Level3ContentButton-Action4Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4Level3ContentButton-Action4Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4ToLevel4ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level4ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4Level4ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4Level4ContentButton-Action4Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4Level4ContentButton-Action4Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4ToLevel5ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level5ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton-Action4Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton-Action4Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Button button.
	buttonid = "tabsMasterView-home-pad-Action5Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level1ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5Level1ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5Level1ContentButton-Action5Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5Level1ContentButton-Action5Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5ToLevel2ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level2ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5Level2ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5Level2ContentButton-Action5Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5Level2ContentButton-Action5Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5ToLevel3ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level3ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5Level3ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5Level3ContentButton-Action5Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5Level3ContentButton-Action5Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5ToLevel4ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level4ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5Level4ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5Level4ContentButton-Action5Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5Level4ContentButton-Action5Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5ToLevel5ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level5ContentButton button.
	buttonid = "tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel-Action5Level5ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.NotJS.GetElementByID("tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel-Action5Level5ContentButton-Action5Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel-Action5Level5ContentButton-Action5Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
}
