package renderer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

const aboutTnameFormat = "about-%s"

// createAboutFiles adds the about files for the generated about section.
func createAboutFiles(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	ids := builder.AboutIDs
	folderpaths := appPaths.GetPaths()
	data := struct {
		ApplicationGitPath                 string
		ButtonName                         string
		ButtonID                           string
		ButtonPanelID                      string
		DefaultPanelID                     string
		ReleasesTabPanelInnerID            string
		ContributorsTabPanelInnerID        string
		CreditsTabPanelInnerID             string
		LicensesTabPanelInnerID            string
		ImportDomainImplementationsCalling string
		ImportDomainTypes                  string
		ImportMainProcessServicesAbout     string
		ImportRendererViewTools            string
	}{
		ApplicationGitPath:                 builder.ImportPath,
		ButtonName:                         ids.ButtonName,
		ButtonID:                           ids.ButtonID,
		DefaultPanelID:                     ids.DefaultPanelID,
		ReleasesTabPanelInnerID:            ids.ReleasesTabPanelInnerID,
		ContributorsTabPanelInnerID:        ids.ContributorsTabPanelInnerID,
		CreditsTabPanelInnerID:             ids.CreditsTabPanelInnerID,
		LicensesTabPanelInnerID:            ids.LicensesTabPanelInnerID,
		ImportDomainImplementationsCalling: folderpaths.ImportDomainImplementationsCalling,
		ImportDomainTypes:                  folderpaths.ImportDomainTypes,
		ImportMainProcessServicesAbout:     folderpaths.ImportMainProcessServicesAbout,
		ImportRendererViewTools:            folderpaths.ImportRendererViewTools,
	}
	dirpath := filepath.Join(folderpaths.OutputRendererPanels, "AboutButton", "AboutPanel")
	if err := os.MkdirAll(dirpath, appPaths.GetDMode()); err != nil {
		return err
	}
	fname := "panel.go"
	tname := fmt.Sprintf(aboutTnameFormat, fname)
	oPath := filepath.Join(dirpath, fname)
	if err := templates.ProcessTemplate(tname, oPath, templates.AboutPanel, data, appPaths); err != nil {
		return err
	}
	fname = "presenter.go"
	tname = fmt.Sprintf(aboutTnameFormat, fname)
	oPath = filepath.Join(dirpath, fname)
	if err := templates.ProcessTemplate(tname, oPath, templates.AboutPanelPresenter, data, appPaths); err != nil {
		return err
	}
	fname = "caller.go"
	tname = fmt.Sprintf(aboutTnameFormat, fname)
	oPath = filepath.Join(dirpath, fname)
	if err := templates.ProcessTemplate(tname, oPath, templates.AboutPanelCaller, data, appPaths); err != nil {
		return err
	}
	// css/about.css
	fname = "about.css"
	oPath = filepath.Join(folderpaths.OutputRendererCSS, fname)
	return appPaths.WriteFile(oPath, []byte(templates.AboutCSS))
}
