package calling

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

// MainProcess is a procedure call from the main process to the renderer and the renderer to the main process.
// MainProcess implements caller.MainProcess with func Process
type MainProcess struct {
	ID                 types.CallID
	mainprocessReceive func(params []byte, call func(params []byte))
}

// NewMainProcess constructs a new call.
func NewMainProcess(
	id types.CallID,
	mainprocessReceive func(params []byte, callback func(params []byte)),
) *MainProcess {
	return &MainProcess{
		ID:                 id,
		mainprocessReceive: mainprocessReceive,
	}
}

// Process passes the params to the main process.
func (call *MainProcess) Process(params []byte, callback func(params []byte)) {
	call.mainprocessReceive(params, callback)
}
