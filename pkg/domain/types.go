package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createTypes(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()

	fname := fileNames.SettingsDotGo
	oPath := filepath.Join(folderpaths.OutputDomainTypes, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.SettingsGo, data, appPaths); err != nil {
		return
	}
	return
}
