package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/paths"
)

// Create the files in the domain/implementations/calling/ folder.
func createCallingGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "mainprocess.go"
	oPath := filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.ImplementationsMainProcessGo, data, appPaths); err != nil {
		return err
	}
	fname = "renderer.go"
	oPath = filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.ImplementationsRendererGo, data, appPaths); err != nil {
		return err
	}
	return nil
}
