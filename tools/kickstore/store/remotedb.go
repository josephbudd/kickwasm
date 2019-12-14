package store

import (
	"fmt"

	"github.com/josephbudd/kickwasm/pkg/domain"
)

// AddRemoteDB add a new store.
func (mngr *Manager) AddRemoteDB(storename string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("AddRemoteDB: %w", err)
		}
	}()

	var storeName string
	if storeName, err = checkName(storename); err != nil {
		return
	}
	if mngr.manifest.HaveDefaultRecord(storeName) {
		err = fmt.Errorf("the local bolt store %q already exists", storeName)
		return
	}
	if mngr.manifest.HaveRemoteDB(storeName) {
		err = fmt.Errorf("the remote database %q already exists", storeName)
		return
	}
	if mngr.manifest.HaveRemoteRecord(storeName) {
		err = fmt.Errorf("the remote database record %q already exists", storeName)
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
			err = fmt.Errorf("DelRemoteDB: %w", err)
		}
	}()

	var storeName string
	if storeName, err = checkName(storename); err != nil {
		return
	}
	if !mngr.manifest.HaveRemoteDB(storeName) {
		err = fmt.Errorf("the remote database %q does not exist", storeName)
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
