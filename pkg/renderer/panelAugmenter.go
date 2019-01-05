package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createPanelHelper(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	// put the example in both interfaces and implementations.
	fileNames := paths.GetFileNames()
	fname := fileNames.ExampleDotTXT
	oPath := filepath.Join(folderpaths.OutputRendererInterfacesPanelHelper, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.PanelHelperExample)); err != nil {
		return
	}
	oPath = filepath.Join(folderpaths.OutputRendererImplementationsPanelHelper, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.PanelHelperExample)); err != nil {
		return
	}
	// interface
	fname = fileNames.HelperDotGo
	oPath = filepath.Join(folderpaths.OutputRendererInterfacesPanelHelper, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.PanelHelperInterface)); err != nil {
		return
	}
	// implementations
	fname = fileNames.NoHelpDotGo
	oPath = filepath.Join(folderpaths.OutputRendererImplementationsPanelHelper, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.PanelHelperImplementation)); err != nil {
		return
	}
	return
}
