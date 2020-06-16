package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsCloserGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	filenames := appPaths.GetFileNames()
	data := &struct {
		ApplicationGitPath              string
		ImportRendererFrameworkCallBack string
		ImportRendererAPIDOM            string
		ImportRendererAPIEvent          string
	}{
		ApplicationGitPath:              builder.ImportPath,
		ImportRendererFrameworkCallBack: folderpaths.ImportRendererFrameworkCallBack,
		ImportRendererAPIDOM:            folderpaths.ImportRendererAPIDOM,
		ImportRendererAPIEvent:          folderpaths.ImportRendererAPIEvent,
	}
	fname := filenames.CloserDotGo //"closer.go"
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsCloser, data, appPaths)
	return
}
