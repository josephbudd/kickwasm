package boltstoring

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
)

const colorBucketName string = "color"

// ColorBoltDB is the bolt db for key codes.
type ColorBoltDB struct {
	DB    *bolt.DB
	path  string
	perms os.FileMode
}

// NewColorBoltDB constructs a new ColorBoltDB.
// Param db [in-out] is an open bolt database.
// Returns a pointer to the new ColorBoltDB.
func NewColorBoltDB(db *bolt.DB, path string, perms os.FileMode) *ColorBoltDB {
	return &ColorBoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
}

// ColorBoltDB implements ColorStorer
// which is defined in domain/types/records.go

// Open opens the database.
// Returns the error.
func (colordb *ColorBoltDB) Open() error {
	// the bolt db is already open
	return nil
}

// Close closes the database.
// Returns the error.
func (colordb *ColorBoltDB) Close() error {
	return colordb.DB.Close()
}

// GetColor retrieves the types.ColorRecord from the db.
// Param id [in] is the record id.
// Returns the record and error.
func (colordb *ColorBoltDB) GetColor(id uint64) (*types.ColorRecord, error) {
	var r types.ColorRecord
	ids := fmt.Sprintf("%d", id)
	er := colordb.DB.View(func(tx *bolt.Tx) error {
		bucketname := []byte(colorBucketName)
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

// GetColors retrieves all of the types.ColorRecord from the db.
// If there are no types.ColorRecords in the db then it calls colordb.initialize().
// See colordb.initialize().
// Returns the records and error.
func (colordb *ColorBoltDB) GetColors() ([]*types.ColorRecord, error) {
	if rr, err := colordb.getColors(); len(rr) > 0 && err == nil {
		return rr, err
	}
	colordb.initialize()
	return colordb.getColors()
}

func (colordb *ColorBoltDB) getColors() ([]*types.ColorRecord, error) {
	rr := make([]*types.ColorRecord, 0, 5)
	er := colordb.DB.View(func(tx *bolt.Tx) error {
		bucketname := []byte(colorBucketName)
		bucket := tx.Bucket(bucketname)
		if bucket != nil {
			c := bucket.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				r := types.NewColorRecord()
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

// UpdateColor updates the types.ColorRecord in the database.
// Param record [in-out] the record to be updated.
// if record is new then record.ID is updated as well.
// Returns the error.
func (colordb *ColorBoltDB) UpdateColor(r *types.ColorRecord) error {
	return colordb.updateColorBucket(r)
}

// RemoveColor removes the types.ColorRecord from the database.
// Param id [in] the key of the record to be removed.
// Returns the error.
func (colordb *ColorBoltDB) RemoveColor(id uint64) error {
	return colordb.DB.Update(func(tx *bolt.Tx) error {
		bucketname := []byte(colorBucketName)
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

// updates the types.ColorRecord in the database.
// Param record [in-out] the record to be updated
func (colordb *ColorBoltDB) updateColorBucket(r *types.ColorRecord) error {
	return colordb.DB.Update(func(tx *bolt.Tx) error {
		bucketname := []byte(colorBucketName)
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
func (colordb *ColorBoltDB) initialize() error {
	/*
		example code:

		defaults := somepackage.GetColorDefaults()
		for _, default := range defaults {
			r := types.NewColorRecord()
			r.Name = default.Name
			r.Price = default.Price
			r.SKU = default.SKU
			err := colordb.updateColorBucket(r)
			if err != nil {
				return err
			}
		}
	*/
	return nil
}
