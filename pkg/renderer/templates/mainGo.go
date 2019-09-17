package templates

// MainGo is the renderer/main.go template.
const MainGo = `package main

import (
	"log"

	"{{.ApplicationGitPath}}{{.ImportDomainLPCMessage}}"
	"{{.ApplicationGitPath}}{{.ImportRendererLPC}}"
	"{{.ApplicationGitPath}}{{.ImportRendererNotJS}}"
	"{{.ApplicationGitPath}}{{.ImportRendererPaneling}}"
	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

/*
	YOU MAY EDIT THIS FILE.

	For example: You may want to redefine the starting.Helper which is passed to your markup panel constructors.
		1. Edit the definition of the renderer/interfaces/starter.Helper interface.
		2. Define a new implementation of starter in the renderer/implementation/starting package.
		3. In func main below, set helper to your new implementation.
		4. Modify the Panel constructors in the markup panel packages
		   in the render/panels/ folder to use your new definition
		   of the starter.Helper interface.

	Rekickwasm will preserve this file for you.

	BUILD INSTRUCTIONS:

		cd renderer
		./build.sh
		cd ..
		go build

*/

func main() {
	sendChan, receiveChan, eojChan := lpc.Channels()
	quitChan := make(chan struct{})
	notJS := notjs.NewNotJS()
	tools := viewtools.NewTools(notJS)
	help := paneling.NewHelp()

	// get the renderer's channels
	host, port := notJS.HostPort()
	client := lpc.NewClient(host, port, tools, quitChan, eojChan, receiveChan, sendChan)

	// finish initializing the caller client.
	err := client.Connect(func() {
		if er := doPanels(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); er != nil {
			errmsg := er.Error()
			tools.ConsoleLog(errmsg)
			tools.Alert(errmsg)
			return
		}
		sendChan <- &message.InitRendererToMainProcess{}
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	// wait for the application to quit.
	<-eojChan
	tools.Quit()
}
`
