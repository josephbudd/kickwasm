package project

// ProofsHomeButtonNames returns the names of each home button.
func (builder *Builder) ProofsHomeButtonNames() (names []string) {
	l := len(builder.Homes)
	names = make([]string, l)
	for i, b := range builder.Homes {
		names[i] = b.ID
	}
	return
}

// ProofsButtonNamePanelNames returns each button id mapped to it's panel ids.
func (builder *Builder) ProofsButtonNamePanelNames() (buttonPanels map[string][]string) {
	buttonPanels = make(map[string][]string, 100)
	for _, b := range builder.Homes {
		proofsButtonNamePanelNames(b, buttonPanels)
	}
	return
}
func proofsButtonNamePanelNames(button *Button, buttonPanels map[string][]string) {
	l := len(button.Panels)
	names := make([]string, l)
	for i, p := range button.Panels {
		names[i] = p.ID
		switch {
		case len(p.Buttons) > 0:
			for _, b := range p.Buttons {
				proofsButtonNamePanelNames(b, buttonPanels)
			}
		case len(p.Tabs) > 0:
			for _, t := range p.Tabs {
				proofsButtonNamePanelNamesTabs(t, buttonPanels)
			}
		}
	}
	buttonPanels[button.ID] = names
}
func proofsButtonNamePanelNamesTabs(t *Tab, buttonPanels map[string][]string) {
	for _, p := range t.Panels {
		for _, b := range p.Buttons {
			proofsButtonNamePanelNames(b, buttonPanels)
		}
	}
}

// ProofsTabNamePanelNames returns each tab id mapped to it's panel ids.
func (builder *Builder) ProofsTabNamePanelNames() (tabPanels map[string][]string) {
	tabPanels = make(map[string][]string, 100)
	for _, b := range builder.Homes {
		for _, p := range b.Panels {
			switch {
			case len(p.Tabs) > 0:
				for _, t := range p.Tabs {
					proofsTabNamePanelNames(t, tabPanels)
				}
			case len(p.Buttons) > 0:
				for _, b := range p.Buttons {
					proofsTabNamePanelNamesButtons(b, tabPanels)
				}
			}
		}
	}
	return
}
func proofsTabNamePanelNames(tab *Tab, tabPanels map[string][]string) {
	l := len(tab.Panels)
	names := make([]string, l)
	for i, p := range tab.Panels {
		names[i] = p.ID
	}
	tabPanels[tab.ID] = names
}
func proofsTabNamePanelNamesButtons(button *Button, tabPanels map[string][]string) {
	for _, p := range button.Panels {
		switch {
		case len(p.Buttons) > 0:
			for _, b := range p.Buttons {
				proofsTabNamePanelNamesButtons(b, tabPanels)
			}
		case len(p.Tabs) > 0:
			for _, t := range p.Tabs {
				proofsTabNamePanelNames(t, tabPanels)
			}
		}
	}
}

// ProofsPanelNameTabNames returns each tab id mapped to it's panel ids.
func (builder *Builder) ProofsPanelNameTabNames() (panelTabs map[string][]string) {
	panelTabs = make(map[string][]string, 100)
	for _, b := range builder.Homes {
		for _, p := range b.Panels {
			switch {
			case len(p.Tabs) > 0:
				proofsPanelNameTabNames(p, panelTabs)
			case len(p.Buttons) > 0:
				for _, b := range p.Buttons {
					proofsPanelNameTabNamesButtons(b, panelTabs)
				}
			}
		}
	}
	return
}
func proofsPanelNameTabNames(panel *Panel, panelTabs map[string][]string) {
	l := len(panel.Tabs)
	names := make([]string, l)
	for i, t := range panel.Tabs {
		names[i] = t.ID
	}
	panelTabs[panel.ID] = names
}
func proofsPanelNameTabNamesButtons(button *Button, panelTabs map[string][]string) {
	for _, p := range button.Panels {
		switch {
		case len(p.Buttons) > 0:
			for _, b := range p.Buttons {
				proofsPanelNameTabNamesButtons(b, panelTabs)
			}
		case len(p.Tabs) > 0:
			proofsPanelNameTabNames(p, panelTabs)
		}
	}
}

// ProofsPanelNameButtonNames returns each button id mapped to it's panel ids.
func (builder *Builder) ProofsPanelNameButtonNames() (panelTabs map[string][]string) {
	panelTabs = make(map[string][]string, 100)
	for _, b := range builder.Homes {
		for _, p := range b.Panels {
			if len(p.Buttons) > 0 {
				proofsPanelNameButtonNames(p, panelTabs)
			}
		}
	}
	return
}
func proofsPanelNameButtonNames(panel *Panel, panelTabs map[string][]string) {
	l := len(panel.Buttons)
	names := make([]string, l)
	for i, b := range panel.Buttons {
		names[i] = b.ID
	}
	panelTabs[panel.ID] = names
}
