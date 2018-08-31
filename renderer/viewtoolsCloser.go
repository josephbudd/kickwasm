package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
)

func createViewToolsCloserGo(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fname := "closer.go"
	oPath := filepath.Join(folderpaths.OutputRendererWASMViewTools, fname)
	return appPaths.WriteFile(oPath, []byte(templates.ViewToolsCloser))
}
