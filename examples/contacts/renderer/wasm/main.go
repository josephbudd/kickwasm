package main

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/transports/calls"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/transports/call"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

// GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go

const (
	host = "127.0.0.1"
	port = uint(9090)
)

func main() {

	quitCh := make(chan struct{})

	// start with kicknotjs
	// only one allowed per application because of the call back registrar.
	// FYI: every call to NewNotJS returns the same exact pointer.
	notjs := kicknotjs.NewNotJS()

	tools := viewtools.NewTools(notjs)

	// get the lpc client.
	caller := call.NewClient(host, port, tools, notjs)
	caller.SetOnConnectionBreak(
		func([]js.Value) {
			quitCh <- struct{}{}
		},
	)
	// get the local procedure calls
	callsStruct, callsMap := calls.NewCallsAndMap(nil, caller.SendPayload)

	// finish initializing the caller client.
	caller.SetLpcMapLpcCalls(callsMap, callsStruct)
	caller.Connect(func() { doPanels(quitCh, tools, callsStruct, notjs) })

	// wait for the application to quit.
	<-quitCh
	tools.Quit()
}
