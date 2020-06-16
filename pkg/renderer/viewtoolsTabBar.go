package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsTabBarGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		TabBarIDs                       []string
		LastPanelID                     string
		LastPanelLevels                 map[string]string
		SelectedTabClass                string
		ApplicationGitPath              string
		ImportRendererFrameworkCallBack string
		ImportRendererAPIEvent          string
	}{
		TabBarIDs:                       builder.GenerateTabBarIDs(),
		LastPanelID:                     "",
		LastPanelLevels:                 builder.GenerateTabBarIDStartPanelIDMap(),
		SelectedTabClass:                builder.Classes.SelectedTab,
		ApplicationGitPath:              builder.ImportPath,
		ImportRendererFrameworkCallBack: folderpaths.ImportRendererFrameworkCallBack,
		ImportRendererAPIEvent:          folderpaths.ImportRendererAPIEvent,
	}
	// execute the template
	fname := "tabBar.go"
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsTabBar, data, appPaths)
}
