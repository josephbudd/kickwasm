package mainprocess

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/tap"
)

const (
	portFolderName = "ports"
)

type templateData struct {
	Port                                  uint
	Host                                  string
	ApplicationName                       string
	ApplicationGitPath                    string
	Repos                                 []string
	ServiceNames                          []string
	LowerCamelCase                        func(string) string
	CamelCase                             func(string) string
	AddAbout                              bool
	ServiceTemplatePanelNames             string
	ServiceEmptyInsidePanelNamePathMap    string
	HeadTemplateFile                      string
	ImportMainProcessBehaviorRepoi        string
	ImportMainProcessDataFilePaths        string
	ImportMainProcessDataRecords          string
	ImportMainProcessRepositoriesBolt     string
	ImportMainProcessServices             string
	ImportMainProcessServicesAbout        string
	ImportMainProcessTransportsCalls      string
	ImportMainProcessTransportsCallServer string

	ImportRendererWASMCall      string
	ImportRendererWASMViewTools string
	ImportRendererWASMPanels    string
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

		ImportMainProcessBehaviorRepoi:        folderpaths.ImportMainProcessBehaviorRepoi,
		ImportMainProcessDataFilePaths:        folderpaths.ImportMainProcessDataFilePaths,
		ImportMainProcessDataRecords:          folderpaths.ImportMainProcessDataRecords,
		ImportMainProcessRepositoriesBolt:     folderpaths.ImportMainProcessRepositoriesBolt,
		ImportMainProcessServices:             folderpaths.ImportMainProcessServices,
		ImportMainProcessServicesAbout:        folderpaths.ImportMainProcessServicesAbout,
		ImportMainProcessTransportsCalls:      folderpaths.ImportMainProcessTransportsCalls,
		ImportMainProcessTransportsCallServer: folderpaths.ImportMainProcessTransportsCallServer,

		ImportRendererWASMCall:      folderpaths.ImportRendererWASMCall,
		ImportRendererWASMViewTools: folderpaths.ImportRendererWASMViewTools,
		ImportRendererWASMPanels:    folderpaths.ImportRendererWASMPanels,
	}
	if err := createMainGo(appPaths, data); err != nil {
		return err
	}
	if err := createPanelMapGo(appPaths, data); err != nil {
		return err
	}
	if addAbout {
		appPaths.CreateAboutFolders()
		if err := createAboutGo(appPaths, data); err != nil {
			return err
		}
	}
	if err := createFilePathsFilePathsGo(appPaths, data); err != nil {
		return err
	}
	if err := createRecordsRecordsGo(appPaths, data); err != nil {
		return err
	}
	if err := createRepoIGo(appPaths, data); err != nil {
		return err
	}
	if err := createBoltDatabaseGo(appPaths, data); err != nil {
		return err
	}
	if err := createCallsGo(appPaths, data); err != nil {
		return err
	}
	if err := createCallServer(appPaths, data); err != nil {
		return err
	}
	return nil
}
