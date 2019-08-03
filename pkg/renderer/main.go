package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createMainGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		ApplicationGitPath      string
		Stores                  []string
		ImportRendererPaneling  string
		ImportRendererNotJS     string
		ImportRendererViewTools string
		ImportRendererLPC       string
	}{
		ApplicationGitPath:      builder.ImportPath,
		Stores:                  builder.Stores,
		ImportRendererPaneling:  folderpaths.ImportRendererPaneling,
		ImportRendererNotJS:     folderpaths.ImportRendererNotJS,
		ImportRendererViewTools: folderpaths.ImportRendererViewTools,
		ImportRendererLPC:       folderpaths.ImportRendererLPC,
	}
	fileNames := paths.GetFileNames()
	fname := fileNames.MainDotGo
	oPath := filepath.Join(folderpaths.OutputRenderer, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths)
}
