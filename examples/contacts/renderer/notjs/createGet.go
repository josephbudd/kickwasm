package notjs

import "syscall/js"

// GetElementByID inovokes the document;s getElementById.
func (notJS *NotJS) GetElementByID(id string) js.Value {
	return notJS.document.Call("getElementById", id)
}

// GetElementsByTagName invokes the document's getElementsByTagName.
func (notJS *NotJS) GetElementsByTagName(tagName string) []js.Value {
	els := notJS.document.Call("getElementsByTagName", tagName)
	l := els.Length()
	tagNames := make([]js.Value, l, l)
	for i := 0; i < l; i++ {
		tagNames[i] = els.Index(i)
	}
	return tagNames
}

// CreateElement invokes the document's createElement.
func (notJS *NotJS) CreateElement(tagName string) js.Value {
	return notJS.document.Call(createElementMethodName, tagName)
}

// CreateElementInput creates a text input element.
func (notJS *NotJS) CreateElementInput() js.Value {
	input := notJS.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, "text")
	return input
}

// CreateElementCheckBoxInGroup creates a checkbox input element.
func (notJS *NotJS) CreateElementCheckBoxInGroup(group string) js.Value {
	input := notJS.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, checkboxTypeName)
	if len(group) > 0 {
		input.Set(groupAttributeName, group)
	}
	return input
}

// CreateElementRadioInGroup creates a radio input element
func (notJS *NotJS) CreateElementRadioInGroup(group string) js.Value {
	input := notJS.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, radioTypeName)
	input.Set(groupAttributeName, group)
	return input
}

// CreateElementCheckBox creates a checkbox input element.
func (notJS *NotJS) CreateElementCheckBox() js.Value {
	input := notJS.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, checkboxTypeName)
	return input
}

// CreateElementRadio creates a radio input element
func (notJS *NotJS) CreateElementRadio() js.Value {
	input := notJS.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, radioTypeName)
	return input
}

// CreateElementSpan invokes the document's createElement.
func (notJS *NotJS) CreateElementSpan() js.Value {
	return notJS.document.Call(createElementMethodName, "span")
}

// CreateElementB invokes the document's createElement.
func (notJS *NotJS) CreateElementB() js.Value {
	return notJS.document.Call(createElementMethodName, "b")
}

// CreateElementBR invokes the document's createElement.
func (notJS *NotJS) CreateElementBR() js.Value {
	return notJS.document.Call(createElementMethodName, "br")
}

// CreateElementP invokes the document's createElement.
func (notJS *NotJS) CreateElementP() js.Value {
	return notJS.document.Call(createElementMethodName, "p")
}

// CreateElementBUTTON invokes the document's createElement.
func (notJS *NotJS) CreateElementBUTTON() js.Value {
	return notJS.document.Call(createElementMethodName, "button")
}

// CreateElementDIV invokes the document's createElement.
func (notJS *NotJS) CreateElementDIV() js.Value {
	return notJS.document.Call(createElementMethodName, "div")
}

// CreateElementH3 invokes the document's createElement.
func (notJS *NotJS) CreateElementH3() js.Value {
	return notJS.document.Call(createElementMethodName, "h3")
}

// CreateElementH4 invokes the document's createElement.
func (notJS *NotJS) CreateElementH4() js.Value {
	return notJS.document.Call(createElementMethodName, "h4")
}

// CreateElementH5 invokes the document's createElement.
func (notJS *NotJS) CreateElementH5() js.Value {
	return notJS.document.Call(createElementMethodName, "h5")
}

// CreateElementH6 invokes the document's createElement.
func (notJS *NotJS) CreateElementH6() js.Value {
	return notJS.document.Call(createElementMethodName, "h6")
}

// CreateTextNode invokes the document's createElement.
func (notJS *NotJS) CreateTextNode(text string) js.Value {
	return notJS.document.Call("createTextNode", text)
}

// CreateElementTH invokes the document's createElement.
func (notJS *NotJS) CreateElementTH() js.Value {
	return notJS.document.Call(createElementMethodName, "th")
}

// CreateElementTR invokes the document's createElement.
func (notJS *NotJS) CreateElementTR() js.Value {
	return notJS.document.Call(createElementMethodName, "tr")
}

// CreateElementTD invokes the document's createElement.
func (notJS *NotJS) CreateElementTD() js.Value {
	return notJS.document.Call(createElementMethodName, "td")
}
