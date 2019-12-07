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
		DashUnderTabBar         string
		DashInner               string
		DashPanelHeading        string
		ApplicationGitPath      string
		SpawnIDReplacePattern   string
		ImportRendererSpawnPack string
		ImportRendererCallBack  string
		ImportRendererEvent     string
	}{
		DashUnderTabBar:         project.DashUnderTabBar,
		DashInner:               project.DashInnerString,
		DashPanelHeading:        project.DashPanelHeading,
		ApplicationGitPath:      builder.ImportPath,
		SpawnIDReplacePattern:   project.SpawnIDReplacePattern,
		ImportRendererSpawnPack: folderpaths.ImportRendererSpawnPack,
		ImportRendererCallBack:  folderpaths.ImportRendererCallBack,
		ImportRendererEvent:     folderpaths.ImportRendererEvent,
	}
	// execute the template
	fname := "spawnTab.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsSpawnTabGo, data, appPaths)
}
