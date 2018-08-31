package templates

// ViewToolsSlider is the renderer/viewtools/slider.go file.
const ViewToolsSlider = `package viewtools

import (
	"fmt"
	"syscall/js"
)

// Back simulates a click on the tall back button at the left of slider panels.
func (tools *Tools) Back() {
	backdiv := tools.backStack[0]
	tools.backStack = tools.backStack[1:]
	tools.HideShow(tools.here, backdiv)
}

func (tools *Tools) showSlider() {
	tools.ElementShow(tools.tabsMasterviewHomeSlider)
}

func (tools *Tools) hideSlider() {
	tools.ElementHide(tools.tabsMasterviewHomeSlider)
}

func (tools *Tools) initializeSlider() {
	notjs := tools.notjs
	buttoncb := notjs.RegisterCallBack(tools.handlePadButtonOnClick)
	divs := notjs.GetElementsByTagName("DIV")
	for _, div := range divs {
		if notjs.ClassListContains(div, SliderButtonPadClassName) {
			children := notjs.ChildrenSlice(div)
			for _, ch := range children {
				if notjs.TagName(ch) == "BUTTON" {
					notjs.SetOnClick(ch, buttoncb)
				}
			}
		} else if div == tools.tabsMasterviewHomeButtonPad {
			children := notjs.ChildrenSlice(div)
			for _, ch := range children {
				if notjs.TagName(ch) == "BUTTON" {
					notjs.SetOnClick(ch, buttoncb)
				}
			}
		}
	}
	backcb := notjs.RegisterCallBack(tools.handleBack)
	notjs.SetOnClick(tools.tabsMasterviewHomeSliderBack, backcb)
}

func (tools *Tools) handlePadButtonOnClick(args []js.Value) {
	// get back div
	notjs := tools.notjs
	target := args[0].Get("target")
	backid := target.Call("getAttribute", BackIDAttribute).String()
	backdiv := notjs.GetElementByID(backid)
	// get forward div
	targetid := notjs.ID(target)
	divs, found := tools.buttonPanelsMap[targetid]
	if !found {
		notjs.Alert(fmt.Sprintf("slider.controler.handlePadButtonOnClick: id %q not found in tools.buttonPanelsMap", targetid))
		return
	}
	for _, div := range divs {
		if notjs.ClassListContains(div, ToBeSeenClassName) {
			tools.here = div
			tools.backStack = append(tools.backStack, backdiv)
			tools.HideShow(backdiv, div)
			return
		}
	}
	notjs.Alert(fmt.Sprintf("slider.controler.handlePadButtonOnClick: tobe-seen not found with button %q", target.Get("innerText")))
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
	p := tools.notjs.ParentNode(tools.here)
	return p == tools.tabsMasterviewHomeSliderCollection
}
`
