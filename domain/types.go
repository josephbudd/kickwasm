package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createTypesCallsGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "calls.go"
	oPath := filepath.Join(folderpaths.OutputDomainTypes, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.TypesCallsGo, data, appPaths); err != nil {
		return err
	}
	fname = "records.go"
	oPath = filepath.Join(folderpaths.OutputDomainTypes, fname)
	return templates.ProcessTemplate(fname, oPath, templates.TypesRecordsGo, data, appPaths)
}
