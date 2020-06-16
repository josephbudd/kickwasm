package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsMarkupGo(appPaths paths.ApplicationPathsI) (err error) {
	fileNames := appPaths.GetFileNames()
	folderpaths := appPaths.GetPaths()
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fileNames.MarkupDotGo)
	err = appPaths.WriteFile(oPath, []byte(templates.ViewToolsMarkupGo))
	return
}
