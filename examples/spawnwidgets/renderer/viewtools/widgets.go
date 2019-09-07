package viewtools

import (
	"syscall/js"
)

// CountWidgetsWaiting returns the number of widget event dispatchers listening to the eoj channel.
func (tools *Tools) CountWidgetsWaiting() (count int) {
	count = tools.countWidgetsWaiting
	return
}

// IncWidgetWaiting increments the number of widget event dispatchers listening to the eoj channel.
func (tools *Tools) IncWidgetWaiting() {
	tools.countWidgetsWaiting++
}

// DecWidgetWaiting decrements the number of widget event dispatchers listening to the eoj channel.
func (tools *Tools) DecWidgetWaiting() {
	if tools.countWidgetsWaiting > 0 {
		tools.countWidgetsWaiting--
	}
}

// Spawned Widgets.

type spawnedWidgetInfo struct {
	element   js.Value
	id        uint64
	unspawnCh chan struct{}
}

// NewSpawnWidgetUniqueID returns a new id for a widget in a spawned panel.
func (tools *Tools) NewSpawnWidgetUniqueID() (spawnWidgetID uint64) {
	spawnWidgetID = tools.newSpawnID()
	return
}

// SpawnWidget spawns a widget.
func (tools *Tools) SpawnWidget(spawnWidgetID uint64, widget, parent js.Value) (unspawnCh chan struct{}) {
	tools.IncWidgetWaiting()
	tools.NotJS.AppendChild(parent, widget)
	unspawnCh = make(chan struct{}, 1)
	tools.spawnedWidgets[spawnWidgetID] = spawnedWidgetInfo{
		element:   widget,
		id:        spawnWidgetID,
		unspawnCh: unspawnCh,
	}
	return
}

// UnSpawnWidget unspawns a widget.
func (tools *Tools) UnSpawnWidget(spawnWidgetID uint64) {
	var info spawnedWidgetInfo
	var found bool
	if info, found = tools.spawnedWidgets[spawnWidgetID]; !found {
		return
	}
	info.unspawnCh <- struct{}{}
	parent := tools.NotJS.ParentNode(info.element)
	tools.NotJS.RemoveChild(parent, info.element)
	tools.DecWidgetWaiting()
	tools.UnRegisterCallBacks(spawnWidgetID)
	delete(tools.spawnedWidgets, spawnWidgetID)
}
