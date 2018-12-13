package notjs

import (
	"syscall/js"
)

// RegisterCallBack converts a go func to a js.Callback and registers it.
func (notJS *NotJS) RegisterCallBack(f func([]js.Value)) js.Callback {
	cb := js.NewCallback(f)
	notJS.jsCallBacks = append(notJS.jsCallBacks, cb)
	return cb
}

// RegisterEventCallBack converts a go func to a js.Callback and registers it.
// Param preventDefault indicates to call event.preventDefault synchronously.
// Param stopPropagation indicates to call event.stopPropagation synchronously.
// Param stopImmediatePropagation indicates to call event.stopImmediatePropagation synchronously.
// Param fn is a function that takes exactly one argument, the event.
func (notJS *NotJS) RegisterEventCallBack(preventDefault, stopPropogation, stopImmediatePropogation bool, fn func(event js.Value)) js.Callback {
	flags := js.EventCallbackFlag(0)
	if preventDefault {
		flags |= js.PreventDefault
	}
	if stopPropogation {
		flags |= js.StopPropagation
	}
	if stopImmediatePropogation {
		flags |= js.StopImmediatePropagation
	}
	cb := js.NewEventCallback(flags, fn)
	notJS.jsCallBacks = append(notJS.jsCallBacks, cb)
	return cb
}

// CloseCallBacks closes every registered call back.
func (notJS *NotJS) CloseCallBacks() int {
	var i int
	var cb js.Callback
	for i, cb = range notJS.jsCallBacks {
		cb.Release()
	}
	return i + 1
}
