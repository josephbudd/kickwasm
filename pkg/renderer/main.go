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
		ApplicationGitPath                       string
		Stores                                   []string
		ImportRendererCallClient                 string
		ImportRendererCalls                      string
		ImportRendererImplementationsPanelHelper string
		ImportRendererNotJS                      string
		ImportRendererViewTools                  string
		ImportDomainImplementationsCalling       string
		ImportDomainDataSettings                 string
		ImportDomainDataCallIDs                  string
		ImportDomainDataLogLevels                string
		ImportDomainTypes                        string
	}{
		ApplicationGitPath: builder.ImportPath,
		Stores:             builder.Stores,
		ImportRendererCallClient:                 folderpaths.ImportRendererCallClient,
		ImportRendererCalls:                      folderpaths.ImportRendererCalls,
		ImportRendererImplementationsPanelHelper: folderpaths.ImportRendererImplementationsPanelHelper,
		ImportRendererNotJS:                      folderpaths.ImportRendererNotJS,
		ImportRendererViewTools:                  folderpaths.ImportRendererViewTools,
		ImportDomainImplementationsCalling:       folderpaths.ImportDomainImplementationsCalling,
		ImportDomainDataSettings:                 folderpaths.ImportDomainDataSettings,
		ImportDomainDataCallIDs:                  folderpaths.ImportDomainDataCallIDs,
		ImportDomainDataLogLevels:                folderpaths.ImportDomainDataLogLevels,
		ImportDomainTypes:                        folderpaths.ImportDomainTypes,
	}
	fname := "main.go"
	oPath := filepath.Join(folderpaths.OutputRenderer, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths)
}
