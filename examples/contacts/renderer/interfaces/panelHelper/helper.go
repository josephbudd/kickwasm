package panelHelper

// Helper is help for setting up the markup panels.
type Helper interface {
	StateAdd() uint64
	StateEdit() uint64
	StateRemove() uint64
	StateAbout() uint64
}
