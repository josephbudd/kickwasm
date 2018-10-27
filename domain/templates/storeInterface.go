package templates

// StorerGo is the template /domain/interfaces/storer/storer.go
const StorerGo = `{{$Dot := .}}package storer

import (
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// {{.Store}}Storer defines the behavior of a {{.Store}} database.
type {{.Store}}Storer interface {

	// Open opens the database.
	// Returns the error.
	Open() error

	// Close closes the database.
	// Returns the error.
	Close() error

	// Get{{.Store}} retrieves one *types.{{.Store}}Record from the db.
	// Param id [in] is the record id.
	// Returns a record pointer and error.
	// Returns (nil, nil) when the record is not found.
	Get{{.Store}}(id uint64) (*types.{{.Store}}Record, error)

	// Get{{.Store}}s retrieves all of the *types.{{.Store}}Records from the db.
	// Returns a slice of record pointers and error.
	// When no records found, the slice length is 0 and error is nil.
	Get{{.Store}}s() ([]*types.{{.Store}}Record, error)

	// Update{{.Store}} updates the *types.{{.Store}}Record in the database.
	// Param r [in-out] the pointer to the record to be updated.
	// If r is a new record then r.ID is updated as well.
	// Returns the error.
	Update{{.Store}}(r *types.{{.Store}}Record) error

	// Remove{{.Store}} removes the types.{{.Store}}Record from the database.
	// Param id [in] the key of the record to be removed.
	// Returns the error.
	Remove{{.Store}}(id uint64) error
}

`
