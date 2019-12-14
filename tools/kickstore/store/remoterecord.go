package store

import (
	"fmt"

	"github.com/josephbudd/kickwasm/pkg/domain"
)

// AddRemoteRecord add a new store.
func (mngr *Manager) AddRemoteRecord(recordname string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("AddRemoteRecord: :%w", err)
		}
	}()

	var recordName string
	if recordName, err = checkName(recordname); err != nil {
		return
	}
	if mngr.manifest.HaveDefaultRecord(recordName) {
		err = fmt.Errorf("the local bolt store %q already exists", recordName)
		return
	}
	if mngr.manifest.HaveRemoteDB(recordName) {
		err = fmt.Errorf("the remote database %q already exists", recordName)
		return
	}
	if mngr.manifest.HaveRemoteRecord(recordName) {
		err = fmt.Errorf("the remote database record %q already exists", recordName)
		return
	}
	// Create the domain/store/record/ store file.
	if err = domain.CreateRemoteDatabaseRecord(mngr.appPaths, mngr.importGitPath, recordName); err != nil {
		return
	}
	mngr.manifest.AddRemoteRecord(recordName)
	return
}

// DelRemoteRecord removes a remote record.
func (mngr *Manager) DelRemoteRecord(recordname string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("DelRemoteRecord: %w", err)
		}
	}()

	var recordName string
	if recordName, err = checkName(recordname); err != nil {
		return
	}
	if !mngr.manifest.HaveRemoteRecord(recordName) {
		err = fmt.Errorf("the remote record %q does not exist", recordName)
		return
	}
	// Delete the domain/store/record/ store file.
	if err = domain.DeleteStoreRecord(mngr.appPaths, recordName); err != nil {
		return
	}
	mngr.manifest.RemoveRemoteRecord(recordName)
	return
}
