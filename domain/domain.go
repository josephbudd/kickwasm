package domain

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/tap"
)

type templateData struct {
	Port                               uint
	Host                               string
	ApplicationName                    string
	ApplicationGitPath                 string
	Stores                             []string
	ServiceNames                       []string
	LowerCamelCase                     func(string) string
	CamelCase                          func(string) string
	AddAbout                           bool
	ServiceTemplatePanelNames          string
	ServiceEmptyInsidePanelNamePathMap string
	HeadTemplateFile                   string

	ImportDomainInterfacesCallers          string
	ImportDomainInterfacesStorers          string
	ImportDomainDataFilepaths              string
	ImportDomainDataCallParams             string
	ImportDomainTypes                      string
	ImportDomainImplementationsCalling     string
	ImportDomainImplementationsStoringBolt string

	ImportMainProcessServicesAbout string

	OutputRendererPanels string
}

// Create creates main process folder files from templates.
func Create(appPaths paths.ApplicationPathsI, builder *tap.Builder, addAbout bool, headTemplateFile string, host string, port uint) error {
	folderpaths := appPaths.GetPaths()
	parts := strings.Split(builder.ImportPath, "/")
	appname := parts[len(parts)-1]
	data := &templateData{
		Port:               port,
		Host:               host,
		ApplicationName:    appname,
		ApplicationGitPath: builder.ImportPath,
		Stores:             builder.Stores,
		ServiceNames:       builder.GenerateServiceNames(),
		LowerCamelCase:     cases.LowerCamelCase,
		CamelCase:          cases.CamelCase,
		AddAbout:           addAbout,
		ServiceEmptyInsidePanelNamePathMap: strings.Replace(fmt.Sprintf("%#v", builder.GenerateServiceEmptyInsidePanelNamePathMap()), ":", ": ", -1),
		ServiceTemplatePanelNames:          fmt.Sprintf("%#v", builder.GenerateServiceTemplatePanelName()),
		HeadTemplateFile:                   headTemplateFile,

		// domain interfaces

		ImportDomainInterfacesCallers:          paths.ImportDomainInterfacesCallers,
		ImportDomainInterfacesStorers:          paths.ImportDomainInterfacesStorers,
		ImportDomainDataFilepaths:              paths.ImportDomainDataFilepaths,
		ImportDomainDataCallParams:             paths.ImportDomainDataCallParams,
		ImportDomainTypes:                      paths.ImportDomainTypes,
		ImportDomainImplementationsCalling:     paths.ImportDomainImplementationsCalling,
		ImportDomainImplementationsStoringBolt: paths.ImportDomainImplementationsStoringBolt,

		ImportMainProcessServicesAbout: paths.ImportMainProcessServicesAbout,

		// output renderer

		OutputRendererPanels: folderpaths.OutputRendererPanels,
	}
	if err := createInterfacesCallInterfaceGo(appPaths); err != nil {
		return err
	}
	if err := createDataFilePathsGo(appPaths, data); err != nil {
		return err
	}
	if err := createInterfacesStoreInterfaceGo(appPaths, data); err != nil {
		return err
	}
	if err := createCallingGo(appPaths, data); err != nil {
		return err
	}
	if err := createStoreBoltStoresGo(appPaths, data); err != nil {
		return err
	}
	if err := createTypesCallsGo(appPaths, data); err != nil {
		return err
	}
	return nil
}
