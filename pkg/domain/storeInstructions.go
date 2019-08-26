package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

// RebuildStoreInstructions rebuilds the store instructions.
func RebuildStoreInstructions(appPaths paths.ApplicationPathsI, boltStoreNames, remoteDBNames, remoteRecordNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	filenames := appPaths.GetFileNames()

	data := &struct {
		BoltStores    []string
		RemoteDBs     []string
		RemoteRecords []string
	}{
		BoltStores:    boltStoreNames,
		RemoteDBs:     remoteDBNames,
		RemoteRecords: remoteRecordNames,
	}
	fname := filenames.InstructionsDotTXT
	oPath := filepath.Join(folderpaths.OutputDomainStore, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.StoreInstructionsTXT, data, appPaths)
	return
}
