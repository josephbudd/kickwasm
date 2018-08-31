package templates

// MainGo is the renderer/main.go template.
const MainGo = `package main

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"{{.ApplicationGitPath}}{{.ImportMainProcessTransportsCalls}}"
	"{{.ApplicationGitPath}}{{.ImportRendererWASMCall}}"
	"{{.ApplicationGitPath}}{{.ImportRendererWASMViewTools}}"
)

// GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go

const (
	host = "{{.Host}}"
	port = uint({{.Port}})
)

func main() {

	quitCh := make(chan struct{})

	// start with kicknotjs
	// only one allowed per application because of the call back registrar.
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
	callsStruct, callsMap := calls.NewCallsAndMap({{range .Repos}}nil, {{end}}caller.SendPayload)

	// finish initializing the caller client.
	caller.SetLpcMapLpcCalls(callsMap, callsStruct)
	caller.Connect(func() { doPanels(quitCh, tools, callsStruct, notjs) })

	// wait for the application to quit.
	<-quitCh
	tools.Quit()
}
`
