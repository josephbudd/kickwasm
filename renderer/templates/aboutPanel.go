package templates

// AboutPanel is the mainprocess about panel template.
const AboutPanel = `package AboutPanel

import (
	"github.com/josephbudd/kicknotjs"

	"{{.ApplicationGitPath}}{{.ImportMainProcessTransportsCalls}}"
	"{{.ApplicationGitPath}}{{.ImportRendererWASMViewTools}}"
)

// Panel is a panel
type Panel struct {
	presenter *Presenter
	caller    *Caller
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notjs *kicknotjs.NotJS, callsStruct *calls.Calls) *Panel {
	v := &Panel{
		presenter: newPresenter(notjs),
		caller:    newCaller(quitCh, callsStruct, tools),
	}
	v.caller.presenter = v.presenter
	return v
}

// InitialCalls runs the first code that the panel needs to run.
func (p *Panel) InitialCalls() {
	p.caller.initialCalls()
}

// newPresenter constructs a new Presenter
// Caller must be set after calling newPresenter.
func newPresenter(notjs *kicknotjs.NotJS) *Presenter {
	v := &Presenter{
		notjs: notjs,
	}
	v.defineMembers()
	return v
}

// newCaller constructs a new Caller.
// Presenter and Controler must be set after this call
func newCaller(quitCh chan struct{}, callsStruct *calls.Calls, tools *viewtools.Tools) *Caller {
	v := &Caller{
		quitCh:   quitCh,
		callsStruct: callsStruct,
		tools:    tools,
	}
	v.addCallBacks()
	return v
}
`
