package slurp

import "github.com/josephbudd/kickwasm/tap"

// GetPanelFilePaths return the path of every panel file.
// Other than the starting yaml file these are the only other yaml files.
// These are full paths not relative paths.
func (sl *Slurper) GetPanelFilePaths() []string {
	return sl.panelFiles
}

// Gulp reads the application yaml file at path and processes it.
// It constructs a slice of tab.Services and uses them to build the tap.Builder.
// It returns the builder and the error.
func (sl *Slurper) Gulp(yamlPath string) (*tap.Builder, error) {
	appInfo, err := sl.slurpApplication(yamlPath)
	if err != nil {
		return nil, err
	}
	// have all the slurp info from the yaml files.
	// convert this slurp info into tap data.
	builder := tap.NewBuilder()
	builder.Title = appInfo.Title
	builder.ImportPath = appInfo.ImportPath
	builder.Stores = appInfo.Stores
	services := make([]*tap.Service, 0, len(appInfo.Services))
	for _, sinfo := range appInfo.Services {
		service := &tap.Service{
			Name: sinfo.Name,
		}
		binfo := sinfo.Button
		button := &tap.Button{
			ID:       binfo.ID,
			Label:    binfo.Label,
			Heading:  binfo.Heading,
			Location: binfo.CC,
			Panels:   make([]*tap.Panel, 0, 5),
		}
		service.Button = button
		for _, pinfo := range binfo.Panels {
			if err := constructButtonPanel(button, pinfo); err != nil {
				return nil, err
			}
		}
		services = append(services, service)
	}
	if err := builder.BuildFromServices(services); err != nil {
		return nil, err
	}
	return builder, nil
}

func constructButtonPanel(button *tap.Button, pinfo *PanelInfo) error {
	panel := &tap.Panel{
		ID:      pinfo.ID,
		Name:    pinfo.Name,
		Note:    pinfo.Note,
		Markup:  pinfo.Markup,
		MyJS:    pinfo.MyJS,
		Buttons: make([]*tap.Button, 0, 5),
		Tabs:    make([]*tap.Tab, 0, 5),
	}
	for _, binfo := range pinfo.Buttons {
		if err := constructButton(panel, binfo); err != nil {
			return err
		}
	}
	for _, tinfo := range pinfo.Tabs {
		if err := constructTab(panel, tinfo); err != nil {
			return err
		}
	}
	button.Panels = append(button.Panels, panel)
	return nil
}

func constructTabPanel(tab *tap.Tab, pinfo *PanelInfo) error {
	panel := &tap.Panel{
		ID:      pinfo.ID,
		Name:    pinfo.Name,
		Note:    pinfo.Note,
		Markup:  pinfo.Markup,
		MyJS:    pinfo.MyJS,
		Buttons: make([]*tap.Button, 0, 5),
		Tabs:    make([]*tap.Tab, 0, 5),
	}
	for _, binfo := range pinfo.Buttons {
		if err := constructButton(panel, binfo); err != nil {
			return err
		}
	}
	for _, tinfo := range pinfo.Tabs {
		if err := constructTab(panel, tinfo); err != nil {
			return err
		}
	}
	tab.Panels = append(tab.Panels, panel)
	return nil
}

func constructButton(panel *tap.Panel, binfo *ButtonInfo) error {
	button := &tap.Button{
		ID:       binfo.ID,
		Label:    binfo.Label,
		Heading:  binfo.Heading,
		Location: binfo.CC,
		Panels:   make([]*tap.Panel, 0, 5),
	}
	for _, pinfo := range binfo.Panels {
		if err := constructButtonPanel(button, pinfo); err != nil {
			return err
		}
	}
	panel.Buttons = append(panel.Buttons, button)
	return nil
}

func constructTab(panel *tap.Panel, t *TabInfo) error {
	tab := &tap.Tab{
		ID:     t.ID,
		Label:  t.Label,
		Panels: make([]*tap.Panel, 0, 5),
	}
	for _, p := range t.Panels {
		if err := constructTabPanel(tab, p); err != nil {
			return err
		}
	}
	panel.Tabs = append(panel.Tabs, tab)
	return nil
}
