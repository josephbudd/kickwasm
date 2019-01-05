package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createDataFilePathsGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fname := "filepaths.go"
	oPath := filepath.Join(folderpaths.OutputDomainDataFilepaths, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DataFilePathsGo, data, appPaths); err != nil {
		return
	}
	fname = "log.go"
	oPath = filepath.Join(folderpaths.OutputDomainDataCallIDs, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DataCallIDsLogGo, data, appPaths); err != nil {
		return
	}
	fname = "misc.go"
	oPath = filepath.Join(folderpaths.OutputDomainDataCallIDs, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DataCallIDsMiscGo, data, appPaths); err != nil {
		return
	}
	fname = "loglevels.go"
	oPath = filepath.Join(folderpaths.OutputDomainDataLogLevels, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DataLogLevelsGo, data, appPaths); err != nil {
		return
	}
	fname = "settings.go"
	oPath = filepath.Join(folderpaths.OutputDomainDataSettings, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DataSettingsGo, data, appPaths); err != nil {
		return
	}
	return
}
