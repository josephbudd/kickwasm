// +build js, wasm

package markup

// Parent Child

// Parent returns the element's parent or nil.
func (e *Element) Parent() (parent *Element) {
	p := e.element.Get("parentNode")
	if p == null {
		return
	}
	parent = &Element{
		element:       p,
		panelUniqueID: e.panelUniqueID,
	}

	return
}

// Children returns the child nodes as a slice of *Elements.
func (e *Element) Children() (children []*Element) {
	chch := e.element.Get(childrenAttributeName)
	l := chch.Length()
	children = make([]*Element, l)
	for i := 0; i < l; i++ {
		ch := chch.Call(itemMethodName, i)
		children[i] = &Element{
			element:       ch,
			panelUniqueID: e.panelUniqueID,
		}
	}
	return
}

// AppendChild appends a child to a parent.
func (e *Element) AppendChild(child *Element) {
	e.element.Call("appendChild", child.element)
}

// AppendText appends a child text node to a parent.
func (e *Element) AppendText(text string) {
	tn := document.Call("createTextNode", text)
	e.element.Call("appendChild", tn)
}

// RemoveChild removes child from parent.
func (e *Element) RemoveChild(child *Element) {
	e.element.Call("removeChild", child.element)
}

// RemoveChildren removes every child node from a parent.
func (e *Element) RemoveChildren() {
	e.element.Set(innerHTMLMemberName, "")
}

// InsertChildBefore inserts newChild before targetChild in parent.
func (e *Element) InsertChildBefore(newChild, targetChild *Element) {
	e.element.Call("insertBefore", newChild.element, targetChild.element)
}
