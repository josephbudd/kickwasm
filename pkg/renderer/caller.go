package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createCallerGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	fname := fileNames.ClientDotGo
	oPath := filepath.Join(folderpaths.OutputRendererCallClient, fname)
	data := struct {
		ApplicationGitPath            string
		ImportRendererNotJS           string
		ImportRendererViewTools       string
		ImportDomainTypes             string
		ImportDomainInterfacesCallers string
	}{
		ApplicationGitPath:            builder.ImportPath,
		ImportRendererNotJS:           folderpaths.ImportRendererNotJS,
		ImportRendererViewTools:       folderpaths.ImportRendererViewTools,
		ImportDomainTypes:             folderpaths.ImportDomainTypes,
		ImportDomainInterfacesCallers: folderpaths.ImportDomainInterfacesCallers,
	}
	return templates.ProcessTemplate(fname, oPath, templates.ClientGo, data, appPaths)
}
