package ftools

import (
	"os"
	"path/filepath"
)

// FMove is a file mover that also can delete an empty source folder.
type FMove struct {
	src, dst             string
	deleteEmptySrcFolder bool
	perm                 os.FileMode
}

// NewFMove constructs a new FMove
func NewFMove(src, dst string, deleteEmptySrcFolder bool, perm os.FileMode) *FMove {
	return &FMove{
		src:                  src,
		dst:                  dst,
		deleteEmptySrcFolder: deleteEmptySrcFolder,
		perm:                 perm,
	}
}

// Move moves a file.
func (fm *FMove) Move() (err error) {
	dir := filepath.Dir(fm.src)
	if err = os.MkdirAll(dir, fm.perm); err != nil {
		return
	}
	if err = CopyFile(fm.src, fm.dst); err != nil {
		return
	}
	var f *os.File
	if f, err = os.Open(dir); err != nil {
		return
	}
	if err = os.Remove(fm.src); err != nil {
		return
	}
	var names []string
	if names, err = f.Readdirnames(-1); err != nil {
		return
	}
	for _, name := range names {
		if name != "." && name != ".." {
			return
		}

	}
	// no files in the folder
	err = os.Remove(dir)
	return
}
