package slurp

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
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

func (sl *Slurper) checkTabName(name, sourcePath string) (string, bool) {
	if sPath, ok := sl.tabNames[name]; ok {
		return sPath, false
	}
	sl.tabNames[name] = sourcePath
	return "", true
}

func (sl *Slurper) checkApplicationInfo(yamlbb []byte, fpath string) (appInfo *ApplicationInfo, err error) {
	var errMessage string
	appInfo = &ApplicationInfo{}
	if err = yaml.Unmarshal(yamlbb, appInfo); err != nil {
		err = errors.New(err.Error() + " in " + fpath)
		return
	}
	appInfo.SourcePath = fpath
	// make sure there is a title.
	if len(appInfo.Title) == 0 {
		err = errors.New("title is required in " + fpath)
		return
	}
	// make sure there is an import path.
	if len(appInfo.ImportPath) == 0 {
		err = errors.New("importPath is required in " + fpath)
		return
	}
	// make sure the application name is correct.
	parts := strings.Split(appInfo.ImportPath, "/")
	appName := parts[len(parts)-1]
	appNameLower := strings.ToLower(appName)
	if appNameLower != appName {
		message := fmt.Sprintf("application folder in importPath should be %q not %q", appNameLower, appName)
		err = errors.New(message)
		return
	}

	// make sure there are homes
	if len(appInfo.Homes) == 0 {
		err = errors.New("homes is missing in " + fpath)
		return
	}
	homeButtonMap := make(map[string]string)
	for _, homeButton := range appInfo.Homes {
		// home name
		homeButton.SourcePath = fpath
		if len(homeButton.ID) == 0 {
			err = errors.New("a home button is missing a name")
			return
		}
		if _, found := homeButtonMap[homeButton.ID]; found {
			errMessage = fmt.Sprintf(`the home button name %q is used more than once`, homeButton.ID)
			err = errors.New(errMessage)
			return
		}
		for _, bn := range homeButtonMap {
			if bn == homeButton.ID {
				errMessage = fmt.Sprintf(`the home button name %q is used more than once`, homeButton.ID)
				err = errors.New(errMessage)
				return
			}
		}
		homeButtonMap[homeButton.ID] = homeButton.ID
		if err = sl.checkButtonInfo(homeButton); err != nil {
			errMessage = fmt.Sprintf(`in the home button named %q, %s`, homeButton.ID, err.Error())
			err = errors.New(errMessage)
			return
		}
	}
	return
}

func (sl *Slurper) checkButtonPanelInfoBB(yamlbb []byte, fpath string, level int) (panel *PanelInfo, err error) {
	panel = &PanelInfo{}
	if err = yaml.Unmarshal(yamlbb, panel); err != nil {
		err = errors.New(err.Error() + " in " + fpath)
		return
	}
	if err = sl.checkButtonPanelInfo(panel, fpath, level); err != nil {
		return
	}
	return
}

func (sl *Slurper) checkButtonPanelInfo(panel *PanelInfo, fpath string, level int) (err error) {
	var errMessage string
	panel.SourcePath = fpath
	if !strings.HasSuffix(panel.Name, "Panel") {
		errMessage = fmt.Sprintf(`the button panel name %q should end with the suffix "Panel" in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	cc := cases.CamelCase(panel.Name)
	if panel.Name != cc {
		errMessage = fmt.Sprintf(`the button panel name %q is not camel cased. It should be %q in %s`, panel.Name, cc, fpath)
		err = errors.New(errMessage)
		return
	}
	if otherPath, ok := sl.checkPanelName(panel.Name, fpath); !ok {
		errMessage = fmt.Sprintf(`the button panel name %q used in %s has already been used in %s`, panel.Name, fpath, otherPath)
		err = errors.New(errMessage)
		return
	}
	if len(panel.Note) == 0 && len(panel.Markup) > 0 {
		errMessage = fmt.Sprintf(`the button panel named %q is missing a note for the markup in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	if len(panel.Markup) > 0 {
		if len(panel.Buttons) > 0 || len(panel.Tabs) > 0 {
			errMessage = fmt.Sprintf(`the button panel named %q must not have a combination of markup, buttons and tabs in %s`, panel.Name, fpath)
			err = errors.New(errMessage)
			return
		}
	}
	if len(panel.Buttons) > 0 && len(panel.Tabs) > 0 {
		errMessage = fmt.Sprintf(`the button panel named %q must not have a combination of markup, buttons and tabs in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	if len(panel.Markup) == 0 && len(panel.Note) == 0 && len(panel.Buttons) == 0 && len(panel.Tabs) == 0 {
		errMessage = fmt.Sprintf(`the button panel named %q must have markup, buttons or tabs in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	panel.Level = level
	if len(panel.Buttons) > 0 {
		for _, b := range panel.Buttons {
			b.SourcePath = fpath
			if err = sl.checkButtonInfo(b); err != nil {
				return
			}
			// make sure that the button id (yaml name) is unique.
			if otherfpath, found := sl.buttonNames[b.ID]; found {
				errMessage = fmt.Sprintf("the button name %q used in %s has already been used in %s", b.ID, b.SourcePath, otherfpath)
				err = errors.New(errMessage)
				return
			}
			// this is a new button id ( yaml name )
			sl.buttonNames[b.ID] = b.SourcePath
			// make sure that the button id (yaml name) is unique for this panel.
			_, found := sl.buttonIDs[panel.Name]
			if !found {
				sl.buttonIDs[panel.Name] = make([]string, 0, 5)
			}
			// add this button id
			sl.buttonIDs[panel.Name] = append(sl.buttonIDs[panel.Name], b.ID)
		}
		return
	}
	l := len(panel.Tabs)
	if l > 0 {
		for _, t := range panel.Tabs {
			t.SourcePath = fpath
			if err = sl.checkTabInfo(t); err != nil {
				return
			}
			// make sure that the tab id (yaml name) is unique.
			if otherfpath, found := sl.tabNames[t.ID]; found {
				errMessage = fmt.Sprintf("the tab name %q used in %s has already been used in %s", t.ID, t.SourcePath, otherfpath)
				err = errors.New(errMessage)
				return
			}
			// this is a new tab id ( yaml name )
			sl.tabNames[t.ID] = t.SourcePath
			// make sure that the tab id (yaml name) is unique for this panel.
			tabIDs, found := sl.tabIDs[panel.Name]
			if found {
				for _, id := range tabIDs {
					if id == t.ID {
						errMessage = fmt.Sprintf("the tab panel named %q has more then one tab named %q", panel.Name, t.ID)
						err = errors.New(errMessage)
						return
					}
				}
			} else {
				sl.tabIDs[panel.Name] = make([]string, 0, 5)
			}
			// add this tab id
			sl.tabIDs[panel.Name] = append(sl.tabIDs[panel.Name], t.ID)
			if t.Spawn {
				l--
			}
		}
		if l == 0 {
			errMessage = fmt.Sprintf("the tab panel named named %q must have at least one real ( not spawned ) tab in %s", panel.Name, fpath)
			err = errors.New(errMessage)
		}
	}
	return
}

func (sl *Slurper) checkButtonInfo(button *ButtonInfo) (err error) {
	var errMessage string
	if len(button.ID) == 0 {
		errMessage = fmt.Sprintf(`a button is missing a name in %s`, button.SourcePath)
		err = errors.New(errMessage)
		return
	}
	if len(button.Label) == 0 {
		errMessage = fmt.Sprintf(`a button named %q is missing a label in %s`, button.ID, button.SourcePath)
		err = errors.New(errMessage)
		return
	}
	if len(button.PanelFiles) == 0 && len(button.Panels) == 0 {
		errMessage = fmt.Sprintf(`a button labeled %q is missing panel files in %s`, button.Label, button.SourcePath)
		err = errors.New(errMessage)
		return
	}
	if len(button.Heading) == 0 {
		button.Heading = button.Label
	}
	if len(button.CC) == 0 {
		button.CC = button.Label
	}
	return
}

func (sl *Slurper) checkTabInfo(tab *TabInfo) (err error) {
	var errMessage string
	if len(tab.ID) == 0 {
		errMessage = fmt.Sprintf(`a tab is missing a name in %s`, tab.SourcePath)
		err = errors.New(errMessage)
		return
	}
	if len(tab.Label) == 0 {
		errMessage = fmt.Sprintf(`a tab named %q is missing a label in %s`, tab.ID, tab.SourcePath)
		err = errors.New(errMessage)
		return
	}
	if len(tab.PanelFiles) == 0 && len(tab.Panels) == 0 {
		errMessage = fmt.Sprintf(`a tab labeled %q is missing panel files in %s`, tab.Label, tab.SourcePath)
		err = errors.New(errMessage)
		return
	}
	if len(tab.Heading) == 0 {
		tab.Heading = tab.Label
	}
	return nil
}

func (sl *Slurper) checkTabPanelInfoBB(yamlbb []byte, fpath string, level int) (panel *PanelInfo, err error) {
	panel = &PanelInfo{}
	if err = yaml.Unmarshal(yamlbb, panel); err != nil {
		err = errors.New(err.Error() + " in " + fpath)
		return
	}
	if err = sl.checkTabPanelInfo(panel, fpath, level); err != nil {
		return
	}
	return
}

func (sl *Slurper) checkTabPanelInfo(panel *PanelInfo, fpath string, level int) (err error) {
	var errMessage string
	panel.SourcePath = fpath
	if !strings.HasSuffix(panel.Name, "Panel") {
		errMessage = fmt.Sprintf(`the tab panel name %q should end with the suffix "Panel" in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	cc := cases.CamelCase(panel.Name)
	if panel.Name != cc {
		errMessage = fmt.Sprintf(`the tab panel name %q is not camel cased. It should be %q in %s`, panel.Name, cc, fpath)
		err = errors.New(errMessage)
		return
	}
	if otherPath, ok := sl.checkPanelName(panel.Name, fpath); !ok {
		errMessage = fmt.Sprintf(`the tab panel name %q used in %s has already been used in %s`, panel.Name, fpath, otherPath)
		err = errors.New(errMessage)
		return
	}
	if len(panel.Note) == 0 {
		errMessage = fmt.Sprintf(`the tab panel named %q is missing a note in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	if len(panel.Markup) == 0 {
		errMessage = fmt.Sprintf(`the tab panel named %q must have markup in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	if len(panel.Buttons) > 0 {
		errMessage = fmt.Sprintf(`the tab panel named %q must not have buttons in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	if len(panel.Tabs) > 0 {
		errMessage = fmt.Sprintf(`the tab panel named %q must not have tabs in %s`, panel.Name, fpath)
		err = errors.New(errMessage)
		return
	}
	panel.Level = level
	return
}
