package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createRecordsRecordsGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "records.go"
	oPath := filepath.Join(folderpaths.OutputMainProcessDataRecords, fname)
	return templates.ProcessTemplate(fname, oPath, templates.RecordsRecordsGo, data, appPaths)
}
