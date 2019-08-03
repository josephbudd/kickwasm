package project

// SpawnFolders indicates is a panel is spawned and it's folder path
type SpawnFolders struct {
	Position int
	Spawn    bool
	Folders  []string
}

// GenerateServiceEmptyInsidePanelNameSpawnedPathMap returns a map of
//   each service name mapped to
//   a map of each markup panel's name mapped to if spawned and a slice of that panel's relevant path
func (builder *Builder) GenerateServiceEmptyInsidePanelNameSpawnedPathMap() map[string]map[string]*SpawnFolders {
	serviceEmptyInsidePanelNamePathMap := make(map[string]map[string]*SpawnFolders)
	for si, s := range builder.Services {
		panelNameSpawnPath := make(map[string]*SpawnFolders)
		for pi, p := range s.Button.Panels {
			folderList := make([]string, 1, 10)
			folderList[0] = s.Button.ID
			spawnPath := &SpawnFolders{
				Position: si + pi,
				Spawn:    false,
				Folders:  folderList,
			}
			generateServiceEmptyInsidePanelNameSpawnedPathMap(p, spawnPath, panelNameSpawnPath)
		}
		serviceEmptyInsidePanelNamePathMap[s.Name] = panelNameSpawnPath
	}
	return serviceEmptyInsidePanelNamePathMap
}
func generateServiceEmptyInsidePanelNameSpawnedPathMap(panel *Panel, spawnPath *SpawnFolders, panelNameSpawnPath map[string]*SpawnFolders) {
	if len(panel.Template) > 0 {
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
			Position: spawnPath.Position + bi,
			Spawn:    spawnPath.Spawn,
			Folders:  newFolderList,
		}
		for _, p := range b.Panels {
			generateServiceEmptyInsidePanelNameSpawnedPathMap(p, newSpawnPath, panelNameSpawnPath)

		}
	}
	for ti, t := range panel.Tabs {
		newFolderList := make([]string, l+1, l*2)
		copy(newFolderList, spawnPath.Folders)
		newFolderList[l] = t.ID
		newSpawnPath := &SpawnFolders{
			Position: spawnPath.Position + ti,
			Spawn:    t.Spawn,
			Folders:  newFolderList,
		}
		for _, p := range t.Panels {
			generateServiceEmptyInsidePanelNameSpawnedPathMap(p, newSpawnPath, panelNameSpawnPath)
		}
	}
}
