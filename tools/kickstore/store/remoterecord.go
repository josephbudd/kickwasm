package store

import (
	"fmt"

	"github.com/josephbudd/kickwasm/pkg/domain"
	"github.com/pkg/errors"
)

// AddRemoteRecord add a new store.
func (mngr *Manager) AddRemoteRecord(recordname string) (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessagef(err, "AddRemoteRecord")
		}
	}()

	var recordName string
	if recordName, err = checkName(recordname); err != nil {
		return
	}
	if mngr.manifest.HaveDefaultRecord(recordName) {
		errmsg := fmt.Sprintf("the local bolt store %q already exists", recordName)
		err = errors.New(errmsg)
		return
	}
	if mngr.manifest.HaveRemoteDB(recordName) {
		errmsg := fmt.Sprintf("the remote database %q already exists", recordName)
		err = errors.New(errmsg)
		return
	}
	if mngr.manifest.HaveRemoteRecord(recordName) {
		errmsg := fmt.Sprintf("the remote database record %q already exists", recordName)
		err = errors.New(errmsg)
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
			err = errors.WithMessagef(err, "DelRemoteRecord")
		}
	}()

	var recordName string
	if recordName, err = checkName(recordname); err != nil {
		return
	}
	if !mngr.manifest.HaveRemoteRecord(recordName) {
		errmsg := fmt.Sprintf("the remote record %q does not exist", recordName)
		err = errors.New(errmsg)
		return
	}
	// Delete the domain/store/record/ store file.
	if err = domain.DeleteStoreRecord(mngr.appPaths, recordName); err != nil {
		return
	}
	mngr.manifest.RemoveRemoteRecord(recordName)
	return
}
