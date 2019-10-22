package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createVSCode(appPaths paths.ApplicationPathsI, appName string) (err error) {
	folderpaths := appPaths.GetPaths()
	filenames := appPaths.GetFileNames()
	var fname, path string

	fname = appName + "_" + filenames.VSCodeMPWorkSpaceJSON
	path = filepath.Join(folderpaths.OutputMainProcess, fname)
	if err = appPaths.WriteFile(path, []byte(templates.VSCodeWorkSpaceJSON)); err != nil {
		return
	}
	return
}
