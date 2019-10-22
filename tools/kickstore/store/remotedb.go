package store

import (
	"fmt"

	"github.com/josephbudd/kickwasm/pkg/domain"
	"github.com/pkg/errors"
)

// AddRemoteDB add a new store.
func (mngr *Manager) AddRemoteDB(storename string) (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessagef(err, "AddRemoteDB")
		}
	}()

	var storeName string
	if storeName, err = checkName(storename); err != nil {
		return
	}
	if mngr.manifest.HaveDefaultRecord(storeName) {
		errmsg := fmt.Sprintf("the local bolt store %q already exists", storeName)
		err = errors.New(errmsg)
		return
	}
	if mngr.manifest.HaveRemoteDB(storeName) {
		errmsg := fmt.Sprintf("the remote database %q already exists", storeName)
		err = errors.New(errmsg)
		return
	}
	if mngr.manifest.HaveRemoteRecord(storeName) {
		errmsg := fmt.Sprintf("the remote database record %q already exists", storeName)
		err = errors.New(errmsg)
		return
	}
	// Create the domain/store/storing/ file.
	if err = domain.CreateRemoteStoreStoring(mngr.appPaths, mngr.importGitPath, storeName); err != nil {
		return
	}
	// Create the domain/store/storer/ store file.
	if err = domain.CreateRemoteStoreStorer(mngr.appPaths, mngr.importGitPath, storeName); err != nil {
		return
	}
	mngr.manifest.AddRemoteDB(storeName)
	return
}

// DelRemoteDB removes a store.
func (mngr *Manager) DelRemoteDB(storename string) (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessagef(err, "DelRemoteDB")
		}
	}()

	var storeName string
	if storeName, err = checkName(storename); err != nil {
		return
	}
	if !mngr.manifest.HaveRemoteDB(storeName) {
		errmsg := fmt.Sprintf("the remote database %q does not exist", storeName)
		err = errors.New(errmsg)
		return
	}
	// Delete the domain/store/storing/ file.
	if err = domain.DeleteStoreStoring(mngr.appPaths, mngr.importGitPath, storeName); err != nil {
		return
	}
	// Delete the domain/store/storer/ store file.
	if err = domain.DeleteStoreStorer(mngr.appPaths, storeName); err != nil {
		return
	}
	mngr.manifest.RemoveRemoteDB(storeName)
	return
}
