package storer

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

// ContactStorer defines the behavior of a Contact database.
type ContactStorer interface {

	// Open opens the database.
	// Returns the error.
	Open() error

	// Close closes the database.
	// Returns the error.
	Close() error

	// GetContact retrieves one *types.ContactRecord from the db.
	// Param id [in] is the record id.
	// Returns a record pointer and error.
	// Returns (nil, nil) when the record is not found.
	GetContact(id uint64) (*types.ContactRecord, error)

	// GetContacts retrieves all of the *types.ContactRecords from the db.
	// Returns a slice of record pointers and error.
	// When no records found, the slice length is 0 and error is nil.
	GetContacts() ([]*types.ContactRecord, error)

	// UpdateContact updates the *types.ContactRecord in the database.
	// Param r [in-out] the pointer to the record to be updated.
	// If r is a new record then r.ID is updated as well.
	// Returns the error.
	UpdateContact(r *types.ContactRecord) error

	// RemoveContact removes the types.ContactRecord from the database.
	// Param id [in] the key of the record to be removed.
	// Returns the error.
	RemoveContact(id uint64) error
}

