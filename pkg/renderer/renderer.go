package renderer

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"golang.org/x/net/html"
)

const (
	mainMasterViewID = "mainMasterView"
	initialIndent    = uint(2)
	favicon          = "favicon.ico"
)

// GetTabsMasterViewID returns the tabs master view id needed for most ids.
func GetTabsMasterViewID() string {
	return mainMasterViewID
}

// GetInitialIndent returns the initial indentaion for the main html template.
func GetInitialIndent() uint {
	return initialIndent
}

// Create creates all of the renderer files.
func Create(appPaths paths.ApplicationPathsI, builder *project.Builder, addLocations bool) (err error) {
	parts := strings.Split(builder.ImportPath, "/")
	appname := parts[len(parts)-1]
	if err = createHTMLTemplates(appPaths, builder, addLocations); err != nil {
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
	if err = createFrameworkDoPanelsGo(appPaths, builder); err != nil {
		return
	}
	if err = createWASMExecJS(appPaths); err != nil {
		return
	}
	if err = createPaneling(appPaths); err != nil {
		return
	}
	// Spawn Tabs.
	if err = createSpawnPack(appPaths); err != nil {
		return
	}
	if err = createSpawnTabHTMLTemplates(appPaths, builder); err != nil {
		return
	}
	if err = createSpawnTabBarFiles(appPaths, builder); err != nil {
		return
	}
	if err = createSpawnTabFiles(appPaths, builder); err != nil {
		return
	}
	if err = createTabSpawnPanels(appPaths, builder); err != nil {
		return
	}
	// LPC
	if err = createLPC(appPaths, builder); err != nil {
		return
	}
	// VSCode
	if err = createVSCode(appPaths, appname); err != nil {
		return
	}
	if err = createCallBack(appPaths, builder); err != nil {
		return
	}
	if err = createDOM(appPaths, builder); err != nil {
		return
	}
	// v 14
	if err = createLocation(appPaths, builder); err != nil {
		return
	}
	if err = createMarkup(appPaths, builder); err != nil {
		return
	}
	if err = createWindow(appPaths); err != nil {
		return
	}
	if err = creatEvent(appPaths); err != nil {
		return
	}
	if err = createDisplay(appPaths, builder); err != nil {
		return
	}
	if err = createProof(appPaths, builder); err != nil {
		return
	}
	if err = createApplication(appPaths, builder); err != nil {
		return
	}
	// v 16
	if err = createJSValue(appPaths, builder); err != nil {
		return
	}

	return
}

// createHTMLTemplates creates the main html template file.
func createHTMLTemplates(appPaths paths.ApplicationPathsI, builder *project.Builder, addLocations bool) (err error) {
	folderpaths := appPaths.GetPaths()
	doc := buildIndexHTMLNode(appPaths, builder, addLocations)
	bbuf := &bytes.Buffer{}
	if err = html.Render(bbuf, doc); err != nil {
		return
	}
	// fix the renderer html
	fileNames := paths.GetFileNames()
	bb := bytes.Replace(bbuf.Bytes(), []byte(`{{template &#34;`), []byte(`{{template "`), -1)
	bb = bytes.Replace(bb, []byte(`&#34;}}`), []byte(`"}}`), -1)
	fpath := filepath.Join(folderpaths.OutputRendererTemplates, fileNames.MainDotTMPL)
	var ofile *os.File
	ofile, err = os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, appPaths.GetFMode())
	if err != nil {
		return
	}
	if _, err = ofile.Write(bb); err != nil {
		return
	}
	if err = ofile.Close(); err != nil {
		return
	}
	homePanelNamePathMap := builder.GenerateHomeEmptyInsidePanelNamePathMap()
	homePanelMap := builder.GenerateHomePanelNameTemplateMap()
	for home, nameMarkup := range homePanelMap {
		panelNamePathMap := homePanelNamePathMap[home]
		for name, markup := range nameMarkup {
			folders := strings.Join(panelNamePathMap[name], string(os.PathSeparator))
			folderPath := filepath.Join(folderpaths.OutputRendererTemplates, folders)
			if err = os.MkdirAll(folderPath, appPaths.GetDMode()); err != nil {
				return
			}
			fname := fmt.Sprintf("%s.tmpl", name)
			fpath := filepath.Join(folderPath, fname)
			ofile, err = os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, appPaths.GetFMode())
			if err != nil {
				return
			}
			_, err = ofile.Write([]byte(markup))
			if err = ofile.Close(); err != nil {
				return
			}
		}
	}
	return
}

// createSpawnTabHTMLTemplates creates the spawn tab html template file.
func createSpawnTabHTMLTemplates(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	markupPath := builder.GenerateSpawnTabMarkupPanelPathMap()
	for markup, shortPath := range markupPath {
		dir := filepath.Dir(shortPath)
		folderPath := filepath.Join(folderpaths.OutputRendererSpawnTemplates, dir)
		if err = os.MkdirAll(folderPath, appPaths.GetDMode()); err != nil {
			return
		}
		fname := filepath.Base(shortPath) + ".tmpl"
		fpath := filepath.Join(folderPath, fname)
		var ofile *os.File
		if ofile, err = os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, appPaths.GetFMode()); err != nil {
			return
		}
		if _, err = ofile.Write([]byte(markup)); err != nil {
			ofile.Close()
			return
		}
		if err = ofile.Close(); err != nil {
			return
		}
	}
	return
}
