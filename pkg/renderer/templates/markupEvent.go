package templates

// MarkupEventGo is the rendererprocess/markup/evnets.go file.
const MarkupEventGo = `// +build js, wasm

package markup

import (
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportRendererAPIEvent}}"
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkCallBack}}"
)

// SetEventHandler sets the event handler for an element.
// Param handler is the element's event handler.
// Param on is the event name. ex: "click".
// Param capturePhase indicates if this happens during the capture phase.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func (e *Element) SetEventHandler(handler func(event.Event) interface{}, on string, capturePhase bool) {
	callback.AddEventHandler(handler, e.element, on, capturePhase, e.eventHandlerID)
}

/*
	A MultipleElement event handler is an event handler for multiple elements.

	For example:
	  Buttons in a virtual list.
	  The handler checks button attributes for the information and acts accordingly.

	  1. handlerID := jsvalue.NewMultipleElementEventHandlerID()
	  2. handler := func(e event.Event) (nilReturn interface{}) {
		  button := document.NewElementFromJSValue(e.JSTarget)
		  text := button.InnerText()
		  log.Println("You clicked the ", text, " button.")
		  return
	  }
	  3. handerFunc := markup.NewMultiElementEventHandler(handler, handlerID)

	  4. button1 := document.ElementByID("myButton1")
	  5. button2 := document.ElementByID("myButton2")
	  6. button1.SetMultipleElementEventHandler(handlerFunc, "click", false)
	  7. button2.SetMultipleElementEventHandler(handlerFunc, "click", false)

	If you need to release the handler then first you must destroy the HTML elements.
	  button1.Parent().RemoveChild(button1)
	  button2.Parent().RemoveChild(button2)
	  markup.ReleaseMultipleElementEventHandler(handlerID)

*/

// NewMultipleElementEventHandlerID returns a new id for a single event handler that multiple elements will use.
// The event handler will be used in the call to func NewMultiElementEventHandler.
// The eventHandlerID can be reused after calling func ReleaseMultipleElementEventHandler.
//   if there is only ever one instance of the widget used at a time.
func NewMultipleElementEventHandlerID() (eventHandlerID uint64) {
	eventHandlerID = callback.NewEventHandlerID()
	return
}

// NewMultiElementEventHandler constructs an event handler to be shared by multiple HTML elements.
// Param handler is a normal event handler func(event.Event) (nilReturn interface{})
// Param eventHandlerID is an id from func NewMultipleElementEventHandlerID()
// Returns a simple (cb js.Func) for func SetMultipleElementEventHandler.
func NewMultiElementEventHandler(handler func(event.Event) (nilReturn interface{}), eventHandlerID uint64) (cb js.Func) {
	cb = callback.NewMultiElementEventHandler(handler, eventHandlerID)
	return
}

// SetMultipleElementEventHandler sets the multiple element event handler for this element.
// Param handler is the event handler from func NewMultiElementEventHandler.
// Param on is the event name. ex: "mousedown".
// Param capturePhase indicates if this happens during the capture phase.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func (e *Element) SetMultipleElementEventHandler(handler js.Func, on string, capturePhase bool) {
	callback.AddMultiElementEventHandler(handler, e.element, on, capturePhase)
}

// ReleaseMultipleElementEventHandler releases the event handlers
//   created with func NewMultiElementEventHandler.
// Params eventHandlerID was generated from func NewMultipleElementEventHandlerID.
//   It is the same eventHandlerID use in func NewMultiElementEventHandler.
// If there is an error it is only because there are no event handlers for this eventHandlerID.
func ReleaseMultipleElementEventHandler(eventHandlerID uint64) (err error) {
	err = callback.UnRegisterCallBacks(eventHandlerID)
	return
}
`
