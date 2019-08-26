package domain

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/format"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createStoreStoresGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	err = RebuildStoreStores(appPaths, data.ApplicationGitPath, nil, nil, nil)
	return
}

// RebuildStoreStores recreates domain/store/stores.go.
func RebuildStoreStores(appPaths paths.ApplicationPathsI, importPath string, boltStoreNames, remoteDBNames, remoteRecordNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()

	data := &struct {
		BoltStores               []string
		RemoteDBs                []string
		RemoteRecords            []string
		LowerCamelCase           func(string) string
		SameWidth                func([]string) []string
		TrimSpace                func(string) string
		ApplicationGitPath       string
		ImportDomainStoreStoring string
	}{
		BoltStores:               boltStoreNames,
		RemoteDBs:                remoteDBNames,
		RemoteRecords:            remoteRecordNames,
		LowerCamelCase:           cases.LowerCamelCase,
		SameWidth:                format.SameWidth,
		TrimSpace:                strings.TrimSpace,
		ApplicationGitPath:       importPath,
		ImportDomainStoreStoring: folderpaths.ImportDomainStoreStoring,
	}
	fname := fileNames.StoresDotGo
	oPath := filepath.Join(folderpaths.OutputDomainStore, fname)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	err = templates.ProcessTemplate(fname, oPath, templates.StoreStoresGo, data, appPaths)
	return
}
