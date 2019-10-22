package gooutputnopack

import (
	"os"
	"path/filepath"
)

var wd string

func init() {
	var err error
	if wd, err = os.Getwd(); err != nil {
		wd = ""
	}
}

// Contents returns the contents of the file at path and if found.
func Contents(path string) (contents []byte, found bool) {
	fullpath := filepath.Join(wd, path)
	var f *os.File
	var err error
	if f, err = os.Open(fullpath); err != nil {
		return
	}
	defer f.Close()
	var stat os.FileInfo
	if stat, err = f.Stat(); err != nil {
		return
	}
	contents = make([]byte, stat.Size())
	if _, err = f.Read(contents); err != nil {
		contents = nil
		return
	}
	found = true
	return
}

// Paths returns a slice of the file paths.
func Paths() (paths []string) {
	paths = make([]string, 0)
	return
}

