package repoi

import (
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
)

// ContactRepoI defines the behavior of a Contact database.
type ContactRepoI interface {

	// Open opens the database.
	// Returns the error.
	Open() error

	// Close closes the database.
	// Returns the error.
	Close() error

	// GetContact retrieves one *records.ContactRecord from the db.
	// Param id [in] is the record id.
	// Returns a record pointer and error.
	// Returns (nil, nil) when the record is not found.
	GetContact(id uint64) (*records.ContactRecord, error)

	// GetContacts retrieves all of the *records.ContactRecords from the db.
	// Returns a slice of record pointers and error.
	// When no records found, the slice length is 0 and error is nil.
	GetContacts() ([]*records.ContactRecord, error)

	// UpdateContact updates the *records.ContactRecord in the database.
	// Param r [in-out] the pointer to the record to be updated.
	// If r is a new record then r.ID is updated as well.
	// Returns the error.
	UpdateContact(r *records.ContactRecord) error

	// RemoveContact removes the records.ContactRecord from the database.
	// Param id [in] the key of the record to be removed.
	// Returns the error.
	RemoveContact(id uint64) error
}
