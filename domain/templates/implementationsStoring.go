package templates

// ImplementationsStoringBoltErrorsGo is the domain/implementations/store/boltstoring/errors.go template.
const ImplementationsStoringBoltErrorsGo = `package boltstoring

import (
	"errors"
)

var (
	errNotFound = errors.New("Not Found")
)
`

// ImplementationsStoringBoltStoringGo is the domain/implementations/store/boltstoring/<store name>.go template.
const ImplementationsStoringBoltStoringGo = `package boltstoring

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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
func New{{.Store}}BoltDB(db *bolt.DB, path string, perms os.FileMode) *{{.Store}}BoltDB {
	return &{{.Store}}BoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
}

// {{.Store}}BoltDB implements {{.Store}}Storer
// which is defined in domain/types/records.go

// Open opens the database.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Open() error {
	// the bolt db is already open
	return nil
}

// Close closes the database.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Close() error {
	return {{call .LowerCamelCase .Store}}db.DB.Close()
}

// Get{{.Store}} retrieves the types.{{.Store}}Record from the db.
// Param id [in] is the record id.
// Returns the record and error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Get{{.Store}}(id uint64) (*types.{{.Store}}Record, error) {
	var r types.{{.Store}}Record
	ids := fmt.Sprintf("%d", id)
	er := {{call .LowerCamelCase .Store}}db.DB.View(func(tx *bolt.Tx) error {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
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

// Get{{.Store}}s retrieves all of the types.{{.Store}}Record from the db.
// If there are no types.{{.Store}}Records in the db then it calls {{call .LowerCamelCase .Store}}db.initialize().
// See {{call .LowerCamelCase .Store}}db.initialize().
// Returns the records and error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Get{{.Store}}s() ([]*types.{{.Store}}Record, error) {
	if rr, err := {{call .LowerCamelCase .Store}}db.get{{.Store}}s(); len(rr) > 0 && err == nil {
		return rr, err
	}
	{{call .LowerCamelCase .Store}}db.initialize()
	return {{call .LowerCamelCase .Store}}db.get{{.Store}}s()
}

func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) get{{.Store}}s() ([]*types.{{.Store}}Record, error) {
	rr := make([]*types.{{.Store}}Record, 0, 5)
	er := {{call .LowerCamelCase .Store}}db.DB.View(func(tx *bolt.Tx) error {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
		bucket := tx.Bucket(bucketname)
		if bucket != nil {
			c := bucket.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				r := types.New{{.Store}}Record()
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

// Update{{.Store}} updates the types.{{.Store}}Record in the database.
// Param record [in-out] the record to be updated.
// if record is new then record.ID is updated as well.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Update{{.Store}}(r *types.{{.Store}}Record) error {
	return {{call .LowerCamelCase .Store}}db.update{{.Store}}Bucket(r)
}

// Remove{{.Store}} removes the types.{{.Store}}Record from the database.
// Param id [in] the key of the record to be removed.
// Returns the error.
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) Remove{{.Store}}(id uint64) error {
	return {{call .LowerCamelCase .Store}}db.DB.Update(func(tx *bolt.Tx) error {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
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

// updates the types.{{.Store}}Record in the database.
// Param record [in-out] the record to be updated
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) update{{.Store}}Bucket(r *types.{{.Store}}Record) error {
	return {{call .LowerCamelCase .Store}}db.DB.Update(func(tx *bolt.Tx) error {
		bucketname := []byte({{call .LowerCamelCase .Store}}BucketName)
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
func ({{call .LowerCamelCase .Store}}db *{{.Store}}BoltDB) initialize() error {
	/*
		example code:

		defaults := somepackage.Get{{.Store}}Defaults()
		for _, default := range defaults {
			r := types.New{{.Store}}Record()
			r.Name = default.Name
			r.Price = default.Price
			r.SKU = default.SKU
			err := {{call .LowerCamelCase .Store}}db.update{{.Store}}Bucket(r)
			if err != nil {
				return err
			}
		}
	*/
	return nil
}
`
