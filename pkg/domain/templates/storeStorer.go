package templates

// StoreStorerGo is the template /domain/interfaces/storer/storer.go
const StoreStorerGo = `{{$Dot := .}}package storer

import (
	"{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"
)

// {{.Store}}Storer defines the behavior (API) of a store of {{.ImportDomainStoreRecord}}.{{.Store}} records.
type {{.Store}}Storer interface {

	// Open opens the data-store.
	// Returns the error.
	Open() (err error)

	// Close closes the data-store.
	// Returns the error.
	Close() (err error)

	// Get retrieves one *record.{{.Store}} from the data-store.
	// Param id is the record ID.
	// Returns a record pointer and error.
	// When no record is found, the returned record pointer is nil and the returned error is nil.
	Get(id uint64) (r *record.{{.Store}}, err error)

	// GetAll retrieves all of the *record.{{.Store}} records from the data-store.
	// Returns a slice of record pointers and error.
	// When no records are found, the returned slice length is 0 and the returned error is nil.
	GetAll() (rr []*record.{{.Store}}, err error)

	// Update updates the *record.{{.Store}} in the data-store.
	// Param r is the pointer to the record to be updated.
	// If r is a new record then r.ID is updated as well.
	// Returns the error.
	Update(r *record.{{.Store}}) (err error)

	// Remove removes the record.{{.Store}} from the data-store.
	// Param id is the record ID of the record to be removed.
	// Returns the error.
	Remove(id uint64) (err error)
}
`
