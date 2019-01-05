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
	notJS := tools.notJS
	isSliderSub = notJS.ParentNode(target) == tools.tabsMasterviewHomeSliderCollection
	// tab sibling panels are in sliders but they are special.
	isTabSibling := notJS.ClassListContains(target, "slider-panel-inner-sibling")
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
		// target is the new here
		tools.here = target
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
	notJS := tools.notJS
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
	notJS := tools.notJS
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
	notJS := tools.notJS
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
	// Service1 Service1Button button.
	buttonid = "tabsMasterView-home-pad-Service1Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ContentButton-Service1Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ContentButton-Service1Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ContentButton-Service1Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ContentButton-Service1Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ContentButton-Service1Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ContentButton-Service1Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ContentButton-Service1Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ContentButton-Service1Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ColorsButton-Service1Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ColorsButton-Service1Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service1 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ColorsButton-Service1Level5ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ColorsButton-Service1Level5ButtonPanel-ContentButton-Service1Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service1Button-Service1Level1ButtonPanel-ColorsButton-Service1Level2ButtonPanel-ColorsButton-Service1Level3ButtonPanel-ColorsButton-Service1Level4ButtonPanel-ColorsButton-Service1Level5ButtonPanel-ContentButton-Service1Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 Service2Button button.
	buttonid = "tabsMasterView-home-pad-Service2Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ContentButton-Service2Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ContentButton-Service2Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ContentButton-Service2Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ContentButton-Service2Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ContentButton-Service2Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ContentButton-Service2Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ContentButton-Service2Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ContentButton-Service2Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ColorsButton-Service2Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ColorsButton-Service2Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service2 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ColorsButton-Service2Level5ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ColorsButton-Service2Level5ButtonPanel-ContentButton-Service2Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service2Button-Service2Level1ButtonPanel-ColorsButton-Service2Level2ButtonPanel-ColorsButton-Service2Level3ButtonPanel-ColorsButton-Service2Level4ButtonPanel-ColorsButton-Service2Level5ButtonPanel-ContentButton-Service2Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 Service3Button button.
	buttonid = "tabsMasterView-home-pad-Service3Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ContentButton-Service3Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ContentButton-Service3Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ContentButton-Service3Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ContentButton-Service3Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ContentButton-Service3Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ContentButton-Service3Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ContentButton-Service3Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ContentButton-Service3Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ColorsButton-Service3Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ColorsButton-Service3Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service3 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ColorsButton-Service3Level5ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ColorsButton-Service3Level5ButtonPanel-ContentButton-Service3Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service3Button-Service3Level1ButtonPanel-ColorsButton-Service3Level2ButtonPanel-ColorsButton-Service3Level3ButtonPanel-ColorsButton-Service3Level4ButtonPanel-ColorsButton-Service3Level5ButtonPanel-ContentButton-Service3Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 Service4Button button.
	buttonid = "tabsMasterView-home-pad-Service4Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ContentButton-Service4Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ContentButton-Service4Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ContentButton-Service4Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ContentButton-Service4Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ContentButton-Service4Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ContentButton-Service4Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ContentButton-Service4Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ContentButton-Service4Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton-Service4Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton-Service4Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service4 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton-Service4Level5ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton-Service4Level5ButtonPanel-ContentButton-Service4Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service4Button-Service4Level1ButtonPanel-ColorsButton-Service4Level2ButtonPanel-ColorsButton-Service4Level3ButtonPanel-ColorsButton-Service4Level4ButtonPanel-ColorsButton-Service4Level5ButtonPanel-ContentButton-Service4Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 Service5Button button.
	buttonid = "tabsMasterView-home-pad-Service5Button"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ContentButton-Service5Level1MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ContentButton-Service5Level1MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ContentButton-Service5Level2MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ContentButton-Service5Level2MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ContentButton-Service5Level3MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ContentButton-Service5Level3MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ContentButton-Service5Level4MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ContentButton-Service5Level4MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ColorsButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ColorsButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ColorsButton-Service5Level5ButtonPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ColorsButton-Service5Level5ButtonPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
	// Service5 ContentButton button.
	buttonid = "tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ColorsButton-Service5Level5ButtonPanel-ContentButton"
	tools.buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = tools.notJS.GetElementByID("tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ColorsButton-Service5Level5ButtonPanel-ContentButton-Service5Level5MarkupPanel")
	if panel == js.Undefined() {
		message := "viewtools.initializeGroups: Cant find #tabsMasterView-home-pad-Service5Button-Service5Level1ButtonPanel-ColorsButton-Service5Level2ButtonPanel-ColorsButton-Service5Level3ButtonPanel-ColorsButton-Service5Level4ButtonPanel-ColorsButton-Service5Level5ButtonPanel-ContentButton-Service5Level5MarkupPanel"
		tools.alert.Invoke(message)
		return
	}
	tools.buttonPanelsMap[buttonid] = append(tools.buttonPanelsMap[buttonid], panel)
}
