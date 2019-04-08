package boltstoring

import (
	"encoding/json"
	"fmt"
	"os"

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
func NewColorBoltDB(db *bolt.DB, path string, perms os.FileMode) (colordb *ColorBoltDB) {
	colordb = &ColorBoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
	return
}

// ColorBoltDB implements ColorStorer
// which is defined in domain/types/records.go

// Open opens the database.
// Returns the error.
func (colordb *ColorBoltDB) Open() (err error) {
	// the bolt db is already open
	return
}

// Close closes the database.
// Returns the error.
func (colordb *ColorBoltDB) Close() (err error) {
	err = colordb.DB.Close()
	return
}

// GetColor retrieves the types.ColorRecord from the db.
// Param id [in] is the record id.
// Returns the record and error.
func (colordb *ColorBoltDB) GetColor(id uint64) (r *types.ColorRecord, err error) {
	ids := fmt.Sprintf("%d", id)
	err = colordb.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(colorBucketName)
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
		r = &types.ColorRecord{}
		if er = json.Unmarshal(rbb, r); er != nil {
			r = nil
			return
		}
		return
	})
	return
}

// GetColors retrieves all of the types.ColorRecord from the db.
// If there are no types.ColorRecords in the db then it calls colordb.initialize().
// See colordb.initialize().
// Returns the records and error.
func (colordb *ColorBoltDB) GetColors() (rr []*types.ColorRecord, err error) {
	if rr, err = colordb.getColors(); len(rr) == 0 && err != nil {
		colordb.initialize()
		rr, err = colordb.getColors()
	}
	return
}

func (colordb *ColorBoltDB) getColors() (rr []*types.ColorRecord, err error) {
	err = colordb.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(colorBucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			er = bolt.ErrBucketNotFound
			return
		}
		c := bucket.Cursor()
		rr = make([]*types.ColorRecord, 0, 1024)
		for k, v := c.First(); k != nil; k, v = c.Next() {
			r := types.NewColorRecord()
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

// UpdateColor updates the types.ColorRecord in the database.
// Param record [in-out] the record to be updated.
// if record is new then record.ID is updated as well.
// Returns the error.
func (colordb *ColorBoltDB) UpdateColor(r *types.ColorRecord) (err error) {
	err = colordb.updateColorBucket(r)
	return
}

// RemoveColor removes the types.ColorRecord from the database.
// Param id [in] the key of the record to be removed.
// If the record is not found the error is nil.
// Returns the error.
func (colordb *ColorBoltDB) RemoveColor(id uint64) (err error) {
	err = colordb.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(colorBucketName)
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

// updates the types.ColorRecord in the database.
// Param record [in-out] the record to be updated.
// If the record is new then it's ID is updated.
// Returns the error.
func (colordb *ColorBoltDB) updateColorBucket(r *types.ColorRecord) (err error) {
	err = colordb.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(colorBucketName)
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

// initialize is only useful if you want to add the default records to the db.
// otherwise you don't need it to do anything.
func (colordb *ColorBoltDB) initialize() (err error) {
	/*
		example code:

		defaults := somepackage.GetColorDefaults()
		for _, default := range defaults {
			r := types.NewColorRecord()
			r.Name = default.Name
			r.Price = default.Price
			r.SKU = default.SKU
			if err = colordb.updateColorBucket(r); err != nil {
				return
			}
		}
	*/
	return
}
