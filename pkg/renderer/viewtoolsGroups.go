package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsGroupsGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	data := &struct {
		ServiceButtonPanelGroups map[string][]*project.ButtonPanelGroup
	}{
		ServiceButtonPanelGroups: builder.GenerateServiceButtonPanelGroups(),
	}
	// execute the template
	folderpaths := appPaths.GetPaths()
	fname := "groups.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsGroups, data, appPaths)
}
