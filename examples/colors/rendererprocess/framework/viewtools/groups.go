// +build js, wasm

package viewtools

import (
	"syscall/js"
)

// ShowPanelInButtonGroup shows a panel in a button pad button group and hides the other panels in the group.
func ShowPanelInButtonGroup(panel js.Value, force bool) {
	if force && hereIsVisible() {
		// show this panel.
		// the app is resized by HideShow.
		HideShow(here, panel)
	} else {
		// not forcing so just set the panel to be visible when the user makes it visible.
		_, isVisible := ShowInGroup(panel, ToBeSeenClassName, ToBeUnSeenClassName)
		if isVisible {
			// this panel is visible anyway so resize the app.
			SizeApp()
		}
	}
}

// ShowPanelInTabGroup shows a panel in a tab button group and hides the other panels in the group.
func ShowPanelInTabGroup(panel js.Value) {
	_, isVisible := ShowInGroup(panel, SeenClassName, UnSeenClassName)
	if isVisible {
		// this tab panel is visible anyway so resize the app.
		SizeApp()
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
func ShowInGroup(target js.Value, showClass, hideClass string) (isSliderSub, isVisible bool) {
	targetParent := target.Get("parentNode")
	isSliderSub = targetParent.Equal(mainMasterviewHomeSliderCollection)
	// tab sibling panels are in sliders but they are special.
	grandParent := targetParent.Get("parentNode")
	classList := grandParent.Get("classList")
	isTabSibling := classList.Call("contains", TabPanelClassName).Bool()
	if !isTabSibling {
		classList = target.Get("classList")
		isTabSibling = classList.Call("contains", SliderPanelInnerSiblingClassName).Bool()
	}
	if !(isSliderSub || isTabSibling) {
		// not in the slider collection
		isSliderSub = (isSliderSub || isTabSibling)
		return
	}
	targetInGroup := false
	var divs []js.Value
	for _, divs = range buttonPanelsMap {
		for _, div := range divs {
			if target.Equal(div) {
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
	setInGroup(divs, target, showClass, hideClass)
	// check for visibility
	for _, div := range divs {
		classList := div.Get("classList")
		if isVisible = !classList.Call("contains", UnSeenClassName).Bool(); isVisible {
			// only really visible if slider is visible
			classList = mainMasterviewHomeSlider.Get("classList")
			isVisible = !classList.Call("contains", UnSeenClassName).Bool()
			break
		}
	}
	if isVisible {
		// set here
		if isTabSibling {
			// here is never a tab panel sibling.
			// here must be a slider panel.
			// Find the correct ancestor.
			sliderPanel := targetParent
			for {
				classList = sliderPanel.Get("classList")
				if classList.Call("contains", SliderPanelClassName).Bool() {
					break
				}
				sliderPanel = sliderPanel.Get("parentNode")
			}
			here = sliderPanel
		} else {
			// target is the new here.
			here = target
		}
		// here is now this slider sub panel.
		// here is never a tab panel sibling.
		if showClass == ToBeSeenClassName {
			// set target for class seen
			// set non targets for class unseen
			setInGroup(divs, target, SeenClassName, UnSeenClassName)
		}
	}
	isSliderSub = isSliderSub || isTabSibling
	return
}

//HideShow hides one panel and shows another panel.
// both panels must have the parentNode == SliderPresenter.sliderCollection.
func HideShow(hideDiv, showDiv js.Value) {
	// hide the hide div
	isSliderH := hideInGroup(hideDiv, SeenClassName, UnSeenClassName)
	// show the show div
	isSliderS, isVisibleS := ShowInGroup(showDiv, SeenClassName, UnSeenClassName)
	if isSliderS {
		// reset the back button's color class.
		backColorLevel := showDiv.Call("getAttribute", "backColorLevel").String()
		classList := mainMasterviewHomeSliderBack.Get("classList")
		firstClass := classList.Call("item", 0).String()
		classList.Call("replace", firstClass, backColorLevel)
	}
	if isSliderH && isSliderS {
		// the slider was visible for the hideDiv and so it still is for the showDiv
		SizeApp()
		return
	}
	// hideDiv and showDiv are not both sliders
	if !isVisibleS {
		// showDiv, the div to show is not visible
		if isSliderH {
			// hideDiv is in the slider collection
			hideSlider()
		} else {
			// hideDiv is not in the slider collection, its a master div or home or some little thing in a panel
			ElementHide(hideDiv)
		}
		if isSliderS {
			// showDiv is in the slider collection
			showSlider()
		} else {
			// showDiv is not in the slider collection
			ElementShow(showDiv)
		}
	}
	SizeApp()
}

// setInGroup works with a group of panels from buttonPanelsMap.
// It sets target's to setClass and removes unSetClass.
// It sets the other panel's to unSetClass and removes setClass.
func setInGroup(group []js.Value, target js.Value, setClass, unSetClass string) {
	var classList js.Value
	for _, panel := range group {
		if !target.Equal(panel) {
			classList = panel.Get("classList")
			if !classList.Call("replace", setClass, unSetClass).Bool() {
				classList.Call("add", unSetClass)
			}
		}
	}
	classList = target.Get("classList")
	if !classList.Call("replace", unSetClass, setClass).Bool() {
		classList.Call("add", setClass)
	}
}

// hideInGroup hides target in a group.
// Returns is the target is a slider sub panel, a child of the slider collection div.
func hideInGroup(target js.Value, showClass, hideClass string) (isSliderSub bool) {
	parentNode := target.Get("parentNode")
	isSliderSub = parentNode.Equal(mainMasterviewHomeSliderCollection)
	if !isSliderSub {
		// not in the slider collection.
		return
	}
	classList := target.Get("classList")
	classList.Call("replace", showClass, hideClass)
	return
}

func initializeGroups() {
	// build the buttonPanelsMap
	var buttonid string
	var panel js.Value
	// Action1Button Action1Button button.
	buttonid = "mainMasterView-home-pad-Action1Button"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level1ContentButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1Level1ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1Level1ContentButton-Action1Level1MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1Level1ContentButton-Action1Level1MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1ToLevel2ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level2ContentButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton-Action1Level2MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1Level2ContentButton-Action1Level2MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1ToLevel3ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level3ContentButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1Level3ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1Level3ContentButton-Action1Level3MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1Level3ContentButton-Action1Level3MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1ToLevel4ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level4ContentButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1Level4ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1Level4ContentButton-Action1Level4MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1Level4ContentButton-Action1Level4MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1ToLevel5ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action1Button Action1Level5ContentButton button.
	buttonid = "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel-Action1Level5ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel-Action1Level5ContentButton-Action1Level5MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action1Button-Action1Level1ButtonPanel-Action1ToLevel2ColorsButton-Action1Level2ButtonPanel-Action1ToLevel3ColorsButton-Action1Level3ButtonPanel-Action1ToLevel4ColorsButton-Action1Level4ButtonPanel-Action1ToLevel5ColorsButton-Action1Level5ButtonPanel-Action1Level5ContentButton-Action1Level5MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Button button.
	buttonid = "mainMasterView-home-pad-Action2Button"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level1ContentButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2Level1ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2Level1ContentButton-Action2Level1MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2Level1ContentButton-Action2Level1MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2ToLevel2ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level2ContentButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2Level2ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2Level2ContentButton-Action2Level2MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2Level2ContentButton-Action2Level2MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2ToLevel3ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level3ContentButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2Level3ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2Level3ContentButton-Action2Level3MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2Level3ContentButton-Action2Level3MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2ToLevel4ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level4ContentButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2Level4ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2Level4ContentButton-Action2Level4MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2Level4ContentButton-Action2Level4MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2ToLevel5ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action2Button Action2Level5ContentButton button.
	buttonid = "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel-Action2Level5ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel-Action2Level5ContentButton-Action2Level5MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action2Button-Action2Level1ButtonPanel-Action2ToLevel2ColorsButton-Action2Level2ButtonPanel-Action2ToLevel3ColorsButton-Action2Level3ButtonPanel-Action2ToLevel4ColorsButton-Action2Level4ButtonPanel-Action2ToLevel5ColorsButton-Action2Level5ButtonPanel-Action2Level5ContentButton-Action2Level5MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Button button.
	buttonid = "mainMasterView-home-pad-Action3Button"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level1ContentButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3Level1ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3Level1ContentButton-Action3Level1MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3Level1ContentButton-Action3Level1MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3ToLevel2ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level2ContentButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3Level2ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3Level2ContentButton-Action3Level2MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3Level2ContentButton-Action3Level2MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3ToLevel3ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level3ContentButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton-Action3Level3MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3Level3ContentButton-Action3Level3MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3ToLevel4ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level4ContentButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3Level4ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3Level4ContentButton-Action3Level4MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3Level4ContentButton-Action3Level4MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3ToLevel5ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action3Button Action3Level5ContentButton button.
	buttonid = "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel-Action3Level5ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel-Action3Level5ContentButton-Action3Level5MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action3Button-Action3Level1ButtonPanel-Action3ToLevel2ColorsButton-Action3Level2ButtonPanel-Action3ToLevel3ColorsButton-Action3Level3ButtonPanel-Action3ToLevel4ColorsButton-Action3Level4ButtonPanel-Action3ToLevel5ColorsButton-Action3Level5ButtonPanel-Action3Level5ContentButton-Action3Level5MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Button button.
	buttonid = "mainMasterView-home-pad-Action4Button"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level1ContentButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4Level1ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4Level1ContentButton-Action4Level1MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4Level1ContentButton-Action4Level1MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4ToLevel2ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level2ContentButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4Level2ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4Level2ContentButton-Action4Level2MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4Level2ContentButton-Action4Level2MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4ToLevel3ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level3ContentButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4Level3ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4Level3ContentButton-Action4Level3MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4Level3ContentButton-Action4Level3MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4ToLevel4ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level4ContentButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4Level4ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4Level4ContentButton-Action4Level4MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4Level4ContentButton-Action4Level4MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4ToLevel5ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action4Button Action4Level5ContentButton button.
	buttonid = "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton-Action4Level5MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action4Button-Action4Level1ButtonPanel-Action4ToLevel2ColorsButton-Action4Level2ButtonPanel-Action4ToLevel3ColorsButton-Action4Level3ButtonPanel-Action4ToLevel4ColorsButton-Action4Level4ButtonPanel-Action4ToLevel5ColorsButton-Action4Level5ButtonPanel-Action4Level5ContentButton-Action4Level5MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Button button.
	buttonid = "mainMasterView-home-pad-Action5Button"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level1ContentButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5Level1ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5Level1ContentButton-Action5Level1MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5Level1ContentButton-Action5Level1MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5ToLevel2ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level2ContentButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5Level2ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5Level2ContentButton-Action5Level2MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5Level2ContentButton-Action5Level2MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5ToLevel3ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level3ContentButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5Level3ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5Level3ContentButton-Action5Level3MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5Level3ContentButton-Action5Level3MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5ToLevel4ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level4ContentButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5Level4ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5Level4ContentButton-Action5Level4MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5Level4ContentButton-Action5Level4MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5ToLevel5ColorsButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
	// Action5Button Action5Level5ContentButton button.
	buttonid = "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel-Action5Level5ContentButton"
	buttonPanelsMap[buttonid] = make([]js.Value, 0, 5)
	panel = getElementByID(document, "mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel-Action5Level5ContentButton-Action5Level5MarkupPanel")
	if panel.IsUndefined() {
		message := "viewtools.initializeGroups: Cant find #mainMasterView-home-pad-Action5Button-Action5Level1ButtonPanel-Action5ToLevel2ColorsButton-Action5Level2ButtonPanel-Action5ToLevel3ColorsButton-Action5Level3ButtonPanel-Action5ToLevel4ColorsButton-Action5Level4ButtonPanel-Action5ToLevel5ColorsButton-Action5Level5ButtonPanel-Action5Level5ContentButton-Action5Level5MarkupPanel"
		alert.Invoke(message)
		return
	}
	buttonPanelsMap[buttonid] = append(buttonPanelsMap[buttonid], panel)
}
