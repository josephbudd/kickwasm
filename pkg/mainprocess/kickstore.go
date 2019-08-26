package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createKickstore(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()

	fname := fileNames.StoresDotYAML
	oPath := filepath.Join(folderpaths.OutputDotKickstore, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.KickStoreStorage, data, appPaths); err != nil {
		return
	}
	return
}
