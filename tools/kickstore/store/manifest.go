package store

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Manifest is the data stored in ./kickwasm/kickwasm/tools/kickstore/stores.yaml.
type Manifest struct {
	path           string      `yaml:",omit"`
	mode           os.FileMode `yaml:",omit"`
	DefaultRecords []string    `yaml:"defaultRecords"`
	RemoteDBs      []string    `yaml:"remoteDBs"`
	RemoteRecords  []string    `yaml:"remoteRecords"`
}

// NewManifest constructs a new Manifest.
func NewManifest(path string, mode os.FileMode) (manifest *Manifest, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		if os.IsNotExist(err) {
			manifest = &Manifest{
				mode:           mode,
				path:           path,
				DefaultRecords: make([]string, 0, 20),
				RemoteDBs:      make([]string, 0, 20),
				RemoteRecords:  make([]string, 0, 20),
			}
			err = manifest.Update()
		}
		return
	}
	var stat os.FileInfo
	if stat, err = f.Stat(); err != nil {
		return
	}
	size := stat.Size()
	bb := make([]byte, size)
	if _, err = f.Read(bb); err != nil {
		f.Close()
		return
	}
	manifest = &Manifest{}
	if err = yaml.Unmarshal(bb, manifest); err != nil {
		f.Close()
		return
	}
	manifest.path = path
	manifest.mode = stat.Mode()
	err = f.Close()
	return
}

// Update saves the manifest back to the disk.
func (manifest *Manifest) Update() (err error) {
	var bb []byte
	if bb, err = yaml.Marshal(manifest); err != nil {
		return
	}
	var f *os.File
	if f, err = os.OpenFile(manifest.path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, manifest.mode); err != nil {
		return
	}
	if _, err = f.Write(bb); err != nil {
		f.Close()
		return
	}
	err = f.Close()
	return
}

// Default Bolt Records.

// AddDefaultRecord adds a default record to manifest.
func (manifest *Manifest) AddDefaultRecord(name string) {
	manifest.DefaultRecords = append(manifest.DefaultRecords, name)
}

// RemoveDefaultRecord adds a default record to manifest.
func (manifest *Manifest) RemoveDefaultRecord(name string) {
	l := len(manifest.DefaultRecords)
	list := make([]string, l-1, l+20)
	var i int
	for _, r := range manifest.DefaultRecords {
		if r != name {
			list[i] = r
			i++
		}
	}
	manifest.DefaultRecords = list
}

// HaveDefaultRecord adds a default record to manifest.
func (manifest *Manifest) HaveDefaultRecord(name string) (have bool) {
	for _, r := range manifest.DefaultRecords {
		if r == name {
			have = true
			break
		}
	}
	return
}

// Remote Database.

// AddRemoteDB adds a default record to manifest.
func (manifest *Manifest) AddRemoteDB(name string) {
	manifest.RemoteDBs = append(manifest.RemoteDBs, name)
}

// RemoveRemoteDB adds a default record to manifest.
func (manifest *Manifest) RemoveRemoteDB(name string) {
	l := len(manifest.RemoteDBs)
	list := make([]string, l-1, l+20)
	var i int
	for _, r := range manifest.RemoteDBs {
		if r != name {
			list[i] = r
			i++
		}
	}
	manifest.RemoteDBs = list
}

// HaveRemoteDB adds a default record to manifest.
func (manifest *Manifest) HaveRemoteDB(name string) (have bool) {
	for _, r := range manifest.RemoteDBs {
		if r == name {
			have = true
			break
		}
	}
	return
}

// RemoteRecord

// AddRemoteRecord adds a default record to manifest.
func (manifest *Manifest) AddRemoteRecord(name string) {
	manifest.RemoteRecords = append(manifest.RemoteRecords, name)
}

// RemoveRemoteRecord adds a default record to manifest.
func (manifest *Manifest) RemoveRemoteRecord(name string) {
	l := len(manifest.RemoteRecords)
	list := make([]string, l-1, l+20)
	var i int
	for _, r := range manifest.RemoteRecords {
		if r != name {
			list[i] = r
			i++
		}
	}
	manifest.RemoteRecords = list
}

// HaveRemoteRecord adds a default record to manifest.
func (manifest *Manifest) HaveRemoteRecord(name string) (have bool) {
	for _, r := range manifest.RemoteRecords {
		if r == name {
			have = true
			break
		}
	}
	return
}
