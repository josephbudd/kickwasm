package project

// GenerateButtonIDsPanelIDs returns button ids mapped to their panel ids.
// Only used by rekickwasm.
func (builder *Builder) GenerateButtonIDsPanelIDs() (buttons map[string][]string) {
	buttons = make(map[string][]string, 100)
	// start with home buttons
	for _, homeButton := range builder.Homes {
		l := len(homeButton.Panels)
		pids := make([]string, l, l)
		for i, p := range homeButton.Panels {
			pids[i] = p.ID
			generateButtonIDsPanelIDs(p, buttons)
		}
		buttons[homeButton.ID] = pids
	}
	return
}
func generateButtonIDsPanelIDs(panel *Panel, buttons map[string][]string) {
	for _, b := range panel.Buttons {
		l := len(b.Panels)
		pids := make([]string, l, l)
		for i, p := range b.Panels {
			pids[i] = p.ID
			generateButtonIDsPanelIDs(p, buttons)
		}
		buttons[b.ID] = pids
	}
	for _, t := range panel.Tabs {
		if !t.Spawn {
			for _, p := range t.Panels {
				generateButtonIDsPanelIDs(p, buttons)
			}
		}
	}
}

// TabSpawnPanelIDs is a tab's spawn and it's panel ids.
type TabSpawnPanelIDs struct {
	Spawn    bool
	PanelIDs []string
}

// GenerateTabIDsPanelIDs returns tab ids mapped to their panel ids.
// Only used by rekickwasm.
func (builder *Builder) GenerateTabIDsPanelIDs() (tabs map[string]TabSpawnPanelIDs) {
	tabs = make(map[string]TabSpawnPanelIDs, 100)
	// start with home buttons
	for _, homeButton := range builder.Homes {
		for _, p := range homeButton.Panels {
			generateTabIDsPanelIDs(p, tabs)
		}
	}
	return
}
func generateTabIDsPanelIDs(panel *Panel, tabs map[string]TabSpawnPanelIDs) {
	if len(panel.Tabs) > 0 {
		for _, t := range panel.Tabs {
			l := len(t.Panels)
			pids := make([]string, l, l)
			for i, p := range t.Panels {
				pids[i] = p.ID
				generateTabIDsPanelIDs(p, tabs)
			}
			// Also need to know if a tab is spawned.
			tabs[t.ID] = TabSpawnPanelIDs{
				Spawn:    t.Spawn,
				PanelIDs: pids,
			}
		}
	}
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			generateTabIDsPanelIDs(p, tabs)
		}
	}
}

// GenerateHomeButtonNames returns the home names.
func (builder *Builder) GenerateHomeButtonNames() (homeButtonNames []string) {
	homeButtonNames = make([]string, len(builder.Homes))
	for i, homeButton := range builder.Homes {
		homeButtonNames[i] = homeButton.ID
	}
	return
}

// GenerateHomeEmptyPanelIDsMap returns
//  each home name mapped to
//  a slice of the ids of the parent divs of the empty divs.
//  Each parent div
//   * is the parent of the panels in a panel group.
//   * is a child of the slider div.
func (builder *Builder) GenerateHomeEmptyPanelIDsMap() (homeEmptyPanelIDsMap map[string][]string) {
	homeEmptyPanelIDsMap = make(map[string][]string, len(builder.Homes))
	for _, homeButton := range builder.Homes {
		homeEmptyPanelIDsMap[homeButton.ID] = make([]string, 0, 5)
		generateHomeEmptyPanelIDsMapButton(homeButton, homeButton.ID, homeEmptyPanelIDsMap)
	}
	return
}
func generateHomeEmptyPanelIDsMapButton(button *Button, homeName string, homePanelIDsMap map[string][]string) {
	if len(button.Panels) == 0 {
		if len(button.PanelInnerHTMLID) > 0 {
			homePanelIDsMap[homeName] = append(homePanelIDsMap[homeName], button.PanelInnerHTMLID)
		}
		return
	}
	for _, panel := range button.Panels {
		switch {
		case len(panel.Tabs) > 0:
			if panel.HasRealTabs {
				for _, tab := range panel.Tabs {
					if !tab.Spawn {
						generateHomeEmptyPanelIDsMapTab(tab, homeName, homePanelIDsMap)
					}
				}
			}
		case len(panel.Buttons) > 0:
			for _, b := range panel.Buttons {
				generateHomeEmptyPanelIDsMapButton(b, homeName, homePanelIDsMap)
			}
		default:
			if len(panel.HTMLID) > 0 {
				homePanelIDsMap[homeName] = append(homePanelIDsMap[homeName], panel.HTMLID)
			}
		}
	}
}
func generateHomeEmptyPanelIDsMapTab(tab *Tab, homeName string, homePanelIDsMap map[string][]string) {
	if len(tab.Panels) == 0 {
		if len(tab.PanelInnerHTMLID) > 0 {
			homePanelIDsMap[homeName] = append(homePanelIDsMap[homeName], tab.PanelInnerHTMLID)
		}
		return
	}
	for _, panel := range tab.Panels {
		homePanelIDsMap[homeName] = append(homePanelIDsMap[homeName], panel.HTMLID)
	}
}

// GenerateTabBarIDStartPanelIDMap returns each tab bar id mapped to the id of its first tab bar panel.
func (builder *Builder) GenerateTabBarIDStartPanelIDMap() (tabBarIDStartPanelIDMap map[string]string) {
	tabBarIDStartPanelIDMap = make(map[string]string)
	for _, homeButton := range builder.Homes {
		for _, p := range homeButton.Panels {
			p.getTabBarLastPanelIDs(tabBarIDStartPanelIDMap)
		}
	}
	return
}
func (panel *Panel) getTabBarLastPanelIDs(tabBarIDStartPanelIDMap map[string]string) {
	if len(panel.Tabs) > 0 {
		tabBarID := panel.TabBarHTMLID
		tabBarIDStartPanelIDMap[tabBarID] = ""
		if panel.HasRealTabs {
			// find the first normal tab
			for _, tab := range panel.Tabs {
				if !tab.Spawn {
					tabBarIDStartPanelIDMap[tabBarID] = tab.PanelHTMLID
					break
				}
			}
			for _, tab := range panel.Tabs {
				if !tab.Spawn {
					for _, p := range tab.Panels {
						if len(p.Tabs) > 0 || len(p.Buttons) > 0 {
							p.getTabBarLastPanelIDs(tabBarIDStartPanelIDMap)
						}
					}
				}
			}
		}
		return
	}
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			p.getTabBarLastPanelIDs(tabBarIDStartPanelIDMap)
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

// GenerateHomeButtonPanelGroups returns
//   each home name mapped to
//   []*ButtonPanelGroup
//      A button panel group struct represents a button-pad or tab-bar button.
//        .IsTabButton : if the button is a tab bar button or not,
//        .ButtonName : the button's name created made from it's html id
//        .ButtonID : the button's html id
//        .PanelNamesIDMap : a map of the button's group of panels
//           where each panel name is mapped to it's Panel struct.
func (builder *Builder) GenerateHomeButtonPanelGroups() (homeButtonPanelGroups map[string][]*ButtonPanelGroup) {
	homeButtonPanelGroups = make(map[string][]*ButtonPanelGroup)
	for _, homeButton := range builder.Homes {
		list := make([]*ButtonPanelGroup, 0, 5)
		g := NewButtonPanelGroup()
		g.ButtonID = homeButton.HTMLID
		g.ButtonName = homeButton.ID
		list = append(list, g)
		for _, p := range homeButton.Panels {
			g.PanelNamesIDMap[p.Name] = p
			p.generateHomeButtonPanelGroups(&list)
		}
		homeButtonPanelGroups[homeButton.ID] = list
	}
	return
}
func (panel *Panel) generateHomeButtonPanelGroups(list *[]*ButtonPanelGroup) {
	for _, b := range panel.Buttons {
		// button panel groups
		g := NewButtonPanelGroup()
		g.ButtonID = b.HTMLID
		g.ButtonName = b.ID
		*list = append(*list, g)
		for _, p := range b.Panels {
			g.PanelNamesIDMap[p.Name] = p
			p.generateHomeButtonPanelGroups(list)
		}
	}
	for _, t := range panel.Tabs {
		if !t.Spawn {
			// tab bar
			g := NewButtonPanelGroup()
			g.IsTabButton = true
			g.ButtonID = t.HTMLID
			g.ButtonName = t.ID
			*list = append(*list, g)
			for _, p := range t.Panels {
				g.PanelNamesIDMap[p.Name] = p
				// go no deeper with tab bars.
			}
		}
	}
}

// GenerateTabBarIDs returns a slice of each tab bar's html id.
func (builder *Builder) GenerateTabBarIDs() (ids []string) {
	// only panels have tabs
	ids = make([]string, 0, 5)
	for _, homeButton := range builder.Homes {
		for _, p := range homeButton.Panels {
			generateTabBarIDs(p, &ids)
		}
	}
	return
}
func generateTabBarIDs(panel *Panel, ids *[]string) {
	if len(panel.Tabs) > 0 {
		*ids = append(*ids, panel.TabBarHTMLID)
		// for _, tab := range panel.Tabs {
		// 	for _, panel := range tab.Panels {
		// 		generateTabBarIDs(panel, ids)
		// 	}
		// }
		return
	}
	for _, button := range panel.Buttons {
		for _, p := range button.Panels {
			generateTabBarIDs(p, ids)
		}
	}
}

// GenerateHomePanelNameTemplateMap returns a map of
//   each home name mapped to
//   a map of each markup panel's name mapped to it's template (note and markup).
// Where each panel is a markup panel.
func (builder *Builder) GenerateHomePanelNameTemplateMap() (homePanelNameTemplateMap map[string]map[string]string) {
	homePanelNameTemplateMap = make(map[string]map[string]string)
	for _, homeButton := range builder.Homes {
		panelNameTemplateMap := make(map[string]string)
		for _, p := range homeButton.Panels {
			generateHomePanelNameTemplateMap(p, panelNameTemplateMap)
		}
		homePanelNameTemplateMap[homeButton.ID] = panelNameTemplateMap
	}
	return
}
func generateHomePanelNameTemplateMap(panel *Panel, panelNameTemplateMap map[string]string) {
	if len(panel.Template) > 0 {
		panelNameTemplateMap[panel.Name] = panel.Template
		return
	}
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			generateHomePanelNameTemplateMap(p, panelNameTemplateMap)
		}
	}
	for _, t := range panel.Tabs {
		if !t.Spawn {
			for _, p := range t.Panels {
				generateHomePanelNameTemplateMap(p, panelNameTemplateMap)
			}
		}
	}
}

// GenerateHomeTemplatePanelName returns a map of
//   each home name mapped to
//   a slice of markup panel names.
func (builder *Builder) GenerateHomeTemplatePanelName() (homePanelNameTemplateMap map[string][]string) {
	homePanelNameTemplateMap = make(map[string][]string)
	for _, homeButton := range builder.Homes {
		panelNameList := make([]string, 0, 5)
		for _, p := range homeButton.Panels {
			generateHomeTemplatePanelName(p, &panelNameList)
		}
		homePanelNameTemplateMap[homeButton.ID] = panelNameList
	}
	return
}
func generateHomeTemplatePanelName(panel *Panel, panelNameList *[]string) {
	if len(panel.Template) > 0 {
		*panelNameList = append(*panelNameList, panel.Name)
		return
	}
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			generateHomeTemplatePanelName(p, panelNameList)
		}
	}
	for _, t := range panel.Tabs {
		if !t.Spawn {
			for _, p := range t.Panels {
				generateHomeTemplatePanelName(p, panelNameList)
			}
		}
	}
}

// GenerateHomeEmptyInsidePanelNamePathMap returns a map of
//   each home name mapped to
//   a map of each markup panel's name mapped to a slice of that panel's full relevant path
func (builder *Builder) GenerateHomeEmptyInsidePanelNamePathMap() (homeEmptyInsidePanelNamePathMap map[string]map[string][]string) {
	homeEmptyInsidePanelNamePathMap = make(map[string]map[string][]string)
	for _, homeButton := range builder.Homes {
		panelNamePathMap := make(map[string][]string)
		for _, p := range homeButton.Panels {
			folderList := make([]string, 1, 10)
			folderList[0] = homeButton.ID
			generateHomeEmptyInsidePanelNamePathMap(p, folderList, panelNamePathMap)
		}
		homeEmptyInsidePanelNamePathMap[homeButton.ID] = panelNamePathMap
	}
	return
}
func generateHomeEmptyInsidePanelNamePathMap(panel *Panel, folderList []string, panelNamePathMap map[string][]string) {
	if len(panel.Template) > 0 {
		panelNamePathMap[panel.Name] = folderList
		return
	}
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	for _, b := range panel.Buttons {
		newFolderList := make([]string, l+1, l*2)
		copy(newFolderList, folderList)
		newFolderList[l] = b.ID
		for _, p := range b.Panels {
			generateHomeEmptyInsidePanelNamePathMap(p, newFolderList, panelNamePathMap)

		}
	}
	for _, t := range panel.Tabs {
		if !t.Spawn {
			newFolderList := make([]string, l+1, l*2)
			copy(newFolderList, folderList)
			newFolderList[l] = t.ID
			for _, p := range t.Panels {
				generateHomeEmptyInsidePanelNamePathMap(p, newFolderList, panelNamePathMap)
			}
		}
	}
}

// GenerateHomePanelNamePanelMap returns a map of
//   each home name mapped to
//   a map of each panel name mapped to it's panel.
func (builder *Builder) GenerateHomePanelNamePanelMap() (homePanelNamePanelMap map[string]map[string]*Panel) {
	homePanelNamePanelMap = make(map[string]map[string]*Panel)
	for _, homeButton := range builder.Homes {
		panelNamePanelMap := make(map[string]*Panel)
		for _, p := range homeButton.Panels {
			panelNamePanelMap[p.Name] = p
			generateHomePanelNamePanelMap(p, panelNamePanelMap)
		}
		homePanelNamePanelMap[homeButton.ID] = panelNamePanelMap
	}
	return
}
func generateHomePanelNamePanelMap(panel *Panel, panelNamePanelMap map[string]*Panel) {
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			panelNamePanelMap[p.Name] = p
			generateHomePanelNamePanelMap(p, panelNamePanelMap)
		}
	}
	for _, t := range panel.Tabs {
		if !t.Spawn {
			for _, p := range t.Panels {
				panelNamePanelMap[p.Name] = p
				generateHomePanelNamePanelMap(p, panelNamePanelMap)
			}
		}
	}
}

const homePanelFolderName = ""

// GenerateHomePanelButtonFolderPathMap returns a map of
//  each home name mapped to
//  each panel name mapped to
//  a slice of each panel button name mapped to its button's full relevant folder path.
func (builder *Builder) GenerateHomePanelButtonFolderPathMap() (homePanelButtonFolderPathMap map[string]map[string]map[string][]string) {
	homePanelButtonFolderPathMap = make(map[string]map[string]map[string][]string)
	for _, homeButton := range builder.Homes {
		// map the home button to its folder.
		panelButtonFolderPathMap := make(map[string]map[string][]string)
		panelButtonFolderPathMap[homePanelFolderName] = make(map[string][]string)
		folderList := make([]string, 1, 20)
		folderList[0] = homeButton.ID
		panelButtonFolderPathMap[homePanelFolderName][homeButton.ID] = folderList
		// the home button is done.
		// now continue with each of the home button's panels
		for _, p := range homeButton.Panels {
			// map this panel to it's folder path.
			if len(p.Buttons) > 0 {
				panelButtonFolderPathMap[p.Name] = make(map[string][]string)
				// create the folder list for the walk.
				// the folder list begins with the home button's id
				folderList := make([]string, 1, 20)
				folderList[0] = homeButton.ID
				generateHomePanelButtonFolderPathMap(p, folderList, panelButtonFolderPathMap)
			}
		}
		homePanelButtonFolderPathMap[homeButton.ID] = panelButtonFolderPathMap
	}
	return
}
func generateHomePanelButtonFolderPathMap(panel *Panel, folderList []string, panelButtonFolderPathMap map[string]map[string][]string) {
	// the folderlist continues with this panel's id
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	for _, b := range panel.Buttons {
		buttonFolderList := make([]string, l+1, 20)
		copy(buttonFolderList, folderList)
		buttonFolderList[l] = b.ID
		panelButtonFolderPathMap[panel.Name][b.ID] = buttonFolderList
		for _, p := range b.Panels {
			if len(p.Buttons) > 0 {
				panelButtonFolderPathMap[p.Name] = make(map[string][]string)
				// create the folder list for the walk.
				newFolderList := make([]string, l+1, 20)
				copy(newFolderList, buttonFolderList)
				newFolderList[l] = b.ID
				panelButtonFolderPathMap[panel.Name][b.ID] = newFolderList
				generateHomePanelButtonFolderPathMap(p, newFolderList, panelButtonFolderPathMap)
			}
		}
	}
}

func boolToInt(b bool) (i int) {
	if b {
		i = 1
	}
	return
}
