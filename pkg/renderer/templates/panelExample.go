package templates

// PanelHelperInstructions is the template for the file renderer/starter/example.go
const PanelHelperInstructions = `
/*

INSTRUCTIONS FOR USING THE Help struct.

STEP 1: Edit the file Helping.go here in renderer/paneling/
		  by completing the definition of the struct Help.
		In the example below I add funcs for getting states.

	// Help helps initialize panels.
	type Help struct{}
	
	func NewHelp() *Help {
		return &Help{}
	}

	// StateLetter returns the state for letters.
	func (help *Help) StateLetter() uint64 { return uint64(1) }
	
	// StateNumber returns the state for numbers.
	func (help *Help) StateNumber() uint64 { return uint64(1 << 1) }
	
	// StatePunctuation returns the state for punctuation.
	func (help *Help) StatePunctuation() uint64 { return uint64(1 << 2) }
	
	// StateSpecial returns the state for special chars.
	func (help *Help) StateSpecial() uint64 { return uint64(1 << 3) }
	
	// StatePractice returns the state for practice.
	func (help *Help) StatePractice() uint64 { return uint64(1 << 4) }
	
	// StateTest returns the state for test.
	func (help *Help) StateTest() uint64 { return uint64(1 << 5) }

	
STEP 2: In renderer/panels/ and renderer/spawnPanels/
		Add any members new members to your controlers, presenters, and callers.
		Set those new members in the panel constructor using help *paneling.Help

	Example: Some of my panel callers send the same message to the main process
			   but the message has a state which represents what the panel's unique purpose.
			   
	2.1: So I add the state member to some of my panel callers.

		state uint64

	2.2: Then in a panel's func NewPanel(..., help *paneling.Help)
	     I set the caller's state member to the correct state for that panel.

		caller := &panelCaller{
			group: group,
			state: help.StatePunctuation(),
		}
`

// PanelHelperImplementation is the template for the file renderer/implementations/starting/noHelp.go
const PanelHelperImplementation = `package paneling

// Help helps initialize panels.
type Help struct{}

func NewHelp() *Help {
	return &Help{}
}
`