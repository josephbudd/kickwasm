// +build js, wasm

package viewtools

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/framework/callback"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/api/event"
)

// Back simulates a click on the tall back button at the left of slider panels.
// TODO: fix this because wrong use of slice.
func Back() {
	if !HandleButtonClick() {
		return
	}
	l := len(backStack) - 1
	backdiv := backStack[l]
	backStack = backStack[:l]
	HideShow(here, backdiv)
}

func showSlider() {
	ElementShow(mainMasterviewHomeSlider)
}

func hideSlider() {
	ElementHide(mainMasterviewHomeSlider)
}

func initializeSlider() {
	divs := document.Call("getElementsByTagName", "DIV")
	ldivs := divs.Length()
	for i := 0; i < ldivs; i++ {
		div := divs.Index(i)
		classList := div.Get("classList")
		if classList.Call("contains", SliderButtonPadClassName).Bool() {
			children := div.Get("children")
			lch := children.Length()
			for j := 0; j < lch; j++ {
				ch := children.Index(j)
				tagname := ch.Get("tagName").String()
				if tagname == "BUTTON" {
					callback.AddEventHandler(handlePadButtonOnClick, ch, "click", false, 0)
				}
			}
		} else if div == mainMasterviewHomeButtonPad {
			children := div.Get("children")
			lch := children.Length()
			for j := 0; j < lch; j++ {
				ch := children.Index(j)
				tagname := ch.Get("tagName").String()
				if tagname == "BUTTON" {
					callback.AddEventHandler(handlePadButtonOnClick, ch, "click", false, 0)
				}
			}
		}
	}
	f := func(e event.Event) (nilReturn interface{}) {
		handleBack(e.JSTarget)
		return
	}
	callback.AddEventHandler(f, mainMasterviewHomeSliderBack, "click", false, 0)
}

func handlePadButtonOnClick(e event.Event) (nilReturn interface{}) {
	// get back div
	target := e.JSTarget
	backid := target.Call("getAttribute", BackIDAttribute).String()
	backdiv := getElementByID(document, backid)
	// get forward div
	targetid := target.Get("id").String()
	divs, found := buttonPanelsMap[targetid]
	if !found {
		alert.Invoke(fmt.Sprintf("slider.controller.handlePadButtonOnClick: id %q not found in buttonPanelsMap", targetid))
		return
	}
	for _, div := range divs {
		classList := div.Get("classList")
		if classList.Call("contains", ToBeSeenClassName).Bool() {
			here = div
			backStack = append(backStack, backdiv)
			HideShow(backdiv, div)
			return
		}
	}
	alert.Invoke(fmt.Sprintf("slider.controller.handlePadButtonOnClick: tobe-seen not found with button %q", target.Get("innerText")))
	return
}

// handleBack provides the behavior for the tall back button at the left of slider panels.
func handleBack(event js.Value) (nilReturn interface{}) {
	Back()
	return
}

// hereIsVisible returns if the current slider panel is actually seen by the user.
func hereIsVisible() (isVisible bool) {
	if here == js.Undefined() {
		return
	}
	p := here.Get("parentNode")
	isVisible = (p == mainMasterviewHomeSliderCollection)
	return
}
