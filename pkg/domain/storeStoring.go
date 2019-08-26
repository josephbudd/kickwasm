package domain

import (
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

// Create the files in the domain/implementations/store/boltstoring/ folder.
func createStoreStoring(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	data2 := struct {
		Store                    string
		ApplicationGitPath       string
		ImportDomainStoreRecord  string
		ImportDomainStoreStorer  string
		ImportDomainStoreStoring string
		LowerCamelCase           func(string) string
	}{
		ApplicationGitPath:       data.ApplicationGitPath,
		ImportDomainStoreRecord:  data.ImportDomainStoreRecord,
		ImportDomainStoreStorer:  data.ImportDomainStoreStorer,
		ImportDomainStoreStoring: data.ImportDomainStoreStoring,
		LowerCamelCase:           data.LowerCamelCase,
	}
	for _, store := range data.Stores {
		data2.Store = store
		// These are editable files so they are capped.
		fname := store + ".go"
		oPath := filepath.Join(folderpaths.OutputDomainStoreStoring, fname)
		if err = templates.ProcessTemplate(fname, oPath, templates.LocalBoltStoreStoringGo, data2, appPaths); err != nil {
			return
		}
	}
	return
}

// CreateStoreStoring creates a single store's storing file in domain/store/storeing/
func CreateStoreStoring(appPaths paths.ApplicationPathsI, importPath string, storeName string) (err error) {
	folderpaths := appPaths.GetPaths()

	data := struct {
		Store                   string
		ApplicationGitPath      string
		ImportDomainStoreRecord string
		ImportDomainStoreStorer string
		LowerCamelCase          func(string) string
	}{
		Store:                   storeName,
		ApplicationGitPath:      importPath,
		ImportDomainStoreRecord: folderpaths.ImportDomainStoreRecord,
		ImportDomainStoreStorer: folderpaths.ImportDomainStoreStorer,
		LowerCamelCase:          cases.LowerCamelCase,
	}
	// This is an editable file so it is capped.
	fname := storeName + ".go"
	oPath := filepath.Join(folderpaths.OutputDomainStoreStoring, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.LocalBoltStoreStoringGo, data, appPaths)
	return
}

// CreateRemoteStoreStoring creates a remote database's storing file in domain/store/storing/
func CreateRemoteStoreStoring(appPaths paths.ApplicationPathsI, importPath string, storeName string) (err error) {
	folderpaths := appPaths.GetPaths()

	data := struct {
		Store                   string
		ApplicationGitPath      string
		ImportDomainStoreRecord string
		ImportDomainStoreStorer string
		LowerCamelCase          func(string) string
	}{
		Store:                   storeName,
		ApplicationGitPath:      importPath,
		ImportDomainStoreRecord: folderpaths.ImportDomainStoreRecord,
		ImportDomainStoreStorer: folderpaths.ImportDomainStoreStorer,
		LowerCamelCase:          cases.LowerCamelCase,
	}
	// This is an editable file so it is capped.
	fname := storeName + ".go"
	oPath := filepath.Join(folderpaths.OutputDomainStoreStoring, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.RemoteStoreStoringGo, data, appPaths)
	return
}

// DeleteStoreStoring removes a single store's storing file in domain/store/storeing/
func DeleteStoreStoring(appPaths paths.ApplicationPathsI, importPath string, storeName string) (err error) {
	folderpaths := appPaths.GetPaths()

	// This is an editable file so it is capped.
	fname := storeName + ".go"
	oPath := filepath.Join(folderpaths.OutputDomainStoreStoring, fname)
	if err = os.Remove(oPath); err != nil && os.IsNotExist(err) {
		err = nil
	}
	return
}
