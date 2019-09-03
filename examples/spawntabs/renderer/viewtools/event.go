package viewtools

import "syscall/js"

// Event an event.
type Event struct {
	Event  js.Value
	Target js.Value
	On     string
}

// SendEvent watches an event and sends it over eventCh.
func (tools *Tools) SendEvent(eventCh chan Event, element js.Value, on string, preventDefault, stopPropagation, stopImmediatePropagation bool) {
	tools.sendEvent(eventCh, element, on, preventDefault, stopPropagation, stopImmediatePropagation, 0)
	return
}

// SendSpawnEvent watches an event and sends it over eventCh.
func (tools *Tools) SendSpawnEvent(eventCh chan Event, element js.Value, on string, preventDefault, stopPropagation, stopImmediatePropagation bool, uniqueID uint64) {
	tools.sendEvent(eventCh, element, on, preventDefault, stopPropagation, stopImmediatePropagation, uniqueID)
	return
}

func (tools *Tools) sendEvent(eventCh chan Event, element js.Value, on string, preventDefault, stopPropagation, stopImmediatePropagation bool, uniqueID uint64) {
	wrapperFn := func(this js.Value, args []js.Value) (nilReturn interface{}) {
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
		channelEvent := Event{
			Event:  event,
			Target: tools.NotJS.GetEventTarget(event),
			On:     on,
		}
		eventCh <- channelEvent
		return
	}
	cb := tools.registerCallBack(wrapperFn, uniqueID)
	element.Set(on, cb)
}

