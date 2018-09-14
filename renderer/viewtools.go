package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

func createViewTools(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
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

func createViewToolsGo(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	data := &struct {
		IDs        *tap.IDs
		Classes    *tap.Classes
		Attributes *tap.Attributes
	}{
		IDs:        builder.IDs,
		Classes:    builder.Classes,
		Attributes: builder.Attributes,
	}
	// execute the template
	folderpaths := appPaths.GetPaths()
	fname := "viewtools.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewTools, data, appPaths)
}
