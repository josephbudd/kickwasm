package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createDataFilePathsGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()

	fname := fileNames.FilePathsDotGo
	oPath := filepath.Join(folderpaths.OutputDomainDataFilepaths, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DataFilePathsGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.LogLevelsDotGo
	oPath = filepath.Join(folderpaths.OutputDomainDataLogLevels, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DataLogLevelsGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.SettingsDotGo
	oPath = filepath.Join(folderpaths.OutputDomainDataSettings, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.DataSettingsGo, data, appPaths); err != nil {
		return
	}
	return
}
