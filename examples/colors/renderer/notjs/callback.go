package notjs

import (
	"syscall/js"
)

// RegisterCallBack converts a go func to a js.Func and registers it.
func (notjs *NotJS) RegisterCallBack(fn func(this js.Value, args []js.Value) interface{}) (cb js.Func) {
	cb = js.FuncOf(fn)
	notjs.jsCallBacks = append(notjs.jsCallBacks, cb)
	return
}

// RegisterEventCallBack converts a go func to a js.Func and registers it.
// Param preventDefault indicates to call event.preventDefault synchronously.
// Param stopPropagation indicates to call event.stopPropagation synchronously.
// Param stopImmediatePropagation indicates to call event.stopImmediatePropagation synchronously.
// Param fn is a function that takes exactly one argument, the event.
func (notjs *NotJS) RegisterEventCallBack(preventDefault, stopPropagation, stopImmediatePropogation bool, fn func(event js.Value)) (cb js.Func) {
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
			if stopImmediatePropogation {
				event.Call("stopImmediatePropogation")
			}
		}
		fn(event)
		return nil
	}
	cb = js.FuncOf(wrapperFn)
	notjs.jsCallBacks = append(notjs.jsCallBacks, cb)
	return
}

// CloseCallBacks closes every registered call back.
func (notjs *NotJS) CloseCallBacks() int {
	var i int
	var cb js.Func
	for i, cb = range notjs.jsCallBacks {
		cb.Release()
	}
	return i + 1
}

