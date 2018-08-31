package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createFilePathsFilePathsGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "filepaths.go"
	oPath := filepath.Join(folderpaths.OutputMainProcessDataFilePaths, fname)
	return templates.ProcessTemplate(fname, oPath, templates.FilePathsFilePathsGo, data, appPaths)
}
