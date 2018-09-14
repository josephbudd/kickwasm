package templates

// StorerGo is the template /domain/interfaces/storer/storer.go
const StorerGo = `{{$Dot := .}}package storer

import (
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// {{.Repo}}Storer defines the behavior of a {{.Repo}} database.
type {{.Repo}}Storer interface {

	// Open opens the database.
	// Returns the error.
	Open() error

	// Close closes the database.
	// Returns the error.
	Close() error

	// Get{{.Repo}} retrieves one *types.{{.Repo}}Record from the db.
	// Param id [in] is the record id.
	// Returns a record pointer and error.
	// Returns (nil, nil) when the record is not found.
	Get{{.Repo}}(id uint64) (*types.{{.Repo}}Record, error)

	// Get{{.Repo}}s retrieves all of the *types.{{.Repo}}Records from the db.
	// Returns a slice of record pointers and error.
	// When no records found, the slice length is 0 and error is nil.
	Get{{.Repo}}s() ([]*types.{{.Repo}}Record, error)

	// Update{{.Repo}} updates the *types.{{.Repo}}Record in the database.
	// Param r [in-out] the pointer to the record to be updated.
	// If r is a new record then r.ID is updated as well.
	// Returns the error.
	Update{{.Repo}}(r *types.{{.Repo}}Record) error

	// Remove{{.Repo}} removes the types.{{.Repo}}Record from the database.
	// Param id [in] the key of the record to be removed.
	// Returns the error.
	Remove{{.Repo}}(id uint64) error
}

`
