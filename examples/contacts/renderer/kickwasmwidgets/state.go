package kickwasmwidgets

const (
	mask       = (1 << 3) - 1
	firstShift = uint(3)
)

// Common vlist states.
var (
	StateInitialize uint64
	StatePrepend    uint64
	StateAppend     uint64
	vListState      *VListState
)

func init() {
	vListState = &VListState{
		shift: firstShift,
	}
	StatePrepend = vListState.GetNextState()
	StateAppend = vListState.GetNextState()
}

// Any state can be ored with any other state.
// States are also intended to be ored with a VList subpanel index < 2^3.
// VList subpanel index values of 0-7 inclusive, are valid.
//
// There are 63 bits for state.
// 3 of those 63 bits are used for a VList subpanel index.
// So that would leave 63 - 3 or 60 bits for 60 states.
// However, 1 state is used for StatePrepend and another state is used for StateAppend.
// So that leaves only 61 bits for 61 more states.
// So you can make 61 calls to VListState.GetNextState() and get 61 states.

// VListState is states for vlists.
type VListState struct {
	shift uint
}

// NewVListState constructs a new VListState.
// It actually only allows for 1 NewVListState per application.
func NewVListState() *VListState {
	return vListState
}

// GetNextState returns a new state.
func (vlstate *VListState) GetNextState() uint64 {
	// math.MaxUint64 = 1<<64 - 1
	if vlstate.shift == 64 {
		panic("VListState is now corrupt.")
	}
	newState := uint64(1) << vlstate.shift
	vlstate.shift++
	return newState
}

// StateToSubPanelIndex converts a state generated by VListState.GetNextState() to an int of it's mask bits.
func StateToSubPanelIndex(state uint64) uint64 {
	return uint64(state & mask)
}
