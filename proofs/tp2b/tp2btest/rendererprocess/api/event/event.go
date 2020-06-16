// +build js, wasm

package event

import (
	"syscall/js"
)

// Event an event.
type Event struct {
	JSEvent      js.Value
	JSTarget     js.Value
}

// BuildEvent constructs an Event.
func BuildEvent(event, target js.Value) (ev Event) {
	ev = Event {
		JSEvent:      event,
		JSTarget:     target,
	}
	return
}

// PreventDefaultBehavior stops the DOM element from executing it's own default behavior.
func (ev *Event) PreventDefaultBehavior() {
	ev.JSEvent.Call("preventDefault")
}

// StopCurrentPhasePropagation stops the events from continuing
//  in the current phase path only.
// If there is a phase that follows this current phase,
//  then the event will continue to propagate through that phase.
// Each phase has it's own path.
// Capture Phase: window -> document -> html -> body ... -> parent of target.
// Target Phase: parent of target -> target -> parent of target.
// Bubble Phase: parent of target -> ... body -> html -> document.
func (ev *Event) StopCurrentPhasePropagation() {
	ev.JSEvent.Call("stopPropagation")
}

// StopAllPhasePropagation stops the events from continuing
//   in the current and following phase paths.
// Each phase has it's own path.
// Capture Phase: window -> document -> html -> body ... -> parent of target.
// Target Phase: parent of target -> target -> parent of target.
// Bubble Phase: parent of target -> ... body -> html -> document.
func (ev *Event) StopAllPhasePropagation() {
	ev.JSEvent.Call("stopImmediatePropagation")
}
