// +build js, wasm

package markup

// SetInnerText sets an element's innerText
func (e *Element) SetInnerText(text string) {
	e.element.Set(innerTextMemberName, text)
}

// SetInnerHTML sets an element's innerHTML
func (e *Element) SetInnerHTML(html string) {
	e.element.Set(innerHTMLMemberName, html)
}

// InnerText returns an element's innerText
// InnerText
func (e *Element) InnerText() (text string) {
	text = e.element.Get(innerTextMemberName).String()
	return
}

// InnerHTML returns an element's innerHTML
// InnerHTML
func (e *Element) InnerHTML() (html string) {
	html = e.element.Get(innerHTMLMemberName).String()
	return
}

// OuterHTML returns an element's outerHTML
// OuterHTML
func (e *Element) OuterHTML() (html string) {
	html = e.element.Get(outerHTMLMemberName).String()
	return
}
