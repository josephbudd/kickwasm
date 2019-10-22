package nopack

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

const (
	emptyString = ""
)

// fMode is the applications mode for files.
var fMode = os.FileMode(0666)

// dMode is the applications mode for folders.
var dMode = os.FileMode(0775)

// write writes each file in pathBytes to the folderpath.
func write(folderpath, packageName string) error {
	if err := os.MkdirAll(folderpath, dMode); err != nil {
		return err
	}
	lines := new(bytes.Buffer)
	lines.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	lines.WriteString(`import (
	"os"
	"path/filepath"
)

var wd string

`)
	lines.WriteString(`func init() {
	var err error
	if wd, err = os.Getwd(); err != nil {
		wd = ""
	}
}

`)
	lines.WriteString(`// Contents returns the contents of the file at path and if found.
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

`)
	fname := packageName + ".go"
	path := filepath.Join(folderpath, fname)
	return writeFile(path, lines.Bytes())
}

func writeFile(filepath string, contents []byte) error {
	ofile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, fMode)
	if err != nil {
		return err
	}
	_, err = ofile.Write(contents)
	ofile.Close()
	return err
}
