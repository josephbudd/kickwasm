// +build js, wasm

package framework

import (
	"fmt"
	"log"


	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework/lpc"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework/viewtools"
	"github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/paneling"
	createpanel "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/panels/TabsButton/TabsButtonTabBarPanel/FirstTab/CreatePanel"
	tabsbuttontabbarpanel "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/spawnPanels/TabsButton/TabsButtonTabBarPanel"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// DoPanels builds and runs the panels.
func DoPanels(quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending,
	help *paneling.Help) (err error) {
	
	defer func() {
		if err != nil {
			err = fmt.Errorf("DoPanels: %w", err)
			log.Println("Error: " + err.Error())
		}
	}()

	// 1. Prepare the spawn panels.
	tabsbuttontabbarpanel.Prepare(quitChan, eojChan, receiveChan, sendChan, help)

	// 2. Construct the panel code.
	var createPanel *createpanel.Panel
	if createPanel, err = createpanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, help); err != nil {
		return
	}

	// 3. Size the app.
	viewtools.SizeApp()

	// 4. Start each panel's message and event dispatchers.
	createPanel.StartDispatchers()

	// 5. Start each panel's initial calls.
	createPanel.InitialJobs()

	return
}