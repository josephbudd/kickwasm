package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsHelpersGo(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()
	fname := fileNames.HelpersDotGo
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	return appPaths.WriteFile(oPath, []byte(templates.ViewtoolsHelpers))
}
