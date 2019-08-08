package createpanel

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/lpc"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/viewtools"
)

/*

	Panel name: CreatePanel

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

	// The framework's renderer API.
	tools *viewtools.Tools

	// Some javascipt like dom functions written in go.
	notJS *notjs.NotJS

	// The javascript null value
	null = js.Null()

	// spawnCount is the number of spawns.
	spawnCount uint
)
