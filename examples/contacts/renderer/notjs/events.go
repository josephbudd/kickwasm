package notjs

import "syscall/js"

// GetEventTarget gets an event's target attribute which is an html element.
func (notJS *NotJS) GetEventTarget(event js.Value) js.Value {
	return event.Get("target")
}

// SetOnClick sets an element's onclick.
func (notJS *NotJS) SetOnClick(element js.Value, cb js.Callback) {
	element.Set("onclick", cb)
}

// SetOnChange sets an element's onchange.
func (notJS *NotJS) SetOnChange(element js.Value, cb js.Callback) {
	element.Set("onchange", cb)
}

// SetOnScroll sets an element's onscroll.
func (notJS *NotJS) SetOnScroll(element js.Value, cb js.Callback) {
	element.Set("onscroll", cb)
}
