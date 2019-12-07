// +build js, wasm

package markup

// Checked returns an element's checked.
func (e *Element) Checked() (checked bool) {
	checked = e.element.Get(checkedAttributeName).Bool()
	return
}

// SetChecked sets an element's checked
func (e *Element) SetChecked(checked bool) {
	e.element.Set(checkedAttributeName, checked)
}
