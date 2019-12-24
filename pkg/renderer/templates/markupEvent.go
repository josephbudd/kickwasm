package templates

// MarkupEventGo is the rendererprocess/markup/evnets.go file.
const MarkupEventGo = `// +build js, wasm

package markup

import (
	"{{.ApplicationGitPath}}{{.ImportRendererCallBack}}"
	"{{.ApplicationGitPath}}{{.ImportRendererEvent}}"
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
	callback.AddEventHandler(handler, e.element, on, capturePhase, e.panelUniqueID)
}
`
