// +build js, wasm

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
