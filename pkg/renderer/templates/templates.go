package templates

import (
	"bytes"
	"text/template"

	"github.com/josephbudd/kickwasm/pkg/paths"
)

// ProcessTemplate executes and writes a template.
func ProcessTemplate(templateName, outputPath, templateString string, templateParams interface{}, appPaths paths.ApplicationPathsI) error {
	bb := new(bytes.Buffer)
	t, err := template.New(templateName).Parse(templateString)
	if err != nil {
		return err
	}
	t, err = t.Parse(templateString)
	if err != nil {
		return err
	}
	if err := t.Execute(bb, templateParams); err != nil {
		return err
	}
	return appPaths.WriteFile(outputPath, bb.Bytes())
}
