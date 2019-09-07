package widgets

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/viewtools"
)

// Button is a spawnable button.
type Button struct {
	notJS         *notjs.NotJS
	tools         *viewtools.Tools
	spawnUniqueID uint64
	eventCh       chan viewtools.Event
	eojCh         chan struct{}
	unspawnCh     chan struct{}
	onclick       func(event js.Value)
}

// NewButton constructs a new button.
func NewButton(tools *viewtools.Tools, notJS *notjs.NotJS, eojCh chan struct{}) (button *Button) {
	button = &Button{
		tools:     tools,
		notJS:     notJS,
		eventCh:   make(chan viewtools.Event, 10),
		eojCh:     eojCh,
		unspawnCh: make(chan struct{}),
	}
	return
}

// Spawn constructs and adds a new spawned button to parent.
func (button *Button) Spawn(parent js.Value, label string, onclick func(event js.Value)) {
	notJS := button.notJS
	tools := button.tools
	// Start the event dispatcher.
	button.onclick = onclick
	button.dispatchEvents()
	// Get a unique id for this spawn.
	button.spawnUniqueID = tools.NewSpawnWidgetUniqueID()
	// Build the widget's DOM element.
	// Don't add the widget's DOM element to the DOM.
	widget := notJS.CreateElementBUTTON()
	notJS.AppendChild(widget, notJS.CreateTextNode(label))
	htmlID := fmt.Sprintf("spawnedButton%d", button.spawnUniqueID)
	notJS.SetID(widget, htmlID)
	// Spawn the widget.
	button.unspawnCh = tools.SpawnWidget(button.spawnUniqueID, widget, parent)
	// Receive the button's onclick event.
	tools.SendSpawnEvent(button.eventCh, widget, "onclick", false, false, false, button.spawnUniqueID)
}

// UnSpawn removes a button.
func (button *Button) UnSpawn() {
	button.tools.UnSpawnWidget(button.spawnUniqueID)
}

// dispatchEvents dispatches events from the button.
// It stops when it receives on the eoj channel.
func (button *Button) dispatchEvents() {
	go func() {
		for {
			select {
			case <-button.eojCh:
				return
			case <-button.unspawnCh:
				return
			case event := <-button.eventCh:
				button.onclick(event.Event)
			}
		}
	}()

	return
}
