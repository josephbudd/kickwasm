package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createResizeGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	data := &struct {
		IDs        *project.IDs
		Classes    *project.Classes
		Attributes *project.Attributes
	}{
		IDs:        builder.IDs,
		Classes:    builder.Classes,
		Attributes: builder.Attributes,
	}
	// execute the template
	folderpaths := appPaths.GetPaths()
	fname := "resize.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsResize, data, appPaths)
}
