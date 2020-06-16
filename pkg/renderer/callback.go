package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createCallBack(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	var fname string
	var oPath string
	data := &struct {
		ApplicationGitPath     string
		ImportRendererAPIEvent string
	}{
		ApplicationGitPath:     builder.ImportPath,
		ImportRendererAPIEvent: folderpaths.ImportRendererAPIEvent,
	}
	fname = fileNames.CallBackDotGo
	oPath = filepath.Join(folderpaths.OutputRendererFrameworkCallBack, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallBackGo, data, appPaths); err != nil {
		return
	}
	return
}
