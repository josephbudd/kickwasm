// +build js, wasm

package main

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/lpc"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/paneling"
	action1level1markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action1Button/Action1Level1ButtonPanel/Action1Level1ContentButton/Action1Level1MarkupPanel"
	action1level2markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action1Button/Action1Level1ButtonPanel/Action1ToLevel2ColorsButton/Action1Level2ButtonPanel/Action1Level2ContentButton/Action1Level2MarkupPanel"
	action1level3markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action1Button/Action1Level1ButtonPanel/Action1ToLevel2ColorsButton/Action1Level2ButtonPanel/Action1ToLevel3ColorsButton/Action1Level3ButtonPanel/Action1Level3ContentButton/Action1Level3MarkupPanel"
	action1level4markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action1Button/Action1Level1ButtonPanel/Action1ToLevel2ColorsButton/Action1Level2ButtonPanel/Action1ToLevel3ColorsButton/Action1Level3ButtonPanel/Action1ToLevel4ColorsButton/Action1Level4ButtonPanel/Action1Level4ContentButton/Action1Level4MarkupPanel"
	action1level5markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action1Button/Action1Level1ButtonPanel/Action1ToLevel2ColorsButton/Action1Level2ButtonPanel/Action1ToLevel3ColorsButton/Action1Level3ButtonPanel/Action1ToLevel4ColorsButton/Action1Level4ButtonPanel/Action1ToLevel5ColorsButton/Action1Level5ButtonPanel/Action1Level5ContentButton/Action1Level5MarkupPanel"
	action2level1markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action2Button/Action2Level1ButtonPanel/Action2Level1ContentButton/Action2Level1MarkupPanel"
	action2level2markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action2Button/Action2Level1ButtonPanel/Action2ToLevel2ColorsButton/Action2Level2ButtonPanel/Action2Level2ContentButton/Action2Level2MarkupPanel"
	action2level3markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action2Button/Action2Level1ButtonPanel/Action2ToLevel2ColorsButton/Action2Level2ButtonPanel/Action2ToLevel3ColorsButton/Action2Level3ButtonPanel/Action2Level3ContentButton/Action2Level3MarkupPanel"
	action2level4markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action2Button/Action2Level1ButtonPanel/Action2ToLevel2ColorsButton/Action2Level2ButtonPanel/Action2ToLevel3ColorsButton/Action2Level3ButtonPanel/Action2ToLevel4ColorsButton/Action2Level4ButtonPanel/Action2Level4ContentButton/Action2Level4MarkupPanel"
	action2level5markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action2Button/Action2Level1ButtonPanel/Action2ToLevel2ColorsButton/Action2Level2ButtonPanel/Action2ToLevel3ColorsButton/Action2Level3ButtonPanel/Action2ToLevel4ColorsButton/Action2Level4ButtonPanel/Action2ToLevel5ColorsButton/Action2Level5ButtonPanel/Action2Level5ContentButton/Action2Level5MarkupPanel"
	action3level1markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action3Button/Action3Level1ButtonPanel/Action3Level1ContentButton/Action3Level1MarkupPanel"
	action3level2markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action3Button/Action3Level1ButtonPanel/Action3ToLevel2ColorsButton/Action3Level2ButtonPanel/Action3Level2ContentButton/Action3Level2MarkupPanel"
	action3level3markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action3Button/Action3Level1ButtonPanel/Action3ToLevel2ColorsButton/Action3Level2ButtonPanel/Action3ToLevel3ColorsButton/Action3Level3ButtonPanel/Action3Level3ContentButton/Action3Level3MarkupPanel"
	action3level4markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action3Button/Action3Level1ButtonPanel/Action3ToLevel2ColorsButton/Action3Level2ButtonPanel/Action3ToLevel3ColorsButton/Action3Level3ButtonPanel/Action3ToLevel4ColorsButton/Action3Level4ButtonPanel/Action3Level4ContentButton/Action3Level4MarkupPanel"
	action3level5markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action3Button/Action3Level1ButtonPanel/Action3ToLevel2ColorsButton/Action3Level2ButtonPanel/Action3ToLevel3ColorsButton/Action3Level3ButtonPanel/Action3ToLevel4ColorsButton/Action3Level4ButtonPanel/Action3ToLevel5ColorsButton/Action3Level5ButtonPanel/Action3Level5ContentButton/Action3Level5MarkupPanel"
	action4level1markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action4Button/Action4Level1ButtonPanel/Action4Level1ContentButton/Action4Level1MarkupPanel"
	action4level2markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action4Button/Action4Level1ButtonPanel/Action4ToLevel2ColorsButton/Action4Level2ButtonPanel/Action4Level2ContentButton/Action4Level2MarkupPanel"
	action4level3markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action4Button/Action4Level1ButtonPanel/Action4ToLevel2ColorsButton/Action4Level2ButtonPanel/Action4ToLevel3ColorsButton/Action4Level3ButtonPanel/Action4Level3ContentButton/Action4Level3MarkupPanel"
	action4level4markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action4Button/Action4Level1ButtonPanel/Action4ToLevel2ColorsButton/Action4Level2ButtonPanel/Action4ToLevel3ColorsButton/Action4Level3ButtonPanel/Action4ToLevel4ColorsButton/Action4Level4ButtonPanel/Action4Level4ContentButton/Action4Level4MarkupPanel"
	action4level5markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action4Button/Action4Level1ButtonPanel/Action4ToLevel2ColorsButton/Action4Level2ButtonPanel/Action4ToLevel3ColorsButton/Action4Level3ButtonPanel/Action4ToLevel4ColorsButton/Action4Level4ButtonPanel/Action4ToLevel5ColorsButton/Action4Level5ButtonPanel/Action4Level5ContentButton/Action4Level5MarkupPanel"
	action5level1markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action5Button/Action5Level1ButtonPanel/Action5Level1ContentButton/Action5Level1MarkupPanel"
	action5level2markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action5Button/Action5Level1ButtonPanel/Action5ToLevel2ColorsButton/Action5Level2ButtonPanel/Action5Level2ContentButton/Action5Level2MarkupPanel"
	action5level3markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action5Button/Action5Level1ButtonPanel/Action5ToLevel2ColorsButton/Action5Level2ButtonPanel/Action5ToLevel3ColorsButton/Action5Level3ButtonPanel/Action5Level3ContentButton/Action5Level3MarkupPanel"
	action5level4markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action5Button/Action5Level1ButtonPanel/Action5ToLevel2ColorsButton/Action5Level2ButtonPanel/Action5ToLevel3ColorsButton/Action5Level3ButtonPanel/Action5ToLevel4ColorsButton/Action5Level4ButtonPanel/Action5Level4ContentButton/Action5Level4MarkupPanel"
	action5level5markuppanel "github.com/josephbudd/kickwasm/examples/colors/rendererprocess/panels/Action5Button/Action5Level1ButtonPanel/Action5ToLevel2ColorsButton/Action5Level2ButtonPanel/Action5ToLevel3ColorsButton/Action5Level3ButtonPanel/Action5ToLevel4ColorsButton/Action5Level4ButtonPanel/Action5ToLevel5ColorsButton/Action5Level5ButtonPanel/Action5Level5ContentButton/Action5Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/rendererprocess/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

func doPanels(quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending,
	tools *viewtools.Tools, notJS *notjs.NotJS, help *paneling.Help) (err error) {
	
	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "doPanels")
			tools.ConsoleLog("Error: " + err.Error())
		}
	}()

	// 1. Prepare the spawn panels.

	// 2. Construct the panel code.
	var action1Level1MarkupPanel *action1level1markuppanel.Panel
	if action1Level1MarkupPanel, err = action1level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action1Level2MarkupPanel *action1level2markuppanel.Panel
	if action1Level2MarkupPanel, err = action1level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action1Level3MarkupPanel *action1level3markuppanel.Panel
	if action1Level3MarkupPanel, err = action1level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action1Level4MarkupPanel *action1level4markuppanel.Panel
	if action1Level4MarkupPanel, err = action1level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action1Level5MarkupPanel *action1level5markuppanel.Panel
	if action1Level5MarkupPanel, err = action1level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action2Level1MarkupPanel *action2level1markuppanel.Panel
	if action2Level1MarkupPanel, err = action2level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action2Level2MarkupPanel *action2level2markuppanel.Panel
	if action2Level2MarkupPanel, err = action2level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action2Level3MarkupPanel *action2level3markuppanel.Panel
	if action2Level3MarkupPanel, err = action2level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action2Level4MarkupPanel *action2level4markuppanel.Panel
	if action2Level4MarkupPanel, err = action2level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action2Level5MarkupPanel *action2level5markuppanel.Panel
	if action2Level5MarkupPanel, err = action2level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action3Level1MarkupPanel *action3level1markuppanel.Panel
	if action3Level1MarkupPanel, err = action3level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action3Level2MarkupPanel *action3level2markuppanel.Panel
	if action3Level2MarkupPanel, err = action3level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action3Level3MarkupPanel *action3level3markuppanel.Panel
	if action3Level3MarkupPanel, err = action3level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action3Level4MarkupPanel *action3level4markuppanel.Panel
	if action3Level4MarkupPanel, err = action3level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action3Level5MarkupPanel *action3level5markuppanel.Panel
	if action3Level5MarkupPanel, err = action3level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action4Level1MarkupPanel *action4level1markuppanel.Panel
	if action4Level1MarkupPanel, err = action4level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action4Level2MarkupPanel *action4level2markuppanel.Panel
	if action4Level2MarkupPanel, err = action4level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action4Level3MarkupPanel *action4level3markuppanel.Panel
	if action4Level3MarkupPanel, err = action4level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action4Level4MarkupPanel *action4level4markuppanel.Panel
	if action4Level4MarkupPanel, err = action4level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action4Level5MarkupPanel *action4level5markuppanel.Panel
	if action4Level5MarkupPanel, err = action4level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action5Level1MarkupPanel *action5level1markuppanel.Panel
	if action5Level1MarkupPanel, err = action5level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action5Level2MarkupPanel *action5level2markuppanel.Panel
	if action5Level2MarkupPanel, err = action5level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action5Level3MarkupPanel *action5level3markuppanel.Panel
	if action5Level3MarkupPanel, err = action5level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action5Level4MarkupPanel *action5level4markuppanel.Panel
	if action5Level4MarkupPanel, err = action5level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}
	var action5Level5MarkupPanel *action5level5markuppanel.Panel
	if action5Level5MarkupPanel, err = action5level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		return
	}

	// 3. Size the app.
	tools.SizeApp()

	// 4. Start each panel's message and event dispatchers.
	action1Level1MarkupPanel.StartDispatchers()
	action1Level2MarkupPanel.StartDispatchers()
	action1Level3MarkupPanel.StartDispatchers()
	action1Level4MarkupPanel.StartDispatchers()
	action1Level5MarkupPanel.StartDispatchers()
	action2Level1MarkupPanel.StartDispatchers()
	action2Level2MarkupPanel.StartDispatchers()
	action2Level3MarkupPanel.StartDispatchers()
	action2Level4MarkupPanel.StartDispatchers()
	action2Level5MarkupPanel.StartDispatchers()
	action3Level1MarkupPanel.StartDispatchers()
	action3Level2MarkupPanel.StartDispatchers()
	action3Level3MarkupPanel.StartDispatchers()
	action3Level4MarkupPanel.StartDispatchers()
	action3Level5MarkupPanel.StartDispatchers()
	action4Level1MarkupPanel.StartDispatchers()
	action4Level2MarkupPanel.StartDispatchers()
	action4Level3MarkupPanel.StartDispatchers()
	action4Level4MarkupPanel.StartDispatchers()
	action4Level5MarkupPanel.StartDispatchers()
	action5Level1MarkupPanel.StartDispatchers()
	action5Level2MarkupPanel.StartDispatchers()
	action5Level3MarkupPanel.StartDispatchers()
	action5Level4MarkupPanel.StartDispatchers()
	action5Level5MarkupPanel.StartDispatchers()

	// 5. Start each panel's initial calls.
	action1Level1MarkupPanel.InitialJobs()
	action1Level2MarkupPanel.InitialJobs()
	action1Level3MarkupPanel.InitialJobs()
	action1Level4MarkupPanel.InitialJobs()
	action1Level5MarkupPanel.InitialJobs()
	action2Level1MarkupPanel.InitialJobs()
	action2Level2MarkupPanel.InitialJobs()
	action2Level3MarkupPanel.InitialJobs()
	action2Level4MarkupPanel.InitialJobs()
	action2Level5MarkupPanel.InitialJobs()
	action3Level1MarkupPanel.InitialJobs()
	action3Level2MarkupPanel.InitialJobs()
	action3Level3MarkupPanel.InitialJobs()
	action3Level4MarkupPanel.InitialJobs()
	action3Level5MarkupPanel.InitialJobs()
	action4Level1MarkupPanel.InitialJobs()
	action4Level2MarkupPanel.InitialJobs()
	action4Level3MarkupPanel.InitialJobs()
	action4Level4MarkupPanel.InitialJobs()
	action4Level5MarkupPanel.InitialJobs()
	action5Level1MarkupPanel.InitialJobs()
	action5Level2MarkupPanel.InitialJobs()
	action5Level3MarkupPanel.InitialJobs()
	action5Level4MarkupPanel.InitialJobs()
	action5Level5MarkupPanel.InitialJobs()

	return
}
