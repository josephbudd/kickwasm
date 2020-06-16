package renderer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

// createTabSpawnPanels creates the rendererprocess/spawnPanels/ go panel files.
// Only for spawned panels.
func createTabSpawnPanels(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	homePanelNamePathMap := builder.GenerateHomeSpawnTabEmptyInsidePanelNamePathMap()
	homeTabPanelGroups := builder.GenerateHomeSpawnTabButtonPanelGroups()
	homeNames := builder.GenerateHomeButtonNames()
	for _, homeName := range homeNames {
		panelNamePathMap := homePanelNamePathMap[homeName]
		tabPanelGroups := homeTabPanelGroups[homeName]
		for _, tabPanelGroup := range tabPanelGroups {
			// make this tab's group
			panelGroup := make([]*project.Panel, 0, 5)
			for _, panel := range tabPanelGroup.PanelNamesIDMap {
				panelGroup = append(panelGroup, panel)
			}
			// template data for each panel file in this group.
			for panelName, panel := range tabPanelGroup.PanelNamesIDMap {
				folders := strings.Join(panelNamePathMap[panelName], string(os.PathSeparator))
				folderpath := filepath.Join(folderpaths.OutputRendererSpawns, folders)
				if err = os.MkdirAll(folderpath, appPaths.GetDMode()); err != nil {
					return
				}
				data := &struct {
					PanelName string

					TabBarID     string
					TabName      string
					TabLabel     string
					PanelHeading string
					PanelGroup   []*project.Panel

					ApplicationGitPath string

					ImportRenderer                   string
					ImportRendererAPIDisplay         string
					ImportRendererAPIDOM             string
					ImportRendererAPIEvent           string
					ImportRendererFrameworkLPC       string
					ImportRendererAPIMarkup          string
					ImportRendererPaneling           string
					ImportRendererFrameworkViewTools string
					ImportDomainDataLogLevels        string
					ImportDomainLPC                  string
					ImportDomainLPCMessage           string
					ImportDomainStoreRecord          string

					CamelCase       func(string) string
					LowerCamelCase  func(string) string
					SplitTabJoin    func(string) string
					PackageNameCase func(string) string

					SpawnID string
				}{
					PanelName:    panel.ID,
					TabName:      tabPanelGroup.TabName,
					TabBarID:     tabPanelGroup.TabBarID,
					TabLabel:     tabPanelGroup.TabLabel,
					PanelHeading: tabPanelGroup.PanelHeading,
					PanelGroup:   panelGroup,

					ApplicationGitPath: builder.ImportPath,

					ImportRenderer:                   folderpaths.ImportRenderer,
					ImportRendererAPIDisplay:         folderpaths.ImportRendererAPIDisplay,
					ImportRendererAPIDOM:             folderpaths.ImportRendererAPIDOM,
					ImportRendererAPIEvent:           folderpaths.ImportRendererAPIEvent,
					ImportRendererFrameworkLPC:       folderpaths.ImportRendererFrameworkLPC,
					ImportRendererAPIMarkup:          folderpaths.ImportRendererAPIMarkup,
					ImportRendererPaneling:           folderpaths.ImportRendererPaneling,
					ImportRendererFrameworkViewTools: folderpaths.ImportRendererFrameworkViewTools,
					ImportDomainDataLogLevels:        folderpaths.ImportDomainDataLogLevels,
					ImportDomainLPC:                  folderpaths.ImportDomainLPC,
					ImportDomainLPCMessage:           folderpaths.ImportDomainLPCMessage,
					ImportDomainStoreRecord:          folderpaths.ImportDomainStoreRecord,

					CamelCase:      cases.CamelCase,
					LowerCamelCase: cases.LowerCamelCase,
					SplitTabJoin: func(s string) string {
						ss := strings.Split(s, "\n")
						return "\t" + strings.Join(ss, "\n\t")
					},
					PackageNameCase: cases.ToGoPackageName,

					SpawnID: "{{.SpawnID}}",
				}
				var fname string
				var oPath string
				fname = fileNames.APIDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabPanelAPI, data, appPaths); err != nil {
					return
				}
				fname = fileNames.MessengerDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabPanelMessenger, data, appPaths); err != nil {
					return
				}
				fname = fileNames.ControllerDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabPanelController, data, appPaths); err != nil {
					return
				}
				fname = fileNames.DataDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabPanelData, data, appPaths); err != nil {
					return
				}
				fname = fileNames.PanelGroupDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabPanelGroup, data, appPaths); err != nil {
					return
				}
				fname = fileNames.PanelDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.SpawnTabPanelPanel, data, appPaths); err != nil {
					return
				}
				fname = fileNames.PresenterDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.SpawnPanelPresenter, data, appPaths); err != nil {
					return
				}
			}
		}
	}
	return
}
