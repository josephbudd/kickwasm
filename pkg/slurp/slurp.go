package slurp

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

func (sl *Slurper) slurpApplication(fpath string) (appInfo *ApplicationInfo, err error) {
	var bb []byte
	if bb, err = getFileBB(fpath); err != nil {
		return
	}
	if appInfo, err = sl.checkApplicationInfo(bb, fpath); err != nil {
		return
	}
	for _, service := range appInfo.Services {
		if err = sl.slurpServiceButtonPanels(fpath, service.Button); err != nil {
			return
		}
	}
	return
}

func (sl *Slurper) slurpServiceButtonPanels(parentFilePath string, button *ButtonInfo) (err error) {
	if button.Panels == nil {
		button.Panels = make([]*PanelInfo, 0, 5)
	}
	for _, panel := range button.Panels {
		if err = sl.checkButtonPanelInfo(panel, parentFilePath, 1); err != nil {
			return
		}
		if err = sl.slurpPanelButtonsTabs(parentFilePath, panel); err != nil {
			return
		}
	}
	dir := path.Dir(parentFilePath)
	for _, fpath := range button.PanelFiles {
		fullpath := filepath.Join(dir, fpath)
		var bb []byte
		bb, err = getFileBB(fullpath)
		if err != nil {
			err = errors.New(err.Error() + " in " + parentFilePath)
			return
		}
		var panel *PanelInfo
		panel, err = sl.checkButtonPanelInfoBB(bb, fullpath, 1)
		if err != nil {
			return
		}
		sl.panelFiles = append(sl.panelFiles, fullpath)
		if err = sl.slurpPanelButtonsTabs(fullpath, panel); err != nil {
			return
		}
		button.Panels = append(button.Panels, panel)
	}
	return
}

func (sl *Slurper) slurpPanelButtonsTabs(parentFilePath string, panel *PanelInfo) (err error) {
	var bb []byte
	var pi *PanelInfo
	nextLevel := panel.Level + 1
	if len(panel.Buttons) > 0 {
		// button panels will be adding a panel level
		if ok := sl.checkLevel(panel.Level); !ok {
			err = fmt.Errorf(`the panel named %q is too deep (level %d) to have buttons in %s`, panel.Name, nextLevel, panel.SourcePath)
			return
		}
	}
	dir := path.Dir(parentFilePath)
	for _, button := range panel.Buttons {
		if button.Panels == nil {
			button.Panels = make([]*PanelInfo, 0, 5)
		}
		for _, p := range button.Panels {
			if err = sl.checkButtonPanelInfo(p, parentFilePath, nextLevel); err != nil {
				return
			}
			if err = sl.slurpPanelButtonsTabs(parentFilePath, p); err != nil {
				return
			}
		}
		for _, fpath := range button.PanelFiles {
			fullpath := filepath.Join(dir, fpath)
			if bb, err = getFileBB(fullpath); err != nil {
				err = errors.New(err.Error() + " in " + parentFilePath)
				return
			}
			if pi, err = sl.checkButtonPanelInfoBB(bb, fullpath, nextLevel); err != nil {
				return
			}
			if err = sl.slurpPanelButtonsTabs(fullpath, pi); err != nil {
				return
			}
			sl.panelFiles = append(sl.panelFiles, fullpath)
			button.Panels = append(button.Panels, pi)
		}
	}
	if len(panel.Tabs) > 0 {
		// button panels will be adding a panel level
		if ok := sl.checkLevel(panel.Level); !ok {
			err = fmt.Errorf(`the panel named %q is too deep (level %d) to have tabs in %s`, panel.Name, nextLevel, panel.SourcePath)
			return
		}
	}
	for _, tab := range panel.Tabs {
		if tab.Panels == nil {
			tab.Panels = make([]*PanelInfo, 0, 5)
		}
		for _, pi = range tab.Panels {
			if err = sl.checkTabPanelInfo(pi, parentFilePath, nextLevel); err != nil {
				return
			}
		}
		for _, fpath := range tab.PanelFiles {
			fullpath := filepath.Join(dir, fpath)
			bb, err = getFileBB(fullpath)
			if err != nil {
				err = errors.New(err.Error() + " in " + parentFilePath)
				return
			}
			panel, err = sl.checkTabPanelInfoBB(bb, fullpath, nextLevel)
			if err != nil {
				return
			}
			sl.panelFiles = append(sl.panelFiles, fullpath)
			tab.Panels = append(tab.Panels, panel)
		}
	}
	return
}

func getFileBB(fpath string) (filebb []byte, err error) {
	var ifile *os.File
	ifile, err = os.Open(fpath)
	if err != nil {
		return
	}
	defer func() {
		// close and check for the error
		if cerr := ifile.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	var stat os.FileInfo
	stat, err = ifile.Stat()
	if err != nil {
		return
	}
	l := int(stat.Size())
	filebb = make([]byte, l, l)
	if _, err = ifile.Read(filebb); err != nil {
		return
	}
	return
}
