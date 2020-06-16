package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createDisplay(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	data := &struct {
		ApplicationGitPath               string
		ImportRendererFrameworkViewTools string
		ImportRendererFrameworkLocation  string
		SpawnIDReplacePattern            string
	}{
		ApplicationGitPath:               builder.ImportPath,
		ImportRendererFrameworkViewTools: folderpaths.ImportRendererFrameworkViewTools,
		ImportRendererFrameworkLocation:  folderpaths.ImportRendererFrameworkLocation,
		SpawnIDReplacePattern:            project.SpawnIDReplacePattern,
	}
	var fname string
	var oPath string
	fname = fileNames.DisplayGo
	oPath = filepath.Join(folderpaths.OutputRendererDisplay, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DisplayGo, data, appPaths); err != nil {
		return
	}
	return
}
