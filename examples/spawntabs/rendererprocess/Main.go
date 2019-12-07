// +build js, wasm

package main

import (
	"log"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawntabs/domain/lpc/message"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework/lpc"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework/location"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/paneling"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework/viewtools"
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
	help := paneling.NewHelp()

	// get the renderer's channels
	host, port := location.HostPort()
	client := lpc.NewClient(host, port, quitChan, eojChan, receiveChan, sendChan)

	// finish initializing the messenger client.
	err := client.Connect(func() {
		if er := framework.DoPanels(quitChan, eojChan, receiveChan, sendChan, help); er != nil {
			errmsg := er.Error()
			log.Println(errmsg)
			js.Global().Get("alert").Invoke(errmsg)
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
	viewtools.Quit()
}
