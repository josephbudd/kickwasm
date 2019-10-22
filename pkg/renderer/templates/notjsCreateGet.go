package templates

// NotJSCreateGetGo is the file renderer/notjs/createGet.go
const NotJSCreateGetGo = `{{ $Dot := . }}// +build js, wasm

package notjs

import "syscall/js"

// GET

// GetElementByID inovokes the document;s getElementById.
func (notjs *NotJS) GetElementByID(id string) js.Value {
	return notjs.document.Call("getElementById", id)
}

// GetElementsByTagName invokes the document's getElementsByTagName.
func (notjs *NotJS) GetElementsByTagName(tagName string) []js.Value {
	els := notjs.document.Call("getElementsByTagName", tagName)
	l := els.Length()
	tagNames := make([]js.Value, l, l)
	for i := 0; i < l; i++ {
		tagNames[i] = els.Index(i)
	}
	return tagNames
}

// VARIOUS CREATES

// CreateTextNode invokes the document's createElement.
func (notjs *NotJS) CreateTextNode(text string) js.Value {
	return notjs.document.Call("createTextNode", text)
}

// CreateElement invokes the document's createElement.
func (notjs *NotJS) CreateElement(tagName string) js.Value {
	return notjs.document.Call(createElementMethodName, tagName)
}

// CREATE FORM INPUT VARIATIONS.

// CreateElementINPUT creates a text input element.
func (notjs *NotJS) CreateElementINPUT() js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, "text")
	return input
}

// CreateElementCheckBoxInGroup creates a checkbox input element.
func (notjs *NotJS) CreateElementCheckBoxInGroup(group string) js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, checkboxTypeName)
	if len(group) > 0 {
		input.Set(groupAttributeName, group)
	}
	return input
}

// CreateElementRadioInGroup creates a radio input element
func (notjs *NotJS) CreateElementRadioInGroup(group string) js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, radioTypeName)
	input.Set(groupAttributeName, group)
	return input
}

// CreateElementCheckBox creates a checkbox input element.
func (notjs *NotJS) CreateElementCheckBox() js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, checkboxTypeName)
	return input
}

// CreateElementRadio creates a radio input element
func (notjs *NotJS) CreateElementRadio() js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, radioTypeName)
	return input
}{{ range .HTMLNames}}

// CreateElement{{ call $Dot.ToUpper . }} invokes the document's createElement.
func (notjs *NotJS) CreateElement{{ call $Dot.ToUpper . }}() js.Value {
	return notjs.document.Call(createElementMethodName, "{{ call $Dot.ToLower . }}")
}{{ end }}
`
