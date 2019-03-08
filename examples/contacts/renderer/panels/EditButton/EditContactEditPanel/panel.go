package editcontacteditpanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/interfaces/panelHelper"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: EditContactEditPanel

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *Controler
	presenter *Presenter
	caller    *Caller
}

// NewPanel constructs a new panel.
func NewPanel(quitCh chan struct{}, tools *viewtools.Tools, notJS *notjs.NotJS, connection map[types.CallID]caller.Renderer, helper panelHelper.Helper) (panel *Panel, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "EditContactEditPanel")
		}
	}()

	panelGroup := &PanelGroup{
		tools: tools,
		notJS: notJS,
	}
	controler := &Controler{
		panelGroup: panelGroup,
		quitCh:     quitCh,
		tools:      tools,
		notJS:      notJS,
	}
	presenter := &Presenter{
		panelGroup: panelGroup,
		tools:      tools,
		notJS:      notJS,
	}
	caller := &Caller{
		panelGroup: panelGroup,
		quitCh:     quitCh,
		connection: connection,
		tools:      tools,
		notJS:      notJS,
		state:      helper.StateEdit(),
	}

	controler.presenter = presenter
	controler.caller = caller
	presenter.controler = controler
	presenter.caller = caller
	caller.controler = controler
	caller.presenter = presenter

	// completions
	if err = panelGroup.defineMembers(); err != nil {
		return
	}
	if err = controler.defineControlsSetHandlers(); err != nil {
		return
	}
	if err = presenter.defineMembers(); err != nil {
		return
	}
	if err = caller.addMainProcessCallBacks(); err != nil {
		return
	}

	// No errors so define the panel.
	panel = &Panel{
		controler: controler,
		presenter: presenter,
		caller:    caller,
	}
	return
}

// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controler.initialCalls()
	panel.caller.initialCalls()
}
