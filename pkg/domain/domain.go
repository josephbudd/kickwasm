package domain

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/pkg"

	"github.com/josephbudd/kickwasm/pkg/cases"
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

	ImportDomainDataFilepaths string

	ImportDomainStore        string
	ImportDomainStoreRecord  string
	ImportDomainStoreStorer  string
	ImportDomainStoreStoring string

	OutputRendererPanels string

	FileNames   *paths.FileNames
	FolderNames *paths.FolderNames

	Host string
	Port uint

	SitePackPackage    string
	SitePackImportPath string
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

		// domain

		ImportDomainDataFilepaths: imports.ImportDomainDataFilepaths,

		ImportDomainStore:        imports.ImportDomainStore,
		ImportDomainStoreRecord:  imports.ImportDomainStoreRecord,
		ImportDomainStoreStorer:  imports.ImportDomainStoreStorer,
		ImportDomainStoreStoring: imports.ImportDomainStoreStoring,

		// output renderer

		OutputRendererPanels: folderpaths.OutputRendererPanels,

		FileNames:   paths.GetFileNames(),
		FolderNames: paths.GetFolderNames(),

		// settings.yaml

		Host: pkg.LocalHost,
		Port: pkg.LocalPort,

		// sitepack

		SitePackPackage:    builder.SitePackPackage,
		SitePackImportPath: builder.SitePackImportPath,
	}
	if err = createDataFilePathsGo(appPaths, data); err != nil {
		return
	}
	if err = createStoreStoring(appPaths, data); err != nil {
		return
	}
	if err = createStoreStorer(appPaths, data); err != nil {
		return
	}
	if err = createStoreStoresGo(appPaths, data); err != nil {
		return
	}
	// if err = RebuildStoreInstructions(appPaths, make([]string, 0)); err != nil {
	if err = RebuildStoreInstructions(appPaths, nil, nil, nil); err != nil {
		return
	}
	if err = createStoreRecord(appPaths, data); err != nil {
		return
	}
	if err = createSettingsYAML(appPaths, data); err != nil {
		return
	}
	if err = createLPC(appPaths, data); err != nil {
		return
	}
	return
}
