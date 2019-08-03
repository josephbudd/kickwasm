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
		DashUnderTabBar           string
		DashInner                 string
		DashPanelHeading          string
		ApplicationGitPath        string
		ImportRendererSpawnPack   string
		ImportRendererSpawnPanels string
		SpawnIDReplacePattern     string
	}{
		DashUnderTabBar:           project.DashUnderTabBar,
		DashInner:                 project.DashInnerString,
		DashPanelHeading:          project.DashPanelHeading,
		ApplicationGitPath:        builder.ImportPath,
		ImportRendererSpawnPack:   folderpaths.ImportRendererSpawnPack,
		ImportRendererSpawnPanels: folderpaths.ImportRendererSpawnPanels,
		SpawnIDReplacePattern:     project.SpawnIDReplacePattern,
	}
	// execute the template
	fname := "spawnTab.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsSpawnTabGo, data, appPaths)
}
