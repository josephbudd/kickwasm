package renderer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

// createGoPanels creates the renderer/panels/ go panel files.
// Only for organic not autogenerated panels.
func createGoPanels(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	servicePanelNamePathMap := builder.GenerateServiceEmptyInsidePanelNamePathMap()
	serviceButtonPanelGroups := builder.GenerateServiceButtonPanelGroups()
	serviceNames := builder.GenerateServiceNames()
	for _, serviceName := range serviceNames {
		panelNamePathMap := servicePanelNamePathMap[serviceName]
		serviceButtonPanelGroup := serviceButtonPanelGroups[serviceName]
		for _, buttonPanelGroups := range serviceButtonPanelGroup {
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
				data := &struct {
					PanelName                           string
					PanelID                             string
					PanelGroup                          []*project.Panel
					IsTabSiblingPanel                   bool
					ApplicationGitPath                  string
					ImportRendererNotJS                 string
					ImportRendererViewTools             string
					ImportRendererInterfacesPanelHelper string
					ImportDomainDataCallIDs             string
					ImportDomainDataLogLevels           string
					ImportDomainTypes                   string
					ImportDomainImplementationsCalling  string
					ImportDomainInterfacesCallers       string

					CamelCase       func(string) string
					LowerCamelCase  func(string) string
					SplitTabJoin    func(string) string
					PackageNameCase func(string) string
				}{
					PanelName:                           panelName,
					PanelID:                             panel.HTMLID,
					PanelGroup:                          panelGroup,
					IsTabSiblingPanel:                   buttonPanelGroups.IsTabButton,
					ApplicationGitPath:                  builder.ImportPath,
					ImportRendererNotJS:                 folderpaths.ImportRendererNotJS,
					ImportRendererViewTools:             folderpaths.ImportRendererViewTools,
					ImportRendererInterfacesPanelHelper: folderpaths.ImportRendererInterfacesPanelHelper,
					ImportDomainDataCallIDs:             folderpaths.ImportDomainDataCallIDs,
					ImportDomainDataLogLevels:           folderpaths.ImportDomainDataLogLevels,
					ImportDomainTypes:                   folderpaths.ImportDomainTypes,
					ImportDomainImplementationsCalling:  folderpaths.ImportDomainImplementationsCalling,
					ImportDomainInterfacesCallers:       folderpaths.ImportDomainInterfacesCallers,

					CamelCase:      cases.CamelCase,
					LowerCamelCase: cases.LowerCamelCase,
					SplitTabJoin: func(s string) string {
						ss := strings.Split(s, "\n")
						return "\t" + strings.Join(ss, "\n\t")
					},
					PackageNameCase: cases.ToGoPackageName,
				}
				fname := fileNames.PanelGroupDotGo
				oPath := filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelGroup, data, appPaths); err != nil {
					return
				}
				fname = fileNames.PanelDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.Panel, data, appPaths); err != nil {
					return
				}
				fname = fileNames.ControlerDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelControler, data, appPaths); err != nil {
					return
				}
				fname = fileNames.PresenterDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelPresenter, data, appPaths); err != nil {
					return
				}
				fname = fileNames.CallerDotGo
				oPath = filepath.Join(folderpath, fname)
				if err = templates.ProcessTemplate(fname, oPath, templates.PanelCaller, data, appPaths); err != nil {
					return
				}
			}
		}
	}
	return
}
