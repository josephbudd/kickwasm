package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsSliderGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fname := "slider.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	data := &struct {
		ApplicationGitPath     string
		ImportRendererCallBack string
		ImportRendererEvent    string
		ImportRendererWindow   string
	}{
		ApplicationGitPath:     builder.ImportPath,
		ImportRendererCallBack: folderpaths.ImportRendererCallBack,
		ImportRendererEvent:    folderpaths.ImportRendererEvent,
		ImportRendererWindow:   folderpaths.ImportRendererWindow,
	}
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsSlider, data, appPaths)
	return
}
