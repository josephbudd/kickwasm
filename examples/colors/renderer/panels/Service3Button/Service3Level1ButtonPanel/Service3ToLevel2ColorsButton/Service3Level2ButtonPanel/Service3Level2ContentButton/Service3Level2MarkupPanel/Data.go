package service3level2markuppanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/colors/renderer/lpc"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/viewtools"
)

/*

	Panel name: Service3Level2MarkupPanel

*/

var (
	// quitCh will close the application
	quitCh chan struct{}

	// eojCh will signal go routines to stop and return becuase the application is ending.
	eojCh chan struct{}

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// The framework's renderer API.
	tools *viewtools.Tools

	// Some javascipt like dom functions written in go.
	notJS *notjs.NotJS

	// The javascript null value
	null = js.Null()
)
