package renderer

import (
	"bytes"
	"text/template"

	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

// Returns the html and the error
func buildIndexHTML(builder *tap.Builder, addLocations bool, headTemplateFile string) ([]byte, error) {
	head, err := buildHead(builder, headTemplateFile)
	if err != nil {
		return nil, err
	}
	data := &struct {
		Classes        *tap.Classes
		Head           string
		TabsMasterView string
	}{
		Classes:        builder.Classes,
		Head:           head,
		TabsMasterView: builder.ToHTML(tabsMasterViewID, initialIndent, addLocations),
	}
	// execute the template
	var bb bytes.Buffer
	t := template.Must(template.New("index.html").Parse(templates.IndexHTML))
	if err := t.Execute(&bb, data); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}
