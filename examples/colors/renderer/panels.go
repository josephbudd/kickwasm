package main

import (
	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/interfaces/panelHelper"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/ColorsButton/Service1Level2ButtonPanel/ColorsButton/Service1Level3ButtonPanel/ColorsButton/Service1Level4ButtonPanel/ColorsButton/Service1Level5ButtonPanel/ContentButton/Service1Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/ColorsButton/Service1Level2ButtonPanel/ColorsButton/Service1Level3ButtonPanel/ColorsButton/Service1Level4ButtonPanel/ContentButton/Service1Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/ColorsButton/Service1Level2ButtonPanel/ColorsButton/Service1Level3ButtonPanel/ContentButton/Service1Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/ColorsButton/Service1Level2ButtonPanel/ContentButton/Service1Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service1Button/Service1Level1ButtonPanel/ContentButton/Service1Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/ColorsButton/Service2Level2ButtonPanel/ColorsButton/Service2Level3ButtonPanel/ColorsButton/Service2Level4ButtonPanel/ColorsButton/Service2Level5ButtonPanel/ContentButton/Service2Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/ColorsButton/Service2Level2ButtonPanel/ColorsButton/Service2Level3ButtonPanel/ColorsButton/Service2Level4ButtonPanel/ContentButton/Service2Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/ColorsButton/Service2Level2ButtonPanel/ColorsButton/Service2Level3ButtonPanel/ContentButton/Service2Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/ColorsButton/Service2Level2ButtonPanel/ContentButton/Service2Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service2Button/Service2Level1ButtonPanel/ContentButton/Service2Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/ColorsButton/Service3Level2ButtonPanel/ColorsButton/Service3Level3ButtonPanel/ColorsButton/Service3Level4ButtonPanel/ColorsButton/Service3Level5ButtonPanel/ContentButton/Service3Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/ColorsButton/Service3Level2ButtonPanel/ColorsButton/Service3Level3ButtonPanel/ColorsButton/Service3Level4ButtonPanel/ContentButton/Service3Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/ColorsButton/Service3Level2ButtonPanel/ColorsButton/Service3Level3ButtonPanel/ContentButton/Service3Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/ColorsButton/Service3Level2ButtonPanel/ContentButton/Service3Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service3Button/Service3Level1ButtonPanel/ContentButton/Service3Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/ColorsButton/Service4Level2ButtonPanel/ColorsButton/Service4Level3ButtonPanel/ColorsButton/Service4Level4ButtonPanel/ColorsButton/Service4Level5ButtonPanel/ContentButton/Service4Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/ColorsButton/Service4Level2ButtonPanel/ColorsButton/Service4Level3ButtonPanel/ColorsButton/Service4Level4ButtonPanel/ContentButton/Service4Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/ColorsButton/Service4Level2ButtonPanel/ColorsButton/Service4Level3ButtonPanel/ContentButton/Service4Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/ColorsButton/Service4Level2ButtonPanel/ContentButton/Service4Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service4Button/Service4Level1ButtonPanel/ContentButton/Service4Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/ColorsButton/Service5Level2ButtonPanel/ColorsButton/Service5Level3ButtonPanel/ColorsButton/Service5Level4ButtonPanel/ColorsButton/Service5Level5ButtonPanel/ContentButton/Service5Level5MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/ColorsButton/Service5Level2ButtonPanel/ColorsButton/Service5Level3ButtonPanel/ColorsButton/Service5Level4ButtonPanel/ContentButton/Service5Level4MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/ColorsButton/Service5Level2ButtonPanel/ColorsButton/Service5Level3ButtonPanel/ContentButton/Service5Level3MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/ColorsButton/Service5Level2ButtonPanel/ContentButton/Service5Level2MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/panels/Service5Button/Service5Level1ButtonPanel/ContentButton/Service5Level1MarkupPanel"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

func doPanels(quitCh chan struct{}, tools *viewtools.Tools, callMap map[types.CallID]caller.Renderer, notJS *notjs.NotJS, helper panelHelper.Helper) (err error) {
	// 1. Construct the panel code.
	var service1Level1MarkupPanel *service1level1markuppanel.Panel
	if service1Level1MarkupPanel, err = service1level1markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service1Level2MarkupPanel *service1level2markuppanel.Panel
	if service1Level2MarkupPanel, err = service1level2markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service1Level3MarkupPanel *service1level3markuppanel.Panel
	if service1Level3MarkupPanel, err = service1level3markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service1Level4MarkupPanel *service1level4markuppanel.Panel
	if service1Level4MarkupPanel, err = service1level4markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service1Level5MarkupPanel *service1level5markuppanel.Panel
	if service1Level5MarkupPanel, err = service1level5markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level1MarkupPanel *service2level1markuppanel.Panel
	if service2Level1MarkupPanel, err = service2level1markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level2MarkupPanel *service2level2markuppanel.Panel
	if service2Level2MarkupPanel, err = service2level2markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level3MarkupPanel *service2level3markuppanel.Panel
	if service2Level3MarkupPanel, err = service2level3markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level4MarkupPanel *service2level4markuppanel.Panel
	if service2Level4MarkupPanel, err = service2level4markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level5MarkupPanel *service2level5markuppanel.Panel
	if service2Level5MarkupPanel, err = service2level5markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level1MarkupPanel *service3level1markuppanel.Panel
	if service3Level1MarkupPanel, err = service3level1markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level2MarkupPanel *service3level2markuppanel.Panel
	if service3Level2MarkupPanel, err = service3level2markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level3MarkupPanel *service3level3markuppanel.Panel
	if service3Level3MarkupPanel, err = service3level3markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level4MarkupPanel *service3level4markuppanel.Panel
	if service3Level4MarkupPanel, err = service3level4markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level5MarkupPanel *service3level5markuppanel.Panel
	if service3Level5MarkupPanel, err = service3level5markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level1MarkupPanel *service4level1markuppanel.Panel
	if service4Level1MarkupPanel, err = service4level1markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level2MarkupPanel *service4level2markuppanel.Panel
	if service4Level2MarkupPanel, err = service4level2markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level3MarkupPanel *service4level3markuppanel.Panel
	if service4Level3MarkupPanel, err = service4level3markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level4MarkupPanel *service4level4markuppanel.Panel
	if service4Level4MarkupPanel, err = service4level4markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level5MarkupPanel *service4level5markuppanel.Panel
	if service4Level5MarkupPanel, err = service4level5markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level1MarkupPanel *service5level1markuppanel.Panel
	if service5Level1MarkupPanel, err = service5level1markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level2MarkupPanel *service5level2markuppanel.Panel
	if service5Level2MarkupPanel, err = service5level2markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level3MarkupPanel *service5level3markuppanel.Panel
	if service5Level3MarkupPanel, err = service5level3markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level4MarkupPanel *service5level4markuppanel.Panel
	if service5Level4MarkupPanel, err = service5level4markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level5MarkupPanel *service5level5markuppanel.Panel
	if service5Level5MarkupPanel, err = service5level5markuppanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}

	// 2. Size the app.
	tools.SizeApp()

	// 3. Start each panel's initial calls.
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
