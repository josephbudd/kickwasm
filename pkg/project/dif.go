package project

import "fmt"

const (
	serviceNamed      = "service named %q"
	buttonLabeled     = "button labeled %q"
	buttonsPanelNamed = "button's panel named %q"
	tabLabeled        = "tab labeled %q"
	panelNamed        = "panel named %q"
	inThe             = "In the %s"
)

// Dif reports on the changes in test.
// It returns removals and additions.
//  removals are the entities contained in control but missing in test.
//  additions are the entities not found in control but found in test.
//  removals and additions are source locations mapped to a slice sub components which are either missing or added.
func (control *Builder) Dif(test *Builder) (removals, additions map[string][]string, somethingMoved bool) {
	removals = make(map[string][]string)
	additions = make(map[string][]string)
	somethingMoved = control.difServices(test, removals, additions)
	return
}

func (control *Builder) difServices(test *Builder, removals, additions map[string][]string) (somethingMoved bool) {
	// report removals
	for ci, cService := range control.Services {
		var matched *Service
		for ti, tService := range test.Services {
			if cService.Name == tService.Name {
				matched = tService
				if ci != ti {
					somethingMoved = true
				}
				break
			}
		}
		if matched == nil {
			removals[fmt.Sprintf(serviceNamed, cService.Name)] = nil
		}
	}
	// report additions
	for _, tService := range test.Services {
		var matched *Service
		for _, cService := range control.Services {
			if cService.Name == tService.Name {
				matched = cService
				break
			}
		}
		if matched == nil {
			additions[fmt.Sprintf(serviceNamed, tService.Name)] = nil
		}
	}
	// walk the unchanged services
	for _, tService := range test.Services {
		var matched *Service
		for _, cService := range control.Services {
			if cService.Name == tService.Name {
				matched = cService
				break
			}
		}
		if matched != nil {
			// check the service button for difs
			if difServiceButton(matched, tService, removals, additions) {
				somethingMoved = true
			}
		}
	}
	return
}

func difServiceButton(control *Service, test *Service, removals, additions map[string][]string) (somethingMoved bool) {
	// button label
	if control.Button.Label != test.Button.Label {
		key := fmt.Sprintf("In the service named %q", control.Name)
		if _, ok := removals[key]; !ok {
			removals[key] = make([]string, 0, 5)
		}
		if _, ok := additions[key]; !ok {
			additions[key] = make([]string, 0, 5)
		}
		removals[key] = append(removals[key], fmt.Sprintf(buttonLabeled, control.Button.Label))
		additions[key] = append(additions[key], fmt.Sprintf(buttonLabeled, test.Button.Label))
		return
	}
	// button panels
	path := fmt.Sprintf("service named %q, button labeled %q", control.Name, control.Button.Label)
	key := fmt.Sprintf(inThe, path)
	// report removals
	for _, cPanel := range control.Button.Panels {
		var matched *Panel
		for _, tPanel := range test.Button.Panels {
			if cPanel.Name == tPanel.Name {
				matched = tPanel
				break
			}
		}
		if matched == nil {
			if _, ok := removals[key]; !ok {
				removals[key] = make([]string, 0, 5)
			}
			removals[key] = append(removals[key], fmt.Sprintf(buttonsPanelNamed, cPanel.Name))
		}
	}
	// report additions
	for _, tPanel := range test.Button.Panels {
		var matched *Panel
		for _, cPanel := range control.Button.Panels {
			if cPanel.Name == tPanel.Name {
				matched = cPanel
				break
			}
		}
		if matched == nil {
			if _, ok := additions[key]; !ok {
				additions[key] = make([]string, 0, 5)
			}
			additions[key] = append(additions[key], fmt.Sprintf(buttonsPanelNamed, tPanel.Name))
		}
	}
	// walk the unchanged button panels
	for ti, tPanel := range test.Button.Panels {
		var matched *Panel
		for ci, cPanel := range control.Button.Panels {
			if cPanel.Name == tPanel.Name {
				matched = cPanel
				if ti != ci {
					somethingMoved = true
				}
				break
			}
		}
		if matched != nil {
			// check the service button for difs
			if difPanels(path, matched, tPanel, removals, additions) {
				somethingMoved = true
			}
		}
	}
	return
}

func difPanels(path string, control, test *Panel, removals, additions map[string][]string) (sometingMoved bool) {
	path = path + fmt.Sprintf(", panel named %q", control.Name)
	key := fmt.Sprintf(inThe, path)
	// the panel names already match
	// report button removals
	for ci, cButton := range control.Buttons {
		var matched *Button
		for ti, tButton := range test.Buttons {
			if cButton.Label == tButton.Label {
				matched = tButton
				if ci != ti {
					sometingMoved = true
				}
				break
			}
		}
		if matched == nil {
			if _, ok := removals[key]; !ok {
				removals[key] = make([]string, 0, 5)
			}
			removals[key] = append(removals[key], fmt.Sprintf(buttonLabeled, cButton.Label))
		}
	}
	// report tab removals
	for ci, cTab := range control.Tabs {
		var matched *Tab
		for ti, tTab := range test.Tabs {
			if cTab.Spawn == tTab.Spawn && cTab.Label == tTab.Label {
				matched = tTab
				if ci != ti {
					sometingMoved = true
				}
				break
			}
		}
		if matched == nil {
			if _, ok := removals[key]; !ok {
				removals[key] = make([]string, 0, 5)
			}
			removals[key] = append(removals[key], fmt.Sprintf(tabLabeled, cTab.Label))
		}
	}

	// report button additions
	for ti, tButton := range test.Buttons {
		var matched *Button
		for ci, cButton := range control.Buttons {
			if cButton.Label == tButton.Label {
				matched = cButton
				if ci != ti {
					sometingMoved = true
				}
				break
			}
		}
		if matched == nil {
			if _, ok := additions[key]; !ok {
				additions[key] = make([]string, 0, 5)
			}
			additions[key] = append(additions[key], fmt.Sprintf(buttonLabeled, tButton.Label))
		}
	}
	// report tab additions
	for ti, tTab := range test.Tabs {
		var matched *Tab
		for ci, cTab := range control.Tabs {
			if cTab.Spawn == tTab.Spawn && cTab.Label == tTab.Label {
				matched = cTab
				if ci != ti {
					sometingMoved = true
				}
				break
			}
		}
		if matched == nil {
			if _, ok := additions[key]; !ok {
				additions[key] = make([]string, 0, 5)
			}
			additions[key] = append(additions[key], fmt.Sprintf(tabLabeled, tTab.Label))
		}
	}

	// walk the unchanged panel buttons
	for ti, tButton := range test.Buttons {
		var matched *Button
		for ci, cButton := range control.Buttons {
			if cButton.Label == tButton.Label {
				matched = cButton
				if ci != ti {
					sometingMoved = true
				}
				break
			}
		}
		if matched != nil {
			// check the service button for difs
			difButtons(path, matched, tButton, removals, additions)
		}
	}
	// walk the unchanged panel tabs
	for ti, tTab := range test.Tabs {
		var matched *Tab
		for ci, cTab := range control.Tabs {
			if cTab.Spawn == tTab.Spawn && cTab.Label == tTab.Label {
				matched = cTab
				if ci != ti {
					sometingMoved = true
				}
				break
			}
		}
		if matched != nil {
			// check the service tabs for difs
			difTabs(path, matched, tTab, removals, additions)
		}
	}
	return
}

func difButtons(path string, control, test *Button, removals, additions map[string][]string) (somethingMoved bool) {
	path = path + fmt.Sprintf(", button labeled %q", control.Label)
	key := fmt.Sprintf(inThe, path)
	// the button labels already match
	// report removals
	for _, cPanel := range control.Panels {
		var matched *Panel
		for _, tPanel := range test.Panels {
			if cPanel.Name == tPanel.Name {
				matched = tPanel
				break
			}
		}
		if matched == nil {
			if _, ok := removals[key]; !ok {
				removals[key] = make([]string, 0, 5)
			}
			removals[key] = append(removals[key], fmt.Sprintf(panelNamed, cPanel.Name))
		}
	}
	// report additions
	for _, tPanel := range test.Panels {
		var matched *Panel
		for _, cPanel := range control.Panels {
			if cPanel.Name == tPanel.Name {
				matched = cPanel
				break
			}
		}
		if matched == nil {
			if _, ok := additions[key]; !ok {
				additions[key] = make([]string, 0, 5)
			}
			additions[key] = append(additions[key], fmt.Sprintf(panelNamed, tPanel.Name))
		}
	}
	// walk the unchanged button panels
	for _, tPanel := range test.Panels {
		var matched *Panel
		for _, cPanel := range control.Panels {
			if cPanel.Name == tPanel.Name {
				matched = cPanel
				break
			}
		}
		if matched != nil {
			// check the button panel for difs
			if difPanels(path, matched, tPanel, removals, additions) {
				somethingMoved = true
			}
		}
	}
	return
}

func difTabs(path string, control, test *Tab, removals, additions map[string][]string) (somethingMoved bool) {
	path = path + fmt.Sprintf(", tab labeled %q", control.Label)
	key := fmt.Sprintf(inThe, path)
	// the tab labels already match
	// report removals
	for _, cPanel := range control.Panels {
		var matched *Panel
		for _, tPanel := range test.Panels {
			if cPanel.Name == tPanel.Name {
				matched = tPanel
				break
			}
		}
		if matched == nil {
			if _, ok := removals[key]; !ok {
				removals[key] = make([]string, 0, 5)
			}
			removals[key] = append(removals[key], fmt.Sprintf(panelNamed, cPanel.Name))
		}
	}
	// report additions
	for _, tPanel := range test.Panels {
		var matched *Panel
		for _, cPanel := range control.Panels {
			if cPanel.Name == tPanel.Name {
				matched = cPanel
				break
			}
		}
		if matched == nil {
			if _, ok := additions[key]; !ok {
				additions[key] = make([]string, 0, 5)
			}
			additions[key] = append(additions[key], fmt.Sprintf(panelNamed, tPanel.Name))
		}
	}
	// walk the unchanged tab panels
	for _, tPanel := range test.Panels {
		var matched *Panel
		for _, cPanel := range control.Panels {
			if cPanel.Name == tPanel.Name {
				matched = cPanel
				break
			}
		}
		if matched != nil {
			// check the tab panel for difs
			if difPanels(path, matched, tPanel, removals, additions) {
				somethingMoved = true
			}
		}
	}
	return
}
