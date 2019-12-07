package templates

// MarkupFocusBlurGo is the renderer/markup/focusblur.go file.
const MarkupFocusBlurGo = `// +build js, wasm

package markup

// Focus sets an element to focused.
func (e *Element) Focus() {
	e.element.Call("focus")
}

// Blur removes an element's focus.
func (e *Element) Blur() {
	e.element.Call("blur")
}
`
