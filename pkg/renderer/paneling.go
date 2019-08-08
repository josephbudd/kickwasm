package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createPaneling(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()

	fname := fileNames.InstructionsDotTXT
	oPath := filepath.Join(folderpaths.OutputRendererPaneling, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.PanelHelperInstructions)); err != nil {
		return
	}
	fname = fileNames.HelpingDotGo
	oPath = filepath.Join(folderpaths.OutputRendererPaneling, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.PanelHelperImplementation)); err != nil {
		return
	}
	return
}
