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
	imports := make([]string, 3, 20)
	imports[0] = builder.ImportPath + folderpaths.ImportMainProcessTransportsCalls
	imports[1] = builder.ImportPath + folderpaths.ImportRendererWASMViewTools
	imports[2] = "github.com/josephbudd/kicknotjs"
	if addAbout {
		importPath := filepath.Join(builder.ImportPath+folderpaths.ImportRendererWASMPanels, "AboutButton", "AboutPanel")
		imports = append(imports, importPath)
	}
	for _, panelNameFolders := range servicePanelNamePathMap {
		for name, folders := range panelNameFolders {
			path := filepath.Join(filepath.Join(folders...), name)
			namePathMap[name] = path
			importPath := filepath.Join(builder.ImportPath+folderpaths.ImportRendererWASMPanels, path)
			imports = append(imports, importPath)
		}
	}
	sort.Strings(imports)
	data := &struct {
		Imports                          []string
		AddAbout                         bool
		ApplicationGitPath               string
		PanelNamePathMap                 map[string]string
		LowerCamelCase                   func(string) string
		ImportMainProcessTransportsCalls string
		ImportRendererWASMViewTools      string
		ImportRendererWASMPanels         string
	}{
		Imports:                          imports,
		AddAbout:                         addAbout,
		ApplicationGitPath:               builder.ImportPath,
		PanelNamePathMap:                 namePathMap,
		LowerCamelCase:                   cases.LowerCamelCase,
		ImportMainProcessTransportsCalls: folderpaths.ImportMainProcessTransportsCalls,
		ImportRendererWASMViewTools:      folderpaths.ImportRendererWASMViewTools,
		ImportRendererWASMPanels:         folderpaths.ImportRendererWASMPanels,
	}
	fname := "panels.go"
	oPath := filepath.Join(folderpaths.OutputRendererWASM, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainDoPanelsGo, data, appPaths)
}
