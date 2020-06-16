package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsSpawnTabGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		DashUnderTabBar                  string
		DashInner                        string
		DashPanelHeading                 string
		ApplicationGitPath               string
		SpawnIDReplacePattern            string
		ImportRendererFrameworkSpawnPack string
		ImportRendererFrameworkCallBack  string
		ImportRendererAPIEvent           string
	}{
		DashUnderTabBar:                  project.DashUnderTabBar,
		DashInner:                        project.DashInnerString,
		DashPanelHeading:                 project.DashPanelHeading,
		ApplicationGitPath:               builder.ImportPath,
		SpawnIDReplacePattern:            project.SpawnIDReplacePattern,
		ImportRendererFrameworkSpawnPack: folderpaths.ImportRendererFrameworkSpawnPack,
		ImportRendererFrameworkCallBack:  folderpaths.ImportRendererFrameworkCallBack,
		ImportRendererAPIEvent:           folderpaths.ImportRendererAPIEvent,
	}
	// execute the template
	fname := "spawnTab.go"
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsSpawnTabGo, data, appPaths)
}
