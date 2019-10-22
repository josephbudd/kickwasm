package slurp

import (
	"strings"

	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

const (
	forwardSlash = "/"
)

// GetApplicationInfo only reads the application info.
func GetApplicationInfo(fpath string) (appInfo *ApplicationInfo, err error) {
	var bb []byte
	if bb, err = getFileBB(fpath); err != nil {
		return
	}
	appInfo = &ApplicationInfo{}
	if err = yaml.Unmarshal(bb, appInfo); err != nil {
		err = errors.New(err.Error() + " in " + fpath)
		return
	}
	appInfo.SourcePath = fpath
	return
}

// GetPanelFilePaths return the path of every panel file.
// Other than the starting yaml file these are the only other yaml files.
// These are full paths not relative paths.
func (sl *Slurper) GetPanelFilePaths() []string {
	return sl.panelFiles
}

// Gulp reads the application yaml file at path and processes it.
// It constructs a slice of tab.Homes and uses them to build the project.Builder.
// It returns the builder and the error.
func (sl *Slurper) Gulp(yamlPath string) (builder *project.Builder, err error) {
	var appInfo *ApplicationInfo
	if appInfo, err = sl.slurpApplication(yamlPath); err != nil {
		return
	}
	// have all the slurp info from the yaml files.
	// convert this slurp info into project data.
	builder = project.NewBuilder()
	builder.Title = appInfo.Title
	builder.ImportPath = appInfo.ImportPath
	i := strings.LastIndex(appInfo.ImportPath, forwardSlash)
	builder.SitePackPackage = appInfo.ImportPath[i+1:] + "sitepack"
	builder.SitePackImportPath = appInfo.ImportPath[:i] + forwardSlash + builder.SitePackPackage
	homes := make([]*project.Button, 0, len(appInfo.Homes))
	for _, hbinfo := range appInfo.Homes {
		homeButton := &project.Button{
			ID:       hbinfo.ID,
			Label:    hbinfo.Label,
			Heading:  hbinfo.Heading,
			Location: hbinfo.CC,
			Panels:   make([]*project.Panel, 0, 5),
		}
		for _, pinfo := range hbinfo.Panels {
			if err = constructButtonPanel(homeButton, pinfo); err != nil {
				return
			}
		}
		homes = append(homes, homeButton)
	}
	if err = builder.BuildFromHomes(homes); err != nil {
		return
	}
	return
}

func constructButtonPanel(button *project.Button, pinfo *PanelInfo) (err error) {
	panel := &project.Panel{
		ID:       pinfo.ID,
		Name:     pinfo.Name,
		Note:     pinfo.Note,
		Markup:   pinfo.Markup,
		HVScroll: pinfo.HVScroll,
		Buttons:  make([]*project.Button, 0, 5),
		Tabs:     make([]*project.Tab, 0, 5),
	}
	for _, binfo := range pinfo.Buttons {
		if err = constructButton(panel, binfo); err != nil {
			return
		}
	}
	for _, tinfo := range pinfo.Tabs {
		if err = constructTab(panel, tinfo); err != nil {
			return
		}
	}
	button.Panels = append(button.Panels, panel)
	return
}

func constructTabPanel(tab *project.Tab, pinfo *PanelInfo) (err error) {
	panel := &project.Panel{
		ID:       pinfo.ID,
		Name:     pinfo.Name,
		Note:     pinfo.Note,
		Markup:   pinfo.Markup,
		HVScroll: pinfo.HVScroll,
		Buttons:  make([]*project.Button, 0, 5),
		Tabs:     make([]*project.Tab, 0, 5),
	}
	for _, binfo := range pinfo.Buttons {
		if err = constructButton(panel, binfo); err != nil {
			return
		}
	}
	for _, tinfo := range pinfo.Tabs {
		if err = constructTab(panel, tinfo); err != nil {
			return
		}
	}
	tab.Panels = append(tab.Panels, panel)
	return
}

func constructButton(panel *project.Panel, binfo *ButtonInfo) (err error) {
	button := &project.Button{
		ID:       binfo.ID,
		Label:    binfo.Label,
		Heading:  binfo.Heading,
		Location: binfo.CC,
		Panels:   make([]*project.Panel, 0, 5),
	}
	for _, pinfo := range binfo.Panels {
		if err = constructButtonPanel(button, pinfo); err != nil {
			return
		}
	}
	panel.Buttons = append(panel.Buttons, button)
	return
}

func constructTab(panel *project.Panel, t *TabInfo) (err error) {
	tab := &project.Tab{
		ID:      t.ID,
		Label:   t.Label,
		Heading: t.Heading,
		Spawn:   t.Spawn,
		Panels:  make([]*project.Panel, 0, 5),
	}
	for _, p := range t.Panels {
		if err = constructTabPanel(tab, p); err != nil {
			return
		}
	}
	panel.Tabs = append(panel.Tabs, tab)
	return
}
