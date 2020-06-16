// +build js, wasm

package jsvalue

import (
	"syscall/js"
)

const (
	hVScrollClassName       = "hvscroll"
	resizeMeWidthClassName  = "resize-me-width"
	resizeMeHeightClassName = "resize-me-height"
)

// AddHorizontalScroll gives the element horizontal scroll.
// Use it when you are using the GO package syscall/js to create DOM elements.
// The element's minimum width must be set in a css file.
func AddHorizontalScroll(e js.Value) {
	classList := e.Get("classList")
	classList.Call("add", hVScrollClassName)
}

// AddHorizontalScrollMinWidth gives the element horizontal scroll.
// Use it when you are using the GO package syscall/js to create DOM elements.
// It also styles it's minimum width.
// Use this if the element's style doesn't set it's min-width.
func AddHorizontalScrollMinWidth(e js.Value, minwidth float64) {
	classList := e.Get("classList")
	if classList.Call("contains", resizeMeWidthClassName).Bool() {
		// This element has resizing so it can't have a min-width or horizontal scrolling.
		return
	}
	classList.Call("add", hVScrollClassName)
	style := e.Get("style")
	style.Set("min-width", minwidth)
}

// AddHorizontalResize allows the block element to be resized horizontally by the framework.
// Use it when you are using the GO package syscall/js to create DOM elements.
func AddHorizontalResize(e js.Value) {
	classList := e.Get("classList")
	if classList.Call("contains", hVScrollClassName).Bool() {
		// This element has horizontal scrolling and or minwidth.
		return
	}
	classList.Call("add", resizeMeWidthClassName)
}

// AddVerticalResize allows the block element to be resized vertically by the framework.
// Use it when you are using the GO package syscall/js to create DOM elements.
func AddVerticalResize(e js.Value) {
	classList := e.Get("classList")
	if classList.Call("contains", hVScrollClassName).Bool() {
		// This element has horizontal scrolling and or minwidth.
		return
	}
	classList.Call("add", resizeMeHeightClassName)
}
