// +build js, wasm

package callback

import (
	"fmt"
	"math"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/event"
)


var (
	jsCallBacks = make(map[uint64][]js.Func, 100)
	global      = js.Global()
	undefined   = js.Undefined()
)

const (
	applicationID = math.MaxUint64
)

// SetApplicationOnCloseHandler sets the window.onclose handler.
// window.onclose is treated differently than other events.
func SetApplicationOnCloseHandler(handler func(event.Event) interface{}) {
	wrapperFn := func(this js.Value, args []js.Value) (nilReturn interface{}) {
		var eventJS js.Value
		if len(args) > 0 {
			eventJS = args[0]
		} else {
			eventJS = this
		}
		channelEvent := event.BuildEvent(eventJS, global, "close", false, applicationID)
		nilReturn = handler(channelEvent)
		return
	}
	cb := registerCallBack(wrapperFn, applicationID)
	global.Set("onclose", cb)
}

// RemoveApplicationOnCloseHandler
// * unsets the window.onclose
// * unsets window.onbeforeunload which is actually set in index.html javascript.
// Call this to close the application automatically.
// Follow a call to RemoveApplicationOnCloseHandler with quitCh <-struct{}{}.
func RemoveApplicationOnCloseHandler() {
	UnRegisterCallBacks(applicationID)
	global.Set("onclose", undefined)
	global.Set("onbeforeunload", undefined)
}

// AddEventHandler adds an event handler to the element.
// Param handler is the element's event handler.
// Param element is the target element.
// Param on is the event name. ex: "click".
// Param capturePhase indicates if this happens during the capture phase.
// Param panelUniqueID is the spawn's unique id.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func AddEventHandler(handler func(event.Event) interface{}, element js.Value, on string, capturePhase bool, panelUniqueID uint64) {
	handleEvent(handler, element, on, capturePhase, panelUniqueID)
	return
}

func handleEvent(handler func(event.Event) interface{}, element js.Value, on string, capturePhase bool, panelUniqueID uint64) {
	wrapperFn := func(this js.Value, args []js.Value) (nilReturn interface{}) {
		var eventJS js.Value
		if len(args) > 0 {
			eventJS = args[0]
		} else {
			eventJS = this
		}
		channelEvent := event.BuildEvent(eventJS, eventJS.Get("target"), on, capturePhase, panelUniqueID)
		nilReturn = handler(channelEvent)
		return
	}
	cb := registerCallBack(wrapperFn, panelUniqueID)
	element.Call("addEventListener", on, cb, capturePhase)
}

// UnRegisterCallBacks deletes the call backs for a panel.
// Param panelUniqueID is the spawn panel's unique id.
func UnRegisterCallBacks(panelUniqueID uint64) (err error) {
	if panelUniqueID == 0 {
		return
	}
	var funcs []js.Func
	var found bool
	if funcs, found = jsCallBacks[panelUniqueID]; !found {
		err = fmt.Errorf("panelUniqueID not found in jsCallBacks")
		return
	}
	for _, f := range funcs {
		f.Release()
	}
	delete(jsCallBacks, panelUniqueID)
	return
}

// CloseCallBacks closes every registered call back.
func CloseCallBacks() (funcCount uint64) {
	for _, funcs := range jsCallBacks {
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
func RegisterCallBack(fn func(this js.Value, args []js.Value) interface{}) (cb js.Func) {
	// panelUniqueID for not spawns is 0.
	cb = registerCallBack(fn, 0)
	return
}

func registerCallBack(f func(this js.Value, args []js.Value) interface{}, panelUniqueID uint64) (jsFunc js.Func) {
	var funcs []js.Func
	var found bool
	if funcs, found = jsCallBacks[panelUniqueID]; !found {
		funcs = make([]js.Func, 0, 20)
	}
	jsFunc = js.FuncOf(f)
	funcs = append(funcs, jsFunc)
	jsCallBacks[panelUniqueID] = funcs
	return
}
