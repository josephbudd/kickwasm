package templates

// TypesCallsGo the file at domain/types/calls.go
const TypesCallsGo = `package types

import "{{.ApplicationGitPath}}{{.ImportDomainInterfacesCallers}}"

// CallID is the unique id for a RendererCallMap or a MainProcessCallsMap
type CallID uint64

// Payload is a the information transported between the main process and the renderer.
type Payload struct {
	Procedure CallID
	Params    string
}

// RendererCallMap is each call id mapped to its Renderer interface.
type RendererCallMap map[CallID]caller.Renderer

// MainProcessCallsMap is each call id mapped to its MainProcessor interfaces.
type MainProcessCallsMap map[CallID]caller.MainProcesser

`

// TypesRecordsGo is the domain/types/records.go template.
const TypesRecordsGo = `{{$Dot := .}}package types

/*

	TODO:

	You need to complete these record definitions.

*/{{range .Repos}}

// {{.}}Record is a {{.}} record.
type {{.}}Record struct {
	ID uint64
}{{end}}{{range .Repos}}

// New{{.}}Record constructs a new {{.}} record.
func New{{.}}Record() *{{.}}Record {
	v := &{{.}}Record{}
	return v
}{{end}}
`
