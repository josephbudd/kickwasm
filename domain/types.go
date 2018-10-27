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
	if err := templates.ProcessTemplate(fname, oPath, templates.TypesRecordsGo, data, appPaths); err != nil {
		return err
	}
	fname = "logCallParams.go"
	oPath = filepath.Join(folderpaths.OutputDomainTypes, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.TypesLogGo, data, appPaths); err != nil {
		return err
	}
	if data.AddAbout {
		fname = "getAboutCallParams.go"
		oPath = filepath.Join(folderpaths.OutputDomainTypes, fname)
		if err := templates.ProcessTemplate(fname, oPath, templates.TypesGetAboutParamsGo, data, appPaths); err != nil {
			return err
		}
	}
	return nil
}
