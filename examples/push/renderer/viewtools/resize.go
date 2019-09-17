package viewtools

import (
	"fmt"
	"syscall/js"
)

func (tools *Tools) initializeResize() {
	cb := tools.RegisterEventCallBack(
		func(event js.Value) interface{} {
			tools.SizeApp()
			return nil
		},
		true, true, true,
	)
	tools.Global.Set("onresize", cb)
}

// SizeApp resizes the app
func (tools *Tools) SizeApp() {
	// begin with the height of the inside of the browser where the window is.
	notJS := tools.NotJS
	windowWidth := notJS.WindowInnerWidth()
	windowHeight := notJS.WindowInnerHeight()
	// and subtract body measurments
	bodies := notJS.GetElementsByTagName("body")
	body := bodies[0]
	xh := notJS.HeightExtras(body)
	xw := notJS.WidthExtras(body)
	windowHeight -= xh
	windowWidth -= xw
	// size each master view
	tools.sizeTabsMasterView(windowWidth, windowHeight)
	tools.sizeModalMasterView(windowWidth, windowHeight)
	tools.sizeCloserMasterView(windowWidth, windowHeight)
}

func (tools *Tools) sizeTabsMasterView(w, h float64) {
	// now set the masterview height
	notJS := tools.NotJS
	if tools.ElementIsShown(tools.tabsMasterview) {
		// tabs masterview is visible
		// subtract extras before setting
		h -= notJS.HeightExtras(tools.tabsMasterview)
		w -= notJS.WidthExtras(tools.tabsMasterview)
		// set master view height, width
		notJS.SetStyleHeight(tools.tabsMasterview, h)
		notJS.SetStyleWidth(tools.tabsMasterview, w)
		// div #tabsMasterview children
		// * H1
		// * div #tabsMasterview-home
		// * div #tabsMasterview-home-slider
		//
		// Process h1
		h1Ht := float64(0)
		children := notJS.ChildrenSlice(tools.tabsMasterview)
		for _, ch := range children {
			if notJS.TagName(ch) == "H1" {
				chwx := notJS.WidthExtras(ch)
				notJS.SetStyleWidth(ch, w-chwx)
				h1Ht = notJS.OuterHeight(ch)
				break
			}
		}
		// h - h1Ht is now the ht for home or slider.
		// home panel or slider is under the h1.
		h -= h1Ht
		// Process Home
		if tools.ElementIsShown(tools.tabsMasterviewHome) {
			// remove extra measurements
			h -= notJS.HeightExtras(tools.tabsMasterviewHome)
			w -= notJS.WidthExtras(tools.tabsMasterviewHome)
			// set the inside height and width
			notJS.SetStyleHeight(tools.tabsMasterviewHome, h)
			notJS.SetStyleWidth(tools.tabsMasterviewHome, w)
			// homepad is the button pad in home.
			h -= notJS.HeightExtras(tools.tabsMasterviewHomeButtonPad)
			w -= notJS.WidthExtras(tools.tabsMasterviewHomeButtonPad)
			//h -= 100
			notJS.SetStyleHeight(tools.tabsMasterviewHomeButtonPad, h)
			notJS.SetStyleWidth(tools.tabsMasterviewHomeButtonPad, w)
			return
		}
		// home is not visible check the slider
		// Process Slider
		if tools.ElementIsShown(tools.tabsMasterviewHomeSlider) {
			// remove extra measurements
			h -= notJS.HeightExtras(tools.tabsMasterviewHomeSlider)
			w -= notJS.WidthExtras(tools.tabsMasterviewHomeSlider)
			// set the inside height and width
			notJS.SetStyleHeight(tools.tabsMasterviewHomeSlider, h)
			notJS.SetStyleWidth(tools.tabsMasterviewHomeSlider, w)
			// slider has a back button
			backOuterWidth := notJS.OuterWidth(tools.tabsMasterviewHomeSliderBack)
			w -= backOuterWidth
			// size slider collection
			h -= notJS.HeightExtras(tools.tabsMasterviewHomeSliderCollection)
			w -= notJS.WidthExtras(tools.tabsMasterviewHomeSliderCollection)
			notJS.SetStyleHeight(tools.tabsMasterviewHomeSliderCollection, h)
			notJS.SetStyleWidth(tools.tabsMasterviewHomeSliderCollection, w)
			// slider collection children
			children := notJS.ChildrenSlice(tools.tabsMasterviewHomeSliderCollection)
			for _, ch := range children {
				if notJS.TagName(ch) == "DIV" && notJS.ClassListContainsAnd(ch, SliderPanelClassName, SeenClassName) {
					tools.sizeSliderPanel(ch, w, h)
					break
				}
			}
		}
	}
}

func (tools *Tools) reSizeSliderBack(height, margintop float64) {
	style := tools.tabsMasterviewHomeSliderBack.Get("style")
	style.Set("height", fmt.Sprintf("%fpx", height))
	style.Set("margin-top", fmt.Sprintf("%fpx", margintop))
}

func (tools *Tools) sizeModalMasterView(w, h float64) {
	// modal master view
	if tools.ElementIsShown(tools.modalMasterView) {
		notJS := tools.NotJS
		// modal view is visible
		w -= notJS.WidthExtras(tools.modalMasterView)
		h -= notJS.HeightExtras(tools.modalMasterView)
		notJS.SetStyleWidth(tools.modalMasterView, w)
		notJS.SetStyleHeight(tools.modalMasterView, h)
		// the center div
		w -= notJS.WidthExtras(tools.modalMasterViewCenter)
		h -= notJS.HeightExtras(tools.modalMasterViewCenter)
		notJS.SetStyleWidth(tools.modalMasterViewCenter, w)
		notJS.SetStyleHeight(tools.modalMasterViewCenter, h)
		// subtract ht of h1 and p > button
		children := notJS.ChildrenSlice(tools.modalMasterViewCenter)
		for _, ch := range children {
			tagName := notJS.TagName(ch)
			if tagName == "H1" || tagName == "P" {
				chwx := notJS.WidthExtras(ch)
				notJS.SetStyleWidth(ch, w-chwx)
				oh := notJS.OuterHeight(ch)
				h -= oh
			}
		}
		// message
		w -= notJS.WidthExtras(tools.modalMasterViewMessage)
		h -= notJS.HeightExtras(tools.modalMasterViewMessage)
		notJS.SetStyleWidth(tools.modalMasterViewMessage, w)
		notJS.SetStyleHeight(tools.modalMasterViewMessage, h)
	}
}

func (tools *Tools) sizeCloserMasterView(w, h float64) {
	if tools.ElementIsShown(tools.closerMasterView) {
		tools.NotJS.SetStyleHeight(tools.closerMasterView, h)
	}
}


func (tools *Tools) resizeMe(mine js.Value, w, h float64) {
	notJS := tools.NotJS
	w = w - notJS.WidthExtras(mine)
	notJS.SetStyleWidth(mine, w)
	children := notJS.ChildrenSlice(mine)
	for _, ch := range children {
		if !notJS.ClassListContains(ch, UnSeenClassName) {
			if notJS.ClassListContains(ch, ResizeMeWidthClassName) {
				tools.resizeMe(ch, w, h)
			}
		}
	}
}
