package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createDataFilePathsGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "filepaths.go"
	oPath := filepath.Join(folderpaths.OutputDomainDataFilepaths, fname)
	return templates.ProcessTemplate(fname, oPath, templates.DataFilePathsGo, data, appPaths)
}
