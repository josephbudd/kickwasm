package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
)

func createViewToolsModalGo(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fname := "modal.go"
	oPath := filepath.Join(folderpaths.OutputRendererWASMViewTools, fname)
	return appPaths.WriteFile(oPath, []byte(templates.ViewToolsModal))
}
