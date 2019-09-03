package viewtools

// CountMarkupPanels returns the number of markup panels.
func (tools *Tools) CountMarkupPanels() (count int) {
	count = tools.countSpawnedMarkupPanels + tools.countMarkupPanels
	return
}

// IncSpawnedPanels increments the number of spawned panels.
func (tools *Tools) IncSpawnedPanels(n int) {
	tools.countSpawnedMarkupPanels += n
}

// DecSpawnedPanels decrements the number of spawned panels.
func (tools *Tools) DecSpawnedPanels(n int) {
	tools.countSpawnedMarkupPanels -= n
	if tools.countSpawnedMarkupPanels < 0 {
		tools.countSpawnedMarkupPanels = 0
	}
}

// CountWidgetsWaiting returns the number of widgets widgets listening to the eoj channel.
func (tools *Tools) CountWidgetsWaiting() (count int) {
	count = tools.countWidgetsWaiting
	return
}

// IncWidgetWaiting increments the number of widgets listening to the eoj channel.
func (tools *Tools) IncWidgetWaiting() {
	tools.countWidgetsWaiting++
}

// DecWidgetWaiting decrements the number of widgets listening to the eoj channel.
func (tools *Tools) DecWidgetWaiting() {
	if tools.countWidgetsWaiting > 0 {
		tools.countWidgetsWaiting--
	}
}
