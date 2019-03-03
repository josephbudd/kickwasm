package viewtools

import (
	"fmt"
	"syscall/js"
)

// Back simulates a click on the tall back button at the left of slider panels.
func (tools *Tools) Back() {
	if !tools.HandleButtonClick() {
		return
	}
	l := len(tools.backStack) - 1
	backdiv := tools.backStack[l]
	tools.backStack = tools.backStack[:l]
	tools.HideShow(tools.here, backdiv)
}

func (tools *Tools) showSlider() {
	tools.ElementShow(tools.tabsMasterviewHomeSlider)
}

func (tools *Tools) hideSlider() {
	tools.ElementHide(tools.tabsMasterviewHomeSlider)
}

func (tools *Tools) initializeSlider() {
	notJS := tools.notJS
	buttoncb := notJS.RegisterCallBack(tools.handlePadButtonOnClick)
	divs := notJS.GetElementsByTagName("DIV")
	for _, div := range divs {
		if notJS.ClassListContains(div, SliderButtonPadClassName) {
			children := notJS.ChildrenSlice(div)
			for _, ch := range children {
				if notJS.TagName(ch) == "BUTTON" {
					notJS.SetOnClick(ch, buttoncb)
				}
			}
		} else if div == tools.tabsMasterviewHomeButtonPad {
			children := notJS.ChildrenSlice(div)
			for _, ch := range children {
				if notJS.TagName(ch) == "BUTTON" {
					notJS.SetOnClick(ch, buttoncb)
				}
			}
		}
	}
	backcb := notJS.RegisterCallBack(tools.handleBack)
	notJS.SetOnClick(tools.tabsMasterviewHomeSliderBack, backcb)
}

func (tools *Tools) handlePadButtonOnClick(args []js.Value) {
	// get back div
	notJS := tools.notJS
	target := args[0].Get("target")
	backid := target.Call("getAttribute", BackIDAttribute).String()
	backdiv := notJS.GetElementByID(backid)
	// get forward div
	targetid := notJS.ID(target)
	divs, found := tools.buttonPanelsMap[targetid]
	if !found {
		notJS.Alert(fmt.Sprintf("slider.controler.handlePadButtonOnClick: id %q not found in tools.buttonPanelsMap", targetid))
		return
	}
	for _, div := range divs {
		if notJS.ClassListContains(div, ToBeSeenClassName) {
			tools.here = div
			tools.backStack = append(tools.backStack, backdiv)
			tools.HideShow(backdiv, div)
			return
		}
	}
	notJS.Alert(fmt.Sprintf("slider.controler.handlePadButtonOnClick: tobe-seen not found with button %q", target.Get("innerText")))
}

// handleBack provides the behavior for the tall back button at the left of slider panels.
func (tools *Tools) handleBack(args []js.Value) {
	tools.Back()
}

// hereIsVisible returns if the current slider panel is actually seen by the user.
func (tools *Tools) hereIsVisible() bool {
	if tools.here == js.Undefined() {
		return false
	}
	p := tools.notJS.ParentNode(tools.here)
	return p == tools.tabsMasterviewHomeSliderCollection
}
