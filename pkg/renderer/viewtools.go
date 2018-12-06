package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewTools(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	if err := createViewToolsGo(appPaths, builder); err != nil {
		return err
	}
	if err := createViewToolsCloserGo(appPaths); err != nil {
		return err
	}
	if err := createViewToolsGroupsGo(appPaths, builder); err != nil {
		return err
	}
	if err := createViewToolsHelpersGo(appPaths); err != nil {
		return err
	}
	if err := createViewToolsHideShowGo(appPaths); err != nil {
		return err
	}
	if err := createViewToolsModalGo(appPaths); err != nil {
		return err
	}
	if err := createResizeGo(appPaths, builder); err != nil {
		return err
	}
	if err := createViewToolsSliderGo(appPaths); err != nil {
		return err
	}
	if err := createViewToolsTabBarGo(appPaths, builder); err != nil {
		return err
	}
	return nil
}

func createViewToolsGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		IDs                 *project.IDs
		Classes             *project.Classes
		Attributes          *project.Attributes
		ApplicationGitPath  string
		ImportRendererNotJS string
	}{
		IDs:                 builder.IDs,
		Classes:             builder.Classes,
		Attributes:          builder.Attributes,
		ApplicationGitPath:  builder.ImportPath,
		ImportRendererNotJS: folderpaths.ImportRendererNotJS,
	}
	// execute the template
	fname := "viewtools.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewTools, data, appPaths)
}