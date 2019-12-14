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
		ApplicationGitPath      string
		ImportRendererCallBack  string
		ImportRendererEvent     string
		ImportRendererWindow    string
		ImportRendererViewTools string
		HVScrollClassName       string
		ResizeMeWidthClassName  string
	}{
		ApplicationGitPath:      builder.ImportPath,
		ImportRendererCallBack:  folderpaths.ImportRendererCallBack,
		ImportRendererEvent:     folderpaths.ImportRendererEvent,
		ImportRendererWindow:    folderpaths.ImportRendererWindow,
		ImportRendererViewTools: folderpaths.ImportRendererViewTools,
		HVScrollClassName:       classes.HVScroll,
		ResizeMeWidthClassName:  classes.ResizeMeWidth,
	}

	// renderer/markup/attributes.go
	fname = fileNames.AttributesDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupAttributeGo)); err != nil {
		return
	}
	// renderer/markup/checked.go
	fname = fileNames.CheckedDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupCheckedGo)); err != nil {
		return
	}
	// renderer/markup/childParent.go
	fname = fileNames.ChildParentDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupChildParentGo)); err != nil {
		return
	}
	// renderer/markup/class.go
	fname = fileNames.ClassDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupClassGo)); err != nil {
		return
	}
	// renderer/markup/data.go
	fname = fileNames.LCDataDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupDataGo)); err != nil {
		return
	}
	// renderer/markup/element.go
	fname = fileNames.ElementDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupElementGo, data, appPaths); err != nil {
		return
	}
	// renderer/markup/event.go
	fname = fileNames.EventDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupEventGo, data, appPaths); err != nil {
		return
	}
	// renderer/markup/focusblur.go
	fname = fileNames.FocusBlurDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupFocusBlurGo)); err != nil {
		return
	}
	// renderer/markup/hideshow.go
	fname = fileNames.HideShowDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupHideShowGo, data, appPaths); err != nil {
		return
	}
	// renderer/markup/scroll.go
	fname = fileNames.ScrollDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupScrollGo, data, appPaths); err != nil {
		return
	}
	// renderer/markup/size.go
	fname = fileNames.SizeDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MarkupSizeGo, data, appPaths); err != nil {
		return
	}
	// renderer/markup/texthtml.go
	fname = fileNames.TextHTMLDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupTextHTML)); err != nil {
		return
	}
	// renderer/markup/value.go
	fname = fileNames.ValueDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupValueGo)); err != nil {
		return
	}
	// renderer/markup/id.go
	fname = fileNames.IDDotGo
	oPath = filepath.Join(folderpaths.OutputRendererMarkup, fname)
	if err = appPaths.WriteFile(oPath, []byte(templates.MarkupIDGo)); err != nil {
		return
	}
	return
}