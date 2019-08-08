package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createResizeGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	data := &struct {
		IDs        *project.IDs
		Classes    *project.Classes
		Attributes *project.Attributes
	}{
		IDs:        builder.IDs,
		Classes:    builder.Classes,
		Attributes: builder.Attributes,
	}
	folderpaths := appPaths.GetPaths()
	var fname string
	var oPath string

	fname = "resize.go"
	oPath = filepath.Join(folderpaths.OutputRendererViewTools, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsResize, data, appPaths); err != nil {
		return
	}

	fname = "resizeSlider.go"
	oPath = filepath.Join(folderpaths.OutputRendererViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsResizeSliderPanel, data, appPaths)
	return
}
