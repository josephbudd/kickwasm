package panelHelping

import (
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/kickwasmwidgets"
)

// Production implements renderer/interfaces/panelHelper.Helper
type Production struct {
	add, edit, remove, about uint64
}

// NewProduction constructs a new Production.
func NewProduction() *Production {
	vliststate := kickwasmwidgets.NewVListState()
	return &Production{
		add:    vliststate.GetNextState(),
		edit:   vliststate.GetNextState(),
		remove: vliststate.GetNextState(),
		about: vliststate.GetNextState(),
	}
}

// StateAdd returns the Add state.
func (p *Production) StateAdd() uint64 {
	return p.add
}

// StateEdit returns the Edit state.
func (p *Production) StateEdit() uint64 {
	return p.edit
}

// StateRemove returns the Remove state.
func (p *Production) StateRemove() uint64 {
	return p.remove
}

// StateAbout returns the Remove state.
func (p *Production) StateAbout() uint64 {
	return p.about
}

