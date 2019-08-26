package templates

// RemoteStoreStoringGo is the domain/store/storing/<store name>.go template for a remote database.
const RemoteStoreStoringGo = `package storing

import (
	"github.com/pkg/errors"
)

// {{.Store}}RemoteDB is the API of the {{.Store}} remote database connection.
// It is the implementation of the interface in {{.ImportDomainStoreStorer}}{{.Store}}.go.
type {{.Store}}RemoteDB struct {}

// New{{.Store}}RemoteDB constructs a new {{.Store}}RemoteDB.
// Returns a pointer to the new {{.Store}}RemoteDB.
func New{{.Store}}RemoteDB() (store *{{.Store}}RemoteDB) {
	store = &{{.Store}}RemoteDB{}
	return
}

// Open opens the connection to the remote database.
// Returns the error.
func (store *{{.Store}}RemoteDB) Open() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "{{.Store}}RemoteDB.Open")
		}
	}()

	return
}

// Close closes the connection to the remote database.
// Returns the error.
func (store *{{.Store}}RemoteDB) Close() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "{{.Store}}RemoteDB.Open")
		}
	}()

	return
}
`

// LocalBoltStoreStoringGo is the domain/store/storing/<store name>.go template.
const LocalBoltStoreStoringGo = `package storing

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"

	"{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"
)

const {{call .LowerCamelCase .Store}}BucketName string = "{{call .LowerCamelCase .Store}}"

// {{.Store}}LocalBoltStore is the API of the {{.Store}} local bolt store.
// It is the implementation of the interface in {{.ImportDomainStoreStorer}}{{.Store}}.go.
type {{.Store}}LocalBoltStore struct {
	DB    *bolt.DB
	path  string
	perms os.FileMode
}

// New{{.Store}}LocalBoltStore constructs a new {{.Store}}LocalBoltStore.
// Param db is an open bolt data-store.
// Returns a pointer to the new {{.Store}}LocalBoltStore.
func New{{.Store}}LocalBoltStore(path string, perms os.FileMode) (store *{{.Store}}LocalBoltStore) {
	store = &{{.Store}}LocalBoltStore{
		path:  path,
		perms: perms,
	}
	return
}

// Open opens the bolt data-store.
// Returns the error.
func (store *{{.Store}}LocalBoltStore) Open() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "{{.Store}}LocalBoltStore.Open")
		}
	}()

	if store.DB, err = bolt.Open(store.path, store.perms, nil); err != nil {
		err = errors.WithMessage(err, "bolt.Open(path, filepaths.GetFmode(), nil)")
	}
	return
}

// Close closes the bolt data-store.
// Returns the error.
func (store *{{.Store}}LocalBoltStore) Close() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "{{.Store}}LocalBoltStore.Close")
		}
	}()

	err = store.DB.Close()
	return
}

// Get retrieves the record.{{.Store}} from the bolt data-store.
// Param id is the record id.
// Returns the record and error.
// If no record is found returns a nil record and a nil error.
func (store *{{.Store}}LocalBoltStore) Get(id uint64) (r *record.{{.Store}}, err error) {
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
// If no record is found then it calls store.initialize() and tries again. See *{{.Store}}LocalBoltStore.initialize().
// Returns the records and error.
// If no record is found returns a zero length records and a nil error.
func (store *{{.Store}}LocalBoltStore) GetAll() (rr []*record.{{.Store}}, err error) {
	if rr, err = store.getAll(); len(rr) == 0 && err == nil {
		store.initialize()
		rr, err = store.getAll()
	}
	return
}

func (store *{{.Store}}LocalBoltStore) getAll() (rr []*record.{{.Store}}, err error) {
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
func (store *{{.Store}}LocalBoltStore) Update(r *record.{{.Store}}) (err error) {
	err = store.update(r)
	return
}

// Remove removes a record.{{.Store}} from the bolt data-store.
// Param id the key of the record to be removed.
// If the record is not found returns a nil error.
// Returns the error.
func (store *{{.Store}}LocalBoltStore) Remove(id uint64) (err error) {
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
func (store *{{.Store}}LocalBoltStore) update(r *record.{{.Store}}) (err error) {
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
func (store *{{.Store}}LocalBoltStore) initialize() (err error) {
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
