package domain

import (
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createStoreRecord(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	rdata := &struct {
		Store                   string
		ApplicationGitPath      string
		ImportDomainStoreRecord string
	}{
		ApplicationGitPath:      data.ApplicationGitPath,
		ImportDomainStoreRecord: folderpaths.ImportDomainStoreRecord,
	}
	for _, store := range data.Stores {
		// These are editable files so they are capped.
		fname := store + ".go"
		oPath := filepath.Join(folderpaths.OutputDomainStoreRecord, fname)
		rdata.Store = store
		if err = templates.ProcessTemplate(fname, oPath, templates.StoreRecordGo, rdata, appPaths); err != nil {
			return
		}
	}
	return
}

// CreateStoreRecord creates a single store's file in domain/store/storer/.
func CreateStoreRecord(appPaths paths.ApplicationPathsI, importPath string, storeName string) (err error) {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		Store                   string
		ApplicationGitPath      string
		ImportDomainStoreRecord string
	}{
		Store:                   storeName,
		ApplicationGitPath:      importPath,
		ImportDomainStoreRecord: folderpaths.ImportDomainStoreRecord,
	}
	// These are editable files so the are capped.
	fname := storeName + ".go"
	oPath := filepath.Join(folderpaths.OutputDomainStoreRecord, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.StoreRecordGo, data, appPaths)
	return
}

// DeleteStoreRecord deletes a single store's file in domain/store/storer/.
func DeleteStoreRecord(appPaths paths.ApplicationPathsI, storeName string) (err error) {
	folderpaths := appPaths.GetPaths()
	// These are editable files so the are capped.
	fname := storeName + ".go"
	oPath := filepath.Join(folderpaths.OutputDomainStoreRecord, fname)
	if err = os.Remove(oPath); err != nil && os.IsNotExist(err) {
		err = nil
	}
	return
}
