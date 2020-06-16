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
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	data := &struct {
		ApplicationGitPath              string
		ImportRendererFrameworkCallBack string
		ImportRendererAPIEvent          string
		ImportRendererAPIWindow         string
	}{
		ApplicationGitPath:              builder.ImportPath,
		ImportRendererFrameworkCallBack: folderpaths.ImportRendererFrameworkCallBack,
		ImportRendererAPIEvent:          folderpaths.ImportRendererAPIEvent,
		ImportRendererAPIWindow:         folderpaths.ImportRendererAPIWindow,
	}
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsSlider, data, appPaths)
	return
}
