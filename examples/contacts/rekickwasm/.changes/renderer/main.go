package main

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/callClient"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/calls"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/implementations/panelHelping"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*
	YOU MAY EDIT THIS FILE.

	For example: You may want to redefine the helper which is passed to your markup panel constructors.
		1. Edit the definition of the renderer/interfaces/panelHelper.Helper interface.
		2. Define a new implementation of panelHelper in the renderer/implementation/panelHelping package.
		3. In func main below, set helper to your new implementation.
		4. Modify the Panel constructors in the markup panel packages
		   in the render/panels/ folder to use your new definition
		   of the panelHelper.Helper interface.

	Rekickwasm will preserve this file for you.

	BUILD INSTRUCTIONS:

		GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go
		cd ..
		go build

*/

func main() {
	quitCh := make(chan struct{})
	notJS := notjs.NewNotJS()
	tools := viewtools.NewTools(notJS)
	helper := &panelHelping.NoHelp{}

	// get the renderer's connection client.
	host, port := notJS.HostPort()
	client := call.NewClient(host, port, tools, notJS)
	client.SetOnConnectionBreak(
		func([]js.Value) {
			quitCh <- struct{}{}
		},
	)
	// get the local procedure calls
	callMap := calls.GetCallMap(client.SendPayload)

	// finish initializing the caller client.
	client.SetCallMap(callMap)
	client.Connect(func() { doPanels(quitCh, tools, callMap, notJS, helper) })

	// wait for the application to quit.
	<-quitCh
	tools.Quit()
}

