package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createPanelMapGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "panelMap.go"
	oPath := filepath.Join(folderpaths.Output, fname)
	return templates.ProcessTemplate(fname, oPath, templates.PanelMapGo, data, appPaths)
}
