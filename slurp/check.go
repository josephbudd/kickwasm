package slurp

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

func (sl *Slurper) checkLevel(level int) bool {
	if level > sl.maxLevel {
		return false
	}
	if level > sl.CurrentLevel {
		sl.CurrentLevel = level
	}
	return true
}

func (sl *Slurper) checkPanelName(name, sourcePath string) (string, bool) {
	if sPath, ok := sl.panelNames[name]; ok {
		return sPath, false
	}
	sl.panelNames[name] = sourcePath
	return "", true
}

func (sl *Slurper) checkApplicationInfo(yamlbb []byte, fpath string) (*ApplicationInfo, error) {
	appInfo := &ApplicationInfo{}
	if err := yaml.Unmarshal(yamlbb, appInfo); err != nil {
		return nil, errors.New(err.Error() + " in " + fpath)
	}
	appInfo.SourcePath = fpath
	// make sure there is a title.
	if len(appInfo.Title) == 0 {
		return nil, errors.New("title is required in " + fpath)
	}
	// make sure there is an import path.
	if len(appInfo.ImportPath) == 0 {
		return nil, errors.New("importPath is required in " + fpath)
	}
	// make sure there are resos
	if len(appInfo.Stores) == 0 {
		return nil, errors.New("stores is missing in " + fpath)
	}
	// make sure there are services
	if len(appInfo.Services) == 0 {
		return nil, errors.New("services is missing in " + fpath)
	}
	for _, service := range appInfo.Services {
		service.SourcePath = fpath
		if len(service.Name) == 0 {
			return nil, errors.New("a service is missing a name")
		}
		if service.Button == nil {
			return nil, fmt.Errorf(`the service named %q is missing a button`, service.Name)
		}
		service.Button.SourcePath = fpath
		if err := sl.checkButtonInfo(service.Button); err != nil {
			return nil, fmt.Errorf(`in the service named %q, %s`, service.Name, err.Error())
		}
	}
	return appInfo, nil
}

func (sl *Slurper) checkButtonPanelInfoBB(yamlbb []byte, fpath string, level int) (*PanelInfo, error) {
	panel := &PanelInfo{}
	if err := yaml.Unmarshal(yamlbb, panel); err != nil {
		return nil, errors.New(err.Error() + " in " + fpath)
	}
	if err := sl.checkButtonPanelInfo(panel, fpath, level); err != nil {
		return nil, err
	}
	return panel, nil
}

func (sl *Slurper) checkButtonPanelInfo(panel *PanelInfo, fpath string, level int) error {
	panel.SourcePath = fpath
	if !strings.HasSuffix(panel.Name, "Panel") {
		return fmt.Errorf(`the button panel name %q should end with the suffix "Panel" in %s`, panel.Name, fpath)
	}
	cc := cases.CamelCase(panel.Name)
	if panel.Name != cc {
		return fmt.Errorf(`the button panel name %q is not camel cased. It should be %q in %s`, panel.Name, cc, fpath)
	}
	if otherPath, ok := sl.checkPanelName(panel.Name, fpath); !ok {
		return fmt.Errorf(`the button panel name %q used in %s has already been used in %s`, panel.Name, fpath, otherPath)
	}
	if len(panel.Note) == 0 && len(panel.Markup) > 0 {
		return fmt.Errorf(`the button panel named %q is missing a note for the markup in %s`, panel.Name, fpath)
	}
	if len(panel.Markup) > 0 {
		if len(panel.Buttons) > 0 || len(panel.Tabs) > 0 {
			return fmt.Errorf(`the button panel named %q must not have a combination of markup, buttons and tabs in %s`, panel.Name, fpath)
		}
	}
	if len(panel.Buttons) > 0 && len(panel.Tabs) > 0 {
		return fmt.Errorf(`the button panel named %q must not have a combination of markup, buttons and tabs in %s`, panel.Name, fpath)
	}
	if len(panel.Markup) == 0 && len(panel.Buttons) == 0 && len(panel.Tabs) == 0 {
		return fmt.Errorf(`the button panel named %q must have markup, buttons or tabs in %s`, panel.Name, fpath)
	}
	panel.Level = level
	for _, b := range panel.Buttons {
		b.SourcePath = fpath
		if err := sl.checkButtonInfo(b); err != nil {
			return err
		}
		// make sure that the button id (yaml name) is unique for this panel.
		buttonIDs, found := sl.buttonIDs[panel.Name]
		if found {
			for _, id := range buttonIDs {
				if id == b.ID {
					return fmt.Errorf("the button panel named %q has more then one button named %q", panel.Name, b.ID)
				}
			}
		} else {
			sl.buttonIDs[panel.Name] = make([]string, 0, 5)
		}
		// add this button id
		sl.buttonIDs[panel.Name] = append(sl.buttonIDs[panel.Name], b.ID)
		return nil
	}
	for _, t := range panel.Tabs {
		t.SourcePath = fpath
		if err := sl.checkTabInfo(t); err != nil {
			return err
		}
		// make sure that the tab id (yaml name) is unique for this panel.
		tabIDs, found := sl.tabIDs[panel.Name]
		if found {
			for _, id := range tabIDs {
				if id == t.ID {
					return fmt.Errorf("the tab panel named %q has more then one tab named %q", panel.Name, t.ID)
				}
			}
		} else {
			sl.tabIDs[panel.Name] = make([]string, 0, 5)
		}
		// add this tab id
		sl.tabIDs[panel.Name] = append(sl.tabIDs[panel.Name], t.ID)
		return nil
	}
	return nil
}

func (sl *Slurper) checkButtonInfo(button *ButtonInfo) error {
	if len(button.ID) == 0 {
		return fmt.Errorf(`a button is missing a name in %s`, button.SourcePath)
	}
	if len(button.Label) == 0 {
		return fmt.Errorf(`a button named %q is missing a label in %s`, button.ID, button.SourcePath)
	}
	if len(button.PanelFiles) == 0 && len(button.Panels) == 0 {
		return fmt.Errorf(`a button labeled %q is missing panel files in %s`, button.Label, button.SourcePath)
	}
	if len(button.Heading) == 0 {
		button.Heading = button.Label
	}
	if len(button.CC) == 0 {
		button.CC = button.Label
	}
	return nil
}

func (sl *Slurper) checkTabInfo(tab *TabInfo) error {
	if len(tab.ID) == 0 {
		return fmt.Errorf(`a tab is missing a name in %s`, tab.SourcePath)
	}
	if len(tab.Label) == 0 {
		return fmt.Errorf(`a tab named %q is missing a label in %s`, tab.ID, tab.SourcePath)
	}
	if len(tab.PanelFiles) == 0 && len(tab.Panels) == 0 {
		return fmt.Errorf(`a tab labeled %q is missing panel files in %s`, tab.Label, tab.SourcePath)
	}
	return nil
}

func (sl *Slurper) checkTabPanelInfoBB(yamlbb []byte, fpath string, level int) (*PanelInfo, error) {
	panel := &PanelInfo{}
	if err := yaml.Unmarshal(yamlbb, panel); err != nil {
		return nil, errors.New(err.Error() + " in " + fpath)
	}
	if err := sl.checkTabPanelInfo(panel, fpath, level); err != nil {
		return nil, err
	}
	return panel, nil
}

func (sl *Slurper) checkTabPanelInfo(panel *PanelInfo, fpath string, level int) error {
	panel.SourcePath = fpath
	if !strings.HasSuffix(panel.Name, "Panel") {
		return fmt.Errorf(`the tab panel name %q should end with the suffix "Panel" in %s`, panel.Name, fpath)
	}
	cc := cases.CamelCase(panel.Name)
	if panel.Name != cc {
		return fmt.Errorf(`the tab panel name %q is not camel cased. It should be %q in %s`, panel.Name, cc, fpath)
	}
	if otherPath, ok := sl.checkPanelName(panel.Name, fpath); !ok {
		return fmt.Errorf(`the tab panel name %q used in %s has already been used in %s`, panel.Name, fpath, otherPath)
	}
	if len(panel.Note) == 0 {
		return fmt.Errorf(`the tab panel named %q is missing a note in %s`, panel.Name, fpath)
	}
	if len(panel.Markup) == 0 {
		return fmt.Errorf(`the tab panel named %q must have markup in %s`, panel.Name, fpath)
	}
	if len(panel.Buttons) > 0 {
		return fmt.Errorf(`the tab panel named %q must not have buttons in %s`, panel.Name, fpath)
	}
	if len(panel.Tabs) > 0 {
		return fmt.Errorf(`the tab panel named %q must not have tabs in %s`, panel.Name, fpath)
	}
	panel.Level = level
	return nil
}
