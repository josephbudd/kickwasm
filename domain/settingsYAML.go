package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createSettingsYAML(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fname := "http.yaml"
	oPath := filepath.Join(folderpaths.Output, fname)
	return templates.ProcessTemplate(fname, oPath, templates.HTTPSettingsYAML, nil, appPaths)
}
