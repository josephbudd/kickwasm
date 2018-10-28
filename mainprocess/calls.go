package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createCalls(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "map.go"
	oPath := filepath.Join(folderpaths.OutputMainProcessCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsMapGo, data, appPaths); err != nil {
		return err
	}
	fname = "log.go"
	oPath = filepath.Join(folderpaths.OutputMainProcessCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsLogGo, data, appPaths); err != nil {
		return err
	}
	fname = "exampleGo.txt"
	oPath = filepath.Join(folderpaths.OutputMainProcessCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsExampleGoTxt, data, appPaths); err != nil {
		return err
	}
	if data.AddAbout {
		fname = "getAbout.go"
		oPath = filepath.Join(folderpaths.OutputMainProcessCalls, fname)
		if err := templates.ProcessTemplate(fname, oPath, templates.CallsGetAboutGo, data, appPaths); err != nil {
			return err
		}
	}
	return nil
}
