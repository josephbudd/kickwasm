package storer

import (
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
)

// ColorStorer defines the behavior of a Color database.
type ColorStorer interface {

	// Open opens the database.
	// Returns the error.
	Open() error

	// Close closes the database.
	// Returns the error.
	Close() error

	// GetColor retrieves one *types.ColorRecord from the db.
	// Param id [in] is the record id.
	// Returns a record pointer and error.
	// Returns (nil, nil) when the record is not found.
	GetColor(id uint64) (*types.ColorRecord, error)

	// GetColors retrieves all of the *types.ColorRecords from the db.
	// Returns a slice of record pointers and error.
	// When no records found, the slice length is 0 and error is nil.
	GetColors() ([]*types.ColorRecord, error)

	// UpdateColor updates the *types.ColorRecord in the database.
	// Param r [in-out] the pointer to the record to be updated.
	// If r is a new record then r.ID is updated as well.
	// Returns the error.
	UpdateColor(r *types.ColorRecord) error

	// RemoveColor removes the types.ColorRecord from the database.
	// Param id [in] the key of the record to be removed.
	// Returns the error.
	RemoveColor(id uint64) error
}

