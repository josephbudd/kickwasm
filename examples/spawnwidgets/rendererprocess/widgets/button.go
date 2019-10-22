// +build js, wasm

package widgets

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/notjs"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/viewtools"
)

// Button is a spawnable button.
type Button struct {
	notJS         *notjs.NotJS
	tools         *viewtools.Tools
	spawnUniqueID uint64
}

// NewButton constructs a new button.
func NewButton(tools *viewtools.Tools, notJS *notjs.NotJS) (button *Button) {
	button = &Button{
		tools: tools,
		notJS: notJS,
	}
	return
}

// Spawn constructs and adds a new spawned button to parent.
func (button *Button) Spawn(parent js.Value, label string, onclick func(event viewtools.Event) interface{}) {
	notJS := button.notJS
	tools := button.tools
	// Get a unique id for this spawn.
	button.spawnUniqueID = tools.NewSpawnWidgetUniqueID()
	// Build the widget's DOM element.
	// Don't add the widget's DOM element to the DOM.
	widget := notJS.CreateElementBUTTON()
	notJS.AppendChild(widget, notJS.CreateTextNode(label))
	htmlID := fmt.Sprintf("spawnedButton%d", button.spawnUniqueID)
	notJS.SetID(widget, htmlID)
	// Spawn the widget.
	tools.SpawnWidget(button.spawnUniqueID, widget, parent)
	// Handle the widget's onclick event.
	tools.AddSpawnEventHandler(onclick, widget, "click", false, button.spawnUniqueID)
}

// UnSpawn removes a button.
func (button *Button) UnSpawn() {
	button.tools.UnSpawnWidget(button.spawnUniqueID)
}
