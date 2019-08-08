package viewtools

import (
	"fmt"
	"syscall/js"
)

func (tools *Tools) sizeSliderPanel(sliderPanel js.Value, w, h float64) {
	// #tabsMasterView-home-slider-collection is the parant of a slider panel
	notJS := tools.NotJS
	w -= notJS.WidthExtras(sliderPanel)
	h -= notJS.HeightExtras(sliderPanel)
	inner := js.Undefined()
	marginHt := float64(0)
	headingHt := float64(0)
	// Find the sliderPanel's cookie crumbs, heading and inner panel.
	// Get height of cookie crumbs and heading. Both use PanelHeadingClassName.
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
	tools.sizeSliderPanelInnerPanel(inner, w, h)
}

func (tools *Tools) sizeSliderPanelInnerPanel(inner js.Value, w, h float64) {
	notJS := tools.NotJS
	// Calculate and set the inside dimensions of the inner panel.
	w -= notJS.WidthExtras(inner)
	h -= notJS.HeightExtras(inner)
	notJS.SetStyleWidth(inner, w)
	notJS.SetStyleHeight(inner, h)
	// Inside the inner panel will be:
	// * button pad
	// * or user content
	// * or tab bar.
	tabbar := js.Undefined()
	underTabbar := js.Undefined()
	children := notJS.ChildrenSlice(inner)
	for _, ch := range children {
		if notJS.ClassListContains(ch, SliderButtonPadClassName) {
			// This inner panel child is a button pad.
			tools.sizeSliderPanelInnerPanelButtonPad(ch, w, h)
			return
		}
		if notJS.ClassListContains(ch, UserContentClassName) {
			// This inner panel child is user content, a markup panel.
			tools.sizeSliderPanelInnerPanelUserContent(ch, w, h)
			return
		}
		if notJS.ClassListContains(ch, TabBarClassName) {
			// This inner panel child is a tab bar.
			tabbar = ch
			// Continue to get the UnderTabBarClassName
		}
		if notJS.ClassListContains(ch, UnderTabBarClassName) {
			// This inner panel wraps a tab bar.
			// Under the tab bar div is the under tab bar div.
			// This inner panel child is the under tab bar div.
			underTabbar = ch
			tools.sizeSliderPanelInnerPanelTabBar(tabbar, underTabbar, w, h)
			return
		}
	}
}

func (tools *Tools) sizeSliderPanelInnerPanelButtonPad(buttonPad js.Value, w, h float64) {
	notJS := tools.NotJS
	// Calculate and set the inside dimensions of the button pad.
	w -= notJS.WidthExtras(buttonPad)
	h -= notJS.HeightExtras(buttonPad)
	notJS.SetStyleHeight(buttonPad, h)
	notJS.SetStyleWidth(buttonPad, w)
}

func (tools *Tools) sizeSliderPanelInnerPanelUserContent(userContent js.Value, w, h float64) {
	// Calculate and set the inside dimensions of the user content div.
	notJS := tools.NotJS
	w -= notJS.WidthExtras(userContent)
	h -= notJS.HeightExtras(userContent)
	notJS.SetStyleWidth(userContent, w)
	notJS.SetStyleHeight(userContent, h)
	// The user content div wraps the template markup.
	children := notJS.ChildrenSlice(userContent)
	markup := children[0]
	// Calculate and set the inside dimensions of the markup div.
	w -= notJS.WidthExtras(markup)
	h -= notJS.HeightExtras(markup)
	notJS.SetStyleWidth(markup, w)
	notJS.SetStyleHeight(markup, h)
	// Check the children of the markup div for whatever needs it's width sized.
	children = notJS.ChildrenSlice(markup)
	for _, ch := range children {
		if !notJS.ClassListContains(ch, UnSeenClassName) {
			if notJS.ClassListContains(ch, ResizeMeWidthClassName) {
				tools.resizeMe(ch, w, h)
			}
		}
	}
}

func (tools *Tools) sizeSliderPanelInnerPanelTabBar(tabbar, underTabbar js.Value, w, h float64) {
	notJS := tools.NotJS
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
			chhx1 := notJS.HeightExtras(ch)
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
}
