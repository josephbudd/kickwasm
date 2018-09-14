package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createAboutGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	// mainprocess/services/about.go
	contents := templates.GetServicesAboutGo()
	fname := "about.go"
	oPath := filepath.Join(folderpaths.OutputMainProcessServicesAbout, fname)
	return appPaths.WriteFile(oPath, []byte(contents))
}
