package templates

// ViewtoolsHelpers is the renderer/viewtools/helpers.go file.
const ViewtoolsHelpers = `package viewtools

import (
	"fmt"
	"syscall/js"
)

// ConsoleLog logs to the console.
func (tools *Tools) ConsoleLog(message string) {
	tools.console.Call("log", message)
}

// Alert inokes the browser's alert.
func (tools *Tools) Alert(message string) {
	tools.alert.Invoke(message)
}

// Success displays a message titled "Success"
func (tools *Tools) Success(message string) {
	tools.GoModal(message, "Success", nil)
}

// Error displays a message titled "Error"
func (tools *Tools) Error(message string) {
	tools.GoModal(message, "Error", nil)
}

// GetElementByID inovokes the document;s getElementById.
func (tools *Tools) GetElementByID(id string) js.Value {
	return tools.Document.Call("getElementById", id)
}

// GetElementsByTagName invokes the document's getElementsByTagName.
func (tools *Tools) GetElementsByTagName(tagName string) []js.Value {
	els := tools.Document.Call("getElementsByTagName", tagName)
	l := els.Length()
	tagNames := make([]js.Value, l, l)
	for i := 0; i < l; i++ {
		tagNames[i] = els.Index(i)
	}
	return tagNames
}

// CreateElement invokes the document's createElement.
func (tools *Tools) CreateElement(tagName string) js.Value {
	return tools.Document.Call("createElement", tagName)
}

// CreateTextNode invokes the document's createElement.
func (tools *Tools) CreateTextNode(text string) js.Value {
	return tools.Document.Call("createTextNode", text)
}

// Children returns an element's children and children length
func (tools *Tools) Children(parent js.Value) (children js.Value, length int) {
	children = parent.Get("children")
	length = children.Length()
	return
}

// ChildrenSlice returns a slice of an element's children
func (tools *Tools) ChildrenSlice(parent js.Value) []js.Value {
	children := parent.Get("children")
	l := children.Length()
	slice := make([]js.Value, l, l)
	for i := 0; i < l; i++ {
		slice[i] = children.Call("item", i)
	}
	return slice
}

// AppendChild appends a child to a parent.
func (tools *Tools) AppendChild(parent, child js.Value) {
	parent.Call("appendChild", child)
}

// ParentNode returns an element's parent node.
func (tools *Tools) ParentNode(child js.Value) js.Value {
	return child.Get("parentNode")
}

// TagName returns an element's tab name.
func (tools *Tools) TagName(el js.Value) string {
	return el.Get("tagName").String()
}

// ID returns an element's id.
func (tools *Tools) ID(el js.Value) string {
	return el.Get("id").String()
}

// ClassListContains return if a class list contains a class.
func (tools *Tools) ClassListContains(element js.Value, class string) bool {
	classList := element.Get("classList")
	return classList.Call("contains", class).Bool()
}

// ClassListContainsAnd returns if a class list contains all of the classes.
func (tools *Tools) ClassListContainsAnd(element js.Value, classes ...string) bool {
	classList := element.Get("classList")
	for _, class := range classes {
		if !classList.Call("contains", class).Bool() {
			return false
		}
	}
	return true
}

// ClassListContainsOr returns if a class list contains any of the classes.
func (tools *Tools) ClassListContainsOr(element js.Value, classes ...string) bool {
	classList := element.Get("classList")
	for _, class := range classes {
		if classList.Call("contains", class).Bool() {
			return true
		}
	}
	return false
}

// ClassListReplaceClass replaces or adds an element's class.
func (tools *Tools) ClassListReplaceClass(element js.Value, old, new string) {
	classList := element.Get("classList")
	if classList.Call("replace", old, new).Bool() {
		return
	}
	if !classList.Call("contains", new).Bool() {
		classList.Call("add", new)
	}
}

// ClassListGetClassAt returns the class in classList at index.
func (tools *Tools) ClassListGetClassAt(element js.Value, index int) string {
	classList := element.Get("classList")
	return classList.Call("item", index).String()
}

// ClassListAddClass adds an element's class.
func (tools *Tools) ClassListAddClass(element js.Value, new string) {
	classList := element.Get("classList")
	classList.Call("add", new)
}

// SetStyleHeight sets an element's style height.
func (tools *Tools) SetStyleHeight(element js.Value, height int) {
	style := element.Get("style")
	style.Set("height", fmt.Sprintf("%dpx", height))
}

// SetStyleWidth sets an element's style width.
func (tools *Tools) SetStyleWidth(element js.Value, width int) {
	style := element.Get("style")
	style.Set("width", fmt.Sprintf("%dpx", width))
}

// c sets an element's href attribute
func (tools *Tools) SetAttributeHref(element js.Value, value string) {
	element.Call("setAttribute", "href", value)
}

// GetEventTarget gets an event's target attribute which is an html element.
func (tools *Tools) GetEventTarget(event js.Value) js.Value {
	return event.Get("target")
}

// SetAttribute sets an element's href attribute
func (tools *Tools) SetAttribute(element js.Value, name, value string) {
	element.Call("setAttribute", name, value)
}

// GetAttribute sets an element's href attribute
func (tools *Tools) GetAttribute(element js.Value, name string) string {
	return element.Call("getAttribute", name).String()
}

// SetOnClick sets an element's onclick
func (tools *Tools) SetOnClick(element js.Value, cb js.Callback) {
	element.Set("onclick", cb)
}

// SetOnChange sets an element's onchange
func (tools *Tools) SetOnChange(element js.Value, cb js.Callback) {
	element.Set("onchange", cb)
}

// SetInnerText sets an element's innerText
func (tools *Tools) SetInnerText(element js.Value, text string) {
	element.Set("innerText", text)
}

// SetInnerHTML sets an element's innerHTML
func (tools *Tools) SetInnerHTML(element js.Value, html string) {
	element.Set("innerHtml", html)
}

// SetValue sets an element's value
func (tools *Tools) SetValue(element js.Value, value string) {
	element.Set("value", value)
}

// GetValue gets an element's value as a string.
func (tools *Tools) GetValue(element js.Value) string {
	return element.Get("value").String()
}

// GetValueInt gets an element's value as an int.
func (tools *Tools) GetValueInt(element js.Value) int {
	return element.Get("value").Int()
}

// GetValueFloat64 gets an element's value
func (tools *Tools) GetValueFloat64(element js.Value) float64 {
	return element.Get("value").Float()
}
`
