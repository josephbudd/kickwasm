package mainprocess

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
)

type templateData struct {
	ApplicationName                    string
	ApplicationGitPath                 string
	Stores                             []string
	ServiceNames                       []string
	LowerCamelCase                     func(string) string
	CamelCase                          func(string) string
	ServiceTemplatePanelNames          string
	ServiceEmptyInsidePanelNamePathMap string

	ImportDomainInterfacesStorers          string
	ImportDomainInterfacesCallers          string
	ImportDomainDataFilepaths              string
	ImportDomainDataCallIDs                string
	ImportDomainDataLogLevels              string
	ImportDomainDataSettings               string
	ImportDomainTypes                      string
	ImportDomainImplementationsCalling     string
	ImportDomainImplementationsStoringBolt string

	ImportMainProcessCalls      string
	ImportMainProcessCallServer string

	FileNames *paths.FileNames
}

// Create creates main process folder files from templates.
func Create(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	parts := strings.Split(builder.ImportPath, "/")
	appname := parts[len(parts)-1]
	data := &templateData{
		ApplicationName:                    appname,
		ApplicationGitPath:                 builder.ImportPath,
		Stores:                             builder.Stores,
		ServiceNames:                       builder.GenerateServiceNames(),
		LowerCamelCase:                     cases.LowerCamelCase,
		CamelCase:                          cases.CamelCase,
		ServiceEmptyInsidePanelNamePathMap: strings.Replace(fmt.Sprintf("%#v", builder.GenerateServiceEmptyInsidePanelNamePathMap()), ":", ": ", -1),
		ServiceTemplatePanelNames:          fmt.Sprintf("%#v", builder.GenerateServiceTemplatePanelName()),

		ImportDomainInterfacesStorers:          folderpaths.ImportDomainInterfacesStorers,
		ImportDomainInterfacesCallers:          folderpaths.ImportDomainInterfacesCallers,
		ImportDomainDataFilepaths:              folderpaths.ImportDomainDataFilepaths,
		ImportDomainDataCallIDs:                folderpaths.ImportDomainDataCallIDs,
		ImportDomainDataLogLevels:              folderpaths.ImportDomainDataLogLevels,
		ImportDomainDataSettings:               folderpaths.ImportDomainDataSettings,
		ImportDomainTypes:                      folderpaths.ImportDomainTypes,
		ImportDomainImplementationsCalling:     folderpaths.ImportDomainImplementationsCalling,
		ImportDomainImplementationsStoringBolt: folderpaths.ImportDomainImplementationsStoringBolt,
		ImportMainProcessCalls:                 folderpaths.ImportMainProcessCalls,
		ImportMainProcessCallServer:            folderpaths.ImportMainProcessCallServer,

		FileNames: paths.GetFileNames(),
	}
	if err = createMainGo(appPaths, data); err != nil {
		return
	}
	if err = createPanelMapGo(appPaths, data); err != nil {
		return
	}
	if err = createServeGo(appPaths, data); err != nil {
		return
	}
	if err = createCallServer(appPaths, data); err != nil {
		return
	}
	if err = createCalls(appPaths, data); err != nil {
		return
	}
	return
}
