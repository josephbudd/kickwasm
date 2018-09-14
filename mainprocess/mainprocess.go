package mainprocess

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
	Repos                              []string
	ServiceNames                       []string
	LowerCamelCase                     func(string) string
	CamelCase                          func(string) string
	AddAbout                           bool
	ServiceTemplatePanelNames          string
	ServiceEmptyInsidePanelNamePathMap string
	HeadTemplateFile                   string

	ImportDomainInterfacesStorers          string
	ImportDomainDataFilepaths              string
	ImportDomainTypes                      string
	ImportDomainImplementationsCalling     string
	ImportDomainImplementationsStoringBolt string

	ImportMainProcessServicesAbout string
	ImportMainProcessCallServer    string
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
		Repos:              builder.Repos,
		ServiceNames:       builder.GenerateServiceNames(),
		LowerCamelCase:     cases.LowerCamelCase,
		CamelCase:          cases.CamelCase,
		AddAbout:           addAbout,
		ServiceEmptyInsidePanelNamePathMap: strings.Replace(fmt.Sprintf("%#v", builder.GenerateServiceEmptyInsidePanelNamePathMap()), ":", ": ", -1),
		ServiceTemplatePanelNames:          fmt.Sprintf("%#v", builder.GenerateServiceTemplatePanelName()),
		HeadTemplateFile:                   headTemplateFile,

		ImportDomainInterfacesStorers:          folderpaths.ImportDomainInterfacesStorers,
		ImportDomainDataFilepaths:              folderpaths.ImportDomainDataFilepaths,
		ImportDomainTypes:                      folderpaths.ImportDomainTypes,
		ImportDomainImplementationsCalling:     folderpaths.ImportDomainImplementationsCalling,
		ImportDomainImplementationsStoringBolt: folderpaths.ImportDomainImplementationsStoringBolt,
		ImportMainProcessServicesAbout:         folderpaths.ImportMainProcessServicesAbout,
		ImportMainProcessCallServer:            folderpaths.ImportMainProcessCallServer,
	}
	if err := createMainGo(appPaths, data); err != nil {
		return err
	}
	if err := createPanelMapGo(appPaths, data); err != nil {
		return err
	}
	if err := createServeGo(appPaths, data); err != nil {
		return err
	}
	if addAbout {
		appPaths.CreateAboutFolders()
		if err := createAboutGo(appPaths, data); err != nil {
			return err
		}
	}
	if err := createCallServer(appPaths, data); err != nil {
		return err
	}
	return nil
}
