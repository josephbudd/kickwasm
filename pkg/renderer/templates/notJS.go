package templates

// NotJSNotJSGo is the file renderer/notjs/notJS.go
var NotJSNotJSGo = `package notjs

import "syscall/js"

var (
	count int
)

// NotJS is a js notjs.
type NotJS struct {
	global   js.Value
	document js.Value
	console  js.Value
	alert    js.Value

	// call backs
	jsCallBacks []js.Func
}

// NewNotJS constructs a new NotJS.
func NewNotJS() *NotJS {
	if count > 0 {
		panic("Tried to construct more than 1 NotJS.")
	}
	count++
	g := js.Global()
	return &NotJS{
		global:      g,
		document:    g.Get("document"),
		alert:       g.Get("alert"),
		console:     g.Get("console"),
		jsCallBacks: make([]js.Func, 0, 100),
	}
}

`
