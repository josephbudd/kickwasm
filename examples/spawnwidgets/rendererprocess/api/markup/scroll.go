// +build js, wasm

package markup

import (
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/window"
)

// scroll

// ScrollTo scrolls element to the left, top.
func (e *Element) ScrollTo(left, top float64) {
	e.element.Call("scrollTo", left, top)
}

// ScrollTop returns an element's scroll top.
func (e *Element) ScrollTop() (top float64) {
	top = e.element.Get("scrollTop").Float()
	return
}

// ScrollLeft returns an element's scroll left.
func (e *Element) ScrollLeft() (left float64) {
	left = e.element.Get("scrollLeft").Float()
	return
}

// Horizontal Scrolling.

// AddHorizontalScroll gives the element horizontal scroll.
// The element's minimum width must be set in a css file.
func (e *Element) AddHorizontalScroll() {
	e.AddClass(hVScrollClassName)
}

// AddHorizontalScrollMinWidth gives the element horizontal scroll.
// It also styles it's minimum width.
// Use this is the element's style doesn't set it's min-width.
func (e *Element) AddHorizontalScrollMinWidth(minwidth float64) {
	e.AddClass(hVScrollClassName)
	window.SetStyleMinWidth(e.element, minwidth)
}
