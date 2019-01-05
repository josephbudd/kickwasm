package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createCalls(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	fname := fileNames.MapDotGo
	oPath := filepath.Join(folderpaths.OutputMainProcessCalls, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallsMapGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.LogDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessCalls, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallsLogGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.ExampleGoDotTXT
	oPath = filepath.Join(folderpaths.OutputMainProcessCalls, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallsExampleGoTxt, data, appPaths); err != nil {
		return
	}
	return nil
}
