package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

// createSpawnTabFiles creates the spawn tab specific files.
// The files are prepare.go and api.go,
func createSpawnTabFiles(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	// tabBarTabPanelGroups:
	//   tabBarName : tabName : *project.TabPanelGroup
	tabBarTabPanelGroups := buildMapTabpanelGroup(builder)
	// map each panel name to its slice of folders.
	panelNamePathMap := buildPanelNamePathMap(builder)

	// Setup vars.
	folderpaths := appPaths.GetPaths()
	folderNames := appPaths.GetFolderNames()
	fileNames := paths.GetFileNames()
	tabPanelPaths := builder.GenerateSpawnTabPanelPathsMap()
	// Build the core imports for each api.go.
	// Let each tab bar add the rest of the imports.
	prepareCoreImports := make([]string, 4)
	prepareCoreImports[0] = folderpaths.ImportRendererLPC
	prepareCoreImports[1] = folderpaths.ImportRendererViewTools
	prepareCoreImports[2] = folderpaths.ImportRendererNotJS
	prepareCoreImports[3] = folderpaths.ImportRendererPaneling

	spawnCoreImports := make([]string, 2)
	spawnCoreImports[0] = folderpaths.ImportRendererLPC
	spawnCoreImports[1] = folderpaths.ImportRendererViewTools
	// Build the tab bar template data.
	// Let each tab bar set the rest of the spawnTabData members.
	data := &spawnTabData{
		ApplicationGitPath: builder.ImportPath,
		CamelCase:          cases.CamelCase,
		LowerCamelCase:     cases.LowerCamelCase,
		LowerCase:          strings.ToLower,
		SplitTabJoin: func(s string) string {
			ss := strings.Split(s, "\n")
			return "\t" + strings.Join(ss, "\n\t")
		},
		PackageNameCase: cases.ToGoPackageName,
		SpawnID:         "{{.SpawnID}}",
	}

	// One tab bar at a time.
	for tabBarName, tabPanelGroups := range tabBarTabPanelGroups {
		// Step 1: For this tab bar
		//         Build tabBarFolderPath ( the tab bar folder path ).
		//         It will be used to make each tab path.
		var tabBarFolderPath string
		var tabBarID string
		for _, panelGroup := range tabPanelGroups {
			firstTabPanelGroup := panelGroup
			tabBarID = firstTabPanelGroup.TabBarID
			for panelName := range firstTabPanelGroup.PanelNamesIDMap {
				ff := panelNamePathMap[panelName]
				for pos, f := range ff {
					if f == tabBarName {
						pos++
						folders := strings.Join(ff[:pos], string(os.PathSeparator))
						tabBarFolderPath = filepath.Join(folderpaths.OutputRendererSpawns, folders)
						break
					}
				}
				break
			}
			break
		}

		// One tab at a time.
		for tabName, panelGroup := range tabPanelGroups {
			// Step 1: For this tab
			//         Build a slice of each of this tab's panel names.
			var panelNames []string
			l := len(panelGroup.PanelNamesIDMap)
			panelNames = make([]string, 0, l)
			for panelName := range panelGroup.PanelNamesIDMap {
				panelNames = append(panelNames, panelName)
			}

			// Step 2: For this tab
			//         Build the prepare imports.
			paths := tabPanelPaths[tabName]
			lp := len(paths)
			markupTemplatePaths := make([]string, l)
			lpci := len(prepareCoreImports)
			prepareImports := make([]string, lp+lpci)
			copy(prepareImports, prepareCoreImports)
			lsci := len(spawnCoreImports)
			spawnImports := make([]string, lp+lsci)
			copy(spawnImports, prepareCoreImports)
			for i, path := range paths {
				markupTemplatePaths[i] = filepath.Join(folderNames.SpawnTemplates, path) + ".tmpl"
				ipath := folderpaths.ImportRendererSpawnPanels + "/" + path
				prepareImports[lpci+i] = ipath
				spawnImports[lsci+i] = ipath
			}
			sort.Strings(prepareImports)
			sort.Strings(spawnImports)

			// Step 3: For this tab
			//         Customize the template data.
			data.TabName = tabName
			data.TabBarID = tabBarID
			data.PanelNames = panelNames
			data.PrepareImports = prepareImports
			data.SpawnImports = spawnImports
			data.MarkupTemplatePaths = fmt.Sprintf("%#v", markupTemplatePaths)
			data.TabLabel = panelGroup.TabLabel
			data.PanelHeading = panelGroup.PanelHeading

			// Step 4: For this tab
			//         Make the tab folder and process the tab template.
			var fname string
			var oPath string
			var tabPath string
			tabPath = filepath.Join(tabBarFolderPath, tabName)
			if err = os.MkdirAll(tabPath, appPaths.GetDMode()); err != nil {
				return
			}
			fname = fileNames.PrepareDotGo
			oPath = filepath.Join(tabPath, fname)
			if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabPrepare, data, appPaths); err != nil {
				return
			}
			fname = fileNames.SpawnDotGo
			oPath = filepath.Join(tabPath, fname)
			if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabSpawn, data, appPaths); err != nil {
				return
			}
		}
	}
	return
}

type spawnTabData struct {
	TabName    string
	TabBarID   string
	PanelNames []string

	MarkupTemplatePaths string

	TabLabel     string
	PanelHeading string

	PrepareImports []string
	SpawnImports   []string

	ApplicationGitPath string

	CamelCase       func(string) string
	LowerCamelCase  func(string) string
	LowerCase       func(string) string
	SplitTabJoin    func(string) string
	PackageNameCase func(string) string

	SpawnID string
}

func buildPanelNamePathMap(builder *project.Builder) (panelNamePathMap map[string][]string) {
	panelNamePathMap = make(map[string][]string, 100)
	serviceNames := builder.GenerateServiceNames()
	servicePanelNamePathMap := builder.GenerateServiceSpawnTabEmptyInsidePanelNamePathMap()
	for _, serviceName := range serviceNames {
		pnpMap := servicePanelNamePathMap[serviceName]
		for panelName, panelPath := range pnpMap {
			panelNamePathMap[panelName] = panelPath
		}
	}
	return
}
