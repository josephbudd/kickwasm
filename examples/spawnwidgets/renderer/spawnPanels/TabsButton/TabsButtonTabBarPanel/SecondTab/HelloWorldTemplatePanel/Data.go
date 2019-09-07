package helloworldtemplatepanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/lpc"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/paneling"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/renderer/viewtools"
)

/*

	Panel name: HelloWorldTemplatePanel

*/

var (
	// quitCh will close the application
	quitCh chan struct{}

	// eojCh will close each caller's go routine.
	eojCh chan struct{}

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// The framework's renderer API.
	tools *viewtools.Tools

	// Some javascipt like dom functions written in go.
	notJS *notjs.NotJS

	// Your customized paneling.Help for initializing your panels.
	help *paneling.Help

	// The javascript null value.
	null = js.Null()
)
