package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

// RebuildStoreInstructions rebuilds the store instructions.
func RebuildStoreInstructions(appPaths paths.ApplicationPathsI, storeNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	filenames := appPaths.GetFileNames()

	data := &struct {
		Stores []string
	}{
		Stores: storeNames,
	}
	fname := filenames.InstructionsDotTXT
	oPath := filepath.Join(folderpaths.OutputDomainStore, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.StoreInstructionsTXT, data, appPaths)
	return
}
