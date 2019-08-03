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

func (tools *Tools) sizeSliderPanel(sliderPanel js.Value, w, h float64) {
	// #tabsMasterView-home-slider-collection is the parant of a slider panel
	// finds and sets the ht of div.slider-panel-inner
	// remove extras.
	notJS := tools.NotJS
	w -= notJS.WidthExtras(sliderPanel)
	h -= notJS.HeightExtras(sliderPanel)
	inner := js.Undefined()
	marginHt := float64(0)
	headingHt := float64(0)
	// get height of headings. the ccs are also headings
	// first the optional ccs
	// then the actual heading
	children := notJS.ChildrenSlice(sliderPanel)
	for _, ch := range children {
		if notJS.ClassListContains(ch, PanelHeadingClassName) {
			chwx := notJS.WidthExtras(ch)
			notJS.SetStyleWidth(ch, w-chwx)
			marginHt = headingHt
			headingHt += notJS.OuterHeight(ch)
		}
		if notJS.ClassListContains(ch, SliderPanelInnerClassName) {
			inner = ch
		}
	}
	// size the back button
	tools.reSizeSliderBack(h-marginHt, marginHt)
	// size this slider panel
	h -= headingHt
	notJS.SetStyleWidth(sliderPanel, w)
	notJS.SetStyleHeight(sliderPanel, h)
	// size slider panel's inner panel
	// inside inner panel
	w -= notJS.WidthExtras(inner)
	h -= notJS.HeightExtras(inner)
	notJS.SetStyleWidth(inner, w)
	notJS.SetStyleHeight(inner, h)
	// inside the inner panel will be:
	// * button pad
	// * or user content
	// * or tab bar.
	buttonPad := js.Undefined()
	userContent := js.Undefined()
	tabbar := js.Undefined()
	underTabbar := js.Undefined()
	children = notJS.ChildrenSlice(inner)
	for _, ch := range children {
		if notJS.ClassListContains(ch, SliderButtonPadClassName) {
			buttonPad = ch
			break
		}
		if notJS.ClassListContains(ch, UserContentClassName) {
			userContent = ch
			break
		}
		if notJS.ClassListContains(ch, TabBarClassName) {
			tabbar = ch
			// continue to get the UnderTabBarClassName
		}
		if notJS.ClassListContains(ch, UnderTabBarClassName) {
			underTabbar = ch
			break
		}
	}
	if buttonPad != js.Undefined() {
		// a button pad is inside the inner panel
		w -= notJS.WidthExtras(buttonPad)
		h -= notJS.HeightExtras(buttonPad)
		notJS.SetStyleHeight(buttonPad, h)
		notJS.SetStyleWidth(buttonPad, w)
		return
	}
	if userContent != js.Undefined() {
		// a user content is inside the inner panel
		w -= notJS.WidthExtras(userContent)
		h -= notJS.HeightExtras(userContent)
		notJS.SetStyleWidth(userContent, w)
		notJS.SetStyleHeight(userContent, h)
		// markup div wraps the template markup.
		children := notJS.ChildrenSlice(userContent)
		markup := children[0]
		w -= notJS.WidthExtras(markup)
		h -= notJS.HeightExtras(markup)
		notJS.SetStyleWidth(markup, w)
		notJS.SetStyleHeight(markup, h)
		// check the template for whatever needs it's width sized.
		children = notJS.ChildrenSlice(markup)
		for _, ch := range children {
			if !notJS.ClassListContains(ch, UnSeenClassName) {
				if notJS.ClassListContains(ch, ResizeMeWidthClassName) {
					tools.resizeMe(ch, w, h)
				}
			}
		}
		return
	}
	if tabbar != js.Undefined() && underTabbar != js.Undefined() {
		// a tab bar is inside the inner panel
		seen := js.Undefined()
		// the tab bar height is already set.
		// remove the height of the tab bar.
		notJS.SetStyleWidth(tabbar, w-notJS.WidthExtras(tabbar))
		h -= notJS.OuterHeight(tabbar)
		// set the under tab bar width
		w -= notJS.WidthExtras(underTabbar)
		notJS.SetStyleWidth(underTabbar, w)
		// set the under tab bar height
		h -= notJS.HeightExtras(underTabbar)
		notJS.SetStyleHeight(underTabbar, h)

		// find the visible panel under the tab bar
		// its class will be "panel-bound-to-tab"
		children := notJS.ChildrenSlice(underTabbar)
		for _, ch := range children {
			if notJS.ClassListContains(ch, SeenClassName) {
				seen = ch
				break
			}
		}
		if seen == js.Undefined() {
			// this will only happen in development and testing of kickwasm.
			message := fmt.Sprintf("missing seen div under %s", underTabbar.Get("id"))
			notJS.Alert(message)
			return
		}
		// size the visible panel inside the under the tab bar
		w -= notJS.WidthExtras(seen)
		notJS.SetStyleWidth(seen, w)
		notJS.SetStyleHeight(seen, h)

		// the visible panel inside the under the tab bar has a heading over its inner panel
		// the inner panel's height is height of the under tab bar - the heading height.
		// h3
		// TabPanelGroupClassName
		//  user-content & seen & scroller
		//    markup > template
		//  user-content & unseen & scroller
		//    markup > template
		children = notJS.ChildrenSlice(seen)
		for _, ch := range children {
			if notJS.ClassListContains(ch, PanelHeadingClassName) {
				// size the heading
				chwx1 := notJS.WidthExtras(ch)
				notJS.SetStyleWidth(ch, w-chwx1)
				h -= notJS.OuterHeight(ch)
			} else if notJS.ClassListContains(ch, TabPanelGroupClassName) {
				// size the panel group panel
				chwx1 := notJS.WidthExtras(ch)
				notJS.SetStyleWidth(ch, w-chwx1)
				chhx1 := notJS.HeightExtras(ch)   // added xxx
				notJS.SetStyleHeight(ch, h-chhx1)
				children2 := notJS.ChildrenSlice(ch)
				// site the visible user content panel in this group.
				for _, ch := range children2 {
					if !notJS.ClassListContains(ch, UnSeenClassName) {
						if notJS.ClassListContains(ch, UserContentClassName) {
							// size the visible user content panel in this group.
							chwx2 := notJS.WidthExtras(ch)
							notJS.SetStyleWidth(ch, w-chwx1-chwx2)
							// size the markup panel containing the template markup
							children3 := notJS.ChildrenSlice(ch)
							chwx3 := notJS.WidthExtras(children3[0])
							notJS.SetStyleWidth(children3[0], w-chwx1-chwx2-chwx3)
							// size all children with the ResizeMeWidthClassName
							children4 := notJS.ChildrenSlice(ch)
							for _, ch := range children4 {
								if !notJS.ClassListContains(ch, UnSeenClassName) {
									if notJS.ClassListContains(ch, ResizeMeWidthClassName) {
										tools.resizeMe(ch, w-chwx1-chwx2-chwx3, h)
									}
								}
							}
						}
					}
				}
				break
			}
		}
		return
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
