package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsHideShowGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	data := struct {
		ApplicationGitPath      string
		ImportRendererAPIMarkup string
	}{
		ApplicationGitPath:      builder.ImportPath,
		ImportRendererAPIMarkup: folderpaths.ImportRendererAPIMarkup,
	}
	fname := "hideshow.go"
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsHideShow, data, appPaths)
	return
}
