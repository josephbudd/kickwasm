package templates

// RebuildDispatchGo is mainprocess/lpc/dispatch/dispatch.go.
const RebuildDispatchGo = `package dispatch

import (
	"log"

	"{{.ApplicationGitPath}}{{.ImportDomainLPCMessage}}"
	"{{.ApplicationGitPath}}{{.ImportDomainStore}}"
	"{{.ApplicationGitPath}}{{.ImportMainProcessLPC}}"
)

/*
	DO NOT EDIT THIS FILE.

	USE THE TOOL kicklpc TO ADD OR REMOVE LPC Messages.

	kicklpc will edit this file for you.

*/

// Do dispatches local process communications messages received from the renderer.
// They are dispatched to the main process handlers here in package dispatch.
// You are required to code the functionality into those handlers.
func Do(cargo interface{}, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
	switch cargo := cargo.(type) {
	case *message.LogRendererToMainProcess:
		// Log does not need a lpc.EOJer because it does not have a go routine.
		handleLog(cargo, sending){{ range .LPCNames }}
	case *message.{{.}}RendererToMainProcess:
		handle{{.}}(cargo, sending, eojing, stores){{ end }}
	default:
		log.Println("dispatch Do: unknown cargo type.")
	}
}
`

// DispatchInstructions is the mainprocess/lpc/dispatch/instructions.txt.
const DispatchInstructions = `
ABOUT THE FILES IN THE FOLDER mainprocess/lpc/dispatch/.

  * dispatch.go contains func Do which dispatches the LPC ( Local Process Communications ) messages received from the renderer.
    Do not edit the file dispatch.go.
    In func Do, the messages are dispatched to the LPC message handlers here in this folder.

  Message handler files:

  * Log.go was generated by kickwasm when this framework was created.
    The file contains func handleLog(rxMessage *message.LogRendererToMainProcess, sending lpc.Sending)
      which processes the log message received from the renderer.
	You may edit the file if you need to.{{ range .LPCNames }}

  * {{.}}.go was generated by kicklpc when you added the {{.}} LPC message.
    The file contains func handle{{.}}(rxmessage *message.{{.}}RendererToMainProcess, sending lpc.Sending, eojing lpc.EOJer)
	  which must process the {{.}} message received from the renderer.
	kicklpc created func handle{{.}} void of any functionality so that it's functionality could to be coded.
	There fore, you may edit {{.}}.go.{{ end }}

ABOUT THE FILES IN THE FOLDER domain/lpc/message/.

* Log.go was generated by kickwasm when this framework was created.
  The file contains the types of the 2 Log messages.
	1. LogRendererToMainProcess is the message that the renderer sends to the main process.
	2. LogMainProcessToRenderer is the message that the main process sends to the renderer.{{ range .LPCNames }}

* {{.}}.go was generated by kicklpc.
  The file contains the types of the 2 {{.}} messages.
	1. {{.}}RendererToMainProcess is the message that the renderer sends to the main process.
	2. {{.}}MainProcessToRenderer is the message that the main process sends to the renderer.
  kicklpc created the 2 {{.}} message types with little or no structure so that their structure could to be completed.
  There fore, you may edit {{.}}.go.{{ end }}

MANAGING LPC MESSAGES WITH kicklpc.

* Use kicklpc in this application's root folder:
  $ cd {{.ApplicationGitPath}}/

* Listing all of the messages:
  $ kicklpc -l
  1. kicklpc would
    * Display the names of each LPC message.

* Adding a message:
  $ kicklpc -add UpdateCustomer
  1. kicklpc would
    * Add the file domain/lpc/UpdateCustomer.go
    * Add the file mainprocess/lpc/dispatch/UpdateCustomer.go
    * Update the file mainprocess/lpc/dispatch/dispatch.go
  2. You would need to
	* Complete the message definitions in domain/lpc/UpdateCustomer.go
	* Complete the func handleUpdateCustomer in mainprocess/lpc/dispatch/UpdateCustomer.go

* Deleting a message:
  $ kicklpc -delete-forever UpdateCustomer
  1. kicklpc would
    * Delete the file domain/lpc/UpdateCustomer.go
    * Delete the file mainprocess/lpc/dispatch/UpdateCustomer.go
    * Update the file mainprocess/lpc/dispatch/dispatch.go
`

// DispatchLogGo is mainprocess/lpc/dispatch/log.go.
const DispatchLogGo = `package dispatch

import (
	"log"

	"{{.ApplicationGitPath}}{{.ImportDomainDataLogLevels}}"
	"{{.ApplicationGitPath}}{{.ImportDomainLPCMessage}}"
	"{{.ApplicationGitPath}}{{.ImportMainProcessLPC}}"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.
	Kicklpc will not edit this file.

*/

// handleLog logs a renderer message to the application log.
// Param rxMessage *message.LogRendererToMainProcess is the params received from the renderer.
// Param sending is the channel to use to send a *message.LogMainProcessToRenderer to the renderer.
func handleLog(rxMessage *message.LogRendererToMainProcess, sending lpc.Sending) {
	var msg string
	var errorMsg = ""
	var isErr bool
	switch rxMessage.Level {
	case loglevels.LogLevelInfo:
		msg = "Renderer Log: Info: " + rxMessage.Message
	case loglevels.LogLevelWarning:
		msg = "Renderer Log: Warning: " + rxMessage.Message
	case loglevels.LogLevelError:
		msg = "Renderer Log: Error: " + rxMessage.Message
	case loglevels.LogLevelFatal:
		msg = "Renderer Log: Fatal: " + rxMessage.Message
	default:
		errorMsg = "Unknown Level"
		isErr = true
		msg = "Renderer Log: ???: " + rxMessage.Message
	}
	// Log the message from the renderer.
	log.Println(msg)
	// Send an update back to the renderer.
	txMessage := &message.LogMainProcessToRenderer{
		Message:      msg,
		ErrorMessage: errorMsg,
		Error:        isErr,
	}
	sending <- txMessage
	return
}
`

// DispatchMessageGo is mainprocess/lpc/dispatch/<Message Name>.go.
// This is only for kickllpc when it adds a new message.
const DispatchMessageGo = `package dispatch

import (
	"{{.ApplicationGitPath}}{{.ImportDomainLPCMessage}}"
	"{{.ApplicationGitPath}}{{.ImportDomainStore}}"
	"{{.ApplicationGitPath}}{{.ImportMainProcessLPC}}"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.
	Kicklpc will not edit this file.

*/

// handle{{.MessageName}} is the *message.{{.MessageName}}RendererToMainProcess handler.
// It's response back to the renderer is the *message.{{.MessageName}}MainProcessToRenderer.
// Param rxmessage *message.{{.MessageName}}RendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.{{.MessageName}}MainProcessToRenderer message back to the renderer.
// Param eojing lpc.EOJer ( End Of Job ) is an interface for your go routine to receive a stop signal.
//   It signals go routines that they must stop because the main process is ending.
//   So only use it inside a go routine if you have one.
//   In your go routine
//     1. Get a channel to listen to with eojing.NewEOJ().
//     2. Before your go routine returns, release that channel with eojing.Release().
func handle{{.MessageName}}(rxmessage *message.{{.MessageName}}RendererToMainProcess, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
	return
}
`