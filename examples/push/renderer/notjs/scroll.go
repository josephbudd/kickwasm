package notjs

import "syscall/js"

// ScrollTo scrolls element to the left, top.
func (notjs *NotJS) ScrollTo(element js.Value, left, top int) {
	element.Call("scrollTo", left, top)
}

// GetScrollTop returns an element's scroll top.
func (notjs *NotJS) GetScrollTop(element js.Value) int {
	return element.Get("scrollTop").Int()
}
