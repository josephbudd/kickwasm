package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createCalls(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
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
	fileNames := paths.GetFileNames()
	fname := fileNames.MapDotGo
	oPath := filepath.Join(folderpaths.OutputRendererCalls, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallsMapGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.LogDotGo
	oPath = filepath.Join(folderpaths.OutputRendererCalls, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallsLogGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.ExampleGoDotTXT
	oPath = filepath.Join(folderpaths.OutputRendererCalls, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallsExampleGoTxt, data, appPaths); err != nil {
		return
	}
	return
}
