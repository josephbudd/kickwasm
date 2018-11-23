package templates

// Panel is the genereric renderer panel template.
const Panel = `{{$Dot := .}}package {{.PanelName}}

import (
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesCallers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
	"{{.ApplicationGitPath}}{{.ImportRendererInterfacesPanelHelper}}"
	"{{.ApplicationGitPath}}{{.ImportRendererNotJS}}"
	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

/*

	Panel name: {{.PanelName}}
	Panel id:   {{.PanelID}}

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
	tools     *viewtools.Tools // see {{.ImportRendererViewTools}}{{range $panel := .PanelGroup}}

	{{call $Dot.LowerCamelCase $panel.Name}} js.Value{{end}}
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notJS *notjs.NotJS, connection map[types.CallID]caller.Renderer, helper panelHelper.Helper) *Panel {
	panel := &Panel{
		tools: tools,
	}{{range $panel := .PanelGroup}}

	panel.{{call $Dot.LowerCamelCase $panel.Name}} = notJS.GetElementByID("{{$panel.HTMLID}}"){{end}}
	// initialize controler, presenter, caller.
	controler := &Controler{
		panel:  panel,
		quitCh: quitCh,
		tools:  tools,
		notJS:  notJS,
	}
	presenter := &Presenter{
		panel:   panel,
		tools:   tools,
		notJS:   notJS,
	}
	caller := &Caller{
		panel:      panel,
		quitCh:     quitCh,
		connection: connection,
		tools:      tools,
		notJS:      notJS,
	}
	// settings
	panel.controler = controler
	panel.presenter = presenter
	panel.caller = caller
	controler.presenter = presenter
	controler.caller = caller
	presenter.controler = controler
	presenter.caller = caller
	caller.controler = controler
	caller.presenter = presenter
	// completions
	controler.defineControlsSetHandlers()
	presenter.defineMembers()
	caller.addMainProcessCallBacks()
	return panel
}

/*
	Show panel funcs.

	Call these from the controler, presenter and caller.
*/{{if .IsTabSiblingPanel}}{{range $panel := .PanelGroup}}

// show{{$panel.Name}} shows the panel you named {{$panel.Name}} while hiding any other panels in it's group.
// The panel will become visible only when this group of panels becomes visible.
/* Your note for this panel is:
{{$panel.Note}}
*/
func (panel *Panel) show{{$panel.Name}}() {
	panel.tools.ShowPanelInTabGroup(panel.{{call $Dot.LowerCamelCase $panel.Name}})
}
{{end}}{{else}}{{range $panel := .PanelGroup}}

// show{{$panel.Name}} shows the panel you named {{$panel.Name}} while hiding any other panels in it's group.
// This panel's id is {{$panel.HTMLID}}.
// This panel either becomes visible immediately or whenever it's panel group is made visible for whatever reason.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
{{$panel.Note}}
*/
func (panel *Panel) show{{$panel.Name}}(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.{{call $Dot.LowerCamelCase $panel.Name}}, force)
}{{end}}{{end}}

// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controler.initialCalls()
	panel.caller.initialCalls()
}
`
