package store

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/domain"
	"github.com/josephbudd/kickwasm/pkg/mainprocess"
)

// List returns the filtered list of store names.
func (mngr *Manager) List() (boltStores, remoteDBs, remoteRecords []string) {
	var l int
	l = len(mngr.manifest.DefaultRecords)
	boltStores = make([]string, l)
	copy(boltStores, mngr.manifest.DefaultRecords)
	l = len(mngr.manifest.RemoteDBs)
	remoteDBs = make([]string, l)
	copy(remoteDBs, mngr.manifest.RemoteDBs)
	l = len(mngr.manifest.RemoteRecords)
	remoteRecords = make([]string, l)
	copy(remoteRecords, mngr.manifest.RemoteRecords)
	return
}

// Finish completes the changes in other application files.
func (mngr *Manager) Finish() (err error) {
	defaultBoltStoreNames := mngr.manifest.DefaultRecords
	remoteDBs := mngr.manifest.RemoteDBs
	remoteRecords := mngr.manifest.RemoteRecords
	if err = mainprocess.RebuildStoresGo(mngr.appPaths, mngr.importGitPath, defaultBoltStoreNames, remoteDBs, remoteRecords); err != nil {
		return
	}
	if err = domain.RebuildStoreStores(mngr.appPaths, mngr.importGitPath, defaultBoltStoreNames, remoteDBs, remoteRecords); err != nil {
		return
	}
	if err = domain.RebuildStoreInstructions(mngr.appPaths, defaultBoltStoreNames, remoteDBs, remoteRecords); err != nil {
		return
	}
	err = mngr.manifest.Update()
	return
}

// filteredList returns the filtered list of store names.
func (mngr *Manager) filteredList() (stores []string, err error) {
	var fnames []string
	if fnames, err = mngr.fileNames(); err != nil {
		return
	}
	stores = make([]string, 0, len(fnames))
	for _, fname := range fnames {
		parts := strings.Split(fname, ".")
		name := parts[0]
		if name != mngr.instructions {
			stores = append(stores, name)
		}
	}
	return
}

// fileNames returns the names of every file in the domain/lpc/store/ folder.
func (mngr *Manager) fileNames() (stores []string, err error) {

	var dir *os.File
	defer func() {
		cerr := dir.Close()
		if err == nil {
			err = cerr
		}
	}()

	// open the store dir.
	if dir, err = os.Open(mngr.storerFolder); err != nil {
		return
	}
	stores, err = dir.Readdirnames(-1)
	return
}

func checkName(storename string) (storeName string, err error) {
	if storeName = cases.CamelCase(storename); storeName != storename {
		msg := fmt.Sprintf("Store names are CamelCased so %q should be %q", storename, storeName)
		err = errors.New(msg)
	}
	return
}
