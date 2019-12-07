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
		TabBarIDs              []string
		LastPanelID            string
		LastPanelLevels        map[string]string
		UnSelectedTabClass     string
		SelectedTabClass       string
		ApplicationGitPath     string
		ImportRendererCallBack string
		ImportRendererEvent    string
	}{
		TabBarIDs:              builder.GenerateTabBarIDs(),
		LastPanelID:            "",
		LastPanelLevels:        builder.GenerateTabBarIDStartPanelIDMap(),
		SelectedTabClass:       builder.Classes.SelectedTab,
		UnSelectedTabClass:     builder.Classes.UnSelectedTab,
		ApplicationGitPath:     builder.ImportPath,
		ImportRendererCallBack: folderpaths.ImportRendererCallBack,
		ImportRendererEvent:    folderpaths.ImportRendererEvent,
	}
	// execute the template
	fname := "tabBar.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsTabBar, data, appPaths)
}
