package templates

// MainGo is the rendererprocess/Main.go template.
const MainGo = `// +build js, wasm

package main

import (
	"context"
	"log"
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportDomainLPCMessage}}"
	"{{.ApplicationGitPath}}{{.ImportRendererFramework}}"
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkLocation}}"
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkLPC}}"
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkViewTools}}"
	"{{.ApplicationGitPath}}{{.ImportRendererPaneling}}"
)

/*
	YOU MAY EDIT THIS FILE.
	Rekickwasm will preserve this file for you.

	You may want to redefine the paneling.Help which is passed to your markup panel constructors.
		1. Edit type Help struct in rendererprocess/paneling/Helping.go.
		2. Modify each package's Panel.go func NewPanel in the markup panel packages
		   in the renderprocess/panels/ folder to use your new definition
		   of paneling.Help.
		3. Modify each package's Panel.go func newPanel in the markup panel packages
			in the renderprocess/spawnPanels/ folder to use your new definition
			of paneling.Help.


*/

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()
	host, port := location.HostPort()
	sendChan, receiveChan := lpc.Channels()

	// Connect the LPC client to the LPC server.
	client := lpc.NewClient(ctx, ctxCancel, host, port, receiveChan, sendChan)
	err := client.Connect(func() {
		help := paneling.NewHelp()
		if er := framework.DoMarkupPanels(ctx, ctxCancel, receiveChan, sendChan, help); er != nil {
			errmsg := er.Error()
			log.Println(errmsg)
			js.Global().Get("alert").Invoke(errmsg)
			return
		}
		// Tell the main process that the renderer process is up and running.
		sendChan <- &message.InitRendererToMainProcess{}
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	// wait for the application to quit.
	select {
	case <-ctx.Done():
		viewtools.Quit()
		return
	}
}
`
