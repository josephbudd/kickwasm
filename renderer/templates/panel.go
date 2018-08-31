package templates

// Panel is the genereric renderer panel template.
const Panel = `{{$Dot := .}}package {{.PanelName}}

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"{{.ApplicationGitPath}}{{.ImportMainProcessTransportsCalls}}"
	"{{.ApplicationGitPath}}{{.ImportRendererWASMViewTools}}"
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
	tools     *viewtools.Tools // see {{.ImportRendererWASMViewTools}}
	notjs     *kicknotjs.NotJS{{range $panel := .PanelGroup}}

	{{call $Dot.LowerCamelCase $panel.Name}} js.Value{{end}}
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notjs *kicknotjs.NotJS, connection *calls.Calls) *Panel {
	panel := &Panel{
		tools: tools,
	}{{range $panel := .PanelGroup}}

	panel.{{call $Dot.LowerCamelCase $panel.Name}} = notjs.GetElementByID("{{$panel.HTMLID}}"){{end}}
	// initialize controler, presenter, caller.
	controler := &Controler{
		panel:  panel,
		quitCh: quitCh,
		tools:  tools,
		notjs:  notjs,
	}
	controler.defineControlsSetHandlers()
	presenter := &Presenter{
		panel:   panel,
		tools:   tools,
		notjs:   notjs,
	}
	presenter.defineMembers()
	caller := &Caller{
		panel:      panel,
		quitCh:     quitCh,
		connection: connection,
		tools:      tools,
		notjs:      notjs,
	}
	caller.addMainProcessCallBacks()
	// finish controler, presenter, caller.
	controler.presenter = presenter
	controler.caller = caller
	presenter.controler = controler
	presenter.caller = caller
	caller.controler = controler
	caller.presenter = presenter
	// finish panel
	panel.controler = controler
	panel.presenter = presenter
	panel.caller = caller
	return panel
}

/*
	Show panel funcs.

	Calls these from the controler, presenter and caller.
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
// The panel will become visible only when this group of panels becomes visible.
// Param force boolean
//  * if force is true and the currently displayed panel is a descendent of div #tabsMasterView-home-slider-collection,
//    ( like a button pad (but not the home button pad), or a tab bar or one of your content panels)
//    Then this function
//     * immediately hides that currently displayed panel.
//     * immediately shows this panels group which means that
//          this panel #{{$panel.HTMLID}}, which you named {{$panel.Name}}, becomes visible.
/* Your note for this panel is:
{{$panel.Note}}
*/
func (panel *Panel) show{{$panel.Name}}(force bool) {
	panel.tools.ShowPanelInButtonGroup(panel.{{call $Dot.LowerCamelCase $panel.Name}}, force)
}{{end}}{{end}}

// InitialCalls runs the first code that the panel needs to run.
func (p *Panel) InitialCalls() {
	p.controler.initialCalls()
	p.caller.initialCalls()
}
`
