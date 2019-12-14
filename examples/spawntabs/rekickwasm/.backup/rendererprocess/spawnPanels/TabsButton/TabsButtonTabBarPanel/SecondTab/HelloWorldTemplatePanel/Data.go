// +build js, wasm

package helloworldtemplatepanel

import (
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework/lpc"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/paneling"
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

	// Your customized paneling.Help for initializing your panels.
	help *paneling.Help
)