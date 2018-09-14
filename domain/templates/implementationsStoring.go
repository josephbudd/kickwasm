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

// ImplementationsStoringBoltStoringGo is the domain/implementations/store/boltstoring/<repo name>.go template.
const ImplementationsStoringBoltStoringGo = `package boltstoring

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

const {{call .LowerCamelCase .Repo}}BucketName string = "{{call .LowerCamelCase .Repo}}"

// {{.Repo}}BoltDB is the bolt db for key codes.
type {{.Repo}}BoltDB struct {
	DB    *bolt.DB
	path  string
	perms os.FileMode
}

// New{{.Repo}}BoltDB constructs a new {{.Repo}}BoltDB.
// Param db [in-out] is an open bolt database.
// Returns a pointer to the new {{.Repo}}BoltDB.
func New{{.Repo}}BoltDB(db *bolt.DB, path string, perms os.FileMode) *{{.Repo}}BoltDB {
	return &{{.Repo}}BoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
}

// {{.Repo}}BoltDB implements {{.Repo}}Storer
// which is defined in domain/types/records.go

// Open opens the database.
// Returns the error.
func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) Open() error {
	// the bolt db is already open
	return nil
}

// Close closes the database.
// Returns the error.
func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) Close() error {
	return {{call .LowerCamelCase .Repo}}db.DB.Close()
}

// Get{{.Repo}} retrieves the types.{{.Repo}}Record from the db.
// Param id [in] is the record id.
// Returns the record and error.
func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) Get{{.Repo}}(id uint64) (*types.{{.Repo}}Record, error) {
	var r types.{{.Repo}}Record
	ids := fmt.Sprintf("%d", id)
	er := {{call .LowerCamelCase .Repo}}db.DB.View(func(tx *bolt.Tx) error {
		bucketname := []byte({{call .LowerCamelCase .Repo}}BucketName)
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

// Get{{.Repo}}s retrieves all of the types.{{.Repo}}Record from the db.
// If there are no types.{{.Repo}}Records in the db then it calls {{call .LowerCamelCase .Repo}}db.initialize().
// See {{call .LowerCamelCase .Repo}}db.initialize().
// Returns the records and error.
func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) Get{{.Repo}}s() ([]*types.{{.Repo}}Record, error) {
	if rr, err := {{call .LowerCamelCase .Repo}}db.get{{.Repo}}s(); len(rr) > 0 && err == nil {
		return rr, err
	}
	{{call .LowerCamelCase .Repo}}db.initialize()
	return {{call .LowerCamelCase .Repo}}db.get{{.Repo}}s()
}

func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) get{{.Repo}}s() ([]*types.{{.Repo}}Record, error) {
	rr := make([]*types.{{.Repo}}Record, 0, 5)
	er := {{call .LowerCamelCase .Repo}}db.DB.View(func(tx *bolt.Tx) error {
		bucketname := []byte({{call .LowerCamelCase .Repo}}BucketName)
		bucket := tx.Bucket(bucketname)
		if bucket != nil {
			c := bucket.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				r := types.New{{.Repo}}Record()
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

// Update{{.Repo}} updates the types.{{.Repo}}Record in the database.
// Param record [in-out] the record to be updated.
// if record is new then record.ID is updated as well.
// Returns the error.
func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) Update{{.Repo}}(r *types.{{.Repo}}Record) error {
	return {{call .LowerCamelCase .Repo}}db.update{{.Repo}}Bucket(r)
}

// Remove{{.Repo}} removes the types.{{.Repo}}Record from the database.
// Param id [in] the key of the record to be removed.
// Returns the error.
func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) Remove{{.Repo}}(id uint64) error {
	return {{call .LowerCamelCase .Repo}}db.DB.Update(func(tx *bolt.Tx) error {
		bucketname := []byte({{call .LowerCamelCase .Repo}}BucketName)
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

// updates the types.{{.Repo}}Record in the database.
// Param record [in-out] the record to be updated
func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) update{{.Repo}}Bucket(r *types.{{.Repo}}Record) error {
	return {{call .LowerCamelCase .Repo}}db.DB.Update(func(tx *bolt.Tx) error {
		bucketname := []byte({{call .LowerCamelCase .Repo}}BucketName)
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
func ({{call .LowerCamelCase .Repo}}db *{{.Repo}}BoltDB) initialize() error {
	/*
		example code:

		defaults := somepackage.Get{{.Repo}}Defaults()
		for _, default := range defaults {
			r := types.New{{.Repo}}Record()
			r.Name = default.Name
			r.Price = default.Price
			r.SKU = default.SKU
			err := {{call .LowerCamelCase .Repo}}db.update{{.Repo}}Bucket(r)
			if err != nil {
				return err
			}
		}
	*/
	return nil
}
`
