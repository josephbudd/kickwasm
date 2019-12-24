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

// createGoPanels creates the rendererprocess/panels/ go panel files.
// Only for organic not spawned panels.
func createGoPanels(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	homePanelNamePathMap := builder.GenerateHomeEmptyInsidePanelNamePathMap()
	homeButtonPanelGroups := builder.GenerateHomeButtonPanelGroups()
	homeNames := builder.GenerateHomeButtonNames()
	for _, homeName := range homeNames {
		panelNamePathMap := homePanelNamePathMap[homeName]
		homeButtonPanelGroup := homeButtonPanelGroups[homeName]
		for _, buttonPanelGroups := range homeButtonPanelGroup {
			// make this panel's group
			panelGroup := make([]*project.Panel, 0, 5)
			for _, panel := range buttonPanelGroups.PanelNamesIDMap {
				panelGroup = append(panelGroup, panel)
			}
			// template data for each panel file in this group.
			for panelName, panel := range buttonPanelGroups.PanelNamesIDMap {
				if len(panel.Buttons) > 0 || len(panel.Tabs) > 0 {
					continue
				}
				folders := strings.Join(panelNamePathMap[panelName], string(os.PathSeparator))
				folderpath := filepath.Join(folderpaths.OutputRendererPanels, folders, panelName)
				if err = os.MkdirAll(folderpath, appPaths.GetDMode()); err != nil {
					return
				}
				var tabButtonID string
				if buttonPanelGroups.IsTabButton {
					tabButtonID = buttonPanelGroups.ButtonID
				}
				data := &struct {
					PanelName                 string
					PanelID                   string
					PanelH3ID                 string
					PanelGroup                []*project.Panel
					IsTabSiblingPanel         bool
					TabButtonID               string
					ApplicationGitPath        string
					ImportRendererDOM         string
					ImportRendererDisplay     string
					ImportRendererEvent       string
					ImportRendererLPC         string
					ImportRendererMarkup      string
					ImportRendererViewTools   string
					ImportRendererPaneling    string
					ImportDomainDataLogLevels string
					ImportDomainStoreRecord   string
					ImportDomainLPCMessage    string

					CamelCase       func(string) string
					LowerCamelCase  func(string) string
					SplitTabJoin    func(string) string
					PackageNameCase func(string) string

					StartBracket string
					EndBracket   string
				}{
					PanelName:                 panelName,
					PanelID:                   panel.HTMLID,
					PanelH3ID:                 panel.H3ID,
					PanelGroup:                panelGroup,
					IsTabSiblingPanel:         buttonPanelGroups.IsTabButton,
					TabButtonID:               tabButtonID,
					ApplicationGitPath:        builder.ImportPath,
					ImportRendererDisplay:     folderpaths.ImportRendererDisplay,
					ImportRendererDOM:         folderpaths.ImportRendererDOM,
					ImportRendererEvent:       folderpaths.ImportRendererEvent,
					ImportRendererLPC:         folderpaths.ImportRendererLPC,
					ImportRendererMarkup:      folderpaths.ImportRendererMarkup,
					ImportRendererViewTools:   folderpaths.ImportRendererViewTools,
					ImportRendererPaneling:    folderpaths.ImportRendererPaneling,
					ImportDomainDataLogLevels: folderpaths.ImportDomainDataLogLevels,
					ImportDomainStoreRecord:   folderpaths.ImportDomainStoreRecord,
					ImportDomainLPCMessage:    folderpaths.ImportDomainLPCMessage,

					CamelCase:      cases.CamelCase,
					LowerCamelCase: cases.LowerCamelCase,
					SplitTabJoin: func(s string) string {
						ss := strings.Split(s, "\n")
						return "\t" + strings.Join(ss, "\n\t")
					},
					PackageNameCase: cases.ToGoPackageName,

					StartBracket: "{",
					EndBracket:   "}",
				}
				var fname string
				var oPath string
				fname = fileNames.DataDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelData, data, appPaths); err != nil {
					return
				}
				fname = fileNames.PanelGroupDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelGroup, data, appPaths); err != nil {
					return
				}
				fname = fileNames.PanelDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.Panel, data, appPaths); err != nil {
					return
				}
				fname = fileNames.ControllerDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelController, data, appPaths); err != nil {
					return
				}
				fname = fileNames.PresenterDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelPresenter, data, appPaths); err != nil {
					return
				}
				fname = fileNames.MessengerDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelMessenger, data, appPaths); err != nil {
					return
				}
			}
		}
	}
	return
}
