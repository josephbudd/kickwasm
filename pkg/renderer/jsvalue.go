package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createJSValue(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	classes := builder.Classes
	var fname string
	var oPath string
	data := &struct {
		ApplicationGitPath               string
		ImportRendererFrameworkCallBack  string
		ImportRendererAPIEvent           string
		ImportRendererAPIWindow          string
		ImportRendererFrameworkViewTools string
		HVScrollClassName                string
		ResizeMeWidthClassName           string
		ResizeMeHeightClassName          string
		DoNotPrintClassName              string
	}{
		ApplicationGitPath:               builder.ImportPath,
		ImportRendererFrameworkCallBack:  folderpaths.ImportRendererFrameworkCallBack,
		ImportRendererAPIEvent:           folderpaths.ImportRendererAPIEvent,
		ImportRendererAPIWindow:          folderpaths.ImportRendererAPIWindow,
		ImportRendererFrameworkViewTools: folderpaths.ImportRendererFrameworkViewTools,
		HVScrollClassName:                classes.HVScroll,
		ResizeMeWidthClassName:           classes.ResizeMeWidth,
		ResizeMeHeightClassName:          classes.ResizeMeHeight,
		DoNotPrintClassName:              classes.DoNotPrint,
	}

	// rendererprocess/api/jsvalue/hideshow.go
	fname = fileNames.HideShowDotGo
	oPath = filepath.Join(folderpaths.OutputRendererAPIJSValue, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.JSValueHideShowGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/jsvalue/width.go
	fname = fileNames.WidthDotGo
	oPath = filepath.Join(folderpaths.OutputRendererAPIJSValue, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.JSValueWidthGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/jsvalue/spawnpanel.go
	fname = fileNames.SpawnPanelDotGo
	oPath = filepath.Join(folderpaths.OutputRendererAPIJSValue, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.JSValueSpawnPanel, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/jsvalue/temporary.go
	fname = fileNames.TemporaryDotGo
	oPath = filepath.Join(folderpaths.OutputRendererAPIJSValue, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.JSValueTemporary, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/jsvalue/permanent.go
	fname = fileNames.PermanentDotGo
	oPath = filepath.Join(folderpaths.OutputRendererAPIJSValue, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.JSValuePermanent, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/jsvalue/print.go
	fname = fileNames.PrintDotGo
	oPath = filepath.Join(folderpaths.OutputRendererAPIJSValue, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.JSValuePrint, data, appPaths); err != nil {
		return
	}
	return
}
