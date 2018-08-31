package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createAboutGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	// mainprocess/about/about.go
	contents := templates.GetAboutAboutGo()
	fname := "about.go"
	oPath := filepath.Join(folderpaths.OutputMainProcessServicesAbout, fname)
	if err := appPaths.WriteFile(oPath, []byte(contents)); err != nil {
		return err
	}
	// mainprocess/calls/about.go
	// execute the template
	oPath = filepath.Join(folderpaths.OutputMainProcessTransportsCalls, fname)
	return templates.ProcessTemplate(fname, oPath, templates.CallsAboutGo, data, appPaths)
}
