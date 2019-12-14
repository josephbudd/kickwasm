package ftools

import (
	"fmt"
	"os"
	"path/filepath"
)

// DCopy allows folders to be copied.
type DCopy struct {
	sourceFolder      string
	destinationFolder string
	offset            int
	clearDestination  bool
	copyHiddenFiles   bool
	skipFolders       []string
}

// NewDCopy constructs a new DCopy.
// Copies the contents of dcopy.sourceFolder to dcopy.destinationFolder.
func NewDCopy(src, dest string, clearDest, copyHiddenFiles bool, skipFolders []string) (dcopy *DCopy, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("NewDCopy: %w", err)
		}
	}()
	s := filepath.Clean(src)
	d := filepath.Clean(dest)
	if s == d {
		err = fmt.Errorf("source and destination %q are the same", s)
		return
	}
	// check the source
	if _, err = os.Stat(s); os.IsNotExist(err) {
		err = fmt.Errorf("source path %q does not exist", s)
		return
	}
	if _, err = os.Stat(d); os.IsNotExist(err) {
		info, _ := os.Stat(s)
		if err = os.MkdirAll(d, info.Mode()); err != nil {
			return
		}
	}
	dcopy = &DCopy{
		sourceFolder:      s,
		destinationFolder: d,
		offset:            len(s),
		clearDestination:  clearDest,
		copyHiddenFiles:   copyHiddenFiles,
		skipFolders:       skipFolders,
	}
	return
}

// Copy does a folder copy
func (dcopy *DCopy) Copy() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("(dcopy *DCopy) Copy(): %w", err)
		}
	}()
	if dcopy.clearDestination {
		if err = os.RemoveAll(dcopy.destinationFolder); err != nil && !os.IsNotExist(err) {
			return
		}
	}
	if err = os.MkdirAll(dcopy.destinationFolder, os.ModePerm); err != nil {
		return
	}
	err = filepath.Walk(
		dcopy.sourceFolder,
		func(path string, info os.FileInfo, er error) error {
			if er != nil {
				return er
			}
			dest := path[dcopy.offset:]
			var isHidden bool
			if len(dest) > 0 {
				isHidden = dest[:1] == "."
			}
			if info.IsDir() {
				// skip the skip folders
				base := filepath.Base(path)
				for _, f := range dcopy.skipFolders {
					if base == f {
						// DEBUG
						// fmt.Printf("skipping %s\n", f)
						return filepath.SkipDir
					}
				}
				// always copy source folder even if its hidden.
				if path != dcopy.sourceFolder && (!dcopy.copyHiddenFiles && isHidden) {
					// skip hidden folders
					return filepath.SkipDir
				}
				// copying this folder.
				destPath := filepath.Join(dcopy.destinationFolder, dest)
				//fmt.Printf("will Mkdir(%q)\n", destPath)
				if err := os.Mkdir(destPath, info.Mode()); err != nil {
					if os.IsExist(err) {
						// the folder exists.
						// not an error
						return nil
					}
					return er
				}
			} else {
				// file not a folder
				if isHidden && !dcopy.copyHiddenFiles {
					// do not copy hidden files
					return nil
				}
				// copying this file.
				destPath := filepath.Join(dcopy.destinationFolder, dest)
				//fmt.Printf("will copy %q to %q\n", path, destPath)
				copyFile(path, destPath)
			}
			return nil
		})
	return
}
