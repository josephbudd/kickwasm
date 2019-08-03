package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createMain(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()

	fname := fileNames.MainDotGo
	oPath := filepath.Join(folderpaths.Output, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths); err != nil {
		return
	}

	fname = fileNames.PanelMapDotGo
	oPath = filepath.Join(folderpaths.Output, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.PanelMapGo, data, appPaths); err != nil {
		return
	}

	fname = fileNames.ServeDotGo
	oPath = filepath.Join(folderpaths.Output, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ServeGo, data, appPaths); err != nil {
		return
	}

	err = RebuildStoresGo(appPaths, data.ApplicationGitPath, data.Stores)

	return
}

// RebuildStoresGo rebuilds stores.go.
func RebuildStoresGo(appPaths paths.ApplicationPathsI, importPath string, storesNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	data := struct {
		ApplicationGitPath        string
		ImportDomainDataFilepaths string
		ImportDomainStore         string
		ImportDomainStoreStoring  string
		Stores                    []string
	}{
		ApplicationGitPath:        importPath,
		ImportDomainDataFilepaths: folderpaths.ImportDomainDataFilepaths,
		ImportDomainStore:         folderpaths.ImportDomainStore,
		ImportDomainStoreStoring:  folderpaths.ImportDomainStoreStoring,
		Stores:                    storesNames,
	}
	fname := fileNames.StoresDotGo
	oPath := filepath.Join(folderpaths.Output, fname)
	var temp string
	if len(storesNames) == 0 {
		temp = templates.NoStoresGo
	} else {
		temp = templates.StoresGo
	}
	err = templates.ProcessTemplate(fname, oPath, temp, data, appPaths)
	return
}
