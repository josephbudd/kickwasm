package secondtab

import (
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/lpc"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/paneling"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab/HelloWorldTemplatePanel"
	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// Prepare initializes this package in preparation for spawning.
func Prepare(cl *lpc.Client, quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending, vtools *viewtools.Tools, njs *notjs.NotJS, help *paneling.Help) {
	client = cl
	tools = vtools

	helloworldtemplatepanel.Prepare(quitChan, eojChan, receiveChan, sendChan, vtools, njs, help)
}