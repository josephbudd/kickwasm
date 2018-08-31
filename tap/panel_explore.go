package tap

import (
	"strings"
)

// GenerateOpeningTabPanelID returns the id of the innermost open tab panel at startup.
func (builder *Builder) GenerateOpeningTabPanelID() string {
	// only panels have tabs and buttons
	if len(builder.panel.Tabs) > 0 {
		tab := builder.panel.Tabs[0]
		if len(tab.Panels) == 0 {
			return tab.HTMLID + suffixPanel
		}
		return tab.Panels[0].generateOpeningTabPanelID()
	}
	return emptyString
}
func (panel *Panel) generateOpeningTabPanelID() string {
	if len(panel.Tabs) > 0 {
		tab := panel.Tabs[0]
		if len(tab.Panels) == 0 {
			return tab.HTMLID + suffixPanel
		}
		return tab.Panels[0].generateOpeningTabPanelID()
	}
	return emptyString
}

// GenerateServiceNames returns the service names from the panel.
// The names come from the first buttons of the panel.
func (builder *Builder) GenerateServiceNames() []string {
	names := make([]string, 0, 5)
	for _, s := range builder.Services {
		names = append(names, s.Name)
	}
	return names
}

// GenerateOrganicServiceNames returns a list of service names
//  none of which are flag generated.
// The names come from the first buttons of the panel.
func (builder *Builder) GenerateOrganicServiceNames() []string {
	names := make([]string, 0, 5)
	for _, s := range builder.Services {
		if !s.Generated {
			names = append(names, s.Name)
		}
	}
	return names
}

// GenerateServiceEmptyPanelIDsMap returns
//  each service name mapped to
//  a slice of the ids of the parent divs of the empty divs.
//  Each parent div
//   * is the parent of the panels in a panel group.
//   * is a child of the slider div.
func (builder *Builder) GenerateServiceEmptyPanelIDsMap() map[string][]string {
	servicePanelIDsMap := make(map[string][]string)
	for _, service := range builder.Services {
		serviceName := service.Name
		servicePanelIDsMap[serviceName] = make([]string, 0, 5)
		generateServiceEmptyPanelIDsMapButton(service.Button, serviceName, servicePanelIDsMap)
	}
	return servicePanelIDsMap
}
func generateServiceEmptyPanelIDsMapButton(button *Button, serviceName string, servicePanelIDsMap map[string][]string) {
	if len(button.Panels) == 0 {
		if len(button.PanelInnerHTMLID) > 0 {
			servicePanelIDsMap[serviceName] = append(servicePanelIDsMap[serviceName], button.PanelInnerHTMLID)
		}
		return
	}
	for _, panel := range button.Panels {
		switch {
		case len(panel.Tabs) > 0:
			for _, tab := range panel.Tabs {
				generateServiceEmptyPanelIDsMapTab(tab, serviceName, servicePanelIDsMap)
			}
		case len(panel.Buttons) > 0:
			for _, b := range panel.Buttons {
				generateServiceEmptyPanelIDsMapButton(b, serviceName, servicePanelIDsMap)
			}
		default:
			if len(panel.HTMLID) > 0 {
				servicePanelIDsMap[serviceName] = append(servicePanelIDsMap[serviceName], panel.HTMLID)
			}
		}
	}
}
func generateServiceEmptyPanelIDsMapTab(tab *Tab, serviceName string, servicePanelIDsMap map[string][]string) {
	if len(tab.Panels) == 0 {
		if len(tab.PanelInnerHTMLID) > 0 {
			servicePanelIDsMap[serviceName] = append(servicePanelIDsMap[serviceName], tab.PanelInnerHTMLID)
		}
		return
	}
	for _, panel := range tab.Panels {
		servicePanelIDsMap[serviceName] = append(servicePanelIDsMap[serviceName], panel.HTMLID)
	}
}

// GenerateServiceEmptyInsidePanelIDsMap returns
//   each service name mapped to
//   a map of each markup panel's name mapped to the html id of it's inner most empty div
//     where the markup panel's template will be included.
func (builder *Builder) GenerateServiceEmptyInsidePanelIDsMap() map[string]map[string]string {
	servicePanelIDsMap := make(map[string]map[string]string)
	for _, s := range builder.Services {
		serviceName := s.Name
		servicePanelIDsMap[serviceName] = make(map[string]string)
		generateServiceEmptyInsidePanelIDsMapButton(s.Button, serviceName, servicePanelIDsMap)
	}
	return servicePanelIDsMap
}
func generateServiceEmptyInsidePanelIDsMapButton(button *Button, serviceName string, servicePanelIDsMap map[string]map[string]string) {
	for _, panel := range button.Panels {
		switch {
		case len(panel.Tabs) > 0:
			for _, tab := range panel.Tabs {
				generateServiceEmptyInsidePanelIDsMapTab(tab, serviceName, servicePanelIDsMap)
			}
		case len(panel.Buttons) > 0:
			for _, b := range panel.Buttons {
				generateServiceEmptyInsidePanelIDsMapButton(b, serviceName, servicePanelIDsMap)
			}
		default:
			if len(panel.HTMLID) > 0 {
				servicePanelIDsMap[serviceName][panel.Name] = panel.innerID()
			}
		}
	}
}
func generateServiceEmptyInsidePanelIDsMapTab(tab *Tab, serviceName string, servicePanelIDsMap map[string]map[string]string) {
	for _, panel := range tab.Panels {
		servicePanelIDsMap[serviceName][panel.Name] = panel.innerID()
	}
}

// GenerateTabBarLevelStartPanelMap returns a level mapped to the id of its first tab bar panel.
func (builder *Builder) GenerateTabBarLevelStartPanelMap() map[string]string {
	tabBarPanelMap := make(map[string]string)
	for _, s := range builder.Services {
		for _, p := range s.Button.Panels {
			p.getTabBarLastPanelIDs(tabBarPanelMap)
		}
	}
	return tabBarPanelMap
}
func (panel *Panel) getTabBarLastPanelIDs(tabBarPanelMap map[string]string) {
	if len(panel.Tabs) > 0 {
		level := strings.Split(panel.TabBarHTMLID, dashString)[0]
		tabBarPanelMap[level] = panel.Tabs[0].PanelHTMLID
		for _, tab := range panel.Tabs {
			for _, p := range tab.Panels {
				if len(p.Tabs) > 0 || len(p.Buttons) > 0 {
					p.getTabBarLastPanelIDs(tabBarPanelMap)
				}
			}
		}
		return
	}
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			p.getTabBarLastPanelIDs(tabBarPanelMap)
		}
	}
}

// ButtonPanelGroup is a button panel group information.
type ButtonPanelGroup struct {
	IsTabButton     bool
	ButtonName      string
	ButtonID        string
	PanelNamesIDMap map[string]*Panel // [panel name]*Panel
}

// NewButtonPanelGroup constructs a new ButtonPanelGroup
func NewButtonPanelGroup() *ButtonPanelGroup {
	return &ButtonPanelGroup{
		PanelNamesIDMap: make(map[string]*Panel),
	}
}

// GenerateServiceButtonPanelGroups returns
//   each service name mapped to
//   []*ButtonPanelGroup
//      A button panel group struct represents a button-pad or tab-bar button.
//        .IsTabButton : if the button is a tab bar button or not,
//        .ButtonName : the button's name created made from it's html id
//        .ButtonID : the button's html id
//        .PanelNamesIDMap : a map of the button's group of panels
//           where each panel name is mapped to it's Panel struct.
func (builder *Builder) GenerateServiceButtonPanelGroups() map[string][]*ButtonPanelGroup {
	serviceButtonPanelGroupMap := make(map[string][]*ButtonPanelGroup)
	for _, s := range builder.Services {
		// service
		serviceName := s.Name
		list := make([]*ButtonPanelGroup, 0, 5)
		// button panel groups
		g := NewButtonPanelGroup()
		list = append(list, g)
		b := s.Button
		g.ButtonID = b.HTMLID
		//g.ButtonName = panelIDToName(b.HTMLID)
		g.ButtonName = b.ID
		for _, p := range b.Panels {
			g.PanelNamesIDMap[p.Name] = p
			p.generateServiceButtonPanelGroups(&list)
		}
		serviceButtonPanelGroupMap[serviceName] = list
	}
	return serviceButtonPanelGroupMap
}
func (panel *Panel) generateServiceButtonPanelGroups(list *[]*ButtonPanelGroup) {
	for _, b := range panel.Buttons {
		// button panel groups
		g := NewButtonPanelGroup()
		*list = append(*list, g)
		g.ButtonID = b.HTMLID
		//g.ButtonName = panelIDToName(b.HTMLID)
		g.ButtonName = b.ID
		for _, p := range b.Panels {
			g.PanelNamesIDMap[p.Name] = p
			p.generateServiceButtonPanelGroups(list)
		}
	}
	for _, t := range panel.Tabs {
		// tab bar
		g := NewButtonPanelGroup()
		g.IsTabButton = true
		*list = append(*list, g)
		g.ButtonID = t.HTMLID
		//g.ButtonName = panelIDToName(t.HTMLID)
		g.ButtonName = t.ID
		for _, p := range t.Panels {
			g.PanelNamesIDMap[p.Name] = p
			// go no deeper with tab bars.
		}
	}
}

// GenerateTabBarIDs returns a slice of each tab bar's html id.
func (builder *Builder) GenerateTabBarIDs() []string {
	// only panels have tabs
	ids := make([]string, 0, 5)
	for _, s := range builder.Services {
		for _, p := range s.Button.Panels {
			generateTabBarIDs(p, &ids)
		}
	}
	return ids
}
func generateTabBarIDs(panel *Panel, ids *[]string) {
	if len(panel.Tabs) > 0 {
		*ids = append(*ids, panel.TabBarHTMLID)
		for _, tab := range panel.Tabs {
			for _, panel := range tab.Panels {
				generateTabBarIDs(panel, ids)
			}
		}
		return
	}
	for _, button := range panel.Buttons {
		for _, p := range button.Panels {
			generateTabBarIDs(p, ids)
		}
	}
}

// GenerateServicePanelNameTemplateMap returns a map of
//   each service name mapped to
//   a map of each markup panel's name mapped to it's template (note and markup).
// Where each panel is a markup panel.
func (builder *Builder) GenerateServicePanelNameTemplateMap() map[string]map[string]string {
	servicePanelNameTemplateMap := make(map[string]map[string]string)
	for _, s := range builder.Services {
		if !s.Generated {
			panelNameTemplateMap := make(map[string]string)
			for _, p := range s.Button.Panels {
				generateServicePanelNameTemplateMap(p, panelNameTemplateMap)
			}
			servicePanelNameTemplateMap[s.Name] = panelNameTemplateMap
		}
	}
	return servicePanelNameTemplateMap
}
func generateServicePanelNameTemplateMap(panel *Panel, panelNameTemplateMap map[string]string) {
	if len(panel.Template) > 0 {
		panelNameTemplateMap[panel.Name] = panel.Template
		return
	}
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			generateServicePanelNameTemplateMap(p, panelNameTemplateMap)
		}
	}
	for _, t := range panel.Tabs {
		for _, p := range t.Panels {
			generateServicePanelNameTemplateMap(p, panelNameTemplateMap)
		}
	}
}

// GenerateServiceTemplatePanelName returns a map of
//   each service name mapped to
//   a slice of markup panel names.
func (builder *Builder) GenerateServiceTemplatePanelName() map[string][]string {
	servicePanelNameTemplateMap := make(map[string][]string)
	for _, s := range builder.Services {
		if !s.Generated {
			panelNameList := make([]string, 0, 5)
			for _, p := range s.Button.Panels {
				generateServiceTemplatePanelName(p, &panelNameList)
			}
			servicePanelNameTemplateMap[s.Name] = panelNameList
		}
	}
	return servicePanelNameTemplateMap
}
func generateServiceTemplatePanelName(panel *Panel, panelNameList *[]string) {
	if len(panel.Template) > 0 {
		*panelNameList = append(*panelNameList, panel.Name)
		return
	}
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			generateServiceTemplatePanelName(p, panelNameList)
		}
	}
	for _, t := range panel.Tabs {
		for _, p := range t.Panels {
			generateServiceTemplatePanelName(p, panelNameList)
		}
	}
}

// GenerateServiceEmptyInsidePanelNamePathMap returns a map of
//   each service name mapped to
//   a map of each markup panel's name mapped to a slice of that panel's full relevant path
func (builder *Builder) GenerateServiceEmptyInsidePanelNamePathMap() map[string]map[string][]string {
	serviceEmptyInsidePanelNamePathMap := make(map[string]map[string][]string)
	for _, s := range builder.Services {
		if !s.Generated {
			panelNamePathMap := make(map[string][]string)
			for _, p := range s.Button.Panels {
				folderList := make([]string, 1, 10)
				folderList[0] = s.Button.ID
				generateServiceEmptyInsidePanelNamePathMap(p, folderList, panelNamePathMap)
			}
			serviceEmptyInsidePanelNamePathMap[s.Name] = panelNamePathMap
		}
	}
	return serviceEmptyInsidePanelNamePathMap
}
func generateServiceEmptyInsidePanelNamePathMap(panel *Panel, folderList []string, panelNamePathMap map[string][]string) {
	if len(panel.Template) > 0 {
		panelNamePathMap[panel.Name] = folderList
		return
	}
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	for _, b := range panel.Buttons {
		newFolderList := make([]string, l+1, l*2)
		for i, f := range folderList {
			newFolderList[i] = f
		}
		newFolderList[l] = b.ID
		for _, p := range b.Panels {
			generateServiceEmptyInsidePanelNamePathMap(p, newFolderList, panelNamePathMap)

		}
	}
	for _, t := range panel.Tabs {
		newFolderList := make([]string, l+1, l*2)
		for i, f := range folderList {
			newFolderList[i] = f
		}
		newFolderList[l] = t.ID
		for _, p := range t.Panels {
			//generateServiceEmptyInsidePanelNamePathMap(p, folderList, panelNamePathMap)
			generateServiceEmptyInsidePanelNamePathMap(p, newFolderList, panelNamePathMap)
		}
	}
}

// GenerateServicePanelNamePanelMap returns a map of
//   each service name mapped to
//   a map of each panel name mapped to it's panel.
func (builder *Builder) GenerateServicePanelNamePanelMap() map[string]map[string]*Panel {
	servicePanelNamePanelMap := make(map[string]map[string]*Panel)
	for _, s := range builder.Services {
		if !s.Generated {
			panelNamePanelMap := make(map[string]*Panel)
			for _, p := range s.Button.Panels {
				panelNamePanelMap[p.Name] = p
				generateServicePanelNamePanelMap(p, panelNamePanelMap)
			}
			servicePanelNamePanelMap[s.Name] = panelNamePanelMap
		}
	}
	return servicePanelNamePanelMap
}
func generateServicePanelNamePanelMap(panel *Panel, panelNamePanelMap map[string]*Panel) {
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			panelNamePanelMap[p.Name] = p
			generateServicePanelNamePanelMap(p, panelNamePanelMap)
		}
	}
	for _, t := range panel.Tabs {
		for _, p := range t.Panels {
			panelNamePanelMap[p.Name] = p
			generateServicePanelNamePanelMap(p, panelNamePanelMap)
		}
	}
}

const servicePanelFolderName = ""

// GenerateServicePanelButtonFolderPathMap returns a map of
//  each service name mapped to
//  each panel name mapped to
//  a slice of each panel button name mapped to its button's full relevant folder path.
func (builder *Builder) GenerateServicePanelButtonFolderPathMap() map[string]map[string]map[string][]string {
	servicePanelButtonFolderPathMap := make(map[string]map[string]map[string][]string)
	for _, s := range builder.Services {
		if !s.Generated {
			// map the home button to its folder.
			panelButtonFolderPathMap := make(map[string]map[string][]string)
			panelButtonFolderPathMap[servicePanelFolderName] = make(map[string][]string)
			folderList := make([]string, 1, 20)
			folderList[0] = s.Button.ID
			panelButtonFolderPathMap[servicePanelFolderName][s.Button.ID] = folderList
			// the service button is done.
			// now continue with each of the service button's panels
			for _, p := range s.Button.Panels {
				// map this panel to it's folder path.
				if len(p.Buttons) > 0 {
					panelButtonFolderPathMap[p.Name] = make(map[string][]string)
					// create the folder list for the walk.
					// the folder list begins with the service button's id
					folderList := make([]string, 1, 20)
					folderList[0] = s.Button.ID
					generateServicePanelButtonFolderPathMap(p, folderList, panelButtonFolderPathMap)
				}
			}
			servicePanelButtonFolderPathMap[s.Name] = panelButtonFolderPathMap
		}
	}
	return servicePanelButtonFolderPathMap
}
func generateServicePanelButtonFolderPathMap(panel *Panel, folderList []string, panelButtonFolderPathMap map[string]map[string][]string) {
	// the folderlist continues with this panel's id
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	for _, b := range panel.Buttons {
		buttonFolderList := make([]string, l+1, 20)
		for i, f := range folderList {
			buttonFolderList[i] = f
		}
		buttonFolderList[l] = b.ID
		panelButtonFolderPathMap[panel.Name][b.ID] = buttonFolderList
		for _, p := range b.Panels {
			if len(p.Buttons) > 0 {
				panelButtonFolderPathMap[p.Name] = make(map[string][]string)
				// create the folder list for the walk.
				newFolderList := make([]string, l+1, 20)
				for i, f := range buttonFolderList {
					newFolderList[i] = f
				}
				newFolderList[l] = b.ID
				panelButtonFolderPathMap[panel.Name][b.ID] = newFolderList
				generateServicePanelButtonFolderPathMap(p, newFolderList, panelButtonFolderPathMap)
			}
		}
	}
}
