package templates

import (
	"bytes"
	"text/template"

	"github.com/josephbudd/kickwasm/paths"
)

// ProcessTemplate processes a template.
func ProcessTemplate(templateName, outputPath, templateString string, templateParams interface{}, appPaths paths.ApplicationPathsI) error {
	bb := new(bytes.Buffer)
	t := template.Must(template.New(templateName).Parse(templateString))
	if err := t.Execute(bb, templateParams); err != nil {
		return err
	}
	return appPaths.WriteFile(outputPath, bb.Bytes())
}
