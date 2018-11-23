package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

var notJSFileNames = map[*string]string{
	&templates.NotJSAttributesGo:  "attributes.go",
	&templates.NotJSCallbackGo:    "callback.go",
	&templates.NotJSClassGo:       "class.go",
	&templates.NotJSCreateGetGo:   "createGet.go",
	&templates.NotJSDataGo:        "data.go",
	&templates.NotJSDocumentGo:    "document.go",
	&templates.NotJSEventsGo:      "events.go",
	&templates.NotJSFormsGo:       "forms.go",
	&templates.NotJSHelpersGo:     "helpers.go",
	&templates.NotJSInnerGo:       "inner.go",
	&templates.NotJSNotJSGo:       "notJS.go",
	&templates.NotJSParentChildGo: "parentChild.go",
	&templates.NotJSScrollGo:      "scroll.go",
	&templates.NotJSSizeGo:        "size.go",
	&templates.NotJSStyleGo:       "style.go",
}

func createNotJS(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	for tptr, fname := range notJSFileNames {
		oPath := filepath.Join(folderpaths.OutputRendererNotJS, fname)
		if err = appPaths.WriteFile(oPath, []byte(*tptr)); err != nil {
			return
		}
	}
	return
}
