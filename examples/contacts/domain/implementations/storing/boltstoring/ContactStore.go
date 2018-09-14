package boltstoring

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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
func NewContactBoltDB(db *bolt.DB, path string, perms os.FileMode) *ContactBoltDB {
	return &ContactBoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
}

// ContactBoltDB implements ContactStorer
// which is defined in domain/types/records.go

// Open opens the database.
// Returns the error.
func (contactdb *ContactBoltDB) Open() error {
	// the bolt db is already open
	return nil
}

// Close closes the database.
// Returns the error.
func (contactdb *ContactBoltDB) Close() error {
	return contactdb.DB.Close()
}

// GetContact retrieves the types.ContactRecord from the db.
// Param id [in] is the record id.
// Returns the record and error.
func (contactdb *ContactBoltDB) GetContact(id uint64) (*types.ContactRecord, error) {
	var r types.ContactRecord
	ids := fmt.Sprintf("%d", id)
	er := contactdb.DB.View(func(tx *bolt.Tx) error {
		bucketname := []byte(contactBucketName)
		bucket := tx.Bucket(bucketname)
		if bucket != nil {
			bb := bucket.Get([]byte(ids))
			if bb != nil {
				// found
				err := json.Unmarshal(bb, &r)
				if err == nil {
					r.ID = id
				}
				return err
			}
		}
		// no bucket or not found
		return errNotFound
	})
	if er == nil {
		// found
		return &r, nil
	} else if er == errNotFound {
		// not found
		return nil, nil
	}
	return nil, er
}

// GetContacts retrieves all of the types.ContactRecord from the db.
// If there are no types.ContactRecords in the db then it calls contactdb.initialize().
// See contactdb.initialize().
// Returns the records and error.
func (contactdb *ContactBoltDB) GetContacts() ([]*types.ContactRecord, error) {
	if rr, err := contactdb.getContacts(); len(rr) > 0 && err == nil {
		return rr, err
	}
	contactdb.initialize()
	return contactdb.getContacts()
}

func (contactdb *ContactBoltDB) getContacts() ([]*types.ContactRecord, error) {
	rr := make([]*types.ContactRecord, 0, 5)
	er := contactdb.DB.View(func(tx *bolt.Tx) error {
		bucketname := []byte(contactBucketName)
		bucket := tx.Bucket(bucketname)
		if bucket != nil {
			c := bucket.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				r := types.NewContactRecord()
				err := json.Unmarshal(v, r)
				if err != nil {
					return err
				}
				r.ID, err = strconv.ParseUint(string(k), 10, 64)
				if err != nil {
					return err
				}
				rr = append(rr, r)
			}
		}
		return nil
	})
	return rr, er
}

// UpdateContact updates the types.ContactRecord in the database.
// Param record [in-out] the record to be updated.
// if record is new then record.ID is updated as well.
// Returns the error.
func (contactdb *ContactBoltDB) UpdateContact(r *types.ContactRecord) error {
	return contactdb.updateContactBucket(r)
}

// RemoveContact removes the types.ContactRecord from the database.
// Param id [in] the key of the record to be removed.
// Returns the error.
func (contactdb *ContactBoltDB) RemoveContact(id uint64) error {
	return contactdb.DB.Update(func(tx *bolt.Tx) error {
		bucketname := []byte(contactBucketName)
		bucket := tx.Bucket(bucketname)
		if bucket != nil {
			idbb := []byte(fmt.Sprintf("%d", id))
			col := bucket.Get(idbb)
			if col != nil {
				err := bucket.Delete(idbb)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// updates the types.ContactRecord in the database.
// Param record [in-out] the record to be updated
func (contactdb *ContactBoltDB) updateContactBucket(r *types.ContactRecord) error {
	return contactdb.DB.Update(func(tx *bolt.Tx) error {
		bucketname := []byte(contactBucketName)
		bucket, err := tx.CreateBucketIfNotExists(bucketname)
		if err == nil {
			if r.ID == 0 {
				id, err := bucket.NextSequence()
				if err == nil {
					r.ID = id
				}
			}
			if err == nil {
				bb, err := json.Marshal(r)
				if err == nil {
					idbb := []byte(fmt.Sprintf("%d", r.ID))
					err = bucket.Put(idbb, bb)
				}
			}
		}
		return err
	})
}

// initialize is only useful if you want to add the default records to the db.
// otherwise you don't need it to do anything.
func (contactdb *ContactBoltDB) initialize() error {
	/*
		example code:

		defaults := somepackage.GetContactDefaults()
		for _, default := range defaults {
			r := types.NewContactRecord()
			r.Name = default.Name
			r.Price = default.Price
			r.SKU = default.SKU
			err := contactdb.updateContactBucket(r)
			if err != nil {
				return err
			}
		}
	*/
	return nil
}
