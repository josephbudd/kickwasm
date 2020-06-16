package project

// SpawnFolders indicates is a panel is spawned and it's folder path
type SpawnFolders struct {
	Position    int
	Spawn       bool
	Folders     []string
	HVScroll    bool
	ParentIsTab bool
}

// GenerateHomeEmptyInsidePanelNameSpawnedPathMap returns a map of
//   each home name mapped to
//   a map of each markup panel's name mapped to if spawned and a slice of that panel's relevant path
func (builder *Builder) GenerateHomeEmptyInsidePanelNameSpawnedPathMap() map[string]map[string]*SpawnFolders {
	homeEmptyInsidePanelNamePathMap := make(map[string]map[string]*SpawnFolders)
	for i, homeButton := range builder.Homes {
		panelNameSpawnPath := make(map[string]*SpawnFolders)
		for j, p := range homeButton.Panels {
			folderList := make([]string, 1, 10)
			folderList[0] = homeButton.ID
			spawnPath := &SpawnFolders{
				Position:    i + j,
				Spawn:       false,
				Folders:     folderList,
				ParentIsTab: false,
			}
			generateHomeEmptyInsidePanelNameSpawnedPathMap(p, spawnPath, panelNameSpawnPath)
		}
		homeEmptyInsidePanelNamePathMap[homeButton.ID] = panelNameSpawnPath
	}
	return homeEmptyInsidePanelNamePathMap
}
func generateHomeEmptyInsidePanelNameSpawnedPathMap(panel *Panel, spawnPath *SpawnFolders, panelNameSpawnPath map[string]*SpawnFolders) {
	if len(panel.Template) > 0 {
		spawnPath.HVScroll = panel.HVScroll
		panelNameSpawnPath[panel.Name] = spawnPath
		return
	}
	spawnPath.Folders = append(spawnPath.Folders, panel.Name)
	l := len(spawnPath.Folders)
	for bi, b := range panel.Buttons {
		newFolderList := make([]string, l+1, l*2)
		copy(newFolderList, spawnPath.Folders)
		newFolderList[l] = b.ID
		newSpawnPath := &SpawnFolders{
			Position:    spawnPath.Position + bi,
			Spawn:       spawnPath.Spawn,
			Folders:     newFolderList,
			ParentIsTab: false,
		}
		for _, p := range b.Panels {
			generateHomeEmptyInsidePanelNameSpawnedPathMap(p, newSpawnPath, panelNameSpawnPath)

		}
	}
	for ti, t := range panel.Tabs {
		newFolderList := make([]string, l+1, l*2)
		copy(newFolderList, spawnPath.Folders)
		newFolderList[l] = t.ID
		newSpawnPath := &SpawnFolders{
			Position:    spawnPath.Position + ti,
			Spawn:       t.Spawn,
			Folders:     newFolderList,
			ParentIsTab: true,
		}
		for _, p := range t.Panels {
			generateHomeEmptyInsidePanelNameSpawnedPathMap(p, newSpawnPath, panelNameSpawnPath)
		}
	}
}
