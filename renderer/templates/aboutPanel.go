package templates

// AboutPanel is the mainprocess about panel template.
const AboutPanel = `package AboutPanel

import (
	"github.com/josephbudd/kicknotjs"

	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

// Panel is a panel
type Panel struct {
	presenter *Presenter
	caller    *Caller
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notjs *kicknotjs.NotJS, connection types.RendererCallMap) *Panel {
	v := &Panel{
		presenter: newPresenter(notjs),
		caller:    newCaller(quitCh, connection, tools),
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
func newCaller(quitCh chan struct{}, connection types.RendererCallMap, tools *viewtools.Tools) *Caller {
	v := &Caller{
		quitCh:   quitCh,
		connection: connection,
		tools:    tools,
	}
	v.addCallBacks()
	return v
}
`
