package renderer

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/format"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

type namePath struct {
	Name, Path string
}

// createFrameworkDoPanelsGo writes panels.go in package main.
func createFrameworkDoPanelsGo(appPaths paths.ApplicationPathsI, builder *project.Builder) error {
	folderpaths := appPaths.GetPaths()
	homePanelNamePathMap := builder.GenerateHomeEmptyInsidePanelNamePathMap()
	imports := make([]string, 3, 100)
	panelNamePath := make(map[string]string)
	imports[0] = builder.ImportPath + folderpaths.ImportRendererPaneling
	imports[1] = builder.ImportPath + folderpaths.ImportRendererFrameworkLPC
	imports[2] = builder.ImportPath + folderpaths.ImportRendererFrameworkViewTools
	for _, panelNameFolders := range homePanelNamePathMap {
		for name, folders := range panelNameFolders {
			path := filepath.Join(filepath.Join(folders...), name)
			importPath := filepath.Join(builder.ImportPath+folderpaths.ImportRendererPanels, path)
			imports = append(imports, importPath)
			panelNamePath[name] = path
		}
	}
	tabBarNamePath := make(map[string]string)
	spawnTabBarPanelPaths := buildSpawnTabBarPackagePanelPath(builder)
	for name, path := range spawnTabBarPanelPaths {
		importPath := filepath.Join(builder.ImportPath+folderpaths.ImportRendererSpawnPanels, path)
		imports = append(imports, importPath)
		tabBarNamePath[name] = path
	}
	sort.Strings(imports)
	data := &struct {
		Imports             []string
		ApplicationGitPath  string
		PanelNamePath       map[string]string
		SpawnTabBarNamePath map[string]string
		LowerCamelCase      func(string) string
		PackageNameCase     func(string) string
	}{
		Imports:             format.FixImports(imports),
		ApplicationGitPath:  builder.ImportPath,
		PanelNamePath:       panelNamePath,
		SpawnTabBarNamePath: tabBarNamePath,
		LowerCamelCase:      cases.LowerCamelCase,
		PackageNameCase:     cases.ToGoPackageName,
	}
	fileNames := paths.GetFileNames()
	fname := fileNames.PanelsDotGo
	oPath := filepath.Join(folderpaths.OutputRendererFramework, fname)
	return templates.ProcessTemplate(fname, oPath, templates.FrameworkDoPanelsGo, data, appPaths)
}

func buildSpawnTabBarPackagePanelPath(builder *project.Builder) (spawnTabBarPanelPaths map[string]string) {
	spawnTabBarPanelPaths = make(map[string]string, 100)
	tabBarPanelPaths := builder.GenerateSpawnTabBarPanelPathsMap()
	for name, folders := range tabBarPanelPaths {
		spawnTabBarPanelPaths[name] = strings.Join(folders, "/")
	}
	return
}
