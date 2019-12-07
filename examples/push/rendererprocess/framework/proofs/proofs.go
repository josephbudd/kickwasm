package proofs

var (
	homeNames = []string{"PushButton"}

	buttonNamePanelNames = map[string][]string{"PushButton":[]string{"PushPanel"}}

	tabNamePanelNames = map[string][]string{}

	panelNameButtonNames = map[string][]string{}

	panelNameTabNames = map[string][]string{}
)

// HomeButtonsNames returns the names of each home button.
func HomeButtonsNames() (names []string) {
	names = make([]string, len(homeNames))
	copy(names, homeNames)
	return
}

// ButtonNamePanelNames returns each button name mapped to the button's panel names.
func ButtonNamePanelNames() (buttonPanels map[string][]string) {
	buttonPanels = make(map[string][]string, len(buttonNamePanelNames))
	for buttonID, panelIDs := range buttonNamePanelNames {
		names := make([]string, len(panelIDs))
		for i, id := range panelIDs {
			names[i] = id
		}
		buttonPanels[buttonID] = names
	}
	return
}

// TabNamePanelNames returns each tab name mapped to the tab's panel name.
func TabNamePanelNames() (tabPanels map[string][]string) {
	tabPanels = make(map[string][]string, len(tabNamePanelNames))
	for tabID, panelIDs := range tabNamePanelNames {
		names := make([]string, len(panelIDs))
		for i, id := range panelIDs {
			names[i] = id
		}
		tabPanels[tabID] = names
	}
	return
}

// PanelNameTabNames returns each panel name mapped to its tab names.
func PanelNameTabNames() (panelTabs map[string][]string) {
	panelTabs = make(map[string][]string, len(panelNameTabNames))
	for panelID, tabIDs := range panelNameTabNames {
		names := make([]string, len(tabIDs))
		for i, id := range tabIDs {
			names[i] = id
		}
		panelTabs[panelID] = names
	}
	return
}

// PanelNameButtonNames returns each panel name mapped to its button names.
func PanelNameButtonNames() (panelButtons map[string][]string) {
	panelButtons = make(map[string][]string, len(panelNameTabNames))
	for panelID, buttonIDs := range panelNameButtonNames {
		names := make([]string, len(buttonIDs))
		for i, id := range buttonIDs {
			names[i] = id
		}
		panelButtons[panelID] = names
	}
	return
}

// PanelsInSameButtonGroup checks to see if the 2 panels are in the same button panel group.
// Returns the button name and if they are in the same group.
func PanelsInSameButtonGroup(p1Name, p2Name string) (buttonName string, inSame bool) {
	var pNames []string
	for buttonName, pNames = range buttonNamePanelNames {
		foundP1 := false
		for _, pName := range pNames {
			if p1Name == pName {
				foundP1 = true
			} 
		}
		if !foundP1 {
			return
		}
		for _, pName := range pNames {
			if p1Name == pName {
				inSame = true
				return
			} 
		}
	}
	return
}

// PanelsInSameTabGroup checks to see if the 2 panels are in the same tab panel group.
// Returns the button name and if they are in the same group.
func PanelsInSameTabGroup(p1Name, p2Name string) (tabName string, inSame bool) {
	var pNames []string
	for tabName, pNames = range tabNamePanelNames {
		foundP1 := false
		for _, pName := range pNames {
			if p1Name == pName {
				foundP1 = true
			} 
		}
		if !foundP1 {
			return
		}
		for _, pName := range pNames {
			if p1Name == pName {
				inSame = true
				return
			} 
		}
	}
	return
}
