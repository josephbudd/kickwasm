package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsWidgetsGo(appPaths paths.ApplicationPathsI) (err error) {
	fileNames := appPaths.GetFileNames()
	folderpaths := appPaths.GetPaths()
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fileNames.WidgetsDotGo)
	err = appPaths.WriteFile(oPath, []byte(templates.ViewToolsWidgetGo))
	return
}
