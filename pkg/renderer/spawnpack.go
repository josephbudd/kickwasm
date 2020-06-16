package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createSpawnPack(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	folderNames := appPaths.GetFolderNames()
	fname := folderNames.SpawnPack + ".go"
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkSpawnPack, fname)
	data := &struct {
		PackageName string
	}{
		PackageName: folderNames.SpawnPack,
	}
	err = templates.ProcessTemplate(fname, oPath, templates.SpawnPack, data, appPaths)
	return
}
