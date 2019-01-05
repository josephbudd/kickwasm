package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createMainGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	fname := fileNames.MainDotGo
	oPath := filepath.Join(folderpaths.Output, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths)
	return
}

func createPanelMapGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	fname := fileNames.PanelMapDotGo
	oPath := filepath.Join(folderpaths.Output, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.PanelMapGo, data, appPaths)
	return
}

func createServeGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	fname := fileNames.ServeDotGo
	oPath := filepath.Join(folderpaths.Output, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ServeGo, data, appPaths)
	return
}
