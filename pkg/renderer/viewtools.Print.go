package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsPrintGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		IDs                  *project.IDs
		Classes              *project.Classes
		Attributes           *project.Attributes
		ApplicationGitPath   string
		ImportRendererWindow string
		ImportRendererEvent  string
	}{
		IDs:                  builder.IDs,
		Classes:              builder.Classes,
		Attributes:           builder.Attributes,
		ApplicationGitPath:   builder.ImportPath,
		ImportRendererWindow: folderpaths.ImportRendererWindow,
		ImportRendererEvent:  folderpaths.ImportRendererEvent,
	}
	var fname string
	var oPath string
	filenames := appPaths.GetFileNames()

	fname = filenames.PrintDotGo
	oPath = filepath.Join(folderpaths.OutputRendererViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsPrint, data, appPaths)
	return
}