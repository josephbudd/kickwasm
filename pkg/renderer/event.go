package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func creatEvent(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	var fname string
	var oPath string
	fname = fileNames.EventDotGo
	oPath = filepath.Join(folderpaths.OutputRendererEvent, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.EventGo)); err != nil {
		return
	}
	return
}
