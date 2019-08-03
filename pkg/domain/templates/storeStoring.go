package templates

// StoreStoringGo is the domain/store/storing/<store name>.go template.
const StoreStoringGo = `package storing

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"
)

/*
	type {{.Store}}BoltDB
	is the implementation of the {{.ImportDomainStoreStorer}}.{{.Store}}Storer interface
	  for the bolt database.
*/

const {{call .LowerCamelCase .Store}}BucketName string = "{{call .LowerCamelCase .Store}}"

// {{.Store}}BoltDB is the bolt db for key codes.
type {{.Store}}BoltDB struct {
	DB    *bolt.DB
	path  string
	perms os.FileMode
}

// New{{.Store}}BoltDB constructs a new {{.Store}}BoltDB.
// Param db is an open bolt data-store.
// Returns a pointer to the new {{.Store}}BoltDB.
func New{{.Store}}BoltDB(db *bolt.DB, path string, perms os.FileMode) (store *{{.Store}}BoltDB) {
	store = &{{.Store}}BoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
	return
}

// Open opens the bolt data-store.
// Returns the error.
func (store *{{.Store}}BoltDB) Open() (err error) {
	// the bolt db is already open
	return
}

// Close closes the bolt data-store.
// Returns the error.
func (store *{{.Store}}BoltDB) Close() (err error) {
	err = store.DB.Close()
	return
}

// Get retrieves the record.{{.Store}} from the bolt data-store.
// Param id is the record id.
// Returns the record and error.
// If no record is found returns a nil record and a nil error.
func (store *{{.Store}}BoltDB) Get(id uint64) (r *record.{{.Store}}, err error) {
	ids := fmt.Sprintf("%d", id)
	err = store.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			return
		}
		var rbb []byte
		if rbb = bucket.Get([]byte(ids)); rbb == nil {
			// not found
			return
		}
		r = &record.{{.Store}}{}
		if er = json.Unmarshal(rbb, r); er != nil {
			r = nil
			return
		}
		return
	})
	return
}

// GetAll retrieves all of the record.{{.Store}} from the bolt data-store.
// If no record is found then it calls store.initialize() and tries again. See *{{.Store}}BoltDB.initialize().
// Returns the records and error.
// If no record is found returns a zero length records and a nil error.
func (store *{{.Store}}BoltDB) GetAll() (rr []*record.{{.Store}}, err error) {
	if rr, err = store.getAll(); len(rr) == 0 && err == nil {
		store.initialize()
		rr, err = store.getAll()
	}
	return
}

func (store *{{.Store}}BoltDB) getAll() (rr []*record.{{.Store}}, err error) {
	err = store.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			return
		}
		c := bucket.Cursor()
		rr = make([]*record.{{.Store}}, 0, 1024)
		for k, v := c.First(); k != nil; k, v = c.Next() {
			r := record.New{{.Store}}()
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

// Update adds or updates a record.{{.Store}} in the bolt data-store.
// Param r is the record to be updated.
// If r is a new record then r.ID is updated with the new record id.
// Returns the error.
func (store *{{.Store}}BoltDB) Update(r *record.{{.Store}}) (err error) {
	err = store.update(r)
	return
}

// Remove removes a record.{{.Store}} from the bolt data-store.
// Param id the key of the record to be removed.
// If the record is not found returns a nil error.
// Returns the error.
func (store *{{.Store}}BoltDB) Remove(id uint64) (err error) {
	err = store.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			return
		}
		idbb := []byte(fmt.Sprintf("%d", id))
		er = bucket.Delete(idbb)
		return
	})
	return
}

// updates the record.{{.Store}} in the bolt data-store.
// Param record the record to be updated.
// If the record is new then it's ID is updated.
// Returns the error.
func (store *{{.Store}}BoltDB) update(r *record.{{.Store}}) (err error) {
	err = store.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
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

// initialize is only useful if you want to add the default records to the bolt data-store.
// otherwise you don't need it to do anything.
func (store *{{.Store}}BoltDB) initialize() (err error) {
	/*
		example code:

		defaults := somepackage.Get{{.Store}}Defaults()
		for _, default := range defaults {
			r := types.New{{.Store}}Record()
			r.Name = default.Name
			r.Price = default.Price
			r.SKU = default.SKU
			if err = store.update(r); err != nil {
				return
			}
		}
	*/
	return
}
`
