package domain

import (
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createStoreStorer(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	rdata := &struct {
		Store                   string
		LowerCamelCase          func(string) string
		ApplicationGitPath      string
		ImportDomainStoreRecord string
	}{
		LowerCamelCase:          data.LowerCamelCase,
		ApplicationGitPath:      data.ApplicationGitPath,
		ImportDomainStoreRecord: folderpaths.ImportDomainStoreRecord,
	}
	for _, store := range data.Stores {
		// These are editable files to the are capped.
		//fname := fmt.Sprintf("%s.go", data.LowerCamelCase(store))
		fname := store + ".go"
		oPath := filepath.Join(folderpaths.OutputDomainStoreStorer, fname)
		rdata.Store = store
		if err = templates.ProcessTemplate(fname, oPath, templates.StoreStorerGo, rdata, appPaths); err != nil {
			return
		}
	}
	return
}

// CreateStoreStorer creates a single store's file in domain/store/storer/.
func CreateStoreStorer(appPaths paths.ApplicationPathsI, importPath string, storeName string) (err error) {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		Store                   string
		LowerCamelCase          func(string) string
		ApplicationGitPath      string
		ImportDomainStoreRecord string
	}{
		Store:                   storeName,
		LowerCamelCase:          cases.LowerCamelCase,
		ApplicationGitPath:      importPath,
		ImportDomainStoreRecord: folderpaths.ImportDomainStoreRecord,
	}
	// These are editable files so the are capped.
	fname := storeName + ".go"
	oPath := filepath.Join(folderpaths.OutputDomainStoreStorer, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.StoreStorerGo, data, appPaths)
	return
}

// DeleteStoreStorer deletes a single store's file in domain/store/storer/.
func DeleteStoreStorer(appPaths paths.ApplicationPathsI, storeName string) (err error) {
	folderpaths := appPaths.GetPaths()
	// These are editable files so the are capped.
	fname := storeName + ".go"
	oPath := filepath.Join(folderpaths.OutputDomainStoreStorer, fname)
	if err = os.Remove(oPath); err != nil && os.IsNotExist(err) {
		err = nil
	}
	return
}
