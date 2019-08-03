package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createSettingsYAML(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()
	fname := fileNames.HTTPDotYAML
	oPath := filepath.Join(folderpaths.Output, fname)
	return templates.ProcessTemplate(fname, oPath, templates.HTTPSettingsYAML, data, appPaths)
}
