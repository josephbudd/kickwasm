package renderer

import (
	"path/filepath"
	"sort"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

type namePath struct {
	Name, Path string
}

// createMainDoPanelsGo writes panels.go in package main.
func createMainDoPanelsGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	servicePanelNamePathMap := builder.GenerateServiceEmptyInsidePanelNamePathMap()
	imports := make([]string, 5, 20)
	panelNamePath := make(map[string]string)
	imports[0] = builder.ImportPath + folderpaths.ImportDomainTypes
	imports[1] = builder.ImportPath + folderpaths.ImportRendererInterfacesPanelHelper
	imports[2] = builder.ImportPath + folderpaths.ImportRendererViewTools
	imports[3] = builder.ImportPath + folderpaths.ImportRendererNotJS
	imports[4] = builder.ImportPath + folderpaths.ImportDomainInterfacesCallers
	for _, panelNameFolders := range servicePanelNamePathMap {
		for name, folders := range panelNameFolders {
			path := filepath.Join(filepath.Join(folders...), name)
			importPath := filepath.Join(builder.ImportPath+folderpaths.ImportRendererPanels, path)
			imports = append(imports, importPath)
			panelNamePath[name] = path
		}
	}
	sort.Strings(imports)
	data := &struct {
		Imports                 []string
		ApplicationGitPath      string
		PanelNamePath           map[string]string
		LowerCamelCase          func(string) string
		PackageNameCase         func(string) string
		ImportRendererViewTools string
		ImportRendererPanels    string
	}{
		Imports:                 imports,
		ApplicationGitPath:      builder.ImportPath,
		PanelNamePath:           panelNamePath,
		LowerCamelCase:          cases.LowerCamelCase,
		PackageNameCase:         cases.ToGoPackageName,
		ImportRendererViewTools: folderpaths.ImportRendererViewTools,
		ImportRendererPanels:    folderpaths.ImportRendererPanels,
	}
	fileNames := paths.GetFileNames()
	fname := fileNames.PanelsDotGo
	oPath := filepath.Join(folderpaths.OutputRenderer, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainDoPanelsGo, data, appPaths)
}
