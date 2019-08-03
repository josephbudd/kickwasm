package sitepack

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/sitepack/templates"
)

// Create creates all of the renderer files.
func Create(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	folderNames := appPaths.GetFolderNames()
	fname := folderNames.SitePack + ".go"
	oPath := filepath.Join(folderpaths.OutputSitePack, fname)
	data := &struct {
		PackageName string
	}{
		PackageName: folderNames.SitePack,
	}
	err = templates.ProcessTemplate(fname, oPath, templates.SitePack, data, appPaths)
	return
}
