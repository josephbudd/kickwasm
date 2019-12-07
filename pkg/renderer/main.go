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
		ImportRendererCallBack  string
		ImportRendererFramework string
		ImportRendererPaneling  string
		ImportRendererViewTools string
		ImportRendererLocation  string
		ImportRendererLPC       string
		ImportDomainLPCMessage  string
	}{
		ApplicationGitPath:      builder.ImportPath,
		ImportRendererCallBack:  folderpaths.ImportRendererCallBack,
		ImportRendererFramework: folderpaths.ImportRendererFramework,
		ImportRendererPaneling:  folderpaths.ImportRendererPaneling,
		ImportRendererViewTools: folderpaths.ImportRendererViewTools,
		ImportRendererLocation:  folderpaths.ImportRendererLocation,
		ImportRendererLPC:       folderpaths.ImportRendererLPC,
		ImportDomainLPCMessage:  folderpaths.ImportDomainLPCMessage,
	}
	fileNames := paths.GetFileNames()
	fname := fileNames.UCMainDotGo
	oPath := filepath.Join(folderpaths.OutputRenderer, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths)
}
