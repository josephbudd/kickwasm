package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsModalGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()
	data := struct {
		ApplicationGitPath     string
		ImportRendererAPIEvent string
	}{
		ApplicationGitPath:     builder.ImportPath,
		ImportRendererAPIEvent: folderpaths.ImportRendererAPIEvent,
	}
	fname := fileNames.ModalDotGo
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsModal, data, appPaths)
	return
}
