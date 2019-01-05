package main

import (
	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
	"github.com/josephbudd/kickwasm/examples/colors/renderer/interfaces/panelHelper"
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
	"github.com/josephbudd/kickwasm/examples/colors/site/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/site/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

func doPanels(quitCh chan struct{}, tools *viewtools.Tools, callMap map[types.CallID]caller.Renderer, notJS *notjs.NotJS, helper panelHelper.Helper) (err error) {
	// 1. Construct the panel code.
	var service1Level1MarkupPanel *Service1Level1MarkupPanel.Panel
	if service1Level1MarkupPanel, err = Service1Level1MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service1Level2MarkupPanel *Service1Level2MarkupPanel.Panel
	if service1Level2MarkupPanel, err = Service1Level2MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service1Level3MarkupPanel *Service1Level3MarkupPanel.Panel
	if service1Level3MarkupPanel, err = Service1Level3MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service1Level4MarkupPanel *Service1Level4MarkupPanel.Panel
	if service1Level4MarkupPanel, err = Service1Level4MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service1Level5MarkupPanel *Service1Level5MarkupPanel.Panel
	if service1Level5MarkupPanel, err = Service1Level5MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level1MarkupPanel *Service2Level1MarkupPanel.Panel
	if service2Level1MarkupPanel, err = Service2Level1MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level2MarkupPanel *Service2Level2MarkupPanel.Panel
	if service2Level2MarkupPanel, err = Service2Level2MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level3MarkupPanel *Service2Level3MarkupPanel.Panel
	if service2Level3MarkupPanel, err = Service2Level3MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level4MarkupPanel *Service2Level4MarkupPanel.Panel
	if service2Level4MarkupPanel, err = Service2Level4MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service2Level5MarkupPanel *Service2Level5MarkupPanel.Panel
	if service2Level5MarkupPanel, err = Service2Level5MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level1MarkupPanel *Service3Level1MarkupPanel.Panel
	if service3Level1MarkupPanel, err = Service3Level1MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level2MarkupPanel *Service3Level2MarkupPanel.Panel
	if service3Level2MarkupPanel, err = Service3Level2MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level3MarkupPanel *Service3Level3MarkupPanel.Panel
	if service3Level3MarkupPanel, err = Service3Level3MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level4MarkupPanel *Service3Level4MarkupPanel.Panel
	if service3Level4MarkupPanel, err = Service3Level4MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service3Level5MarkupPanel *Service3Level5MarkupPanel.Panel
	if service3Level5MarkupPanel, err = Service3Level5MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level1MarkupPanel *Service4Level1MarkupPanel.Panel
	if service4Level1MarkupPanel, err = Service4Level1MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level2MarkupPanel *Service4Level2MarkupPanel.Panel
	if service4Level2MarkupPanel, err = Service4Level2MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level3MarkupPanel *Service4Level3MarkupPanel.Panel
	if service4Level3MarkupPanel, err = Service4Level3MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level4MarkupPanel *Service4Level4MarkupPanel.Panel
	if service4Level4MarkupPanel, err = Service4Level4MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service4Level5MarkupPanel *Service4Level5MarkupPanel.Panel
	if service4Level5MarkupPanel, err = Service4Level5MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level1MarkupPanel *Service5Level1MarkupPanel.Panel
	if service5Level1MarkupPanel, err = Service5Level1MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level2MarkupPanel *Service5Level2MarkupPanel.Panel
	if service5Level2MarkupPanel, err = Service5Level2MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level3MarkupPanel *Service5Level3MarkupPanel.Panel
	if service5Level3MarkupPanel, err = Service5Level3MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level4MarkupPanel *Service5Level4MarkupPanel.Panel
	if service5Level4MarkupPanel, err = Service5Level4MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}
	var service5Level5MarkupPanel *Service5Level5MarkupPanel.Panel
	if service5Level5MarkupPanel, err = Service5Level5MarkupPanel.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
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
