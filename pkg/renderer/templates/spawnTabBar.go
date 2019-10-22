package templates

// SpawnTabBarPrepare is the genereric renderer spawn tab bar prepare.go template.
const SpawnTabBarPrepare = `{{$Dot := .}}// +build js, wasm

package {{call .PackageNameCase .TabBarName}}

import (
{{ range .Imports }}	"{{$Dot.ApplicationGitPath}}{{.}}"
{{end}})

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

const (
	tabBarID = "{{.TabBarID}}"
)

// Prepare prepares the tab bar and it's panels and templates to be spawned.
// This is called once by package main when the application starts.
func Prepare(quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending,
	vtools *viewtools.Tools, njs *notjs.NotJS, help *paneling.Help) {
{{ range .TabNames }}
	{{ call $Dot.PackageNameCase . }}.Prepare(quitChan, eojChan, receiveChan, sendChan, vtools, njs, help){{end}}
}
`
