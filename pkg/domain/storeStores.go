package domain

import (
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createStoreStoresGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	err = RebuildStoreStores(appPaths, data.ApplicationGitPath, data.Stores)
	return
}

// RebuildStoreStores recreates domain/store/stores.go.
func RebuildStoreStores(appPaths paths.ApplicationPathsI, importPath string, storeNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()

	data := &struct {
		Stores                  []string
		LowerCamelCase          func(string) string
		ApplicationGitPath      string
		ImportDomainStoreStorer string
	}{
		Stores:                  storeNames,
		LowerCamelCase:          cases.LowerCamelCase,
		ApplicationGitPath:      importPath,
		ImportDomainStoreStorer: folderpaths.ImportDomainStoreStorer,
	}
	fname := fileNames.StoresDotGo
	oPath := filepath.Join(folderpaths.OutputDomainStore, fname)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	err = templates.ProcessTemplate(fname, oPath, templates.StoreStoresGo, data, appPaths)
	return
}
