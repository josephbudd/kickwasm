// +build js, wasm

package viewtools

// CountMarkupPanels returns the number of markup panels.
func CountMarkupPanels() (count int) {
	count = countSpawnedMarkupPanels + countMarkupPanels
	return
}

// IncSpawnedPanels increments the number of spawned panels.
func IncSpawnedPanels(n int) {
	countSpawnedMarkupPanels += n
}

// DecSpawnedPanels decrements the number of spawned panels.
func DecSpawnedPanels(n int) {
	countSpawnedMarkupPanels -= n
	if countSpawnedMarkupPanels < 0 {
		countSpawnedMarkupPanels = 0
	}
}
