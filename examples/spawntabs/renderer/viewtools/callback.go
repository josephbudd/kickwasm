package viewtools

import (
	"syscall/js"

	"github.com/pkg/errors"
)

// UnRegisterCallBacks deletes the call backs for a panel.
// Param uniqueID is the spawn panel's unique id.
func (tools *Tools) UnRegisterCallBacks(uniqueID uint64) (err error) {
	if uniqueID == 0 {
		return
	}
	var funcs []js.Func
	var found bool
	if funcs, found = tools.jsCallBacks[uniqueID]; !found {
		err = errors.New("uniqueID not found in tools.jsCallBacks")
		return
	}
	for _, f := range funcs {
		f.Release()
	}
	delete(tools.jsCallBacks, uniqueID)
	return
}

// CloseCallBacks closes every registered call back.
func (tools *Tools) CloseCallBacks() (funcCount uint64) {
	for _, funcs := range tools.jsCallBacks {
		for _, f := range funcs {
			funcCount++
			f.Release()
		}
	}
	return
}

// RegisterCallBack converts a go func to a js.Func and registers it.
// Call this from your normal markup panels.
// Returns the call back as a js.Func
func (tools *Tools) RegisterCallBack(fn func(this js.Value, args []js.Value) interface{}) (cb js.Func) {
	// uniqueID for not spawns is 0.
	cb = tools.registerCallBack(fn, 0)
	return
}

// RegisterSpawnCallBack converts a go func to a js.Func and registers it.
// Call this from your spawned markup panels.
// Param uniqueID is the spawn panel's unique id.
// Returns the call back as a js.Func
func (tools *Tools) RegisterSpawnCallBack(fn func(this js.Value, args []js.Value) interface{}, uniqueID uint64) (cb js.Func) {
	cb = tools.registerCallBack(fn, uniqueID)
	return
}

// RegisterEventCallBack converts a go func to a js.Func and registers it.
// Call this from your normal markup panels.
// Param preventDefault indicates to call event.preventDefault synchronously.
// Param stopPropagation indicates to call event.stopPropagation synchronously.
// Param stopImmediatePropagation indicates to call event.stopImmediatePropagation synchronously.
// Param fn is a function that takes exactly one argument, the event.
func (tools *Tools) RegisterEventCallBack(fn func(event js.Value) interface{}, preventDefault, stopPropagation, stopImmediatePropagation bool) (cb js.Func) {
	cb = tools.registerEventCallBack(fn, preventDefault, stopPropagation, stopImmediatePropagation, 0)
	return
}

// RegisterSpawnEventCallBack converts a go func to a js.Func and registers it.
// Call this from your spawned markup panels.
// Param preventDefault indicates to call event.preventDefault synchronously.
// Param stopPropagation indicates to call event.stopPropagation synchronously.
// Param stopImmediatePropagation indicates to call event.stopImmediatePropagation synchronously.
// Param fn is a function that takes exactly one argument, the event.
// Param uniqueID is the spawn panel's unique id.
// Returns the call back as a js.Func
func (tools *Tools) RegisterSpawnEventCallBack(fn func(event js.Value) interface{}, preventDefault, stopPropagation, stopImmediatePropagation bool, uniqueID uint64) (cb js.Func) {
	cb = tools.registerEventCallBack(fn, preventDefault, stopPropagation, stopImmediatePropagation, uniqueID)
	return
}

// registerEventCallBack converts a go func to a js.Func and registers it.
// Param preventDefault indicates to call event.preventDefault synchronously.
// Param stopPropagation indicates to call event.stopPropagation synchronously.
// Param stopImmediatePropagation indicates to call event.stopImmediatePropagation synchronously.
// Param fn is a function that takes exactly one argument, the event.
// Param uniqueID is the spawn panel's unique id.
// Returns the call back as a js.Func
func (tools *Tools) registerEventCallBack(fn func(event js.Value) interface{}, preventDefault, stopPropagation, stopImmediatePropagation bool, uniqueID uint64) (cb js.Func) {
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
	cb = tools.registerCallBack(wrapperFn, uniqueID)
	return
}

func (tools *Tools) registerCallBack(f func(this js.Value, args []js.Value) interface{}, uniqueID uint64) (jsFunc js.Func) {
	var funcs []js.Func
	var found bool
	if funcs, found = tools.jsCallBacks[uniqueID]; !found {
		funcs = make([]js.Func, 0, 20)
	}
	jsFunc = js.FuncOf(f)
	funcs = append(funcs, jsFunc)
	tools.jsCallBacks[uniqueID] = funcs
	return
}
