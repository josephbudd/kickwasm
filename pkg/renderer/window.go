package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createWindow(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	var fname string
	var oPath string
	// renderer/markup/attributes.go
	fname = fileNames.WindowDotGo
	oPath = filepath.Join(folderpaths.OutputRendererWindow, fname)
	err = appPaths.WriteFile(oPath, []byte(templates.WindowGo))
	return
}
