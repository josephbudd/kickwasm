package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

func createViewToolsTabBarGo(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	data := &struct {
		TabBarIDs          []string
		LastPanelID        string
		LastPanelLevels    map[string]string
		UnSelectedTabClass string
		SelectedTabClass   string
	}{
		TabBarIDs:          builder.GenerateTabBarIDs(),
		LastPanelID:        builder.GenerateOpeningTabPanelID(),
		LastPanelLevels:    builder.GenerateTabBarLevelStartPanelMap(),
		SelectedTabClass:   builder.Classes.SelectedTab,
		UnSelectedTabClass: builder.Classes.UnSelectedTab,
	}
	// execute the template
	folderpaths := appPaths.GetPaths()
	fname := "tabBar.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsTabBar, data, appPaths)
}
