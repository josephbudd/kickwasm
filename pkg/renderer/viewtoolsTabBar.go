package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsTabBarGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	data := &struct {
		TabBarIDs          []string
		LastPanelID        string
		LastPanelLevels    map[string]string
		UnSelectedTabClass string
		SelectedTabClass   string
	}{
		TabBarIDs:          builder.GenerateTabBarIDs(),
		LastPanelID:        "",
		LastPanelLevels:    builder.GenerateTabBarIDStartPanelIDMap(),
		SelectedTabClass:   builder.Classes.SelectedTab,
		UnSelectedTabClass: builder.Classes.UnSelectedTab,
	}
	// execute the template
	folderpaths := appPaths.GetPaths()
	fname := "tabBar.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsTabBar, data, appPaths)
}
