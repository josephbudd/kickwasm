// +build js, wasm

package viewtools

import (
	"fmt"

	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/framework/callback"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/event"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/window"
)

func initializeResize() {
	f := func(e event.Event) (nilReturn interface{}) {
		SizeApp()
		return
	}
	callback.AddEventHandler(f, global, "resize", false, 0)
}

// SizeApp resizes the app
func SizeApp() {
	extraHeight = 0.0
	// begin with the height of the inside of the browser where the window is.
	windowWidth := window.WindowInnerWidth()
	windowHeight := window.WindowInnerHeight()
	// and subtract body measurments
	xh := window.HeightExtras(body)
	xw := window.WidthExtras(body)
	appWidth := windowWidth - xw
	appHeight := windowHeight - xh
	// size each master view
	sizeTabsMasterView(appWidth, appHeight)
	sizeModalMasterView(appWidth, appHeight)
	sizeBlackMasterView(windowWidth, windowHeight)
}

// overSizeApp resizes the app
func overSizeApp() {
	// begin with the height of the inside of the browser where the window is.
	windowWidth := window.WindowInnerWidth()
	windowHeight := window.WindowInnerHeight() + extraHeight
	// and subtract body measurments
	xw := window.WidthExtras(body)
	xh := window.HeightExtras(body)
	appWidth := windowWidth - xw
	appHeight := windowHeight - xh
	// size each master view
	sizeTabsMasterView(appWidth, appHeight)
	sizeModalMasterView(appWidth, appHeight)
	sizeBlackMasterView(windowWidth, windowHeight)
}

func sizeTabsMasterView(w, h float64) {
	// now set the masterview height
	if ElementIsShown(tabsMasterview) {
		// tabs masterview is visible
		// subtract extras before setting
		h -= window.HeightExtras(tabsMasterview)
		w -= window.WidthExtras(tabsMasterview)
		// set master view height, width
		window.SetStyleHeight(tabsMasterview, h)
		window.SetStyleWidth(tabsMasterview, w)
		// div #tabsMasterview children
		// * H1
		// * div #tabsMasterview-home
		// * div #tabsMasterview-home-slider
		//
		// Process h1
		h1Ht := float64(0)

		children := tabsMasterview.Get("children")
		l := children.Length()
		for i := 0; i < l; i++ {
			ch := children.Index(i)
			if ch.Get("tagName").String() == "H1" {
				chwx := window.WidthExtras(ch)
				window.SetStyleWidth(ch, w-chwx)
				h1Ht = window.OuterHeight(ch)
				break
			}
		}
		// h - h1Ht is now the ht for home or slider.
		// home panel or slider is under the h1.
		h -= h1Ht
		// Process Home
		if ElementIsShown(tabsMasterviewHome) {
			// remove extra measurements
			h -= window.HeightExtras(tabsMasterviewHome)
			w -= window.WidthExtras(tabsMasterviewHome)
			// set the inside height and width
			window.SetStyleHeight(tabsMasterviewHome, h)
			window.SetStyleWidth(tabsMasterviewHome, w)
			// homepad is the button pad in home.
			h -= window.HeightExtras(tabsMasterviewHomeButtonPad)
			w -= window.WidthExtras(tabsMasterviewHomeButtonPad)
			//h -= 100
			window.SetStyleHeight(tabsMasterviewHomeButtonPad, h)
			window.SetStyleWidth(tabsMasterviewHomeButtonPad, w)
			return
		}
		// home is not visible check the slider
		// Process Slider
		if ElementIsShown(tabsMasterviewHomeSlider) {
			// remove extra measurements
			h -= window.HeightExtras(tabsMasterviewHomeSlider)
			w -= window.WidthExtras(tabsMasterviewHomeSlider)
			// set the inside height and width
			window.SetStyleHeight(tabsMasterviewHomeSlider, h)
			window.SetStyleWidth(tabsMasterviewHomeSlider, w)
			// slider has a back button
			backOuterWidth := window.OuterWidth(tabsMasterviewHomeSliderBack)
			w -= backOuterWidth
			// size slider collection
			h -= window.HeightExtras(tabsMasterviewHomeSliderCollection)
			w -= window.WidthExtras(tabsMasterviewHomeSliderCollection)
			window.SetStyleHeight(tabsMasterviewHomeSliderCollection, h)
			window.SetStyleWidth(tabsMasterviewHomeSliderCollection, w)
			// slider collection children

			children := tabsMasterviewHomeSliderCollection.Get("children")
			l := children.Length()
			for i := 0; i < l; i++ {
				ch := children.Index(i)
				if tagName := ch.Get("tagName").String(); tagName == "DIV" {
					// Is this div a visible slider panel?
					isvisibleSlider := false
					classList := ch.Get("classList")
					if isvisibleSlider = classList.Call("contains", SliderPanelClassName).Bool(); isvisibleSlider {
						isvisibleSlider = classList.Call("contains", SeenClassName).Bool()
					}
					if isvisibleSlider {
						sizeSliderPanel(ch, w, h)
						break
					}
				}
			}
		}
	}
}

func reSizeSliderBack(height, margintop float64) {
	style := tabsMasterviewHomeSliderBack.Get("style")
	style.Set("height", fmt.Sprintf("%fpx", height))
	style.Set("margin-top", fmt.Sprintf("%fpx", margintop))
}

func sizeModalMasterView(w, h float64) {
	// modal master view
	if ElementIsShown(modalMasterView) {
		// modal view is visible
		w -= window.WidthExtras(modalMasterView)
		h -= window.HeightExtras(modalMasterView)
		window.SetStyleWidth(modalMasterView, w)
		window.SetStyleHeight(modalMasterView, h)
		// the center div
		w -= window.WidthExtras(modalMasterViewCenter)
		h -= window.HeightExtras(modalMasterViewCenter)
		window.SetStyleWidth(modalMasterViewCenter, w)
		window.SetStyleHeight(modalMasterViewCenter, h)
		// subtract ht of h1 and p > button
		children := modalMasterViewCenter.Get("children")
		l := children.Length()
		for i := 0; i < l; i++ {
			ch := children.Index(i)
			tagName := ch.Get("tagName").String()
			if tagName == "H1" || tagName == "P" {
				chwx := window.WidthExtras(ch)
				window.SetStyleWidth(ch, w-chwx)
				oh := window.OuterHeight(ch)
				h -= oh
			}
		}
		// message
		w -= window.WidthExtras(modalMasterViewMessage)
		h -= window.HeightExtras(modalMasterViewMessage)
		window.SetStyleWidth(modalMasterViewMessage, w)
		window.SetStyleHeight(modalMasterViewMessage, h)
	}
}

func sizeBlackMasterView(w, h float64) {
	if ElementIsShown(blackMasterView) {
		window.SetStyleWidth(blackMasterView, w)
		window.SetStyleHeight(blackMasterView, h)
	}
}
