// +build js, wasm

package markup

// SetNotPrintable sets the element to not be printed.
func (e *Element) SetNotPrintable() {
	e.AddClass(doNotPrintClassName)
}
