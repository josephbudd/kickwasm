package renderer

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

type headTemplateData struct {
	Title    string
	UserHead string
}

func buildHead(builder *tap.Builder, headTemplateName string) (string, error) {
	data := &struct {
		Title    string
		UserHead string
	}{
		Title:    builder.Title,
		UserHead: fmt.Sprintf(`{{ template "%s" }}`, headTemplateName),
	}

	// execute the template
	var bb bytes.Buffer
	t := template.Must(template.New("head").Parse(templates.HeadTemplate))
	err := t.Execute(&bb, data)
	if err != nil {
		return tap.EmptyString, err
	}
	return bb.String(), nil
}
