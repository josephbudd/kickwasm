package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createApplication(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	filenames := appPaths.GetFileNames()
	data := struct {
		Title                  string
		ApplicationGitPath     string
		ImportRendererDisplay  string
		ImportRendererEvent    string
		ImportRendererCallBack string
	}{
		Title:                  builder.Title,
		ApplicationGitPath:     builder.ImportPath,
		ImportRendererDisplay:  folderpaths.ImportRendererDisplay,
		ImportRendererEvent:    folderpaths.ImportRendererEvent,
		ImportRendererCallBack: folderpaths.ImportRendererCallBack,
	}
	fname := filenames.ApplicationDotGo
	oPath := filepath.Join(folderpaths.OutputRendererApplication, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ApplicationGo, data, appPaths)
	return

}
