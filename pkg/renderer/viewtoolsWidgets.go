package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsWidgetsGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	fileNames := appPaths.GetFileNames()
	folderpaths := appPaths.GetPaths()

	data := &struct {
		ApplicationGitPath     string
		ImportRendererCallBack string
	}{
		ApplicationGitPath:     builder.ImportPath,
		ImportRendererCallBack: folderpaths.ImportRendererCallBack,
	}
	fname := fileNames.WidgetsDotGo
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsWidgetGo, data, appPaths)
	return
}
