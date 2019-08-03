package main

import (
	"github.com/josephbudd/kickwasm/examples/colors/renderer/lpc"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/paneling"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/Service1Level1ContentButton/Service1Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/Service1ToLevel2ColorsButton/Service1Level2ButtonPanel/Service1Level2ContentButton/Service1Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/Service1ToLevel2ColorsButton/Service1Level2ButtonPanel/Service1ToLevel3ColorsButton/Service1Level3ButtonPanel/Service1Level3ContentButton/Service1Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/Service1ToLevel2ColorsButton/Service1Level2ButtonPanel/Service1ToLevel3ColorsButton/Service1Level3ButtonPanel/Service1ToLevel4ColorsButton/Service1Level4ButtonPanel/Service1Level4ContentButton/Service1Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/Service1ToLevel2ColorsButton/Service1Level2ButtonPanel/Service1ToLevel3ColorsButton/Service1Level3ButtonPanel/Service1ToLevel4ColorsButton/Service1Level4ButtonPanel/Service1ToLevel5ColorsButton/Service1Level5ButtonPanel/Service1Level5ContentButton/Service1Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/Service2Level1ContentButton/Service2Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/Service2ToLevel2ColorsButton/Service2Level2ButtonPanel/Service2Level2ContentButton/Service2Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/Service2ToLevel2ColorsButton/Service2Level2ButtonPanel/Service2ToLevel3ColorsButton/Service2Level3ButtonPanel/Service2Level3ContentButton/Service2Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/Service2ToLevel2ColorsButton/Service2Level2ButtonPanel/Service2ToLevel3ColorsButton/Service2Level3ButtonPanel/Service2ToLevel4ColorsButton/Service2Level4ButtonPanel/Service2Level4ContentButton/Service2Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/Service2ToLevel2ColorsButton/Service2Level2ButtonPanel/Service2ToLevel3ColorsButton/Service2Level3ButtonPanel/Service2ToLevel4ColorsButton/Service2Level4ButtonPanel/Service2ToLevel5ColorsButton/Service2Level5ButtonPanel/Service2Level5ContentButton/Service2Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/Service3Level1ContentButton/Service3Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/Service3ToLevel2ColorsButton/Service3Level2ButtonPanel/Service3Level2ContentButton/Service3Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/Service3ToLevel2ColorsButton/Service3Level2ButtonPanel/Service3ToLevel3ColorsButton/Service3Level3ButtonPanel/Service3Level3ContentButton/Service3Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/Service3ToLevel2ColorsButton/Service3Level2ButtonPanel/Service3ToLevel3ColorsButton/Service3Level3ButtonPanel/Service3ToLevel4ColorsButton/Service3Level4ButtonPanel/Service3Level4ContentButton/Service3Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/Service3ToLevel2ColorsButton/Service3Level2ButtonPanel/Service3ToLevel3ColorsButton/Service3Level3ButtonPanel/Service3ToLevel4ColorsButton/Service3Level4ButtonPanel/Service3ToLevel5ColorsButton/Service3Level5ButtonPanel/Service3Level5ContentButton/Service3Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/Service4Level1ContentButton/Service4Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/Service4ToLevel2ColorsButton/Service4Level2ButtonPanel/Service4Level2ContentButton/Service4Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/Service4ToLevel2ColorsButton/Service4Level2ButtonPanel/Service4ToLevel3ColorsButton/Service4Level3ButtonPanel/Service4Level3ContentButton/Service4Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/Service4ToLevel2ColorsButton/Service4Level2ButtonPanel/Service4ToLevel3ColorsButton/Service4Level3ButtonPanel/Service4ToLevel4ColorsButton/Service4Level4ButtonPanel/Service4Level4ContentButton/Service4Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/Service4ToLevel2ColorsButton/Service4Level2ButtonPanel/Service4ToLevel3ColorsButton/Service4Level3ButtonPanel/Service4ToLevel4ColorsButton/Service4Level4ButtonPanel/Service4ToLevel5ColorsButton/Service4Level5ButtonPanel/Service4Level5ContentButton/Service4Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/Service5Level1ContentButton/Service5Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/Service5ToLevel2ColorsButton/Service5Level2ButtonPanel/Service5Level2ContentButton/Service5Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/Service5ToLevel2ColorsButton/Service5Level2ButtonPanel/Service5ToLevel3ColorsButton/Service5Level3ButtonPanel/Service5Level3ContentButton/Service5Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/Service5ToLevel2ColorsButton/Service5Level2ButtonPanel/Service5ToLevel3ColorsButton/Service5Level3ButtonPanel/Service5ToLevel4ColorsButton/Service5Level4ButtonPanel/Service5Level4ContentButton/Service5Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/Service5ToLevel2ColorsButton/Service5Level2ButtonPanel/Service5ToLevel3ColorsButton/Service5Level3ButtonPanel/Service5ToLevel4ColorsButton/Service5Level4ButtonPanel/Service5ToLevel5ColorsButton/Service5Level5ButtonPanel/Service5Level5ContentButton/Service5Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

func doPanels(client *lpc.Client, quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending,
	tools *viewtools.Tools, notJS *notjs.NotJS, help *paneling.Help) (err error) {

	// 1. Prepare the spawn panels.

	// 2. Construct the panel code.
	var service1Level1MarkupPanel *service1level1markuppanel.Panel
	if service1Level1MarkupPanel, err = service1level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service1Level2MarkupPanel *service1level2markuppanel.Panel
	if service1Level2MarkupPanel, err = service1level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service1Level3MarkupPanel *service1level3markuppanel.Panel
	if service1Level3MarkupPanel, err = service1level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service1Level4MarkupPanel *service1level4markuppanel.Panel
	if service1Level4MarkupPanel, err = service1level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service1Level5MarkupPanel *service1level5markuppanel.Panel
	if service1Level5MarkupPanel, err = service1level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service2Level1MarkupPanel *service2level1markuppanel.Panel
	if service2Level1MarkupPanel, err = service2level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service2Level2MarkupPanel *service2level2markuppanel.Panel
	if service2Level2MarkupPanel, err = service2level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service2Level3MarkupPanel *service2level3markuppanel.Panel
	if service2Level3MarkupPanel, err = service2level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service2Level4MarkupPanel *service2level4markuppanel.Panel
	if service2Level4MarkupPanel, err = service2level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service2Level5MarkupPanel *service2level5markuppanel.Panel
	if service2Level5MarkupPanel, err = service2level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service3Level1MarkupPanel *service3level1markuppanel.Panel
	if service3Level1MarkupPanel, err = service3level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service3Level2MarkupPanel *service3level2markuppanel.Panel
	if service3Level2MarkupPanel, err = service3level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service3Level3MarkupPanel *service3level3markuppanel.Panel
	if service3Level3MarkupPanel, err = service3level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service3Level4MarkupPanel *service3level4markuppanel.Panel
	if service3Level4MarkupPanel, err = service3level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service3Level5MarkupPanel *service3level5markuppanel.Panel
	if service3Level5MarkupPanel, err = service3level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service4Level1MarkupPanel *service4level1markuppanel.Panel
	if service4Level1MarkupPanel, err = service4level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service4Level2MarkupPanel *service4level2markuppanel.Panel
	if service4Level2MarkupPanel, err = service4level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service4Level3MarkupPanel *service4level3markuppanel.Panel
	if service4Level3MarkupPanel, err = service4level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service4Level4MarkupPanel *service4level4markuppanel.Panel
	if service4Level4MarkupPanel, err = service4level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service4Level5MarkupPanel *service4level5markuppanel.Panel
	if service4Level5MarkupPanel, err = service4level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service5Level1MarkupPanel *service5level1markuppanel.Panel
	if service5Level1MarkupPanel, err = service5level1markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service5Level2MarkupPanel *service5level2markuppanel.Panel
	if service5Level2MarkupPanel, err = service5level2markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service5Level3MarkupPanel *service5level3markuppanel.Panel
	if service5Level3MarkupPanel, err = service5level3markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service5Level4MarkupPanel *service5level4markuppanel.Panel
	if service5Level4MarkupPanel, err = service5level4markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}
	var service5Level5MarkupPanel *service5level5markuppanel.Panel
	if service5Level5MarkupPanel, err = service5level5markuppanel.NewPanel(quitChan, eojChan, receiveChan, sendChan, tools, notJS, help); err != nil {
		tools.ConsoleLog("doPanels: erorr is " + err.Error())
		return
	}

	// No errors so continue.
	tools.ConsoleLog("doPanels: no erorrs")

	// 3. Size the app.
	tools.SizeApp()

	// 4. Start each panel's listening for the main process.
	service1Level1MarkupPanel.Listen()
	service1Level2MarkupPanel.Listen()
	service1Level3MarkupPanel.Listen()
	service1Level4MarkupPanel.Listen()
	service1Level5MarkupPanel.Listen()
	service2Level1MarkupPanel.Listen()
	service2Level2MarkupPanel.Listen()
	service2Level3MarkupPanel.Listen()
	service2Level4MarkupPanel.Listen()
	service2Level5MarkupPanel.Listen()
	service3Level1MarkupPanel.Listen()
	service3Level2MarkupPanel.Listen()
	service3Level3MarkupPanel.Listen()
	service3Level4MarkupPanel.Listen()
	service3Level5MarkupPanel.Listen()
	service4Level1MarkupPanel.Listen()
	service4Level2MarkupPanel.Listen()
	service4Level3MarkupPanel.Listen()
	service4Level4MarkupPanel.Listen()
	service4Level5MarkupPanel.Listen()
	service5Level1MarkupPanel.Listen()
	service5Level2MarkupPanel.Listen()
	service5Level3MarkupPanel.Listen()
	service5Level4MarkupPanel.Listen()
	service5Level5MarkupPanel.Listen()

	// 5. Start each panel's initial calls.
	service1Level1MarkupPanel.InitialCalls()
	service1Level2MarkupPanel.InitialCalls()
	service1Level3MarkupPanel.InitialCalls()
	service1Level4MarkupPanel.InitialCalls()
	service1Level5MarkupPanel.InitialCalls()
	service2Level1MarkupPanel.InitialCalls()
	service2Level2MarkupPanel.InitialCalls()
	service2Level3MarkupPanel.InitialCalls()
	service2Level4MarkupPanel.InitialCalls()
	service2Level5MarkupPanel.InitialCalls()
	service3Level1MarkupPanel.InitialCalls()
	service3Level2MarkupPanel.InitialCalls()
	service3Level3MarkupPanel.InitialCalls()
	service3Level4MarkupPanel.InitialCalls()
	service3Level5MarkupPanel.InitialCalls()
	service4Level1MarkupPanel.InitialCalls()
	service4Level2MarkupPanel.InitialCalls()
	service4Level3MarkupPanel.InitialCalls()
	service4Level4MarkupPanel.InitialCalls()
	service4Level5MarkupPanel.InitialCalls()
	service5Level1MarkupPanel.InitialCalls()
	service5Level2MarkupPanel.InitialCalls()
	service5Level3MarkupPanel.InitialCalls()
	service5Level4MarkupPanel.InitialCalls()
	service5Level5MarkupPanel.InitialCalls()

	return
}
