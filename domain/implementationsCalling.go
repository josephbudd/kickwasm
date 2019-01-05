package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

// Create the files in the domain/implementations/calling/ folder.
func createCallingGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fname := "mainprocess.go"
	oPath := filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ImplementationsMainProcessGo, data, appPaths); err != nil {
		return
	}
	fname = "renderer.go"
	oPath = filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ImplementationsRendererGo, data, appPaths); err != nil {
		return
	}
	return
}
