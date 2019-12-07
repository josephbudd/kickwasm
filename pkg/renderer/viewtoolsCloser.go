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
		ApplicationGitPath     string
		ImportRendererCallBack string
		ImportRendererDOM      string
		ImportRendererEvent    string
	}{
		ApplicationGitPath:     builder.ImportPath,
		ImportRendererCallBack: folderpaths.ImportRendererCallBack,
		ImportRendererDOM:      folderpaths.ImportRendererDOM,
		ImportRendererEvent:    folderpaths.ImportRendererEvent,
	}
	fname := filenames.CloserDotGo //"closer.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsCloser, data, appPaths)
	return
}
