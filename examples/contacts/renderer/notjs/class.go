package notjs

import "syscall/js"

// ClassListContains return if a class list contains a class.
func (notjs *NotJS) ClassListContains(element js.Value, class string) bool {
	classList := element.Get(classListAttributeName)
	return classList.Call(containsMethodName, class).Bool()
}

// ClassListContainsAnd returns if a class list contains all of the classes.
func (notjs *NotJS) ClassListContainsAnd(element js.Value, classes ...string) bool {
	classList := element.Get(classListAttributeName)
	for _, class := range classes {
		if !classList.Call(containsMethodName, class).Bool() {
			return false
		}
	}
	return true
}

// ClassListContainsOr returns if a class list contains any of the classes.
func (notjs *NotJS) ClassListContainsOr(element js.Value, classes ...string) bool {
	classList := element.Get(classListAttributeName)
	for _, class := range classes {
		if classList.Call(containsMethodName, class).Bool() {
			return true
		}
	}
	return false
}

// ClassListReplaceClass replaces or adds an element's class.
func (notjs *NotJS) ClassListReplaceClass(element js.Value, old, new string) {
	classList := element.Get(classListAttributeName)
	if classList.Call("replace", old, new).Bool() {
		return
	}
	if !classList.Call(containsMethodName, new).Bool() {
		classList.Call(addMethodName, new)
	}
}

// ClassListGetClassAt returns the class in classList at index.
func (notjs *NotJS) ClassListGetClassAt(element js.Value, index int) string {
	classList := element.Get(classListAttributeName)
	return classList.Call(itemMethodName, index).String()
}

// ClassListAddClass adds an element's class.
func (notjs *NotJS) ClassListAddClass(element js.Value, new string) {
	classList := element.Get(classListAttributeName)
	classList.Call(addMethodName, new)
}
