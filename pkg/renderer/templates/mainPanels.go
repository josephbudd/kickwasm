package templates

// MainDoPanelsGo is ./panels.go.
const MainDoPanelsGo = `{{$Dot := .}}package main

import ({{range .Imports}}
	"{{.}}"{{end}}
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

func doPanels(quitCh chan struct{}, tools *viewtools.Tools, callMap map[types.CallID]caller.Renderer, notJS *notjs.NotJS, helper panelHelper.Helper) (err error) {
	// 1. Construct the panel code.{{range $name, $path := .PanelNamePath}}
	var {{call $Dot.LowerCamelCase $name}} *{{call $Dot.PackageNameCase $name}}.Panel
	if {{call $Dot.LowerCamelCase $name}}, err = {{call $Dot.PackageNameCase $name}}.NewPanel(quitCh, tools, notJS, callMap, helper); err != nil {
		return
	}{{end}}

	// 2. Size the app.
	tools.SizeApp()

	// 3. Start each panel's initial calls.{{range $name, $path := .PanelNamePath}}
	{{call $Dot.LowerCamelCase $name}}.InitialCalls(){{end}}

	return
}
`
