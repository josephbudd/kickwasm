package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsEvent(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()

	var fname string
	var oPath string

	// build.sh
	fname = fileNames.EventDotGo
	oPath = filepath.Join(folderpaths.OutputRendererViewTools, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ViewToolsEventGo, nil, appPaths)
	return
}
