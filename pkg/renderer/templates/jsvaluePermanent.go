package templates

// JSValuePermanent is the rendererprocess/api/jsvalue/permanent.go template.
const JSValuePermanent = `// +build js, wasm

package jsvalue

import (
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportRendererAPIEvent}}"
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkCallBack}}"
)

/*
	Permanent event handlers are handlers for your normal markup panels and their widgets.
	Permanent event handlers are not for spawn panels or their widgets.
	Permanent event handlers are not for temporary widgets.
	For example:
	  Your panel has an application form.
	  All of the text input elements and buttons are permanent parts of your application.
	  If you create their DOM elements using the go package syscall/js
		then you will set their event handlers with func SetPermanentEventHandler.
*/

// SetPermanentEventHandler sets the event handler for an element.
// It is for your normal (not spawned) markup panels and their widgets,
//  when you are using the GO package syscall/js to create DOM elements.
// Param e is the html element that will have the event.
// Param handler is the event handler.
// Param on is the event name. ex: "mousedown".
// Param capturePhase indicates if this happens during the capture phase.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func SetPermanentEventHandler(e js.Value, handler func(event.Event) interface{}, on string, capturePhase bool) {
	callback.AddEventHandler(handler, e, on, capturePhase, callback.ApplicationEventHandlerID)
}
`
