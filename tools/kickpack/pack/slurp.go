package pack

import (
	"fmt"
	"os"
	"path/filepath"
)

// slurp constructs a new map of file paths to their contents.
// Param path is the path to the folder.
// If the file or folder does not exist and MustExistFlag is true returns error.
// Returns the file map and the error.
func slurp(path string, mustExist bool, pathBytes map[string][]byte) (err error) {
	var info os.FileInfo
	if info, err = os.Stat(path); err != nil {
		if !mustExist && os.IsNotExist(err) {
			err = nil
			pathBytes[path] = make([]byte, 0, 0)
		}
		return
	}
	if info.IsDir() {
		// folder
		err = slurpFolder(path, pathBytes)
		return
	}
	if info.Mode()&os.ModeType != 0 {
		// not a folder, not regular file
		err = fmt.Errorf("%s is not a file or folder", path)
		return
	}
	// regular file
	var filebb []byte
	if filebb, err = slurpFile(path); err != nil {
		return
	}
	pathBytes[path] = filebb
	return
}

func slurpFolder(src string, pathBytes map[string][]byte) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	// interate through the dir contents.
	infos, err := f.Readdir(-1)
	if err != nil {
		return err
	}
	for _, info := range infos {
		path := filepath.Join(src, info.Name())
		if info.IsDir() {
			if err := slurpFolder(path, pathBytes); err != nil {
				return err
			}
		} else if info.Mode()&os.ModeType == 0 {
			// regular file
			filebb, err := slurpFile(path)
			if err != nil {
				return err
			}
			pathBytes[path] = filebb
		}
	}
	return nil
}

func slurpFile(path string) ([]byte, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	ifile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	size := info.Size()
	buf := make([]byte, size, size)
	defer ifile.Close()
	if _, err := ifile.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}
