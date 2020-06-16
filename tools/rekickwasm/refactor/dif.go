package refactor

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/project"
)

// DifButtonPanelPositions discovers button panel placement changes.
// Returns true if any button panel position has changed.
func DifButtonPanelPositions(changes, original *project.Builder) (changed bool) {
	// Buttons and their panels.
	chButtonPanelIDs := changes.GenerateButtonIDsPanelIDs()
	orButtonPanelIDs := original.GenerateButtonIDsPanelIDs()
	if changed = (len(chButtonPanelIDs) != len(orButtonPanelIDs)); changed {
		return
	}
	// Check each button's panel group.
	for chButtonID, chPanelPids := range chButtonPanelIDs {
		var orPanelIDs []string
		var found bool
		if orPanelIDs, found = orButtonPanelIDs[chButtonID]; !found {
			// A button has been removed.
			changed = true
			return
		}
		if len(orPanelIDs) != len(chPanelPids) {
			// A panel has been added or removed.
			changed = true
			return
		}
		for i, chPanelPID := range chPanelPids {
			if chPanelPID != orPanelIDs[i] {
				// Panels withing this panel group have moved around.
				changed = true
				return
			}
		}
	}
	return
}

// DifTabPanelPositions discovers tab panel placement changes.
// Returns true if any tab panel position has changed.
func DifTabPanelPositions(changes, original *project.Builder) (changed bool) {
	// Tabs and their spawn and panels.
	chTabIDSpawnPanelIDs := changes.GenerateTabIDsPanelIDs()
	orTabIDSpawnPanelIDs := original.GenerateTabIDsPanelIDs()
	if changed = (len(chTabIDSpawnPanelIDs) != len(orTabIDSpawnPanelIDs)); changed {
		return
	}
	// Check each tab's panel group.
	for chTabID, chSpawnPanelIDs := range chTabIDSpawnPanelIDs {
		var orSpawnPanelIDs project.TabSpawnPanelIDs
		var found bool
		if orSpawnPanelIDs, found = orTabIDSpawnPanelIDs[chTabID]; !found {
			changed = true
			return
		}
		if len(chSpawnPanelIDs.PanelIDs) != len(orSpawnPanelIDs.PanelIDs) {
			// A panel has been added or removed.
			changed = true
			return
		}
		for i, chSpawnPanelID := range chSpawnPanelIDs.PanelIDs {
			if chSpawnPanelID != orSpawnPanelIDs.PanelIDs[i] {
				// Panels withing this panel group have moved around.
				changed = true
				return
			}
		}
	}
	return
}

// DifHomePositionsButtons returns is differenct home buttons.
func DifHomePositionsButtons(changes, original *project.Builder) (changed bool) {
	chHomes := changes.Homes
	for i, oHomeButton := range original.Homes {
		chHomeButton := chHomes[i]
		if oHomeButton.ID != chHomeButton.ID {
			changed = true
			break
		}
	}
	return
}

// MoveSpawnPath represents the source and destination of a move.
type MoveSpawnPath struct {
	From SpawnPath
	To   SpawnPath
}

// SpawnPath indicates a path and if it is spawned.
type SpawnPath struct {
	Spawn       bool
	Path        string
	HVScroll    bool
	ParentIsTab bool
}

// DifPanelPaths returns panel names mapped to their relative paths.
// removals are what must be removed from original.
// additions are what must be added to original from changes.
// moves are what must be moved.
//  * added panel in the group.
//  * move caused an added or removed panel in the group.
//  * the first panel in the group was switched with another in the same group. ( new default panel )
// positionChanged indicates a change in some panel's position in it's group.
func DifPanelPaths(changes, original *project.Builder) (removals, additions map[string]SpawnPath, moves map[string]MoveSpawnPath, positionChanged, scrollChanged bool) {

	// original panels mapped to their paths
	orPanelNamePathMap := buildPanelNamePathMap(original)
	// original paths mapped to their panels
	orPathPanelNamesMap := make(map[SpawnPath][]string)
	for name, spawnFolders := range orPanelNamePathMap {
		spawnPath := SpawnPath{
			Spawn:       spawnFolders.Spawn,
			Path:        filepath.Join(spawnFolders.Folders...),
			HVScroll:    spawnFolders.HVScroll,
			ParentIsTab: spawnFolders.ParentIsTab,
		}
		if _, found := orPathPanelNamesMap[spawnPath]; !found {
			orPathPanelNamesMap[spawnPath] = make([]string, 0, 5)
		}
		orPathPanelNamesMap[spawnPath] = append(orPathPanelNamesMap[spawnPath], name)
	}

	// changes panels mapped to their paths.
	chPanelNamePathMap := buildPanelNamePathMap(changes)
	removals = make(map[string]SpawnPath, 100)
	additions = make(map[string]SpawnPath, 100)
	moves = make(map[string]MoveSpawnPath, 100)
	// additions, moves and scrollChanged
	// for chName, chPath := range chPanelNamePathMap {
	for chName, chSpawnFolders := range chPanelNamePathMap {
		chSpawnPath := SpawnPath{
			Spawn:       chSpawnFolders.Spawn,
			Path:        filepath.Join(chSpawnFolders.Folders...),
			HVScroll:    chSpawnFolders.HVScroll,
			ParentIsTab: chSpawnFolders.ParentIsTab,
		}
		orSpawnFolders, found := orPanelNamePathMap[chName]
		if !found {
			// addition
			additions[chName] = chSpawnPath
		} else {
			// possible move
			orSpawnPath := SpawnPath{
				Spawn:       orSpawnFolders.Spawn,
				Path:        filepath.Join(orSpawnFolders.Folders...),
				HVScroll:    orSpawnFolders.HVScroll,
				ParentIsTab: orSpawnFolders.ParentIsTab,
			}
			if chSpawnPath.Path != orSpawnPath.Path {
				// an attempt to move
				if chSpawnPath.ParentIsTab == orSpawnPath.ParentIsTab {
					// Panel moves from one tab to another tab or from one button to another button.
					if chSpawnFolders.Spawn == orSpawnFolders.Spawn {
						// ok to move because spawn is unchanged.
						// move
						moves[chName] = MoveSpawnPath{
							From: orSpawnPath,
							To:   chSpawnPath,
						}
					} else {
						// spawn changed so this is a removal and addition.
						// remove from original
						removals[chName] = orSpawnPath
						// add to refactor
						additions[chName] = chSpawnPath
					}
				} else {
					// Panel moves from a tab to a button or from a button to a tab.
					// So this is a removal and addition.
					// remove from original
					removals[chName] = orSpawnPath
					// add to refactor
					additions[chName] = chSpawnPath
				}
			} else if chSpawnPath.HVScroll != orSpawnPath.HVScroll {
				// a panel's scroll changed.
				scrollChanged = true
			}
		}
	}
	// removals
	// panels in groups where a panel was removed will need new group files.
	for orName, orSpawnFolders := range orPanelNamePathMap {
		_, found := chPanelNamePathMap[orName]
		if !found {
			path := filepath.Join(orSpawnFolders.Folders...)
			orSpawnPath := SpawnPath{
				Spawn:       orSpawnFolders.Spawn,
				Path:        path,
				ParentIsTab: orSpawnFolders.ParentIsTab,
			}
			removals[orName] = orSpawnPath
		}
	}
	// check for any button, tab, panel position changes.
	for name, orSpawnFolders := range orPanelNamePathMap {
		if chSpawnFolders, found := chPanelNamePathMap[name]; !found {
			positionChanged = true
			break
		} else {
			if orSpawnFolders.Position != chSpawnFolders.Position {
				positionChanged = true
				break
			}
		}
	}
	return
}

func buildPanelNamePathMap(b *project.Builder) (namePathMap map[string]*project.SpawnFolders) {
	namePathMap = make(map[string]*project.SpawnFolders)
	homeEmptyInsidePanelNamePathMap := b.GenerateHomeEmptyInsidePanelNameSpawnedPathMap()
	// homeEmptyInsidePanelNamePathMap := b.GenerateHomeEmptyInsidePanelNamePathMap()
	// One home at a time.
	for _, panelNamePathMap := range homeEmptyInsidePanelNamePathMap {
		// Ignore home name.
		for name, spawnFolders := range panelNamePathMap {
			namePathMap[name] = spawnFolders
		}
	}
	return
}

// MovePath represents a required path move.
type MovePath struct {
	From string
	To   string
}
