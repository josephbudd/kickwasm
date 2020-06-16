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
		ApplicationGitPath               string
		Stores                           []string
		ImportRendererFrameworkCallBack  string
		ImportRendererFramework          string
		ImportRendererPaneling           string
		ImportRendererFrameworkViewTools string
		ImportRendererFrameworkLocation  string
		ImportRendererFrameworkLPC       string
		ImportDomainLPCMessage           string
	}{
		ApplicationGitPath:               builder.ImportPath,
		ImportRendererFrameworkCallBack:  folderpaths.ImportRendererFrameworkCallBack,
		ImportRendererFramework:          folderpaths.ImportRendererFramework,
		ImportRendererPaneling:           folderpaths.ImportRendererPaneling,
		ImportRendererFrameworkViewTools: folderpaths.ImportRendererFrameworkViewTools,
		ImportRendererFrameworkLocation:  folderpaths.ImportRendererFrameworkLocation,
		ImportRendererFrameworkLPC:       folderpaths.ImportRendererFrameworkLPC,
		ImportDomainLPCMessage:           folderpaths.ImportDomainLPCMessage,
	}
	fileNames := paths.GetFileNames()
	fname := fileNames.UCMainDotGo
	oPath := filepath.Join(folderpaths.OutputRenderer, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths)
}
