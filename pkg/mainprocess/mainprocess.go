package mainprocess

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
)

type templateData struct {
	ApplicationName                 string
	ApplicationGitPath              string
	Stores                          []string
	LowerCamelCase                  func(string) string
	CamelCase                       func(string) string
	HomeTemplatePanelNames          string
	HomeEmptyInsidePanelNamePathMap string

	ImportDomainDataFilepaths string
	ImportDomainDataLogLevels string
	ImportDomainDataSettings  string

	ImportDomainStore        string
	ImportDomainStoreStoring string
	ImportDomainStoreStorer  string
	ImportDomainStoreRecord  string

	ImportRendererSpawnPanels string

	ImportDomainLPC              string
	ImportDomainLPCMessage       string
	ImportRendererLPC            string
	ImportMainProcessLPC         string
	ImportMainProcessLPCDispatch string

	FileNames *paths.FileNames

	SitePackImportPath string
	SitePackPackage    string
}

// Create creates main process folder files from templates.
func Create(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	parts := strings.Split(builder.ImportPath, "/")
	appname := parts[len(parts)-1]
	data := &templateData{
		ApplicationName:                 appname,
		ApplicationGitPath:              builder.ImportPath,
		LowerCamelCase:                  cases.LowerCamelCase,
		CamelCase:                       cases.CamelCase,
		HomeEmptyInsidePanelNamePathMap: strings.ReplaceAll(fmt.Sprintf("%#v", builder.GenerateHomeEmptyInsidePanelNamePathMap()), ":", ": "),
		HomeTemplatePanelNames:          fmt.Sprintf("%#v", builder.GenerateHomeTemplatePanelName()),

		ImportDomainDataFilepaths: folderpaths.ImportDomainDataFilepaths,
		ImportDomainDataLogLevels: folderpaths.ImportDomainDataLogLevels,
		ImportDomainDataSettings:  folderpaths.ImportDomainDataSettings,

		ImportDomainStore:        folderpaths.ImportDomainStore,
		ImportDomainStoreStoring: folderpaths.ImportDomainStoreStoring,
		ImportDomainStoreStorer:  folderpaths.ImportDomainStoreStorer,
		ImportDomainStoreRecord:  folderpaths.ImportDomainStoreRecord,

		ImportRendererSpawnPanels: folderpaths.ImportRendererSpawnPanels,

		ImportDomainLPC:              folderpaths.ImportDomainLPC,
		ImportDomainLPCMessage:       folderpaths.ImportDomainLPCMessage,
		ImportRendererLPC:            folderpaths.ImportRendererLPC,
		ImportMainProcessLPC:         folderpaths.ImportMainProcessLPC,
		ImportMainProcessLPCDispatch: folderpaths.ImportMainProcessLPCDispatch,

		FileNames: paths.GetFileNames(),

		SitePackImportPath: builder.SitePackImportPath,
		SitePackPackage:    builder.SitePackPackage,
	}
	if err = createMain(appPaths, data); err != nil {
		return
	}
	if err = createLPC(appPaths, data); err != nil {
		return
	}
	if err = createDispatch(appPaths, data); err != nil {
		return
	}
	if err = createKickstore(appPaths, data); err != nil {
		return
	}
	if err = createVSCode(appPaths, appname); err != nil {
		return
	}
	return
}
