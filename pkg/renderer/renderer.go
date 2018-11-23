package renderer

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/tap"
	"golang.org/x/net/html"
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
func Create(appPaths paths.ApplicationPathsI, builder *tap.Builder, addLocations bool, headTemplateFile string) (err error) {
	if err = createHTMLTemplates(appPaths, builder, addLocations, headTemplateFile); err != nil {
		return
	}
	if err = createCSS(appPaths, builder); err != nil {
		return
	}
	if err = createViewTools(appPaths, builder); err != nil {
		return
	}
	if err = createGoPanels(appPaths, builder); err != nil {
		return
	}
	if err = createMainGo(appPaths, builder); err != nil {
		return
	}
	if err = createMainDoPanelsGo(appPaths, builder); err != nil {
		return
	}
	if err = createCallerGo(appPaths, builder); err != nil {
		return
	}
	if err = createCalls(appPaths, builder); err != nil {
		return
	}
	if err = createWASMExecJS(appPaths); err != nil {
		return
	}
	if err = createPanelHelper(appPaths); err != nil {
		return
	}
	if err = createNotJS(appPaths); err != nil {
		return
	}
	return
}

// createHTMLTemplates creates the main html template file.
func createHTMLTemplates(appPaths paths.ApplicationPathsI, builder *tap.Builder, addLocations bool, headTemplateFile string) error {
	folderpaths := appPaths.GetPaths()
	doc := buildIndexHTMLNode(builder, addLocations, headTemplateFile)
	bbuf := &bytes.Buffer{}
	if err := html.Render(bbuf, doc); err != nil {
		return err
	}
	// fix the renderer html
	bb := bytes.Replace(bbuf.Bytes(), []byte(`{{template &#34;`), []byte(`{{template "`), -1)
	bb = bytes.Replace(bb, []byte(`&#34;}}`), []byte(`"}}`), -1)
	fpath := filepath.Join(folderpaths.OutputRendererTemplates, "main.tmpl")
	ofile, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, appPaths.GetFMode())
	if err != nil {
		return err
	}
	if _, err = ofile.Write(bb); err != nil {
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
