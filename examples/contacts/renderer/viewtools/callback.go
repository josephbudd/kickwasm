package viewtools

import (
	"syscall/js"
)

// RegisterCallBack converts a go func to a js.Func and registers it.
func (tools *Tools) RegisterCallBack(fn func(this js.Value, args []js.Value) interface{}) (cb js.Func) {
	cb = js.FuncOf(fn)
	tools.jsCallBacks = append(tools.jsCallBacks, cb)
	return
}

// RegisterEventCallBack converts a go func to a js.Func and registers it.
// Param preventDefault indicates to call event.preventDefault synchronously.
// Param stopPropagation indicates to call event.stopPropagation synchronously.
// Param stopImmediatePropagation indicates to call event.stopImmediatePropagation synchronously.
// Param fn is a function that takes exactly one argument, the event.
func (tools *Tools) RegisterEventCallBack(fn func(event js.Value) interface{}, preventDefault, stopPropagation, stopImmediatePropagation bool) (cb js.Func) {
	wrapperFn := func(this js.Value, args []js.Value) interface{} {
		var event js.Value
		if len(args) > 0 {
			event = args[0]
		} else {
			event = this
		}
		if event != js.Undefined() && event != js.Null() {
			if preventDefault {
				event.Call("preventDefault")
			}
			if stopPropagation {
				event.Call("stopPropagation")
			}
			if stopImmediatePropagation {
				event.Call("stopImmediatePropagation")
			}
		}
		return fn(event)
	}
	cb = js.FuncOf(wrapperFn)
	tools.jsCallBacks = append(tools.jsCallBacks, cb)
	return
}

// CloseCallBacks closes every registered call back.
func (tools *Tools) CloseCallBacks() int {
	var i int
	var cb js.Func
	for i, cb = range tools.jsCallBacks {
		cb.Release()
	}
	return i + 1
}

