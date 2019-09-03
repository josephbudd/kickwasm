package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/project"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsMarkupGo(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	fileNames := appPaths.GetFileNames()
	folderpaths := appPaths.GetPaths()
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fileNames.MarkupDotGo)
	err = appPaths.WriteFile(oPath, []byte(templates.ViewToolsMarkupGo))
	return
}
