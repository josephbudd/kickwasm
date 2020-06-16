package templates

// JSValueSpawnPanel is the rendererprocess/api/jsvalue/spawnpanel.go file template.
const JSValueSpawnPanel = `// +build js, wasm

package jsvalue

import (
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportRendererAPIEvent}}"
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkCallBack}}"
)

// SetSpawnPanelEventHandler sets the event handler for elements in spawned panels and their widgets.
// It is for spawned markup panels and their widgets,
//   when you are using the GO package syscall/js to create DOM elements.
// Param e is the html element that will have the event.
// Param handler is the event handler.
// Param on is the event name. ex: "mousedown".
// Param capturePhase indicates if this happens during the capture phase.
// Param spawnPanelUniqueID is the spawn panel's unique id.
// Each phase has it's own path.
//   Capture Phase: window -> document -> html -> body ... -> parent of target.
//   Target Phase: parent of target -> target -> parent of target.
//   Bubble Phase: parent of target -> ... body -> html -> document.
func SetSpawnPanelEventHandler(e js.Value, handler func(event.Event) interface{}, on string, capturePhase bool, spawnPanelUniqueID uint64) {
	callback.AddEventHandler(handler, e, on, capturePhase, spawnPanelUniqueID)
}
`
