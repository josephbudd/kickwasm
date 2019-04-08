package templates

// ImplementationsStoringBoltStoringGo is the domain/implementations/store/boltstoring/<store name>.go template.
const ImplementationsStoringBoltStoringGo = `package boltstoring

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

const {{call .LowerCamelCase .Store}}BucketName string = "{{call .LowerCamelCase .Store}}"

// {{.Store}}BoltDB is the bolt db for key codes.
type {{.Store}}BoltDB struct {
	DB    *bolt.DB
	path  string
	perms os.FileMode
}

// New{{.Store}}BoltDB constructs a new {{.Store}}BoltDB.
// Param db [in-out] is an open bolt database.
// Returns a pointer to the new {{.Store}}BoltDB.
func New{{.Store}}BoltDB(db *bolt.DB, path string, perms os.FileMode) ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) {
	{{call .LowerCamelCase .Store}}db = &{{.Store}}BoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
	return
}

// {{.Store}}BoltDB implements {{.Store}}Storer
// which is defined in domain/types/records.go

// Open opens the database.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Open() (err error) {
	// the bolt db is already open
	return
}

// Close closes the database.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Close() (err error) {
	err = {{call .LowerCamelCase .Store}}db.DB.Close()
	return
}

// Get{{.Store}} retrieves the types.{{.Store}}Record from the db.
// Param id [in] is the record id.
// Returns the record and error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Get{{.Store}}(id uint64) (r *types.{{.Store}}Record, err error) {
	ids := fmt.Sprintf("%d", id)
	err = {{call .LowerCamelCase .Store}}db.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
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
		r = &types.{{.Store}}Record{}
		if er = json.Unmarshal(rbb, r); er != nil {
			r = nil
			return
		}
		return
	})
	return
}

// Get{{.Store}}s retrieves all of the types.{{.Store}}Record from the db.
// If there are no types.{{.Store}}Records in the db then it calls {{call .LowerCamelCase .Store}}db.initialize().
// See {{call .LowerCamelCase .Store}}db.initialize().
// Returns the records and error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Get{{.Store}}s() (rr []*types.{{.Store}}Record, err error) {
	if rr, err = {{call .LowerCamelCase .Store}}db.get{{.Store}}s(); len(rr) == 0 && err != nil {
		{{call .LowerCamelCase .Store}}db.initialize()
		rr, err = {{call .LowerCamelCase .Store}}db.get{{.Store}}s()
	}
	return
}

func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) get{{.Store}}s() (rr []*types.{{.Store}}Record, err error) {
	err = {{call .LowerCamelCase .Store}}db.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			er = bolt.ErrBucketNotFound
			return
		}
		c := bucket.Cursor()
		rr = make([]*types.{{.Store}}Record, 0, 1024)
		for k, v := c.First(); k != nil; k, v = c.Next() {
			r := types.New{{.Store}}Record()
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

// Update{{.Store}} updates the types.{{.Store}}Record in the database.
// Param record [in-out] the record to be updated.
// if record is new then record.ID is updated as well.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Update{{.Store}}(r *types.{{.Store}}Record) (err error) {
	err = {{call .LowerCamelCase .Store}}db.update{{.Store}}Bucket(r)
	return
}

// Remove{{.Store}} removes the types.{{.Store}}Record from the database.
// Param id [in] the key of the record to be removed.
// If the record is not found the error is nil.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Remove{{.Store}}(id uint64) (err error) {
	err = {{call .LowerCamelCase .Store}}db.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
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

// updates the types.{{.Store}}Record in the database.
// Param record [in-out] the record to be updated.
// If the record is new then it's ID is updated.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) update{{.Store}}Bucket(r *types.{{.Store}}Record) (err error) {
	err = {{call .LowerCamelCase .Store}}db.DB.Update(func(tx *bolt.Tx) (er error) {
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

// initialize is only useful if you want to add the default records to the db.
// otherwise you don't need it to do anything.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) initialize() (err error) {
	/*
		example code:

		defaults := somepackage.Get{{.Store}}Defaults()
		for _, default := range defaults {
			r := types.New{{.Store}}Record()
			r.Name = default.Name
			r.Price = default.Price
			r.SKU = default.SKU
			if err = {{call .LowerCamelCase .Store}}db.update{{.Store}}Bucket(r); err != nil {
				return
			}
		}
	*/
	return
}
`
