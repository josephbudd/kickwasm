package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createDOM(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	data := &struct {
		ApplicationGitPath      string
		ImportRendererAPIMarkup string
	}{
		ApplicationGitPath:      builder.ImportPath,
		ImportRendererAPIMarkup: folderpaths.ImportRendererAPIMarkup,
	}
	var fname string
	var oPath string
	fname = fileNames.DOMDotGo
	oPath = filepath.Join(folderpaths.OutputRendererDOM, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DOMGo, data, appPaths); err != nil {
		return
	}
	return
}
