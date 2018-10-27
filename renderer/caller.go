package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

func createCallerGo(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	folderpaths := appPaths.GetPaths()
	fname := "client.go"
	oPath := filepath.Join(folderpaths.OutputRendererCallClient, fname)
	data := struct {
		ApplicationGitPath            string
		ImportRendererViewTools       string
		ImportDomainTypes             string
		ImportDomainInterfacesCallers string
	}{
		ApplicationGitPath:            builder.ImportPath,
		ImportRendererViewTools:       folderpaths.ImportRendererViewTools,
		ImportDomainTypes:             folderpaths.ImportDomainTypes,
		ImportDomainInterfacesCallers: folderpaths.ImportDomainInterfacesCallers,
	}
	return templates.ProcessTemplate(fname, oPath, templates.ClientGo, data, appPaths)
}
