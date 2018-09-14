package types

/*

	TODO:

	You need to complete these record definitions.

*/

// ContactRecord is a Contact record.
type ContactRecord struct {
	ID uint64
}

// NewContactRecord constructs a new Contact record.
func NewContactRecord() *ContactRecord {
	v := &ContactRecord{}
	return v
}
