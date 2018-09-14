package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/paths"
)

// Create the files in the domain/implementations/calling/ folder.
func createCallingGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "map.go"
	oPath := filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.ImplementationsCallingMapGo, data, appPaths); err != nil {
		return err
	}
	fname = "call.go"
	oPath = filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.ImplementationsCallGo, data, appPaths); err != nil {
		return err
	}
	fname = "log.go"
	oPath = filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallingLogGo, data, appPaths); err != nil {
		return err
	}
	fname = "example_go.txt"
	oPath = filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallingExampleGoTxt, data, appPaths); err != nil {
		return err
	}
	if data.AddAbout {
		fname = "getAbout.go"
		oPath = filepath.Join(folderpaths.OutputDomainImplementationsCalling, fname)
		if err := templates.ProcessTemplate(fname, oPath, templates.CallingGetAboutGo, data, appPaths); err != nil {
			return err
		}
	}
	return nil
}
