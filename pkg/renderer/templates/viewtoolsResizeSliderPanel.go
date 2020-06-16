package templates

// ViewToolsResizeSliderPanel is the renerer/viewtools/resizeSliderPanel.go file.
const ViewToolsResizeSliderPanel = `// +build js, wasm

package viewtools

import (
	"fmt"
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportRendererAPIWindow}}"
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
	// This is how the panels are layed out.
	// tabbar -> undertabbar -> panelBoundToTab (the tab's panel) -> group (the collection of panels. only 1 is visible.) -> usercontent -> markup (template wrapper) -> template.

	// the tab bar height is already set.
	// remove the height of the tab bar.
	window.SetStyleWidth(tabbar, w-window.WidthExtras(tabbar))
	formatTabBarTabs(tabbar)
	h -= window.OuterHeight(tabbar)
	// Under tab bar size.
	w -= window.WidthExtras(underTabbar)
	h -= window.HeightExtras(underTabbar)
	window.SetStyleWidth(underTabbar, w)
	window.SetStyleHeight(underTabbar, h)

	// Panel bound to tab size.
	panelBoundToTab := js.Undefined()
	children := underTabbar.Get("children")
	l := children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if !classList.Call("contains", UnSeenClassName).Bool() {
			panelBoundToTab = ch
			break
		}
	}
	if panelBoundToTab.IsUndefined() {
		// this will only happen in development and testing of kickwasm.
		message := fmt.Sprintf("missing panelBoundToTab div under %s", underTabbar.Get("id").String())
		alert.Invoke(message)
		return
	}
	w -= window.WidthExtras(panelBoundToTab)
	h -= window.HeightExtras(panelBoundToTab)
	window.SetStyleWidth(panelBoundToTab, w)
	window.SetStyleHeight(panelBoundToTab, h)

	// The heading (h3) and group div inside the panel bound to tab.
	children = panelBoundToTab.Get("children")
	l = children.Length()
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if classList.Call("contains", PanelHeadingClassName).Bool() {
			// size the heading
			window.SetStyleWidth(ch, w-window.WidthExtras(ch))
			h -= window.OuterHeight(ch)
		} else if classList.Call("contains", TabPanelGroupClassName).Bool() {
			// size the panel group panel
			w -= window.WidthExtras(ch)
			h -= window.HeightExtras(ch)
			window.SetStyleWidth(ch, w)
			window.SetStyleHeight(ch, h)
			children2 := ch.Get("children")
			l2 := children2.Length()
			for i2 := 0; i2 < l2; i2++ {
				userContentDiv := children2.Index(i2)
				// find the visible user content panel in this group.
				userContentClassList := userContentDiv.Get("classList")
				if !userContentClassList.Call("contains", UnSeenClassName).Bool() {
					if userContentClassList.Call("contains", UserContentClassName).Bool() {
						// User content panel size.
						w -= window.WidthExtras(userContentDiv)
						h -= window.HeightExtras(userContentDiv)
						window.SetStyleWidth(userContentDiv, w)
						window.SetStyleHeight(userContentDiv, h)
						// Markup (template wrapper) size.
						markupPanel := userContentDiv.Get("children").Index(0)
						markupPanelClassList := markupPanel.Get("classList")
						markupPanelStyles := markupPanel.Get("style")

						// // Check the template elements.
						// // If any elements require a height resize
						// //  then the user content panel needs to be modified.
						templateElements := markupPanel.Get("children")
						templateElementsLength := templateElements.Length()
						for i4 := 0; i4 < templateElementsLength; i4++ {
							el := templateElements.Index(i4)
							classList := el.Get("classList")
							if !classList.Call("contains", UnSeenClassName).Bool() {
								if classList.Call("contains", ResizeMeHeightClassName).Bool() {
									// At least one of the template elements requires height resizing.
									// Don't allow vertical scrolling in the user content panel.
									userContentClassList.Call("remove", VScrollClassName)
									markupPanelClassList.Call("add", ResizeMeHeightClassName)
									markupPanelStyles.Set("overflow", "hidden")
									break
								}
							}
						}
						resizeTemplateMarkupWidth(markupPanel, w)
						resizeTabTemplateWrapperHeight(markupPanel, h)
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

func resizeTabTemplateWrapperHeight(wrapper js.Value, h float64) {
	h = h - window.HeightExtras(wrapper)
	children := wrapper.Get("children")
	l := children.Length()
	resizeMes := make([]js.Value, 0, l)
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if !classList.Call("contains", UnSeenClassName).Bool() {
			if classList.Call("contains", ResizeMeHeightClassName).Bool() {
				resizeMes = append(resizeMes, ch)
			} else {
				h = h - window.OuterHeight(ch)
			}
		}
	}
	h = h / float64(len(resizeMes))
	for _, ch := range resizeMes {
		resizeTemplateMarkupHeight(ch, h)
	}
}

func resizeTemplateMarkupHeight(mine js.Value, h float64) {
	h = h - window.HeightExtras(mine)
	window.SetStyleHeight(mine, h)
	children := mine.Get("children")
	l := children.Length()
	resizeMes := make([]js.Value, 0, l)
	for i := 0; i < l; i++ {
		ch := children.Index(i)
		classList := ch.Get("classList")
		if !classList.Call("contains", UnSeenClassName).Bool() {
			if classList.Call("contains", ResizeMeHeightClassName).Bool() {
				resizeMes = append(resizeMes, ch)
			} else {
				h = h - window.OuterHeight(ch)
			}
		}
	}
	h = h / float64(len(resizeMes))
	for _, ch := range resizeMes {
		resizeTemplateMarkupHeight(ch, h)
	}
}
`
