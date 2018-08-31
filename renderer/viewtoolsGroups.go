package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
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
	oPath := filepath.Join(folderpaths.OutputRendererWASMViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsGroups, data, appPaths)
}
