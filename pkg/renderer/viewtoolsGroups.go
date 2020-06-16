package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsGroupsGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		HomeButtonPanelGroups   map[string][]*project.ButtonPanelGroup
		ApplicationGitPath      string
		ImportRendererAPIMarkup string
	}{
		HomeButtonPanelGroups:   builder.GenerateHomeButtonPanelGroups(),
		ApplicationGitPath:      builder.ImportPath,
		ImportRendererAPIMarkup: folderpaths.ImportRendererAPIMarkup,
	}
	// execute the template
	filenames := appPaths.GetFileNames()
	fname := filenames.GroupsDotGo
	oPath := filepath.Join(folderpaths.OutputRendererFrameworkViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsGroups, data, appPaths)
}
