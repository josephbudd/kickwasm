package domain

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/pkg"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
)

type templateData struct {
	BackTick string

	ApplicationName                    string
	ApplicationGitPath                 string
	Stores                             []string
	ServiceNames                       []string
	LowerCamelCase                     func(string) string
	CamelCase                          func(string) string
	ServiceTemplatePanelNames          string
	ServiceEmptyInsidePanelNamePathMap string

	ImportDomainInterfacesCallers          string
	ImportDomainInterfacesStorers          string
	ImportDomainDataFilepaths              string
	ImportDomainDataCallParams             string
	ImportDomainTypes                      string
	ImportDomainImplementationsCalling     string
	ImportDomainImplementationsStoringBolt string

	OutputRendererPanels string

	FileNames   *paths.FileNames
	FolderNames *paths.FolderNames

	Host string
	Port uint
}

// Create creates main process folder files from templates.
func Create(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	parts := strings.Split(builder.ImportPath, "/")
	appname := parts[len(parts)-1]
	imports := paths.GetImports()
	data := &templateData{
		BackTick: "`",

		ApplicationName:                    appname,
		ApplicationGitPath:                 builder.ImportPath,
		Stores:                             builder.Stores,
		ServiceNames:                       builder.GenerateServiceNames(),
		LowerCamelCase:                     cases.LowerCamelCase,
		CamelCase:                          cases.CamelCase,
		ServiceEmptyInsidePanelNamePathMap: strings.Replace(fmt.Sprintf("%#v", builder.GenerateServiceEmptyInsidePanelNamePathMap()), ":", ": ", -1),
		ServiceTemplatePanelNames:          fmt.Sprintf("%#v", builder.GenerateServiceTemplatePanelName()),

		// domain interfaces

		ImportDomainInterfacesCallers:          imports.ImportDomainInterfacesCallers,
		ImportDomainInterfacesStorers:          imports.ImportDomainInterfacesStorers,
		ImportDomainDataFilepaths:              imports.ImportDomainDataFilepaths,
		ImportDomainDataCallParams:             imports.ImportDomainDataCallParams,
		ImportDomainTypes:                      imports.ImportDomainTypes,
		ImportDomainImplementationsCalling:     imports.ImportDomainImplementationsCalling,
		ImportDomainImplementationsStoringBolt: imports.ImportDomainImplementationsStoringBolt,

		// output renderer

		OutputRendererPanels: folderpaths.OutputRendererPanels,

		FileNames:   paths.GetFileNames(),
		FolderNames: paths.GetFolderNames(),

		// settings.yaml

		Host: pkg.LocalHost,
		Port: pkg.LocalPort,
	}
	if err = createInterfacesCallInterfaceGo(appPaths); err != nil {
		return
	}
	if err = createDataFilePathsGo(appPaths, data); err != nil {
		return
	}
	if err = createInterfacesStoreInterfaceGo(appPaths, data); err != nil {
		return
	}
	if err = createCallingGo(appPaths, data); err != nil {
		return
	}
	if err = createStoreBoltStoresGo(appPaths, data); err != nil {
		return
	}
	if err = createTypes(appPaths, data); err != nil {
		return
	}
	if err = createSettingsYAML(appPaths, data); err != nil {
		return
	}
	return
}
