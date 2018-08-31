package templates

// RecordsRecordsGo is the mainprocess/ports/records/records.go template.
const RecordsRecordsGo = `{{$Dot := .}}package records

// These records are here in this package because they do not belong to any database.

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
