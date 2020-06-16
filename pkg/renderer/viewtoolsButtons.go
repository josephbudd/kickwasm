package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsLocksGo(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, "locks.go")
	return appPaths.WriteFile(oPath, []byte(templates.ViewToolsLocks))
}
