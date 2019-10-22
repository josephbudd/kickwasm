// +build js, wasm

package notjs

import "syscall/js"

// GetEventTarget gets an event's target attribute which is an html element.
func (notjs *NotJS) GetEventTarget(event js.Value) js.Value {
	return event.Get("target")
}

// SetOnClick sets an element's onclick.
func (notjs *NotJS) SetOnClick(element js.Value, cb js.Func) {
	element.Set("onclick", cb)
}

// SetOnChange sets an element's onchange.
func (notjs *NotJS) SetOnChange(element js.Value, cb js.Func) {
	element.Set("onchange", cb)
}

// SetOnScroll sets an element's onscroll.
func (notjs *NotJS) SetOnScroll(element js.Value, cb js.Func) {
	element.Set("onscroll", cb)
}
