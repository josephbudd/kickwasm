package states

import (
	"github.com/josephbudd/kickwasmwidgets"
)

// States is all of the states needed for panels.
type States struct {
	Add    uint64
	Edit   uint64
	Remove uint64
}

var states *States

func init() {
	vListState := kickwasmwidgets.NewVListState()
	states = &States{
		Add:    vListState.GetNextState(),
		Edit:   vListState.GetNextState(),
		Remove: vListState.GetNextState(),
	}
}

// NewStates constructs a States.
func NewStates() *States {
	return &States{
		Add:    states.Add,
		Edit:   states.Edit,
		Remove: states.Remove,
	}
}
