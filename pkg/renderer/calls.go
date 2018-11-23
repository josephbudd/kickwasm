package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
	"github.com/josephbudd/kickwasm/pkg/tap"
)

func createCalls(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	folderpaths := appPaths.GetPaths()
	data := struct {
		ApplicationGitPath                 string
		ImportDomainDataCallIDs            string
		ImportDomainImplementationsCalling string
		ImportDomainInterfacesStorers      string
		ImportDomainInterfacesCallers      string
		ImportDomainTypes                  string
	}{
		ApplicationGitPath:                 builder.ImportPath,
		ImportDomainDataCallIDs:            folderpaths.ImportDomainDataCallIDs,
		ImportDomainImplementationsCalling: folderpaths.ImportDomainImplementationsCalling,
		ImportDomainInterfacesStorers:      folderpaths.ImportDomainInterfacesStorers,
		ImportDomainInterfacesCallers:      folderpaths.ImportDomainInterfacesCallers,
		ImportDomainTypes:                  folderpaths.ImportDomainTypes,
	}
	fname := "map.go"
	oPath := filepath.Join(folderpaths.OutputRendererCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsMapGo, data, appPaths); err != nil {
		return err
	}
	fname = "log.go"
	oPath = filepath.Join(folderpaths.OutputRendererCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsLogGo, data, appPaths); err != nil {
		return err
	}
	fname = "exampleGo.txt"
	oPath = filepath.Join(folderpaths.OutputRendererCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsExampleGoTxt, data, appPaths); err != nil {
		return err
	}
	return nil
}
