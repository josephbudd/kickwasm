package notjs

import "syscall/js"

// SetInnerText sets an element's innerText
func (notjs *NotJS) SetInnerText(element js.Value, text string) {
	element.Set(innerTextMemberName, text)
}

// SetInnerHTML sets an element's innerHTML
func (notjs *NotJS) SetInnerHTML(element js.Value, html string) {
	element.Set(innerHTMLMemberName, html)
}

// GetInnerText sets an element's innerText
func (notjs *NotJS) GetInnerText(element js.Value) string {
	return element.Get(innerTextMemberName).String()
}

// GetInnerHTML sets an element's innerHTML
func (notjs *NotJS) GetInnerHTML(element js.Value) string {
	return element.Get(innerHTMLMemberName).String()
}

// GetOuterHTML sets an element's outerHTML
func (notjs *NotJS) GetOuterHTML(element js.Value) string {
	return element.Get(outerHTMLMemberName).String()
}
