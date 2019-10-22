package project

import (
	"os"
	"strings"
)

// GenerateHomeSpawnTabEmptyInsidePanelNamePathMap returns a map of
//   each home name mapped to
//   a map of each markup panel's name mapped to a slice of that panel's full relevant path
func (builder *Builder) GenerateHomeSpawnTabEmptyInsidePanelNamePathMap() map[string]map[string][]string {
	homeEmptyInsidePanelNamePathMap := make(map[string]map[string][]string)
	for _, homeButton := range builder.Homes {
		panelNamePathMap := make(map[string][]string)
		for _, p := range homeButton.Panels {
			folderList := make([]string, 1, 10)
			folderList[0] = homeButton.ID
			generateHomeSpawnTabEmptyInsidePanelNamePathMap(p, folderList, panelNamePathMap)
		}
		homeEmptyInsidePanelNamePathMap[homeButton.ID] = panelNamePathMap
	}
	return homeEmptyInsidePanelNamePathMap
}
func generateHomeSpawnTabEmptyInsidePanelNamePathMap(panel *Panel, folderList []string, panelNamePathMap map[string][]string) {
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	m := l + 1
	if len(panel.Buttons) > 0 {
		for _, b := range panel.Buttons {
			newFolderList := make([]string, m, m*2)
			copy(newFolderList, folderList)
			newFolderList[l] = b.ID
			for _, p := range b.Panels {
				generateHomeSpawnTabEmptyInsidePanelNamePathMap(p, newFolderList, panelNamePathMap)

			}
		}
		return
	}
	for _, t := range panel.Tabs {
		if t.Spawn {
			// newFolderList := make([]string, m+1)
			// copy(newFolderList, folderList)
			// newFolderList[l] = t.ID
			for _, p := range t.Panels {
				newFolderList := make([]string, m+1)
				copy(newFolderList, folderList)
				newFolderList[l] = t.ID
				newFolderList[m] = p.Name
				panelNamePathMap[p.Name] = newFolderList
			}
			// return
		}
	}
}

// TabPanelGroup is a button panel group information.
type TabPanelGroup struct {
	TabBarID        string
	TabBarName      string
	TabName         string
	TabID           string
	TabLabel        string
	PanelHeading    string
	PanelNamesIDMap map[string]*Panel // [panel name]*Panel
}

// NewTabPanelGroup constructs a new TabPanelGroup
func NewTabPanelGroup() *TabPanelGroup {
	return &TabPanelGroup{
		PanelNamesIDMap: make(map[string]*Panel),
	}
}

// GenerateHomeSpawnTabButtonPanelGroups returns
//   each home name mapped to
//   []*TabPanelGroup
//      A tab panel panel group struct represents
//        a single one of the tab-bar's spawn tabs.
//        .TabBarID : the tab's div id.
//        .TabName : the button's name created made from it's html id.
//        .TabID : the tab's html id.
//        .TabLabel : the tab's label.
//        .PanelNamesIDMap : a map of the spawn tab's group of panels
//           where each panel name is mapped to it's Panel struct.
func (builder *Builder) GenerateHomeSpawnTabButtonPanelGroups() (homeSpawnTabButtonPanelGroups map[string][]*TabPanelGroup) {
	homeSpawnTabButtonPanelGroups = make(map[string][]*TabPanelGroup)
	for _, homeButton := range builder.Homes {
		tpgList := make([]*TabPanelGroup, 0, 5)
		// button panel groups
		for _, p := range homeButton.Panels {
			p.generateHomeSpawnTabButtonPanelGroups(&tpgList)
		}
		homeSpawnTabButtonPanelGroups[homeButton.ID] = tpgList
	}
	return
}
func (panel *Panel) generateHomeSpawnTabButtonPanelGroups(tpgList *[]*TabPanelGroup) {
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			p.generateHomeSpawnTabButtonPanelGroups(tpgList)
		}
	}
	for _, t := range panel.Tabs {
		if t.Spawn {
			// spawn tab
			g := NewTabPanelGroup()
			*tpgList = append(*tpgList, g)
			g.TabBarID = panel.TabBarHTMLID
			g.TabBarName = panel.ID
			g.TabID = t.HTMLID
			g.TabName = t.ID
			g.TabLabel = t.Label
			g.PanelHeading = t.Heading
			for _, p := range t.Panels {
				g.PanelNamesIDMap[p.Name] = p
				// go no deeper with tab bars.
			}
		}
	}
}

// GenerateSpawnTabPanelPathsMap returns a map of
//   each tab name ( .ID ) mapped to
//   a slice of that tab's panel's full relevant paths
func (builder *Builder) GenerateSpawnTabPanelPathsMap() (tabPanelPaths map[string][]string) {
	tabPanelPaths = make(map[string][]string, 100)
	for _, homeButton := range builder.Homes {
		for _, p := range homeButton.Panels {
			folderList := make([]string, 1, 10)
			folderList[0] = homeButton.ID
			generateSpawnTabPanelPathsMap(p, folderList, tabPanelPaths)
		}
	}
	return
}
func generateSpawnTabPanelPathsMap(panel *Panel, folderList []string, tabPanelPaths map[string][]string) {
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	m := l + 1
	for _, b := range panel.Buttons {
		newFolderList := make([]string, l+1, l*2)
		copy(newFolderList, folderList)
		newFolderList[l] = b.ID
		for _, p := range b.Panels {
			generateSpawnTabPanelPathsMap(p, newFolderList, tabPanelPaths)

		}
	}
	for _, t := range panel.Tabs {
		if t.Spawn {
			newFolderList := make([]string, l+2)
			copy(newFolderList, folderList)
			newFolderList[l] = t.ID
			panelPaths := make([]string, len(t.Panels))
			for i, p := range t.Panels {
				newFolderList[m] = p.Name
				panelPaths[i] = strings.Join(newFolderList, string(os.PathSeparator))
			}
			tabPanelPaths[t.ID] = panelPaths
		}
	}
}

// GenerateSpawnTabMarkupPanelPathMap returns a map of
//   each spawn tab panels's markup mapped to it's path.
func (builder *Builder) GenerateSpawnTabMarkupPanelPathMap() (markupPath map[string]string) {
	markupPath = make(map[string]string, 100)
	for _, homeButton := range builder.Homes {
		for _, p := range homeButton.Panels {
			folderList := make([]string, 1, 10)
			folderList[0] = homeButton.ID
			generateSpawnTabMarkupPanelPathsMap(p, folderList, markupPath)
		}
	}
	return
}
func generateSpawnTabMarkupPanelPathsMap(panel *Panel, folderList []string, markupPath map[string]string) {
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	m := l + 1
	for _, b := range panel.Buttons {
		newFolderList := make([]string, m, m*2)
		copy(newFolderList, folderList)
		newFolderList[l] = b.ID
		for _, p := range b.Panels {
			generateSpawnTabMarkupPanelPathsMap(p, newFolderList, markupPath)

		}
	}
	for _, t := range panel.Tabs {
		if t.Spawn {
			newFolderList := make([]string, m+1)
			copy(newFolderList, folderList)
			newFolderList[l] = t.ID
			for _, p := range t.Panels {
				newFolderList[m] = p.Name
				markupPath[p.Template] = strings.Join(newFolderList, string(os.PathSeparator))
			}
		}
	}
}

// GenerateSpawnTabBarPanelPathsMap returns a map of
//   each tab bar name ( .ID ) mapped to
//   a slice of that tab bar panel's full relevant path
func (builder *Builder) GenerateSpawnTabBarPanelPathsMap() (tabBarPanelPaths map[string][]string) {
	tabBarPanelPaths = make(map[string][]string, 100)
	for _, homeButton := range builder.Homes {
		for _, p := range homeButton.Panels {
			folderList := make([]string, 1, 10)
			folderList[0] = homeButton.ID
			generateSpawnTabBarPanelPathsMap(p, folderList, tabBarPanelPaths)
		}
	}
	return
}
func generateSpawnTabBarPanelPathsMap(panel *Panel, folderList []string, tabBarPanelPaths map[string][]string) {
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	var spawns bool
	if len(panel.Tabs) > 0 {
		for _, t := range panel.Tabs {
			if t.Spawn {
				spawns = true
				break
			}
		}
		if spawns {
			newFolderList := make([]string, l)
			copy(newFolderList, folderList)
			tabBarPanelPaths[panel.Name] = newFolderList
		}
		return
	}
	for _, b := range panel.Buttons {
		newFolderList := make([]string, l+1, l*2)
		copy(newFolderList, folderList)
		newFolderList[l] = b.ID
		for _, p := range b.Panels {
			generateSpawnTabBarPanelPathsMap(p, newFolderList, tabBarPanelPaths)

		}
	}
}

// GenerateSpawnTabBarTabPanelPathsMap returns a map of
//   each tab bar name ( .ID ) mapped to
//   a slice of that tab bar's tabs' full relevant paths
func (builder *Builder) GenerateSpawnTabBarTabPanelPathsMap() (tabBarPanelTabPaths map[string][][]string) {
	tabBarPanelTabPaths = make(map[string][][]string, 100)
	for _, homeButton := range builder.Homes {
		for _, p := range homeButton.Panels {
			folderList := make([]string, 1, 10)
			folderList[0] = homeButton.ID
			generateSpawnTabBarTabPanelPathsMap(p, folderList, tabBarPanelTabPaths)
		}
	}
	return
}
func generateSpawnTabBarTabPanelPathsMap(panel *Panel, folderList []string, tabBarPanelTabPaths map[string][][]string) {
	folderList = append(folderList, panel.Name)
	l := len(folderList)
	if len(panel.Tabs) > 0 {
		for _, t := range panel.Tabs {
			if t.Spawn {
				newFolderList := make([]string, l+1)
				copy(newFolderList, folderList)
				newFolderList[l] = t.ID
				var paths [][]string
				var found bool
				if paths, found = tabBarPanelTabPaths[panel.Name]; !found {
					paths = make([][]string, 0, 20)
				}
				paths = append(paths, newFolderList)
				tabBarPanelTabPaths[panel.Name] = paths
			}
		}
		return
	}
	for _, b := range panel.Buttons {
		newFolderList := make([]string, l+1)
		copy(newFolderList, folderList)
		newFolderList[l] = b.ID
		for _, p := range b.Panels {
			generateSpawnTabBarTabPanelPathsMap(p, newFolderList, tabBarPanelTabPaths)

		}
	}
}

// GenerateHomeSpawnPanelNamePanelMap returns a map of
//   each home name mapped to
//   a map of each panel name mapped to it's panel.
func (builder *Builder) GenerateHomeSpawnPanelNamePanelMap() map[string]map[string]*Panel {
	homePanelNamePanelMap := make(map[string]map[string]*Panel)
	for _, homeButton := range builder.Homes {
		panelNamePanelMap := make(map[string]*Panel)
		for _, p := range homeButton.Panels {
			generateHomeSpawnPanelNamePanelMap(p, panelNamePanelMap)
		}
		homePanelNamePanelMap[homeButton.ID] = panelNamePanelMap
	}
	return homePanelNamePanelMap
}
func generateHomeSpawnPanelNamePanelMap(panel *Panel, panelNamePanelMap map[string]*Panel) {
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			generateHomeSpawnPanelNamePanelMap(p, panelNamePanelMap)
		}
	}
	for _, t := range panel.Tabs {
		if t.Spawn {
			for _, p := range t.Panels {
				panelNamePanelMap[p.Name] = p
				generateHomeSpawnPanelNamePanelMap(p, panelNamePanelMap)
			}
		}
	}
}
