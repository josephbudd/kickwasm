package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createLocation(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	var fname string
	var oPath string
	fname = fileNames.HostPortDotGo
	oPath = filepath.Join(folderpaths.OutputRendererLocation, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.HostPortGo)); err != nil {
		return
	}
	return
}
