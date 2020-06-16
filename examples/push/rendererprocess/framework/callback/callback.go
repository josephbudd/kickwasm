// +build js, wasm

package callback

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/event"
)


var (
	jsCallBacks = make(map[uint64][]js.Func, 100)
	global      = js.Global()
	undefined   = js.Undefined()
	eventHandlerID = uint64(1)
)

// ApplicationEventHandlerID is the id for all permanent event handlers.
const (
	ApplicationEventHandlerID = uint64(0)
)

// NewEventHandlerID returns a new event handler id for a spawned panel, widget or temporary widget.
func NewEventHandlerID() (id uint64) {
	id = eventHandlerID
	eventHandlerID++
	return
}

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
		channelEvent := event.BuildEvent(eventJS, global)
		nilReturn = handler(channelEvent)
		return
	}
	cb := registerCallBack(wrapperFn, ApplicationEventHandlerID)
	global.Set("onclose", cb)
}

// RemoveApplicationOnCloseHandler
// * unsets the window.onclose
// * unsets window.onbeforeunload which is actually set in index.html javascript.
// Call this to close the application automatically.
// Follow a call to RemoveApplicationOnCloseHandler with quitCh <-struct{}{}.
func RemoveApplicationOnCloseHandler() {
	UnRegisterCallBacks(ApplicationEventHandlerID)
	global.Set("onclose", undefined)
	global.Set("onbeforeunload", undefined)
}

// AddEventHandler adds an event handler to the element.
// Param handler is the element's event handler.
// Param element is the target element.
// Param on is the event name. ex: "click".
// Param capturePhase indicates if this happens during the capture phase.
// Param eventHandlerID is the event handler id for the spawn panel or spawn widget or temporary widget.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func AddEventHandler(handler func(event.Event) interface{}, element js.Value, on string, capturePhase bool, eventHandlerID uint64) {
	handleEvent(handler, element, on, capturePhase, eventHandlerID)
	return
}

func handleEvent(handler func(event.Event) interface{}, element js.Value, on string, capturePhase bool, eventHandlerID uint64) {
	wrapperFn := func(this js.Value, args []js.Value) (nilReturn interface{}) {
		var eventJS js.Value
		if len(args) > 0 {
			eventJS = args[0]
		} else {
			eventJS = this
		}
		channelEvent := event.BuildEvent(eventJS, eventJS.Get("target"))
		nilReturn = handler(channelEvent)
		return
	}
	cb := registerCallBack(wrapperFn, eventHandlerID)
	element.Call("addEventListener", on, cb, capturePhase)
}

// UnRegisterCallBacks deletes the call backs for a panel.
// Param eventHandlerID is event handler id for the spawn panel or spawn widget or temporary widget.
func UnRegisterCallBacks(eventHandlerID uint64) (err error) {
	if eventHandlerID == ApplicationEventHandlerID {
		return
	}
	var funcs []js.Func
	var found bool
	if funcs, found = jsCallBacks[eventHandlerID]; !found {
		err = fmt.Errorf("eventHandlerID not found in jsCallBacks")
		return
	}
	for _, f := range funcs {
		f.Release()
	}
	delete(jsCallBacks, eventHandlerID)
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
	// eventHandlerID for not spawns is ApplicationEventHandlerID.
	cb = registerCallBack(fn, ApplicationEventHandlerID)
	return
}

func registerCallBack(f func(this js.Value, args []js.Value) interface{}, eventHandlerID uint64) (jsFunc js.Func) {
	var funcs []js.Func
	var found bool
	if funcs, found = jsCallBacks[eventHandlerID]; !found {
		funcs = make([]js.Func, 0, 20)
	}
	jsFunc = js.FuncOf(f)
	funcs = append(funcs, jsFunc)
	jsCallBacks[eventHandlerID] = funcs
	return
}

// NewMultiElementEventHandler constructs an event handler to be shared by multiple HTML elements.
func NewMultiElementEventHandler(handler func(event.Event) interface{}, eventHandlerID uint64) (cb js.Func) {
	wrapperFunc := func(this js.Value, args []js.Value) (nilReturn interface{}) {
		var eventJS js.Value
		if len(args) > 0 {
			eventJS = args[0]
		} else {
			eventJS = this
		}
		channelEvent := event.BuildEvent(eventJS, eventJS.Get("target"))
		nilReturn = handler(channelEvent)
		return
	}
	cb = registerCallBack(wrapperFunc, eventHandlerID)
	return
}

// AddMultiElementEventHandler adds an event handler wrapper func to the element.
// Use it to add the same event handler wrapper func to multiple HTML elements.
// Param handler was generated by func NewMultiElementEventHandler.
//   It is the element's event handler.
// Param element is the target element.
// Param on is the event name. ex: "click".
// Param capturePhase indicates if this happens during the capture phase.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func AddMultiElementEventHandler(cb js.Func, element js.Value, on string, capturePhase bool) {
	element.Call("addEventListener", on, cb, capturePhase)
	return
}
