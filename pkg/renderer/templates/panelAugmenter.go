package templates

// PanelHelperExample is the template for the file renderer/panelHelper/example.go
const PanelHelperExample = `
/*

STEP 1: REFACTOR THE Helper INTERFACE IN renderer/interfaces/panelHelper/helper.go

	Example:

	package panelHelper

	// Helper is help for setting up the markup panels.
	type Helper interface {
		StateAdd() uint64
		StateEdit() uint64
		StateRemove() uint64
	}

STEP 2: REFACTOR THE Helper IMPLEMENTATION IN renderer/implementations/panelHelping/helping.go

	Example: renderer/implementations/panelHelping/production.go

	func NewProductionHelper() *ProductionHelper {
		v := &ProductionHelper{}
		return v
	}

	type ProductionHelper struct{}

	func (helper *ProductionHelper) StateAdd() { return uint64(1) }

	func (helper *ProductionHelper) StateEdit() { return uint64(2) }

	func (helper *ProductionHelper) StateRemove() { return uint64(3) }


STEP 3: IN func main(), SET panelHelper TO YOUR OWN PRODUCTION HELPER.

	Example:

	panelHelper := panelHelper.NewProductionHelper()


STEP 4: REFACTOR YOUR PANEL AND CONSTRUCTOR AND/OR CONTROLERS AND/OR PRESENTERS, AND/OR CALLERS.

	Example:

	In a panel's func NewPanel(..., helper panelHelper.Helper)

	// initialize controler, presenter, caller.
	controler := &Controler{
		panel:  panel,
		quitCh: quitCh,
		tools:  tools,
		notJS:  notJS,
	}
	presenter := &Presenter{
		panel: panel,
		tools: tools,
		notJS: notJS,
	}
	caller := &Caller{
		panel:      panel,
		quitCh:     quitCh,
		connection: connection,
		tools:      tools,
		notJS:      notJS,
		state:      helper.StateEdit(),
	}

	Then in a caller call back handler func ...

	if params.State & panelCaller.state == panelCaller.state {
		...
	}
`

// PanelHelperInterface is the template for the file renderer/interfaces/panelHelper/helper.go
const PanelHelperInterface = `package panelHelper

// Helper is help for setting up the markup panels.
type Helper interface{}
`

// PanelHelperImplementation is the template for the file renderer/implementations/panelHelping/noHelp.go
const PanelHelperImplementation = `package panelHelping

// NoHelp implements the default empty renderer/interfaces/panelHelper.Helper
type NoHelp struct{}
`
