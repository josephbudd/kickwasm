package templates

// PanelGroup is the genereric renderer panelGroup template.
const PanelGroup = `{{$Dot := .}}{{$lpg := len .PanelGroup}}package {{.PanelName}}

import (
	"syscall/js"

	"github.com/pkg/errors"

	"{{.ApplicationGitPath}}{{.ImportRendererNotJS}}"
	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

// PanelGroup is a group of {{$lpg}} panel{{if gt $lpg 1}}s{{end}}.
// It also has {{if eq $lpg 1}}a {{end}}show panel func{{if gt $lpg 1}}s{{end}} for each panel in this panel group.
type PanelGroup struct {
	tools *viewtools.Tools
	notJS *notjs.NotJS
{{range $panel := .PanelGroup}}
	{{call $Dot.LowerCamelCase $panel.Name}} js.Value{{end}}
}

func (panelGroup *PanelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelGroup *PanelGroup) defineMembers()")
		}
	}()

	notJS := panelGroup.notJS
	null := js.Null()

{{range $panel := .PanelGroup}}	if panelGroup.{{call $Dot.LowerCamelCase $panel.Name}} = notJS.GetElementByID("{{$panel.HTMLID}}"); panelGroup.{{call $Dot.LowerCamelCase $panel.Name}} == null {
		err = errors.New("unable to find #{{$panel.HTMLID}}")
		return
	}
{{end}}

	return
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/{{if .IsTabSiblingPanel}}{{range $panel := .PanelGroup}}

// show{{$panel.Name}} shows the panel you named {{$panel.Name}} while hiding any other panels in this panel group.
// {{if eq $Dot.PanelName $panel.Name}}This{{else}}That{{end}} panel will become visible only when this group of panels becomes visible.
/* Your note for {{if eq $Dot.PanelName $panel.Name}}this{{else}}that{{end}} panel is:
{{$panel.Note}}
*/
func (panelGroup *PanelGroup) show{{$panel.Name}}() {
	panelGroup.tools.ShowPanelInTabGroup(panelGroup.{{call $Dot.LowerCamelCase $panel.Name}})
}
{{end}}{{else}}{{range $panel := .PanelGroup}}

// show{{$panel.Name}} shows the panel you named {{$panel.Name}} while hiding any other panels in this panel group.
// {{if eq $Dot.PanelName $panel.Name}}This{{else}}That{{end}} panel's id is {{$panel.HTMLID}}.
// {{if eq $Dot.PanelName $panel.Name}}This{{else}}That{{end}} panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when {{if eq $Dot.PanelName $panel.Name}}this{{else}}that{{end}} panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for {{if eq $Dot.PanelName $panel.Name}}this{{else}}that{{end}} panel is:
{{$panel.Note}}
*/
func (panelGroup *PanelGroup) show{{$panel.Name}}(force bool) {
	panelGroup.tools.ShowPanelInButtonGroup(panelGroup.{{call $Dot.LowerCamelCase $panel.Name}}, force)
}{{end}}{{end}}

`
