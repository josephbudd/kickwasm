package types

/*

	TODO:

	You need to complete these record definitions.

*/

// ColorRecord is a Color record.
type ColorRecord struct {
	ID uint64
}

// NewColorRecord constructs a new Color record.
func NewColorRecord() *ColorRecord {
	v := &ColorRecord{}
	return v
}

