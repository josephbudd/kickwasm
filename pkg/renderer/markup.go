package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createMarkup(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
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

	// rendererprocess/api/markup/attributes.go
	fname = fileNames.AttributesDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupAttributeGo)); err != nil {
		return
	}
	// rendererprocess/api/markup/checked.go
	fname = fileNames.CheckedDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupCheckedGo)); err != nil {
		return
	}
	// rendererprocess/api/markup/childParent.go
	fname = fileNames.ChildParentDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupChildParentGo)); err != nil {
		return
	}
	// rendererprocess/api/markup/class.go
	fname = fileNames.ClassDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupClassGo)); err != nil {
		return
	}
	// rendererprocess/api/markup/data.go
	fname = fileNames.LCDataDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupDataGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/markup/element.go
	fname = fileNames.ElementDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupElementGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/markup/event.go
	fname = fileNames.EventDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupEventGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/markup/focusblur.go
	fname = fileNames.FocusBlurDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupFocusBlurGo)); err != nil {
		return
	}
	// rendererprocess/api/markup/hideshow.go
	fname = fileNames.HideShowDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupHideShowGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/markup/scroll.go
	fname = fileNames.ScrollDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupScrollGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/markup/metrics.go
	fname = fileNames.MetricsDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupMetricsGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/markup/size.go
	fname = fileNames.SizeDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupSizeGo, data, appPaths); err != nil {
		return
	}
	// rendererprocess/api/markup/texthtml.go
	fname = fileNames.TextHTMLDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupTextHTML)); err != nil {
		return
	}
	// rendererprocess/api/markup/value.go
	fname = fileNames.ValueDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupValueGo)); err != nil {
		return
	}
	// rendererprocess/api/markup/id.go
	fname = fileNames.IDDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupIDGo)); err != nil {
		return
	}
	// rendererprocess/api/markup/print.go
	fname = fileNames.PrintDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupPrintGo, data, appPaths); err != nil {
		return
	}
	return
}
