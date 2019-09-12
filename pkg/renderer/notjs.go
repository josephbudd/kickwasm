package renderer

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createNotJS(appPaths paths.ApplicationPathsI) (err error) {
	fileNames := paths.GetFileNames()
	notJSFileNames := map[*string]string{
		&templates.NotJSAttributesGo: fileNames.AttributesDotGo,
		//&templates.NotJSCallbackGo:    fileNames.CallBackDotGo,
		&templates.NotJSClassGo: fileNames.ClassDotGo,
		// &templates.NotJSCreateGetGo:   fileNames.CreateGetDotGo,
		&templates.NotJSDataGo:        fileNames.LCDataDotGo,
		&templates.NotJSDocumentGo:    fileNames.DocumentDotGo,
		&templates.NotJSEventsGo:      fileNames.EventsDotGo,
		&templates.NotJSFormsGo:       fileNames.FormsDotGo,
		&templates.NotJSHelpersGo:     fileNames.HelpersDotGo,
		&templates.NotJSInnerGo:       fileNames.InnerDotGo,
		&templates.NotJSNotJSGo:       fileNames.NotJSDotGo,
		&templates.NotJSParentChildGo: fileNames.ParentChildDotGo,
		&templates.NotJSScrollGo:      fileNames.ScrollDotGo,
		&templates.NotJSSizeGo:        fileNames.SizeDotGo,
		&templates.NotJSStyleGo:       fileNames.StyleDotGo,
	}
	folderpaths := appPaths.GetPaths()
	for tptr, fname := range notJSFileNames {
		oPath := filepath.Join(folderpaths.OutputRendererNotJS, fname)
		if err = appPaths.WriteFile(oPath, []byte(*tptr)); err != nil {
			return
		}
	}

	// The create-get funcs.
	data := &struct {
		HTMLNames []string
		ToUpper   func(string) string
		ToLower   func(string) string
	}{
		ToUpper: strings.ToUpper,
		ToLower: strings.ToLower,
		HTMLNames: []string{
			"style",
			"address",
			"article",
			"aside",
			"footer",
			"header",
			"h1",
			"h2",
			"h3",
			"h4",
			"h5",
			"h6",
			"hgroup",
			"main",
			"section",
			"blockquote",
			"dd",
			"div",
			"dl",
			"dt",
			"figcaption",
			"figure",
			"hr",
			"li",
			"ol",
			"p",
			"pre",
			"ul",
			"a",
			"abbr",
			"b",
			"bdo",
			"bdi",
			"br",
			"cite",
			"data",
			"code",
			"dfn",
			"em",
			"i",
			"kbd",
			"mark",
			"q",
			"rb",
			"rp",
			"rt",
			"rtc",
			"ruby",
			"s",
			"samp",
			"small",
			"span",
			"strong",
			"sub",
			"sup",
			"time",
			"tt",
			"u",
			"var",
			"wbr",
			"area",
			"audio",
			"img",
			"map",
			"track",
			"video",
			"embed",
			"iframe",
			"object",
			"param",
			"picture",
			"source",
			"canvas",
			"del",
			"ins",
			"caption",
			"col",
			"colgroup",
			"table",
			"tbody",
			"td",
			"tfoot",
			"th",
			"thead",
			"tr",
			"button",
			"datalist",
			"fieldset",
			"label",
			"legend",
			"meter",
			"optgroup",
			"option",
			"output",
			"progress",
			"select",
			"textarea",
			"details",
			"dialog",
			"menu",
			"summary",
			"element",
			"slot",
			"template",
		},
	}
	sort.Strings(data.HTMLNames)
	fname := fileNames.CreateGetDotGo
	oPath := filepath.Join(folderpaths.OutputRendererNotJS, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.NotJSCreateGetGo, data, appPaths)
	return
}
