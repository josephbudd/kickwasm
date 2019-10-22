package renderer

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

// createSpawnTabBarFiles creates the tab bar specific files.
// The files are api.go,
func createSpawnTabBarFiles(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	spawnTabBarTabPanelPathsMap := builder.GenerateSpawnTabBarTabPanelPathsMap()
	// tabBarTabPanelGroups:
	//   tabBarName : tabName : []*project.TabPanelGroup
	tabBarTabPanelGroups := buildMapTabpanelGroup(builder)
	// map each panel name to its slice of folders.
	spawnTabBarNamePath := buildSpawnTabBarPanelNamePathMap(builder)
	// Setup vars.
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	// Build the core imports for each api.go.
	// Let each tab bar add the rest of the imports.
	coreImports := make([]string, 4)
	coreImports[0] = folderpaths.ImportRendererLPC
	coreImports[1] = folderpaths.ImportRendererPaneling
	coreImports[2] = folderpaths.ImportRendererNotJS
	coreImports[3] = folderpaths.ImportRendererViewTools
	// Build the tab bar template data.
	data := &spawnTabBarData{
		ApplicationGitPath: builder.ImportPath,

		CamelCase:      cases.CamelCase,
		LowerCamelCase: cases.LowerCamelCase,
		SplitTabJoin: func(s string) string {
			ss := strings.Split(s, "\n")
			return "\t" + strings.Join(ss, "\n\t")
		},
		PackageNameCase: cases.ToGoPackageName,

		SpawnID: "{{.SpawnID}}",
	}

	// One tab bar at a time.
	for tabBarName, tabPanelGroups := range tabBarTabPanelGroups {
		// Step 1: For this tab bar
		//         Get the tab bar id.
		var tabBarID string
		for _, panelGroup := range tabPanelGroups {
			tabBarID = panelGroup.TabBarID
			break
		}

		// Step 2: For this tab bar
		//         Build a slice of each of this tab bar's spawn tab names.
		var l int
		var tabNames []string
		l = len(tabPanelGroups)
		tabNames = make([]string, 0, l)
		for tabName := range tabPanelGroups {
			tabNames = append(tabNames, tabName)
		}

		// Step 3.a: For this tab bar
		//           Build the imports.
		// Step 3.b: For this tab bar
		//           Build tabBarFolderPath ( the tab bar folder path ).
		//           Then make the path on the drive.
		var imports []string
		var tabBarFolderPath string
		tabPaths := spawnTabBarTabPanelPathsMap[tabBarName]
		l = len(coreImports)
		imports = make([]string, l+len(tabPaths))
		copy(imports, coreImports)
		for i, tabPath := range tabPaths {
			path := strings.Join(tabPath, "/")
			imports[l+i] = folderpaths.ImportRendererSpawnPanels + "/" + path
		}
		sort.Strings(imports)
		tabBarFolderPath = filepath.Join(folderpaths.OutputRendererSpawns, spawnTabBarNamePath[tabBarName])
		if err = os.MkdirAll(tabBarFolderPath, appPaths.GetDMode()); err != nil {
			return
		}

		// Step 4: For this tab bar
		//         Customize the template data.
		data.TabBarID = tabBarID
		data.TabBarName = tabBarName
		data.TabNames = tabNames
		data.Imports = imports

		// Step 5: For this tab bar
		//         Process the tab bar prepare template.
		var fname string
		var oPath string
		fname = fileNames.PrepareDotGo
		oPath = filepath.Join(tabBarFolderPath, fname)
		if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabBarPrepare, data, appPaths); err != nil {
			return
		}
	}
	return
}

type spawnTabBarData struct {
	TabBarID   string
	TabBarName string
	TabNames   []string

	Imports []string

	ApplicationGitPath string

	CamelCase       func(string) string
	LowerCamelCase  func(string) string
	SplitTabJoin    func(string) string
	PackageNameCase func(string) string

	SpawnID string
}

// tabBarTabPanelGroups map[string]map[string][]*project.TabPanelGroup
// map tab bar name : tab name : *project.TabPanelGroup
func buildMapTabpanelGroup(builder *project.Builder) (tabBarTabPanelGroups map[string]map[string]*project.TabPanelGroup) {
	homeNames := builder.GenerateHomeButtonNames()
	homeTabPanelGroups := builder.GenerateHomeSpawnTabButtonPanelGroups()
	l := len(homeNames)
	tabBarTabPanelGroups = make(map[string]map[string]*project.TabPanelGroup, l)
	for _, homeName := range homeNames {
		// tabPanelGroups will be used to construct a map
		//  of tab bars mapped to their panel groups.
		groups := homeTabPanelGroups[homeName]
		// groups is []*project.TabPanelGroup
		for _, group := range groups {
			var found bool
			// map each tab bar to it's tabs.
			var tabBarTabs map[string]*project.TabPanelGroup
			if tabBarTabs, found = tabBarTabPanelGroups[group.TabBarName]; !found {
				tabBarTabs = make(map[string]*project.TabPanelGroup, 10)
				tabBarTabPanelGroups[group.TabBarName] = tabBarTabs
			}
			// map each tab to its panel group.
			tabBarTabs[group.TabName] = group
		}
	}
	return
}

func buildSpawnTabBarPanelNamePathMap(builder *project.Builder) (spawnTabBarPanelNamePathMap map[string]string) {
	namePaths := builder.GenerateSpawnTabBarPanelPathsMap()
	spawnTabBarPanelNamePathMap = make(map[string]string, len(namePaths))
	for name, folders := range namePaths {
		path := strings.Join(folders, "/")
		spawnTabBarPanelNamePathMap[name] = path
	}
	return
}
