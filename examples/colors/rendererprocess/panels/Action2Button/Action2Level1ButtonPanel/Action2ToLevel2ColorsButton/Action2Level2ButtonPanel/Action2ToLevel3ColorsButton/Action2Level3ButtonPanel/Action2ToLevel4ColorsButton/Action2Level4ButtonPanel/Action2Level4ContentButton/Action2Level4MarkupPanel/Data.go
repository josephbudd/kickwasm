// +build js, wasm

package action2level4markuppanel

import (
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/dom"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/framework/lpc"
)

/*

	Panel name: Action2Level4MarkupPanel

*/

var (
	// quitCh will close the application
	quitCh chan struct{}

	// eojCh will signal go routines to stop and return because the application is ending.
	eojCh chan struct{}

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// The document object module.
	document *dom.DOM
)
