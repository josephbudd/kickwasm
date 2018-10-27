package templates

// MainGo is the renderer/main.go template.
const MainGo = `package main

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"{{.ApplicationGitPath}}{{.ImportRendererCallClient}}"
	"{{.ApplicationGitPath}}{{.ImportRendererCalls}}"
	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
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

	// get the renderer's connection client.
	client := call.NewClient(host, port, tools, notjs)
	client.SetOnConnectionBreak(
		func([]js.Value) {
			quitCh <- struct{}{}
		},
	)
	// get the local procedure calls
	callMap := calls.GetCallMap(client.SendPayload)

	// finish initializing the caller client.
	client.SetCallMap(callMap)
	client.Connect(func() { doPanels(quitCh, tools, callMap, notjs) })

	// wait for the application to quit.
	<-quitCh
	tools.Quit()
}

`
