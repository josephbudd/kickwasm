package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
	"github.com/josephbudd/kickwasm/pkg/tap"
)

func createViewToolsGroupsGo(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	data := &struct {
		ServiceButtonPanelGroups map[string][]*tap.ButtonPanelGroup
	}{
		ServiceButtonPanelGroups: builder.GenerateServiceButtonPanelGroups(),
	}
	// execute the template
	folderpaths := appPaths.GetPaths()
	fname := "groups.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsGroups, data, appPaths)
}
