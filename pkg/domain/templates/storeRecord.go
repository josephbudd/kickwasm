package templates

// StoreRecordGo is the domain/records/ <store>.go template.
const StoreRecordGo = `{{$Dot := .}}package record

/*

	TODO:

	You need to complete this record definition.

*/

// {{.Store}} is a {{.Store}} record.
type {{.Store}} struct {
	ID uint64
}

// New{{.Store}} constructs a new {{.Store}} record.
func New{{.Store}}() *{{.Store}} {
	v := &{{.Store}}{}
	return v
}
`
