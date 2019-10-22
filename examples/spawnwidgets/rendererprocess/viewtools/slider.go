// +build js, wasm

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
	notJS := tools.NotJS
	divs := notJS.GetElementsByTagName("DIV")
	for _, div := range divs {
		if notJS.ClassListContains(div, SliderButtonPadClassName) {
			children := notJS.ChildrenSlice(div)
			for _, ch := range children {
				if notJS.TagName(ch) == "BUTTON" {
					tools.AddEventHandler(tools.handlePadButtonOnClick, ch, "click", false)
				}
			}
		} else if div == tools.tabsMasterviewHomeButtonPad {
			children := notJS.ChildrenSlice(div)
			for _, ch := range children {
				if notJS.TagName(ch) == "BUTTON" {
					tools.AddEventHandler(tools.handlePadButtonOnClick, ch, "click", false)
				}
			}
		}
	}
	f := func(e Event) (nilReturn interface{}) {
		tools.handleBack(e.Target)
		return
	}
	tools.AddEventHandler(f, tools.tabsMasterviewHomeSliderBack, "click", false)
}

func (tools *Tools) handlePadButtonOnClick(e Event) interface{} {
	// get back div
	notJS := tools.NotJS
	target := e.Target
	backid := target.Call("getAttribute", BackIDAttribute).String()
	backdiv := notJS.GetElementByID(backid)
	// get forward div
	targetid := notJS.ID(target)
	divs, found := tools.buttonPanelsMap[targetid]
	if !found {
		notJS.Alert(fmt.Sprintf("slider.controller.handlePadButtonOnClick: id %q not found in tools.buttonPanelsMap", targetid))
		return nil
	}
	for _, div := range divs {
		if notJS.ClassListContains(div, ToBeSeenClassName) {
			tools.here = div
			tools.backStack = append(tools.backStack, backdiv)
			tools.HideShow(backdiv, div)
			return nil
		}
	}
	notJS.Alert(fmt.Sprintf("slider.controller.handlePadButtonOnClick: tobe-seen not found with button %q", target.Get("innerText")))
	return nil
}

// handleBack provides the behavior for the tall back button at the left of slider panels.
func (tools *Tools) handleBack(event js.Value) interface{} {
	tools.Back()
	return nil
}

// hereIsVisible returns if the current slider panel is actually seen by the user.
func (tools *Tools) hereIsVisible() bool {
	if tools.here == js.Undefined() {
		return false
	}
	p := tools.NotJS.ParentNode(tools.here)
	return p == tools.tabsMasterviewHomeSliderCollection
}
