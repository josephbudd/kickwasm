// +build js, wasm

package viewtools

import "syscall/js"

// Event an event.
type Event struct {
	Event        js.Value
	Target       js.Value
	CapturePhase bool
}

// PreventDefaultBehavior stops the DOM element from executing it's own default behavior.
func (event *Event) PreventDefaultBehavior() {
	event.Event.Call("preventDefault")
}

// StopCurrentPhasePropagation stops the events from continuing
//  in the current phase path only.
// If there is a phase that follows this current phase,
//  then the event will continue to propagate through that phase.
// Each phase has it's own path.
// Capture Phase: window -> document -> html -> body ... -> parent of target.
// Target Phase: parent of target -> target -> parent of target.
// Bubble Phase: parent of target -> ... body -> html -> document.
func (event *Event) StopCurrentPhasePropagation() {
	event.Event.Call("stopPropagation")
}

// StopAllPhasePropagation stops the events from continuing
//   in the current and following phase paths.
// Each phase has it's own path.
// Capture Phase: window -> document -> html -> body ... -> parent of target.
// Target Phase: parent of target -> target -> parent of target.
// Bubble Phase: parent of target -> ... body -> html -> document.
func (event *Event) StopAllPhasePropagation() {
	event.Event.Call("stopImmediatePropagation")
}

// AddEventHandler adds an event handler to the element.
// Param handler is the element's event handler.
// Param element is the target element.
// Param on is the event name. ex: "click".
// Param capturePhase indicates if this happens during the capture phase.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func (tools *Tools) AddEventHandler(handler func(Event) interface{}, element js.Value, on string, capturePhase bool) {
	tools.handleEvent(handler, element, on, capturePhase, 0)
	return
}

// AddSpawnEventHandler adds an event handler to the element.
// Param handler is the element's event handler.
// Param element is the target element.
// Param on is the event name. ex: "click".
// Param capturePhase indicates if this happens during the capture phase.
// Param uniqueID is the spawn's unique id.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func (tools *Tools) AddSpawnEventHandler(handler func(Event) interface{}, element js.Value, on string, capturePhase bool, uniqueID uint64) {
	tools.handleEvent(handler, element, on, capturePhase, uniqueID)
	return
}

func (tools *Tools) handleEvent(handler func(Event) interface{}, element js.Value, on string, capturePhase bool, uniqueID uint64) {
	wrapperFn := func(this js.Value, args []js.Value) (nilReturn interface{}) {
		var event js.Value
		if len(args) > 0 {
			event = args[0]
		} else {
			event = this
		}
		channelEvent := Event{
			Event:        event,
			Target:       tools.NotJS.GetEventTarget(event),
			CapturePhase: capturePhase,
		}
		nilReturn = handler(channelEvent)
		return
	}
	cb := tools.registerCallBack(wrapperFn, uniqueID)
	element.Call("addEventListener", on, cb, capturePhase)
}
