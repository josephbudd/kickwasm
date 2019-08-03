package notjs

import "syscall/js"

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

// CreateElement invokes the document's createElement.
func (notjs *NotJS) CreateElement(tagName string) js.Value {
	return notjs.document.Call(createElementMethodName, tagName)
}

// CreateElementInput creates a text input element.
func (notjs *NotJS) CreateElementInput() js.Value {
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
}

// CreateElementSpan invokes the document's createElement.
func (notjs *NotJS) CreateElementSpan() js.Value {
	return notjs.document.Call(createElementMethodName, "span")
}

// CreateElementB invokes the document's createElement.
func (notjs *NotJS) CreateElementB() js.Value {
	return notjs.document.Call(createElementMethodName, "b")
}

// CreateElementBR invokes the document's createElement.
func (notjs *NotJS) CreateElementBR() js.Value {
	return notjs.document.Call(createElementMethodName, "br")
}

// CreateElementP invokes the document's createElement.
func (notjs *NotJS) CreateElementP() js.Value {
	return notjs.document.Call(createElementMethodName, "p")
}

// CreateElementBUTTON invokes the document's createElement.
func (notjs *NotJS) CreateElementBUTTON() js.Value {
	return notjs.document.Call(createElementMethodName, "button")
}

// CreateElementDIV invokes the document's createElement.
func (notjs *NotJS) CreateElementDIV() js.Value {
	return notjs.document.Call(createElementMethodName, "div")
}

// CreateElementH3 invokes the document's createElement.
func (notjs *NotJS) CreateElementH3() js.Value {
	return notjs.document.Call(createElementMethodName, "h3")
}

// CreateElementH4 invokes the document's createElement.
func (notjs *NotJS) CreateElementH4() js.Value {
	return notjs.document.Call(createElementMethodName, "h4")
}

// CreateElementH5 invokes the document's createElement.
func (notjs *NotJS) CreateElementH5() js.Value {
	return notjs.document.Call(createElementMethodName, "h5")
}

// CreateElementH6 invokes the document's createElement.
func (notjs *NotJS) CreateElementH6() js.Value {
	return notjs.document.Call(createElementMethodName, "h6")
}

// CreateTextNode invokes the document's createElement.
func (notjs *NotJS) CreateTextNode(text string) js.Value {
	return notjs.document.Call("createTextNode", text)
}

// CreateElementTH invokes the document's createElement.
func (notjs *NotJS) CreateElementTH() js.Value {
	return notjs.document.Call(createElementMethodName, "th")
}

// CreateElementTR invokes the document's createElement.
func (notjs *NotJS) CreateElementTR() js.Value {
	return notjs.document.Call(createElementMethodName, "tr")
}

// CreateElementTD invokes the document's createElement.
func (notjs *NotJS) CreateElementTD() js.Value {
	return notjs.document.Call(createElementMethodName, "td")
}
