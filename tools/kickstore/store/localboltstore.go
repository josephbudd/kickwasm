package store

import (
	"fmt"

	"github.com/josephbudd/kickwasm/pkg/domain"
)

// Add add a new store.
func (mngr *Manager) Add(storename string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("Add: %w", err)
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
	if err = domain.CreateStoreStoring(mngr.appPaths, mngr.importGitPath, storeName); err != nil {
		return
	}
	// Create the domain/store/storer/ store file.
	if err = domain.CreateStoreStorer(mngr.appPaths, mngr.importGitPath, storeName); err != nil {
		return
	}
	// Create the domain/store/record/ store file.
	if err = domain.CreateStoreRecord(mngr.appPaths, mngr.importGitPath, storeName); err != nil {
		return
	}
	mngr.manifest.AddDefaultRecord(storeName)
	return
}

// Del removes a store.
func (mngr *Manager) Del(storename string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("Del: %w", err)
		}
	}()

	var storeName string
	if storeName, err = checkName(storename); err != nil {
		return
	}
	if !mngr.manifest.HaveDefaultRecord(storeName) {
		err = fmt.Errorf("the local bolt store %q does not exist", storeName)
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
	// Delete the domain/store/record/ store file.
	if err = domain.DeleteStoreRecord(mngr.appPaths, storeName); err != nil {
		return
	}
	mngr.manifest.RemoveDefaultRecord(storeName)
	return
}
