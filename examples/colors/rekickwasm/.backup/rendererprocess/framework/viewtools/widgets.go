// +build js, wasm

package viewtools

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/framework/callback"
)

// CountWidgetsWaiting returns the number of widget event dispatchers listening to the eoj channel.
func CountWidgetsWaiting() (count int) {
	count = countWidgetsWaiting
	return
}

// IncWidgetWaiting increments the number of widget event dispatchers listening to the eoj channel.
func IncWidgetWaiting() {
	countWidgetsWaiting++
}

// DecWidgetWaiting decrements the number of widget event dispatchers listening to the eoj channel.
func DecWidgetWaiting() {
	if countWidgetsWaiting > 0 {
		countWidgetsWaiting--
	}
}

// Spawned Widgets.

type spawnedWidgetInfo struct {
	jsWidget js.Value
	id       uint64
}

// NewSpawnWidgetUniqueID returns a new id for a widget in a spawned panel.
func NewSpawnWidgetUniqueID() (spawnWidgetID uint64) {
	spawnWidgetID = newSpawnID()
	return
}

// SpawnWidget spawns a widget.
func SpawnWidget(spawnWidgetID uint64, widget, parent js.Value) {
	IncWidgetWaiting()
	parent.Call("appendChild", widget)
	spawnedWidgets[spawnWidgetID] = spawnedWidgetInfo{
		jsWidget: widget,
		id:       spawnWidgetID,
	}
	return
}

// UnSpawnWidget unspawns a widget.
func UnSpawnWidget(spawnWidgetID uint64) {
	var info spawnedWidgetInfo
	var found bool
	if info, found = spawnedWidgets[spawnWidgetID]; !found {
		return
	}
	parent := info.jsWidget.Get("parentNode")
	parent.Call("removeChild", info.jsWidget)
	DecWidgetWaiting()
	callback.UnRegisterCallBacks(spawnWidgetID)
	delete(spawnedWidgets, spawnWidgetID)
}
