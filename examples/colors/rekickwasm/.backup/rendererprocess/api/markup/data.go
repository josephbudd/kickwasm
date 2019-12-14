// +build js, wasm

package markup

import (
	"syscall/js"
)

var (
	global      js.Value
	document    js.Value
	alert       js.Value
	null        js.Value
	jsCallBacks map[uint64][]js.Func

	idAttributeName        = "id"
	checkedAttributeName   = "checked"
	childrenAttributeName  = "children"
	classListAttributeName = "classList"
	groupAttributeName     = "group"
	typeAttributeName      = "type"
	valueAttributeName     = "value"

	addMethodName           = "add"
	containsMethodName      = "contains"
	itemMethodName          = "item"
	getAttributeMethodName  = "getAttribute"
	setAttributeMethodName  = "setAttribute"

	innerHTMLMemberName = "innerHTML"
	outerHTMLMemberName = "outerHTML"
	innerTextMemberName = "innerText"

	inputTypeName    = "input"
	checkboxTypeName = "checkbox"
	radioTypeName    = "radio"

	hVScrollClassName = "{{.HVScrollClassName}}"
	resizeMeWidthClassName = "{{.ResizeMeWidthClassName}}"
)

func init() {
	global = js.Global()
	document = global.Get("document")
	alert = global.Get("alert")
	null = js.Null()
	jsCallBacks = make(map[uint64][]js.Func, 100)
}
