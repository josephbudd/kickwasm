package boltstoring

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

const contactBucketName string = "contact"

// ContactBoltDB is the bolt db for key codes.
type ContactBoltDB struct {
	DB    *bolt.DB
	path  string
	perms os.FileMode
}

// NewContactBoltDB constructs a new ContactBoltDB.
// Param db [in-out] is an open bolt database.
// Returns a pointer to the new ContactBoltDB.
func NewContactBoltDB(db *bolt.DB, path string, perms os.FileMode) (contactdb *ContactBoltDB) {
	contactdb = &ContactBoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
	return
}

// ContactBoltDB implements ContactStorer
// which is defined in domain/types/records.go

// Open opens the database.
// Returns the error.
func (contactdb *ContactBoltDB) Open() (err error) {
	// the bolt db is already open
	return
}

// Close closes the database.
// Returns the error.
func (contactdb *ContactBoltDB) Close() (err error) {
	err = contactdb.DB.Close()
	return
}

// GetContact retrieves the types.ContactRecord from the db.
// Param id [in] is the record id.
// Returns the record and error.
func (contactdb *ContactBoltDB) GetContact(id uint64) (r *types.ContactRecord, err error) {
	ids := fmt.Sprintf("%d", id)
	err = contactdb.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(contactBucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			r = nil
			er = bolt.ErrBucketNotFound
			return
		}
		var rbb []byte
		if rbb = bucket.Get([]byte(ids)); rbb == nil {
			// not found
			r = nil
			er = nil
			return
		}
		r = &types.ContactRecord{}
		if er = json.Unmarshal(rbb, r); er != nil {
			r = nil
			return
		}
		return
	})
	return
}

// GetContacts retrieves all of the types.ContactRecord from the db.
// If there are no types.ContactRecords in the db then it calls contactdb.initialize().
// See contactdb.initialize().
// Returns the records and error.
func (contactdb *ContactBoltDB) GetContacts() (rr []*types.ContactRecord, err error) {
	if rr, err = contactdb.getContacts(); len(rr) == 0 && err != nil {
		contactdb.initialize()
		rr, err = contactdb.getContacts()
	}
	return
}

func (contactdb *ContactBoltDB) getContacts() (rr []*types.ContactRecord, err error) {
	err = contactdb.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(contactBucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			er = bolt.ErrBucketNotFound
			return
		}
		c := bucket.Cursor()
		rr = make([]*types.ContactRecord, 0, 1024)
		for k, v := c.First(); k != nil; k, v = c.Next() {
			r := types.NewContactRecord()
			if er = json.Unmarshal(v, r); er != nil {
				rr = nil
				return
			}
			rr = append(rr, r)
		}
		return
	})
	return
}

// UpdateContact updates the types.ContactRecord in the database.
// Param record [in-out] the record to be updated.
// if record is new then record.ID is updated as well.
// Returns the error.
func (contactdb *ContactBoltDB) UpdateContact(r *types.ContactRecord) (err error) {
	err = contactdb.updateContactBucket(r)
	return
}

// RemoveContact removes the types.ContactRecord from the database.
// Param id [in] the key of the record to be removed.
// If the record is not found the error is nil.
// Returns the error.
func (contactdb *ContactBoltDB) RemoveContact(id uint64) (err error) {
	err = contactdb.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(contactBucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			er = bolt.ErrBucketNotFound
			return
		}
		idbb := []byte(fmt.Sprintf("%d", id))
		er = bucket.Delete(idbb)
		return
	})
	return
}

// updates the types.ContactRecord in the database.
// Param record [in-out] the record to be updated.
// If the record is new then it's ID is updated.
// Returns the error.
func (contactdb *ContactBoltDB) updateContactBucket(r *types.ContactRecord) (err error) {
	err = contactdb.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(contactBucketName)
		var bucket *bolt.Bucket
		if bucket, er = tx.CreateBucketIfNotExists(bucketname); er != nil {
			return
		}
		if r.ID == 0 {
			if r.ID, er = bucket.NextSequence(); er != nil {
				return
			}
		}
		var rbb []byte
		if rbb, er = json.Marshal(r); er != nil {
			return
		}
		idbb := []byte(fmt.Sprintf("%d", r.ID))
		er = bucket.Put(idbb, rbb)
		return
	})
	return
}

// initialize add 2 default records to the db.
func (contactdb *ContactBoltDB) initialize() (err error) {
	fakes := []*types.ContactRecord{
		{
			Name:     "Joseph Budd",
			Address1: "123 Main St.",
			Address2: "Apt. 1",
			City:     "Small Town",
			State:    "MI",
			Zip:      "12345",
			Phone:    "123-123-1234",
			Email:    "joebudd@smalltown.com",
			Social:   "http://fakebook.com/joebudd\nhttp://www.twitcher.com/joebudd",
		},
		{
			Name:     "Josette Flower",
			Address1: "123 Main St.",
			Address2: "Apt. 2",
			City:     "Small Town",
			State:    "MI",
			Zip:      "12345",
			Phone:    "123-123-2345",
			Email:    "josetteflower@smalltown.com",
			Social:   "http://fakebook.com/josetteflower\nhttp://www.twitcher.com/josetteflower",
		},
		{
			Name:     "Juan Carpenter",
			Address1: "234 Main St.",
			Address2: "Apt. 1",
			City:     "Big Town",
			State:    "MI",
			Zip:      "23456",
			Phone:    "234-123-1234",
			Email:    "johncarpenter@smalltown.com",
			Social:   "http://fakebook.com/johncarpenter\nhttp://www.twitcher.com/johncarpenter",
		},
		{
			Name:     "Juanita Gardner",
			Address1: "234 Main St.",
			Address2: "Apt. 2",
			City:     "Big Town",
			State:    "MI",
			Zip:      "23456",
			Phone:    "234-123-2345",
			Email:    "juanitagardner@smalltown.com",
			Social:   "http://fakebook.com/juanitagardner\nhttp://www.twitcher.com/juanitagardner",
		},
		{
			Name:     "James Driver",
			Address1: "345 Main St.",
			Address2: "Apt. 1",
			City:     "Tiny Town",
			State:    "MI",
			Zip:      "34567",
			Phone:    "345-123-1234",
			Email:    "jamesdriver@smalltown.com",
			Social:   "http://fakebook.com/jamesdriver\nhttp://www.twitcher.com/jamesdriver",
		},
		{
			Name:     "Jamie Navigator",
			Address1: "345 Main St.",
			Address2: "Apt. 2",
			City:     "Tiny Town",
			State:    "MI",
			Zip:      "34567",
			Phone:    "345-123-2345",
			Email:    "jamienavigator@smalltown.com",
			Social:   "http://fakebook.com/jamienavigator\nhttp://www.twitcher.com/jamienavigator",
		},
		{
			Name:     "Jack Drummer",
			Address1: "456 Main St.",
			Address2: "Apt. 1",
			City:     "No Town",
			State:    "MI",
			Zip:      "45678",
			Phone:    "456-123-1234",
			Email:    "jackdrummer@smalltown.com",
			Social:   "http://fakebook.com/jackdrummer\nhttp://www.twitcher.com/jackdrummer",
		},
		{
			Name:     "Jackie Singer",
			Address1: "456 Main St.",
			Address2: "Apt. 2",
			City:     "No Town",
			State:    "MI",
			Zip:      "45678",
			Phone:    "456-123-2345",
			Email:    "jackiesinger@smalltown.com",
			Social:   "http://fakebook.com/jackiesinger\nhttp://www.twitcher.com/jackiesinger",
		},
	}
	for _, fake := range fakes {
		if err = contactdb.updateContactBucket(fake); err != nil {
			return
		}
	}
	return
}
