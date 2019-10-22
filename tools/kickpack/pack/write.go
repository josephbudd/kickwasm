package pack

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
func write(pathBytes map[string][]byte, folderpath, packageName string) error {
	if err := os.MkdirAll(folderpath, dMode); err != nil {
		return err
	}
	lines := new(bytes.Buffer)
	lines.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	lines.WriteString(`// Contents returns the contents of the file at path and if found.
func Contents(path string) (contents []byte, found bool) {
	contents, found = fileStore[path]
	return
}

// Paths returns a slice of the file paths.
func Paths() (paths []string) {
	l := len(fileStore)
	paths = make([]string, 0, l)
	for k := range fileStore {
		paths = append(paths, k)
	}
	return
}

`)
	lines.WriteString("// fileStore is a store of various files.\n")
	lines.WriteString("var fileStore =  map[string][]byte{\n")
	for fpath, fcontents := range pathBytes {
		if fpath[:2] == "./" {
			fpath = fpath[2:]
		}
		lines.WriteString(fmt.Sprintf(`    "%s":%#v,`, fpath, fcontents))
		lines.WriteString("\n")
	}
	lines.WriteString("}\n")
	fname := packageName + ".go"
	path := filepath.Join(folderpath, fname)
	return writeFile(path, []byte(lines.String()))
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
