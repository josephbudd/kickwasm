package renderer

import (
	"path/filepath"
	"sort"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

// createMainDoPanelsGo writes panels.go in package main.
func createMainDoPanelsGo(addAbout bool, appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	folderpaths := appPaths.GetPaths()
	servicePanelNamePathMap := builder.GenerateServiceEmptyInsidePanelNamePathMap()
	namePathMap := make(map[string]string)
	imports := make([]string, 4, 20)
	imports[0] = builder.ImportPath + folderpaths.ImportDomainTypes
	imports[1] = builder.ImportPath + folderpaths.ImportRendererViewTools
	imports[2] = "github.com/josephbudd/kicknotjs"
	imports[3] = builder.ImportPath + folderpaths.ImportDomainInterfacesCallers

	if addAbout {
		importPath := filepath.Join(builder.ImportPath+folderpaths.ImportRendererPanels, "AboutButton", "AboutPanel")
		imports = append(imports, importPath)
	}
	for _, panelNameFolders := range servicePanelNamePathMap {
		for name, folders := range panelNameFolders {
			path := filepath.Join(filepath.Join(folders...), name)
			namePathMap[name] = path
			importPath := filepath.Join(builder.ImportPath+folderpaths.ImportRendererPanels, path)
			imports = append(imports, importPath)
		}
	}
	sort.Strings(imports)
	data := &struct {
		Imports                 []string
		AddAbout                bool
		ApplicationGitPath      string
		PanelNamePathMap        map[string]string
		LowerCamelCase          func(string) string
		ImportRendererViewTools string
		ImportRendererPanels    string
	}{
		Imports:                 imports,
		AddAbout:                addAbout,
		ApplicationGitPath:      builder.ImportPath,
		PanelNamePathMap:        namePathMap,
		LowerCamelCase:          cases.LowerCamelCase,
		ImportRendererViewTools: folderpaths.ImportRendererViewTools,
		ImportRendererPanels:    folderpaths.ImportRendererPanels,
	}
	fname := "panels.go"
	oPath := filepath.Join(folderpaths.OutputRenderer, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainDoPanelsGo, data, appPaths)
}
