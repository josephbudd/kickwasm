package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createPanelHelper(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	// put the example in both interfaces and implementations.
	fname := "example.txt"
	oPath := filepath.Join(folderpaths.OutputRendererInterfacesPanelHelper, fname)
	if err := appPaths.WriteFile(oPath, []byte(templates.PanelHelperExample)); err != nil {
		return err
	}
	oPath = filepath.Join(folderpaths.OutputRendererImplementationsPanelHelper, fname)
	if err := appPaths.WriteFile(oPath, []byte(templates.PanelHelperExample)); err != nil {
		return err
	}
	// interface
	fname = "helper.go"
	oPath = filepath.Join(folderpaths.OutputRendererInterfacesPanelHelper, fname)
	if err := appPaths.WriteFile(oPath, []byte(templates.PanelHelperInterface)); err != nil {
		return err
	}
	// implementations
	fname = "noHelp.go"
	oPath = filepath.Join(folderpaths.OutputRendererImplementationsPanelHelper, fname)
	if err := appPaths.WriteFile(oPath, []byte(templates.PanelHelperImplementation)); err != nil {
		return err
	}
	return nil
}
