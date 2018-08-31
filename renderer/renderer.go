package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/tap"
)

const (
	tabsMasterViewID = "tabsMasterView"
	initialIndent    = uint(2)
	favicon          = "favicon.ico"
)

// GetTabsMasterViewID returns the tabs master view id needed for most ids.
func GetTabsMasterViewID() string {
	return tabsMasterViewID
}

// GetInitialIndent returns the initial indentaion for the main html template.
func GetInitialIndent() uint {
	return initialIndent
}

// Create creates all of the renderer files.
func Create(appPaths paths.ApplicationPathsI, builder *tap.Builder, addAbout, addLocations bool, headTemplateFile string, host string, port uint) error {
	if err := createHTMLTemplates(appPaths, builder, addAbout, addLocations, headTemplateFile); err != nil {
		return err
	}
	if err := createCSS(appPaths, builder); err != nil {
		return err
	}
	if err := createViewTools(appPaths, builder); err != nil {
		return err
	}
	if err := createGoPanels(appPaths, builder); err != nil {
		return err
	}
	if addAbout {
		if err := createAboutFiles(appPaths, builder); err != nil {
			return err
		}
	}
	if err := createMainGo(host, port, appPaths, builder); err != nil {
		return err
	}
	if err := createMainDoPanelsGo(addAbout, appPaths, builder); err != nil {
		return err
	}
	if err := createCallerGo(appPaths, builder); err != nil {
		return err
	}
	if err := createWASMExecJS(appPaths); err != nil {
		return err
	}
	return nil
}

// createHTMLTemplates creates the main html template file.
func createHTMLTemplates(appPaths paths.ApplicationPathsI, builder *tap.Builder, addAbout, addLocations bool, headTemplateFile string) error {
	if err := appPaths.CreateTemplateFolder(); err != nil {
		return err
	}
	folderpaths := appPaths.GetPaths()
	if addAbout {
		builder.AddAbout()
	}
	contents, err := buildIndexHTML(builder, addLocations, headTemplateFile)
	if err != nil {
		return err
	}
	fpath := filepath.Join(folderpaths.OutputRendererTemplates, "main.tmpl")
	ofile, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, appPaths.GetFMode())
	if err != nil {
		return err
	}
	if _, err = ofile.Write(contents); err != nil {
		return err
	}
	if err := ofile.Close(); err != nil {
		return err
	}
	servicePanelNamePathMap := builder.GenerateServiceEmptyInsidePanelNamePathMap()
	servicePanelMap := builder.GenerateServicePanelNameTemplateMap()
	for service, nameMarkup := range servicePanelMap {
		panelNamePathMap := servicePanelNamePathMap[service]
		for name, markup := range nameMarkup {
			folders := strings.Join(panelNamePathMap[name], string(os.PathSeparator))
			folderPath := filepath.Join(folderpaths.OutputRendererTemplates, folders)
			if err := os.MkdirAll(folderPath, appPaths.GetDMode()); err != nil {
				return err
			}
			fname := fmt.Sprintf("%s.tmpl", name)
			fpath := filepath.Join(folderPath, fname)
			ofile, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, appPaths.GetFMode())
			if err != nil {
				return err
			}
			_, err = ofile.Write([]byte(markup))
			if err := ofile.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}
