package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createVSCode(appPaths paths.ApplicationPathsI, appName string) (err error) {
	folderpaths := appPaths.GetPaths()
	filenames := appPaths.GetFileNames()
	var fname, path string
	fname = appName + "_" + filenames.VSCodeRPWorkSpaceJSON
	path = filepath.Join(folderpaths.OutputRenderer, fname)
	if err = appPaths.WriteFile(path, []byte(templates.VSCodeWorkSpaceJSON)); err != nil {
		return
	}
	return
}
