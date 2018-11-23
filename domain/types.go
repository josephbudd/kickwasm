package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createTypes(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "calls.go"
	oPath := filepath.Join(folderpaths.OutputDomainTypes, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.TypesCallsGo, data, appPaths); err != nil {
		return err
	}
	fname = "records.go"
	oPath = filepath.Join(folderpaths.OutputDomainTypes, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.TypesRecordsGo, data, appPaths); err != nil {
		return err
	}
	fname = "logCallParams.go"
	oPath = filepath.Join(folderpaths.OutputDomainTypes, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.TypesLogGo, data, appPaths); err != nil {
		return err
	}
	fname = "settings.go"
	oPath = filepath.Join(folderpaths.OutputDomainTypes, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.SettingsGo, data, appPaths); err != nil {
		return err
	}
	return nil
}
