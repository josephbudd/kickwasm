package mainprocess

import (
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/format"
	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createMain(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()

	fname := fileNames.MainDotGo
	oPath := filepath.Join(folderpaths.OutputMainProcess, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths); err != nil {
		return
	}

	fname = fileNames.PanelMapDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcess, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.PanelMapGo, data, appPaths); err != nil {
		return
	}

	fname = fileNames.ServeDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcess, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ServeGo, data, appPaths); err != nil {
		return
	}

	err = RebuildStoresGo(appPaths, data.ApplicationGitPath, nil, nil, nil)

	return
}

// RebuildStoresGo rebuilds stores.go.
func RebuildStoresGo(appPaths paths.ApplicationPathsI, importPath string, boltStoreNames, remoteDBNames, remoteRecordNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	data := struct {
		ApplicationGitPath        string
		ImportDomainDataFilepaths string
		ImportDomainStore         string
		ImportDomainStoreStoring  string
		BoltStores                []string
		RemoteDBs                 []string
		RemoteRecords             []string
		SameWidth                 func([]string) []string
		TrimSpace                 func(string) string
	}{
		ApplicationGitPath:        importPath,
		ImportDomainDataFilepaths: folderpaths.ImportDomainDataFilepaths,
		ImportDomainStore:         folderpaths.ImportDomainStore,
		ImportDomainStoreStoring:  folderpaths.ImportDomainStoreStoring,
		BoltStores:                boltStoreNames,
		RemoteDBs:                 remoteDBNames,
		RemoteRecords:             remoteRecordNames,
		SameWidth:                 format.SameWidth,
		TrimSpace:                 strings.TrimSpace,
	}
	fname := fileNames.StoresDotGo
	oPath := filepath.Join(folderpaths.OutputMainProcess, fname)
	var temp string
	if len(boltStoreNames) == 0 && len(remoteDBNames) == 0 {
		temp = templates.NoStoresGo
	} else {
		temp = templates.StoresGo
	}
	err = templates.ProcessTemplate(fname, oPath, temp, data, appPaths)
	return
}
