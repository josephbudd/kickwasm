package refactor

import (
	"errors"
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/project"
)

// ButtonPanelErrors discovers button panel placement errors.
// Returns the errors as a single error.
func ButtonPanelErrors(changes, original *project.Builder) (err error) {

	errList := make([]string, 0, 50)
	var errmsg string
	defer func() {
		if len(errList) > 0 {
			errmsg = strings.Join(errList, "\n")
			err = errors.New(errmsg)
		}
	}()

	// Buttons and their panels.
	chButtonPanelIDs := changes.GenerateButtonIDsPanelIDs()
	orButtonPanelIDs := original.GenerateButtonIDsPanelIDs()
	chTabIDSpawnPanelIDs := changes.GenerateTabIDsPanelIDs()
	var orPanelIDs []string
	var orPanelID string
	var found bool
	for buttonID, chPanelIDs := range chButtonPanelIDs {
		for _, chPanelID := range chPanelIDs {
			// Assume that the panel did not move.
			if orPanelIDs, found = orButtonPanelIDs[buttonID]; found {
				// Check that same button.
				found = false
				for _, orPanelID = range orPanelIDs {
					if orPanelID == chPanelID {
						found = true
						break
					}
				}
			}
			if !found {
				// The panel is not with the same button.
				// Check every other button.
				var pIDs []string
				var pID string
				for _, pIDs = range chButtonPanelIDs {
					for _, pID = range pIDs {
						if pID == chPanelID {
							// The panel was moved from one panel to another button.
							found = true
							break
						}
					}
				}
			}
			if !found {
				// The panel is not with the same button.
				// The panel is not with any other button.
				// Check the tabs.
				for chTabID, chSpawnPanelIDs := range chTabIDSpawnPanelIDs {
					for _, chTabPanelPid := range chSpawnPanelIDs.PanelIDs {
						if chTabPanelPid == chPanelID {
							// The panel was moved from a button to a tab.
							if chSpawnPanelIDs.Spawn {
								// A panel can not be moved to a spawn tab panel group.
								errmsg = fmt.Sprintf("the button panel named %q was moved to the spawned tab named %q", chPanelID, chTabID)
								errList = append(errList, errmsg)
							}
							break
						}
					}
				}
			}
			// If the panel is not ever found then it was removed.
			// No error in removing a panel.
		}
	}
	return
}

// TabPanelErrors discovers tab panel placement errors.
// Returns the errors as a single error.
func TabPanelErrors(changes, original *project.Builder) (err error) {

	errList := make([]string, 0, 50)
	var errmsg string
	defer func() {
		if len(errList) > 0 {
			errmsg = strings.Join(errList, "\n")
			err = errors.New(errmsg)
		}
	}()

	// Check for tabs with toggled spawns.
	chTabIDSpawnPanelIDs := changes.GenerateTabIDsPanelIDs()
	orTabIDSpawnPanelIDs := original.GenerateTabIDsPanelIDs()
	for chTabID, chTabIDSpawnPanelID := range chTabIDSpawnPanelIDs {
		if orSpawnPanelIDs, found := orTabIDSpawnPanelIDs[chTabID]; found {
			if orSpawnPanelIDs.Spawn != chTabIDSpawnPanelID.Spawn {
				if orSpawnPanelIDs.Spawn {
					errmsg = fmt.Sprintf("the spawned tab named %q has been changed to unspawned", chTabID)
					errList = append(errList, errmsg)
				} else {
					errmsg = fmt.Sprintf("the tab named %q has been changed to spawned", chTabID)
					errList = append(errList, errmsg)
				}
			}
		}
	}
	if len(errList) > 0 {
		// Only list the toggled tab spawn errors.
		// Don't combine them with other errors.
		return
	}

	// Check for tab panel placement into other panel groups.
	for chTabID, chSpawnPanelIDs := range chTabIDSpawnPanelIDs {
		for _, chSpawnPanelID := range chSpawnPanelIDs.PanelIDs {
			for orTabID, orSpawnPanelIDs := range orTabIDSpawnPanelIDs {
				for _, orSpawnPanelID := range orSpawnPanelIDs.PanelIDs {
					if found := (orSpawnPanelID == chSpawnPanelID); found {
						// The panel was moved inside a tab's panel group.
						// Make sure the tab is still spawned or unspawned.
						if orSpawnPanelIDs.Spawn != chSpawnPanelIDs.Spawn {
							if orSpawnPanelIDs.Spawn {
								errmsg = fmt.Sprintf("the panel named %q has been moved from the spawned tab named %q to the normal tab named %q", orSpawnPanelID, orTabID, chTabID)
								errList = append(errList, errmsg)
							} else {
								errmsg = fmt.Sprintf("the panel named %q has been moved from the normal tab named %q to the spawned tab named %q", orSpawnPanelID, orTabID, chTabID)
								errList = append(errList, errmsg)
							}
						}
						break
					} else {
						// The panel is not inside another tab's panel group.
						if chSpawnPanelIDs.Spawn {
							// The panel is from a spawned tab.
							// It may not be moved to a button's panel group.
							chButtonPanelIDs := changes.GenerateButtonIDsPanelIDs()
							for chButtonID, chButtonPanelIDs := range chButtonPanelIDs {
								for _, chButtonPanelID := range chButtonPanelIDs {
									if chButtonPanelID == chSpawnPanelID {
										// The panel was moved from a spawned tab to a button.
										errmsg = fmt.Sprintf("the panel named %q has been moved from the spawn tab named %q to the button named %q", chSpawnPanelID, chTabID, chButtonID)
										errList = append(errList, errmsg)
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
