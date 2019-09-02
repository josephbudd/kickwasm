package templates

// LocalBoltStoreRecordGo is the domain/records/ <store>.go template.
const LocalBoltStoreRecordGo = `{{$Dot := .}}package record

/*

	TODO:

	You need to complete this local bolt store record definition.

*/

// {{.Store}} is the local bolt store {{.Store}} record.
type {{.Store}} struct {
	ID uint64
}

// New{{.Store}} constructs a new local bolt store {{.Store}}.
func New{{.Store}}() *{{.Store}} {
	v := &{{.Store}}{}
	return v
}
`

// RemoteDatabaseRecordGo is the domain/records/ <store>.go template.
const RemoteDatabaseRecordGo = `{{$Dot := .}}package record

/*

	TODO:

	You need to complete this remote record definition.

*/

// {{.Store}} is the remote {{.Store}} record.
type {{.Store}} struct {}

// New{{.Store}} constructs a new remote {{.Store}} record.
func New{{.Store}}() *{{.Store}} {
	v := &{{.Store}}{}
	return v
}
`
