// +build js, wasm

package helloworldtemplatepanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/lpc"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/notjs"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/paneling"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/viewtools"
)

/*

	Panel name: HelloWorldTemplatePanel

*/

var (
	// quitCh will close the application
	quitCh chan struct{}

	// eojCh will close each panel messenger's message dispatcher go routine.
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
