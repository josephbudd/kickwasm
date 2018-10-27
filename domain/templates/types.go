package templates

// TypesCallsGo the file at domain/types/calls.go
const TypesCallsGo = `package types

import "{{.ApplicationGitPath}}{{.ImportDomainInterfacesCallers}}"

// CallID is the unique id for a map[CallID]caller.Renderer or a map[CallID]caller.MainProcessor
type CallID uint64

// Payload is a the information transported between the main process and the renderer.
type Payload struct {
	Procedure CallID
	Params    string
}

`

// TypesRecordsGo is the domain/types/records.go template.
const TypesRecordsGo = `{{$Dot := .}}package types

/*

	TODO:

	You need to complete these record definitions.

*/{{range .Stores}}

// {{.}}Record is a {{.}} record.
type {{.}}Record struct {
	ID uint64
}{{end}}{{range .Stores}}

// New{{.}}Record constructs a new {{.}} record.
func New{{.}}Record() *{{.}}Record {
	v := &{{.}}Record{}
	return v
}{{end}}

`

// TypesLogGo is the domain/types/log.go template.
const TypesLogGo = `package types

// RendererToMainProcessLogCallParams are the Log function parameters that the renderer sends to the main process.
type RendererToMainProcessLogCallParams struct {
	Level   uint64
	Message string
}

// MainProcessToRendererLogCallParams are the Log function parameters that the main process sends to the renderer.
type MainProcessToRendererLogCallParams struct {
	Error        bool
	ErrorMessage string
	Level        uint64
}

`

// TypesGetAboutParamsGo is the domain/types/getAboutCallParams.go template.
const TypesGetAboutParamsGo = `package types

import (
	"{{.ApplicationGitPath}}{{.ImportMainProcessServicesAbout}}"
)

// MainProcessToRendererGetAboutParams are the GetAbout function parameters that the main process sends to the renderer.
type MainProcessToRendererGetAboutParams struct {
	Error        bool
	ErrorMessage string
	Version      string
	Releases     map[string]map[string][]string
	Contributors map[string]string
	Credits      []about.Credit
	Licenses     []about.License
}

/*

	For renderer code see {{.ApplicationGitPath}}{{.OutputRendererPanels}}/AboutButton/AboutPanel/caller.go

*/

`
