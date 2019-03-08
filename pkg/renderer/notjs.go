package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createNotJS(appPaths paths.ApplicationPathsI) (err error) {
	fileNames := paths.GetFileNames()
	notJSFileNames := map[*string]string{
		&templates.NotJSAttributesGo: fileNames.AttributesDotGo,
		//&templates.NotJSCallbackGo:    fileNames.CallBackDotGo,
		&templates.NotJSClassGo:       fileNames.ClassDotGo,
		&templates.NotJSCreateGetGo:   fileNames.CreateGetDotGo,
		&templates.NotJSDataGo:        fileNames.DataDotGo,
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
	return
}
