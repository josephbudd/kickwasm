package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
)

func createViewToolsHelpersGo(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fname := "helpers.go"
	oPath := filepath.Join(folderpaths.OutputRendererWASMViewTools, fname)
	return appPaths.WriteFile(oPath, []byte(templates.ViewtoolsHelpers))
}
