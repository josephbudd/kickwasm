package templates

// MarkupElementGo is the rendererprocess/markup/element.go file.
const MarkupElementGo = `// +build js, wasm

package markup

import (
	"syscall/js"
)

// Element represents an HTML element.
type Element struct {
	element       js.Value
	eventHandlerID uint64
}

func NewElement(el js.Value, eventHandlerID uint64) (e *Element) {
	e = &Element{
		element:       el,
		eventHandlerID: eventHandlerID,
	}
	return
}

// JSValue returns the syscall/js element.
func (e *Element) JSValue() (jsValue js.Value) {
	jsValue = e.element
	return
}

// Is returns if the 2 elements e and check, have the same exact javascript value.
func (e *Element) Is(check *Element) (is bool) {
	is = e.element.Equal(check.element)
	return	
}

// ISJSValue returns if the 2 elements e and check, have the same exact javascript value.
func (e *Element) IsJSValue(check js.Value) (is bool) {
	is = e.element.Equal(check)
	return	
}

// TagName returns the element's tag name.
func (e *Element) TagName() (tagName string) {
	tagName = e.element.Get("tagName").String()
	return
}
`
