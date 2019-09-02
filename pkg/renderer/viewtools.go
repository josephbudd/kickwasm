package renderer

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewTools(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	if err = createViewToolsGo(appPaths, builder); err != nil {
		return
	}
	if err = createViewToolsCallBackGo(appPaths); err != nil {
		return
	}
	if err = createViewToolsCloserGo(appPaths); err != nil {
		return
	}
	if err = createViewToolsGroupsGo(appPaths, builder); err != nil {
		return
	}
	if err = createViewToolsHelpersGo(appPaths); err != nil {
		return
	}
	if err = createViewToolsHideShowGo(appPaths); err != nil {
		return
	}
	if err = createViewToolsModalGo(appPaths); err != nil {
		return
	}
	if err = createResizeGo(appPaths, builder); err != nil {
		return
	}
	if err = createViewToolsSliderGo(appPaths); err != nil {
		return
	}
	if err = createViewToolsTabBarGo(appPaths, builder); err != nil {
		return
	}
	if err = createViewToolsLocksGo(appPaths); err != nil {
		return
	}
	if err = createViewToolsSpawnTabGo(appPaths, builder); err != nil {
		return
	}
	if err = createViewToolsEvent(appPaths); err != nil {
		return
	}
	if err = createViewToolsMarkupGo(appPaths, builder); err != nil {
		return
	}
	return
}

func createViewToolsGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	panelNameHVScroll := make(map[string]bool, 100)
	servicePanelNamePanelMap := builder.GenerateServicePanelNamePanelMap()
	for _, panelNamePanelMap := range servicePanelNamePanelMap {
		for panelName, panel := range panelNamePanelMap {
			panelNameHVScroll[panelName] = panel.HVScroll
		}
	}
	// GenerateServiceSpawnPanelNamePanelMap
	servicePanelNamePanelMap = builder.GenerateServiceSpawnPanelNamePanelMap()
	for _, panelNamePanelMap := range servicePanelNamePanelMap {
		for panelName, panel := range panelNamePanelMap {
			panelNameHVScroll[panelName] = panel.HVScroll
		}
	}
	data := &struct {
		IDs                   *project.IDs
		Classes               *project.Classes
		Attributes            *project.Attributes
		ApplicationGitPath    string
		ImportRendererNotJS   string
		SpawnIDReplacePattern string
		PanelNameHVScroll     string
		NumberOfMarkupPanels  uint64
	}{
		IDs:                   builder.IDs,
		Classes:               builder.Classes,
		Attributes:            builder.Attributes,
		ApplicationGitPath:    builder.ImportPath,
		ImportRendererNotJS:   folderpaths.ImportRendererNotJS,
		SpawnIDReplacePattern: project.SpawnIDReplacePattern,
		PanelNameHVScroll:     fmt.Sprintf("%#v", panelNameHVScroll),
		NumberOfMarkupPanels:  builder.MarkupPanelCount,
	}
	// execute the template
	fileNames := paths.GetFileNames()
	fname := fileNames.ViewToolsDotGo
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewTools, data, appPaths)
}
