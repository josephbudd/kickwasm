package notjs

import (
	"syscall/js"
)

// ParentNode returns an element's parent node.
func (notJS *NotJS) ParentNode(child js.Value) js.Value {
	return child.Get("parentNode")
}

// Children returns an element's children and children length
func (notJS *NotJS) Children(parent js.Value) (children js.Value, length int) {
	children = parent.Get(childrenAttributeName)
	length = children.Length()
	return
}

// ChildrenSlice returns a slice of an element's children
func (notJS *NotJS) ChildrenSlice(parent js.Value) []js.Value {
	children := parent.Get(childrenAttributeName)
	l := children.Length()
	slice := make([]js.Value, l, l)
	for i := 0; i < l; i++ {
		slice[i] = children.Call(itemMethodName, i)
	}
	return slice
}

// FirstChild returns the first child of parent.
func (notJS *NotJS) FirstChild(parent js.Value) js.Value {
	return parent.Get("firstChild")
}

// LastChild returns the last child of parent.
func (notJS *NotJS) LastChild(parent js.Value) js.Value {
	return parent.Get("lastChild")
}

// AppendChild appends a child to a parent.
func (notJS *NotJS) AppendChild(parent, child js.Value) {
	parent.Call("appendChild", child)
}

// RemoveChild removes child from parent.
func (notJS *NotJS) RemoveChild(parent, child js.Value) {
	parent.Call("removeChild", child)
}

// RemoveChildNodes removes every child node from a parent.
func (notJS *NotJS) RemoveChildNodes(parent js.Value) {
	children := notJS.ChildrenSlice(parent)
	for _, ch := range children {
		notJS.RemoveChild(parent, ch)
	}
}

// InsertChildBefore inserts newChild before targetChild in parent.
func (notJS *NotJS) InsertChildBefore(parent, newChild, targetChild js.Value) {
	parent.Call("insertBefore", newChild, targetChild)
}
