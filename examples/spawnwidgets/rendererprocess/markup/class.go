// +build js, wasm

package markup

// Class

// ContainsClass return if a class list contains a class.
func (e *Element) ContainsClass(class string) (contains bool) {
	classList := e.element.Get(classListAttributeName)
	contains = classList.Call(containsMethodName, class).Bool()
	return
}

// ContainsAnyClass returns if a class list contains any of the classes.
func (e *Element) ContainsAnyClass(classes ...string) (contains bool) {
	classList := e.element.Get(classListAttributeName)
	for _, class := range classes {
		if contains = classList.Call(containsMethodName, class).Bool(); contains {
			return
		}
	}
	return
}

// ContainsAllClasses returns if a class list contains all of the classes.
func (e *Element) ContainsAllClasses(classes ...string) (contains bool) {
	classList := e.element.Get(classListAttributeName)
	for _, class := range classes {
		if !classList.Call(containsMethodName, class).Bool() {
			return
		}
	}
	contains = true
	return
}

// AddClass adds an element's class.
func (e *Element) AddClass(class string) {
	classList := e.element.Get(classListAttributeName)
	classList.Call(addMethodName, class)
}

// RemoveClass removes an element's class.
func (e *Element) RemoveClass(class string) {
	classList := e.element.Get(classListAttributeName)
	classList.Call("remove", class)
}

// ReplaceClass replaces or adds an element's class.
func (e *Element) ReplaceClass(oldClass, newClass string) {
	classList := e.element.Get(classListAttributeName)
	if classList.Call("replace", oldClass, newClass).Bool() {
		return
	}
	if !classList.Call(containsMethodName, newClass).Bool() {
		classList.Call(addMethodName, newClass)
	}
}

// ClassAt returns the class in classList at index.
func (e *Element) ClassAt(index int) (attribute string) {
	classList := e.element.Get(classListAttributeName)
	attribute = classList.Call(itemMethodName, index).String()
	return
}

// AddClasses adds an element's classes.
func (e *Element) AddClasses(classes ...string) {
	classList := e.element.Get(classListAttributeName)
	for _, class := range classes {
		classList.Call(addMethodName, class)
	}
}

// RemoveClass removes an element's classes.
func (e *Element) RemoveClasses(classes ...string) {
	classList := e.element.Get(classListAttributeName)
	for _, class := range classes {
		classList.Call("remove", class)
	}
}
