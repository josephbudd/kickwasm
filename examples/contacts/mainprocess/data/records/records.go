package records

// These records are here in this package because they do not belong to any database.

/*

	TODO:

	You need to complete these record definitions.

*/

// ContactRecord is a Contact record.
type ContactRecord struct {
	ID       uint64
	Name     string
	Address1 string
	Address2 string
	City     string
	State    string
	Zip      string
	Phone    string
	Email    string
	Social   string
}

// NewContactRecord constructs a new Contact record.
func NewContactRecord() *ContactRecord {
	v := &ContactRecord{}
	return v
}
