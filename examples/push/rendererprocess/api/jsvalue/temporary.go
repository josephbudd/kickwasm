// +build js, wasm

package jsvalue

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/event"
	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/framework/callback"
)

/*
	Temporary event handlers are handlers for widgets that are not permanent.
	For example:
	  A popup image magnifier widget.
	  When the user clicks on some image the magnifier near the image.
	  The magnifier displays a temporary magnified version of the original image.
	  The magnifier, it's events and it's event handlers are all temporary.
	  The magnifier widget goes away when the user releases the mouse button or moves the mouse off the image.

	  1. We are using the GO package syscall/js to create and manage DOM elements.
	  2. var magnifierID = NewTemporaryEventHandlerID() is called to create the event id for the magnifier widget.
	  3. The widget's html elements are created and inserted into the DOM using the GO syscall/js package.
	  4. func SetTemporaryEventHandler(..) is called when setting the "mouseup" and "mouseleave" event handlers.
	  5. The widget will end when the user releases the mouse button or moves the mouse off the image.
	     * The widget's html elements are removed from the DOM.
	     * func ReleaseTemporaryEventHandlers(magnifierID) is called by the "mouseup" or "mouseleave" event handlers
		     to release the event handlers.
	  4. The same magnifierID can be reused for the next time the magnifier is started up with some other picture
	       because there is only ever one instance of the magnifier.
*/

// NewTemporaryEventHandlerID returns a new id for temporary event handlers.
// It is for your all temporary widgets,
//   when you are using the GO package syscall/js to create DOM elements.
// It could also be used to make html element ids unique for widgets that can have multiple instances.
// ex: id := fmt.Sprintf("myButton%d", eventHandlerID)
// 1. The elements that use it are removed from the DOM.
// 2. The event handlers that used it are deleted with func ReleaseTemporaryEventHandlers.
// The eventHandlerID can be reused after calling ReleaseTemporaryEventHandlers
//   if there is only ever one instance of the widget used at a time.
func NewTemporaryEventHandlerID() (eventHandlerID uint64) {
	eventHandlerID = callback.NewEventHandlerID()
	return
}

// SetTemporaryEventHandler sets the event handler for an element.
// It is for your all temporary widgets,
//   when you are using the GO package syscall/js to create DOM elements.
// Param e is the html element that will have the event.
// Param handler is the event handler.
// Param on is the event name. ex: "mousedown".
// Param capturePhase indicates if this happens during the capture phase.
// Param eventUniqueID is the temporary event handler id.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func SetTemporaryEventHandler(e js.Value, handler func(event.Event) interface{}, on string, capturePhase bool, eventUniqueID uint64) {
	callback.AddEventHandler(handler, e, on, capturePhase, eventUniqueID)
}

// ReleaseTemporaryEventHandlers releases the event handlers added with SetEventHandler.
// If there is an error it is only because there are no event handlers for this eventUniqueID.
func ReleaseTemporaryEventHandlers(eventUniqueID uint64) (err error) {
	err = callback.UnRegisterCallBacks(eventUniqueID)
	return
}
