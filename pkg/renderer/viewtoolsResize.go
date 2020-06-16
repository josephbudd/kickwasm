package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createResizeGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		IDs                             *project.IDs
		Classes                         *project.Classes
		Attributes                      *project.Attributes
		ApplicationGitPath              string
		ImportRendererAPIWindow         string
		ImportRendererFrameworkCallBack string
		ImportRendererAPIEvent          string
	}{
		IDs:                             builder.IDs,
		Classes:                         builder.Classes,
		Attributes:                      builder.Attributes,
		ApplicationGitPath:              builder.ImportPath,
		ImportRendererAPIWindow:         folderpaths.ImportRendererAPIWindow,
		ImportRendererFrameworkCallBack: folderpaths.ImportRendererFrameworkCallBack,
		ImportRendererAPIEvent:          folderpaths.ImportRendererAPIEvent,
	}
	var fname string
	var oPath string
	filenames := appPaths.GetFileNames()

	fname = filenames.ResizeDotGo
	oPath = filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsResize, data, appPaths); err != nil {
		return
	}

	fname = filenames.ResizeSliderDotGo
	oPath = filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsResizeSliderPanel, data, appPaths)
	return
}
