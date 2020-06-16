package proofs

var (
	homeNames = []string{"Action1Button", "Action2Button", "Action3Button", "Action4Button", "Action5Button"}

	buttonNamePanelNames = map[string][]string{"Action1Button":[]string{"Action1Level1ButtonPanel"}, "Action1Level1ContentButton":[]string{"Action1Level1MarkupPanel"}, "Action1Level2ContentButton":[]string{"Action1Level2MarkupPanel"}, "Action1Level3ContentButton":[]string{"Action1Level3MarkupPanel"}, "Action1Level4ContentButton":[]string{"Action1Level4MarkupPanel"}, "Action1Level5ContentButton":[]string{"Action1Level5MarkupPanel"}, "Action1ToLevel2ColorsButton":[]string{"Action1Level2ButtonPanel"}, "Action1ToLevel3ColorsButton":[]string{"Action1Level3ButtonPanel"}, "Action1ToLevel4ColorsButton":[]string{"Action1Level4ButtonPanel"}, "Action1ToLevel5ColorsButton":[]string{"Action1Level5ButtonPanel"}, "Action2Button":[]string{"Action2Level1ButtonPanel"}, "Action2Level1ContentButton":[]string{"Action2Level1MarkupPanel"}, "Action2Level2ContentButton":[]string{"Action2Level2MarkupPanel"}, "Action2Level3ContentButton":[]string{"Action2Level3MarkupPanel"}, "Action2Level4ContentButton":[]string{"Action2Level4MarkupPanel"}, "Action2Level5ContentButton":[]string{"Action2Level5MarkupPanel"}, "Action2ToLevel2ColorsButton":[]string{"Action2Level2ButtonPanel"}, "Action2ToLevel3ColorsButton":[]string{"Action2Level3ButtonPanel"}, "Action2ToLevel4ColorsButton":[]string{"Action2Level4ButtonPanel"}, "Action2ToLevel5ColorsButton":[]string{"Action2Level5ButtonPanel"}, "Action3Button":[]string{"Action3Level1ButtonPanel"}, "Action3Level1ContentButton":[]string{"Action3Level1MarkupPanel"}, "Action3Level2ContentButton":[]string{"Action3Level2MarkupPanel"}, "Action3Level3ContentButton":[]string{"Action3Level3MarkupPanel"}, "Action3Level4ContentButton":[]string{"Action3Level4MarkupPanel"}, "Action3Level5ContentButton":[]string{"Action3Level5MarkupPanel"}, "Action3ToLevel2ColorsButton":[]string{"Action3Level2ButtonPanel"}, "Action3ToLevel3ColorsButton":[]string{"Action3Level3ButtonPanel"}, "Action3ToLevel4ColorsButton":[]string{"Action3Level4ButtonPanel"}, "Action3ToLevel5ColorsButton":[]string{"Action3Level5ButtonPanel"}, "Action4Button":[]string{"Action4Level1ButtonPanel"}, "Action4Level1ContentButton":[]string{"Action4Level1MarkupPanel"}, "Action4Level2ContentButton":[]string{"Action4Level2MarkupPanel"}, "Action4Level3ContentButton":[]string{"Action4Level3MarkupPanel"}, "Action4Level4ContentButton":[]string{"Action4Level4MarkupPanel"}, "Action4Level5ContentButton":[]string{"Action4Level5MarkupPanel"}, "Action4ToLevel2ColorsButton":[]string{"Action4Level2ButtonPanel"}, "Action4ToLevel3ColorsButton":[]string{"Action4Level3ButtonPanel"}, "Action4ToLevel4ColorsButton":[]string{"Action4Level4ButtonPanel"}, "Action4ToLevel5ColorsButton":[]string{"Action4Level5ButtonPanel"}, "Action5Button":[]string{"Action5Level1ButtonPanel"}, "Action5Level1ContentButton":[]string{"Action5Level1MarkupPanel"}, "Action5Level2ContentButton":[]string{"Action5Level2MarkupPanel"}, "Action5Level3ContentButton":[]string{"Action5Level3MarkupPanel"}, "Action5Level4ContentButton":[]string{"Action5Level4MarkupPanel"}, "Action5Level5ContentButton":[]string{"Action5Level5MarkupPanel"}, "Action5ToLevel2ColorsButton":[]string{"Action5Level2ButtonPanel"}, "Action5ToLevel3ColorsButton":[]string{"Action5Level3ButtonPanel"}, "Action5ToLevel4ColorsButton":[]string{"Action5Level4ButtonPanel"}, "Action5ToLevel5ColorsButton":[]string{"Action5Level5ButtonPanel"}}

	tabNamePanelNames = map[string][]string{}

	panelNameButtonNames = map[string][]string{"Action1Level1ButtonPanel":[]string{"Action1Level1ContentButton", "Action1ToLevel2ColorsButton"}, "Action2Level1ButtonPanel":[]string{"Action2Level1ContentButton", "Action2ToLevel2ColorsButton"}, "Action3Level1ButtonPanel":[]string{"Action3Level1ContentButton", "Action3ToLevel2ColorsButton"}, "Action4Level1ButtonPanel":[]string{"Action4Level1ContentButton", "Action4ToLevel2ColorsButton"}, "Action5Level1ButtonPanel":[]string{"Action5Level1ContentButton", "Action5ToLevel2ColorsButton"}}

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
