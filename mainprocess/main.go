package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createMainGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "main.go"
	oPath := filepath.Join(folderpaths.Output, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths)
}
