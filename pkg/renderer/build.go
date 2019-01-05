package renderer

import (
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createBuildSH(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	fname := fileNames.BuildDotSH
	oPath := filepath.Join(folderpaths.OutputRenderer, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.BuildDotSH, nil, appPaths)
	os.Chmod(oPath, 0744)
	return
}
