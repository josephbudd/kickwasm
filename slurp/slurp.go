package slurp

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

func (sl *Slurper) slurpApplication(fpath string) (*ApplicationInfo, error) {
	bb, err := getFileBB(fpath)
	if err != nil {
		return nil, err
	}
	appInfo, err := sl.checkApplicationInfo(bb, fpath)
	if err != nil {
		return nil, err
	}
	for _, service := range appInfo.Services {
		if err := sl.slurpServiceButtonPanels(fpath, service.Button); err != nil {
			return nil, err
		}
	}
	return appInfo, nil
}

func (sl *Slurper) slurpServiceButtonPanels(parentFilePath string, button *ButtonInfo) error {
	if button.Panels == nil {
		button.Panels = make([]*PanelInfo, 0, 5)
	}
	for _, panel := range button.Panels {
		if err := sl.checkButtonPanelInfo(panel, parentFilePath, 1); err != nil {
			return err
		}
		if err := sl.slurpPanelButtonsTabs(parentFilePath, panel); err != nil {
			return err
		}
	}
	dir := path.Dir(parentFilePath)
	for _, fpath := range button.PanelFiles {
		fullpath := filepath.Join(dir, fpath)
		bb, err := getFileBB(fullpath)
		if err != nil {
			return errors.New(err.Error() + " in " + parentFilePath)
		}
		panel, err := sl.checkButtonPanelInfoBB(bb, fullpath, 1)
		if err != nil {
			return err
		}
		sl.panelFiles = append(sl.panelFiles, fullpath)
		if err := sl.slurpPanelButtonsTabs(fullpath, panel); err != nil {
			return err
		}
		button.Panels = append(button.Panels, panel)
	}
	return nil
}

func (sl *Slurper) slurpPanelButtonsTabs(parentFilePath string, panel *PanelInfo) error {
	nextLevel := panel.Level + 1
	if len(panel.Buttons) > 0 {
		// button panels will be adding a panel level
		if ok := sl.checkLevel(panel.Level); !ok {
			return fmt.Errorf(`the panel named %q is too deep (level %d) to have buttons in %s`, panel.Name, nextLevel, panel.SourcePath)
		}
	}
	dir := path.Dir(parentFilePath)
	for _, button := range panel.Buttons {
		if button.Panels == nil {
			button.Panels = make([]*PanelInfo, 0, 5)
		}
		for _, p := range button.Panels {
			if err := sl.checkButtonPanelInfo(p, parentFilePath, nextLevel); err != nil {
				return err
			}
			if err := sl.slurpPanelButtonsTabs(parentFilePath, p); err != nil {
				return err
			}
		}
		for _, fpath := range button.PanelFiles {
			fullpath := filepath.Join(dir, fpath)
			bb, err := getFileBB(fullpath)
			if err != nil {
				return errors.New(err.Error() + " in " + parentFilePath)
			}
			panel, err := sl.checkButtonPanelInfoBB(bb, fullpath, nextLevel)
			if err != nil {
				return err
			}
			if err := sl.slurpPanelButtonsTabs(fullpath, panel); err != nil {
				return err
			}
			sl.panelFiles = append(sl.panelFiles, fullpath)
			button.Panels = append(button.Panels, panel)
		}
	}
	if len(panel.Tabs) > 0 {
		// button panels will be adding a panel level
		if ok := sl.checkLevel(panel.Level); !ok {
			return fmt.Errorf(`the panel named %q is too deep (level %d) to have tabs in %s`, panel.Name, nextLevel, panel.SourcePath)
		}
	}
	for _, tab := range panel.Tabs {
		if tab.Panels == nil {
			tab.Panels = make([]*PanelInfo, 0, 5)
		}
		for _, p := range tab.Panels {
			if err := sl.checkTabPanelInfo(p, parentFilePath, nextLevel); err != nil {
				return err
			}
		}
		for _, fpath := range tab.PanelFiles {
			fullpath := filepath.Join(dir, fpath)
			bb, err := getFileBB(fullpath)
			if err != nil {
				return errors.New(err.Error() + " in " + parentFilePath)
			}
			panel, err := sl.checkTabPanelInfoBB(bb, fullpath, nextLevel)
			if err != nil {
				return err
			}
			sl.panelFiles = append(sl.panelFiles, fullpath)
			tab.Panels = append(tab.Panels, panel)
		}
	}
	return nil
}

func getFileBB(fpath string) ([]byte, error) {
	ifile, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer func() {
		// close and check for the error
		if cerr := ifile.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	stat, err := ifile.Stat()
	if err != nil {
		return nil, err
	}
	l := int(stat.Size())
	ifilebb := make([]byte, l, l)
	if _, err := ifile.Read(ifilebb); err != nil {
		return nil, err
	}
	return ifilebb, nil
}
