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
	if err := templates.ProcessTemplate(fname, oPath, templates.DataFilePathsGo, data, appPaths); err != nil {
		return err
	}
	if data.AddAbout {
		fname = "getAbout.go"
		oPath = filepath.Join(folderpaths.OutputDomainDataCallIDs, fname)
		if err := templates.ProcessTemplate(fname, oPath, templates.DataCallIDsGetAboutGo, data, appPaths); err != nil {
			return err
		}
	}
	fname = "log.go"
	oPath = filepath.Join(folderpaths.OutputDomainDataCallIDs, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.DataCallIDsLogGo, data, appPaths); err != nil {
		return err
	}
	fname = "misc.go"
	oPath = filepath.Join(folderpaths.OutputDomainDataCallIDs, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.DataCallIDsMiscGo, data, appPaths); err != nil {
		return err
	}
	fname = "loglevels.go"
	oPath = filepath.Join(folderpaths.OutputDomainDataLogLevels, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.DataLogLevelsGo, data, appPaths); err != nil {
		return err
	}
	return nil
}
