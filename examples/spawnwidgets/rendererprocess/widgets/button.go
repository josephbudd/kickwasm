// +build js, wasm

package widgets

import (
	"context"
	"fmt"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/display"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/dom"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/event"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/markup"
)

// Button is a spawnable button.
type Button struct {
	spawnUniqueID uint64
	ctx           context.Context
}

// SpawnButton constructs and adds a new spawned button to parent.
// A spawned widget unspawns itself.
func SpawnButton(ctx context.Context, document *dom.DOM, parent *markup.Element, label string, onclick func(event.Event) interface{}) (widget *Button) {
	// Get a unique id for this spawn.
	widget = &Button{
		spawnUniqueID: display.NewSpawnWidgetUniqueID(),
		ctx:           ctx,
	}
	// Build the spawning widget.
	button := document.NewBUTTON()
	button.SetInnerText(label)
	button.SetID(fmt.Sprintf("spawnedButton%d", widget.spawnUniqueID))
	button.SetEventHandler(onclick, "click", false)

	// Register and Spawn the widget.
	display.SpawnWidget(widget.spawnUniqueID, button.JSValue(), parent.JSValue())

	// Unspawn the widget when the panel closes.
	go func(wgt *Button) {
		for {
			select {
			case <-wgt.ctx.Done():
				wgt.UnSpawn()
				return
			}
		}
	}(widget)

	return
}

// Unregister and unspawn the widget.
func (widget *Button) UnSpawn() {
	display.UnSpawnWidget(widget.spawnUniqueID)
}
