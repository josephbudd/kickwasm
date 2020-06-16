// +build js, wasm

package viewtools

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/api/window"
)

func sizeSliderPanel(sliderPanel js.Value, w, h float64) {
	// #mainMasterView-home-slider-collection is the parant of a slider panel
	w -= window.WidthExtras(sliderPanel)
	h -= window.HeightExtras(sliderPanel)
	inner := js.Undefined()
	marginHt := float64(0)
	headingHt := float64(0)
	// Find the sliderPanel's cookie crumbs, heading and inner panel.
	// Get height of cookie crumbs and heading. Both use PanelHeadingClassName.
	children := sliderPanel.Get("children")
	l := children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if classList.Call("contains", PanelHeadingClassName).Bool() {
			chwx := window.WidthExtras(ch)
			window.SetStyleWidth(ch, w-chwx)
			marginHt = headingHt
			headingHt += window.OuterHeight(ch)
		}
		if classList.Call("contains", SliderPanelInnerClassName).Bool() {
			inner = ch
		}
	}
	// size the back button
	reSizeSliderBack(h-marginHt, marginHt)
	// size this slider panel
	h -= headingHt
	window.SetStyleWidth(sliderPanel, w)
	window.SetStyleHeight(sliderPanel, h)
	// size slider panel's inner panel
	sizeSliderPanelInnerPanel(inner, w, h)
}

func sizeSliderPanelInnerPanel(inner js.Value, w, h float64) {
	// Calculate and set the inside dimensions of the inner panel.
	w -= window.WidthExtras(inner)
	h -= window.HeightExtras(inner)
	window.SetStyleWidth(inner, w)
	window.SetStyleHeight(inner, h)
	// Inside the inner panel will be:
	// * button pad
	// * or user content
	// * or tab bar.
	tabbar := js.Undefined()
	underTabbar := js.Undefined()
	children := inner.Get("children")
	l := children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if classList.Call("contains", SliderButtonPadClassName).Bool() {
			// This inner panel child is a button pad.
			sizeSliderPanelInnerPanelButtonPad(ch, w, h)
			return
		}
		if classList.Call("contains", UserContentClassName).Bool() {
			// This inner panel child is user content, a markup panel.
			sizeSliderPanelInnerPanelUserContent(ch, w, h)
			return
		}
		if classList.Call("contains", TabBarClassName).Bool() {
			// This inner panel child is a tab bar.
			tabbar = ch
			// Continue to get the UnderTabBarClassName
		}
		if classList.Call("contains", UnderTabBarClassName).Bool() {
			// This inner panel wraps a tab bar.
			// Under the tab bar div is the under tab bar div.
			// This inner panel child is the under tab bar div.
			underTabbar = ch
			sizeSliderPanelInnerPanelTabBar(tabbar, underTabbar, w, h)
			return
		}
	}
}

func sizeSliderPanelInnerPanelButtonPad(buttonPad js.Value, w, h float64) {
	// Calculate and set the inside dimensions of the button pad.
	w -= window.WidthExtras(buttonPad)
	h -= window.HeightExtras(buttonPad)
	window.SetStyleHeight(buttonPad, h)
	window.SetStyleWidth(buttonPad, w)
}

func sizeSliderPanelInnerPanelUserContent(userContent js.Value, w, h float64) {
	userContentChildren := userContent.Get("children")
	userContentClassList := userContent.Get("classList")
	templateWrapper := userContentChildren.Index(0)
	templateWrapperChildren := templateWrapper.Get("children")
	template := templateWrapperChildren.Index(0)
	templateClassList := template.Get("classList")
	if templateClassList.Call("contains", ResizeMeHeightClassName).Bool() {
		userContentClassList.Call("remove", VScrollClassName)
	}
	// Calculate and set the inside dimensions of the user content div.
	// A user content panel has scroll bars to it's height must be set.
	w -= window.WidthExtras(userContent)
	h -= window.HeightExtras(userContent)
	window.SetStyleWidth(userContent, w)
	window.SetStyleHeight(userContent, h)
	// The user content div's child div wraps the template markup.
	w -= window.WidthExtras(templateWrapper)
	h -= window.HeightExtras(templateWrapper)
	if !userContentClassList.Call("contains", VScrollClassName).Bool() {
		// The user content div does not contain the vertical only scroll class.
		// So size the markup width.
		// Calculate and set the inside dimensions of the markup div.
		w -= window.WidthExtras(template)
		h -= window.HeightExtras(template)
		window.SetStyleWidth(template, w)
		window.SetStyleHeight(template, h)
	}
	// Check the children of the markup div for whatever needs it's width sized.
	children := template.Get("children")
	l := children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if !classList.Call("contains", UnSeenClassName).Bool() {
			if classList.Call("contains", ResizeMeWidthClassName).Bool() {
				resizeTemplateMarkupWidth(ch, w)
			}
			if classList.Call("contains", ResizeMeHeightClassName).Bool() {
				resizeTemplateMarkupHeight(ch, h)
			}
		}
	}
}

func sizeSliderPanelInnerPanelTabBar(tabbar, underTabbar js.Value, w, h float64) {
	seen := js.Undefined()
	// the tab bar height is already set.
	// remove the height of the tab bar.
	window.SetStyleWidth(tabbar, w-window.WidthExtras(tabbar))
	formatTabBarTabs(tabbar)
	h -= window.OuterHeight(tabbar)
	// set the under tab bar width
	w -= window.WidthExtras(underTabbar)
	window.SetStyleWidth(underTabbar, w)
	// set the under tab bar height
	h -= window.HeightExtras(underTabbar)
	window.SetStyleHeight(underTabbar, h)

	// find the visible panel under the tab bar
	// its class will be "panel-bound-to-tab"
	children := underTabbar.Get("children")
	l := children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if !classList.Call("contains", UnSeenClassName).Bool() {
			seen = ch
			break
		}
	}
	if seen.IsUndefined() {
		// this will only happen in development and testing of kickwasm.
		message := fmt.Sprintf("missing seen div under %s", underTabbar.Get("id").String())
		alert.Invoke(message)
		return
	}
	// size the visible panel inside the under the tab bar
	w -= window.WidthExtras(seen)
	window.SetStyleWidth(seen, w)
	window.SetStyleHeight(seen, h)

	// the visible panel inside the under the tab bar has a heading over its inner panel
	// the inner panel's height is height of the under tab bar - the heading height.
	// h3
	// TabPanelGroupClassName
	//  user-content & seen & scroller
	//    markup > template
	//  user-content & unseen & scroller
	//    markup > template
	children = seen.Get("children")
	l = children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if classList.Call("contains", PanelHeadingClassName).Bool() {
			// size the heading
			chwx1 := window.WidthExtras(ch)
			window.SetStyleWidth(ch, w-chwx1)
			h -= window.OuterHeight(ch)
		} else if classList.Call("contains", TabPanelGroupClassName).Bool() {
			// size the panel group panel
			chwx1 := window.WidthExtras(ch)
			window.SetStyleWidth(ch, w-chwx1)
			chhx1 := window.HeightExtras(ch)
			window.SetStyleHeight(ch, h-chhx1)
			children2 := ch.Get("children")
			l2 := children2.Length()
			for i2 := 0; i2 < l2; i2++ {
				ch2 := children2.Index(i2)
				// site the visible user content panel in this group.
				classList := ch2.Get("classList")
				if !classList.Call("contains", UnSeenClassName).Bool() {
					if classList.Call("contains", UserContentClassName).Bool() {
						// size the visible user content panel in this group.
						// A user content panel has scroll bars so it height must be set.
						chwx2 := window.WidthExtras(ch2)
						chhx2 := window.HeightExtras(ch2)
						window.SetStyleWidth(ch2, w-chwx1-chwx2)
						window.SetStyleHeight(ch2, h-chhx1-chhx2)
						// size the markup panel containing the template markup
						children3 := ch2.Get("children")
						ch3 := children3.Index(0)
						chwx3 := window.WidthExtras(ch3)
						window.SetStyleWidth(ch3, w-chwx1-chwx2-chwx3)
						// size all children with the ResizeMeWidthClassName
						children4 := ch3.Get("children")
						l4 := children4.Length()
						width4 := w - chwx1 - chwx2 - chwx3
						for i4 := 0; i4 < l4; i4++ {
							ch4 := children4.Index(i4)
							classList := ch4.Get("classList")
							if !classList.Call("contains", UnSeenClassName).Bool() {
								if classList.Call("contains", ResizeMeWidthClassName).Bool() {
									resizeTemplateMarkupWidth(ch4, width4)
								}
								if classList.Call("contains", ResizeMeHeightClassName).Bool() {
									resizeTemplateMarkupHeight(ch, h)
								}
							}
						}
						return
					}
				}
			}
			break
		}
	}
}

func resizeTemplateMarkupWidth(mine js.Value, w float64) {
	w = w - window.WidthExtras(mine)
	window.SetStyleWidth(mine, w)
	children := mine.Get("children")
	l := children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if !classList.Call("contains", UnSeenClassName).Bool() {
			if classList.Call("contains", ResizeMeWidthClassName).Bool() {
				resizeTemplateMarkupWidth(ch, w)
			}
		}
	}
}

func resizeTemplateMarkupHeight(mine js.Value, h float64) {
	h = h - window.HeightExtras(mine)
	window.SetStyleHeight(mine, h)
	children := mine.Get("children")
	l := children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if !classList.Call("contains", UnSeenClassName).Bool() {
			if classList.Call("contains", ResizeMeHeightClassName).Bool() {
				resizeTemplateMarkupHeight(ch, h)
			} else {
				h = h - window.OuterHeight(ch)
			}
		}
	}
}
