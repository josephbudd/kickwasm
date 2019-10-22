package ftools

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// DMove moves a folder and its contents.
type DMove struct {
	sourceFolder, destinationFolder   string
	clearDestination, copyHiddenFiles bool
	skipFolders                       []string
	offset                            int
}

// NewDMove constructs a new DMove.
func NewDMove(src, dest string, clearDest, copyHiddenFiles bool, skipFolders []string) (dmove *DMove, err error) {
	defer func() {
		if err != nil {
			err = errors.WithMessage(err, `NewDMove`)
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
	dmove = &DMove{
		sourceFolder:      s,
		destinationFolder: d,
		copyHiddenFiles:   copyHiddenFiles,
		clearDestination:  clearDest,
		offset:            len(s),
		skipFolders:       skipFolders,
	}
	return
}

// Move moves a folder
func (dmove *DMove) Move() (err error) {
	defer func() {
		if err != nil {
			err = errors.WithMessage(err, `(dmove *DMove) Move()`)
		}
	}()
	if dmove.clearDestination {
		if err = os.RemoveAll(dmove.destinationFolder); err != nil && !os.IsNotExist(err) {
			return
		}
	}
	if err = os.MkdirAll(dmove.destinationFolder, os.ModePerm); err != nil {
		return
	}
	err = filepath.Walk(
		dmove.sourceFolder,
		func(path string, info os.FileInfo, er error) error {
			if er != nil {
				return er
			}
			dest := path[dmove.offset:]
			if info.IsDir() {
				// skip the skip folders
				base := filepath.Base(path)
				for _, f := range dmove.skipFolders {
					if base == f {
						// DEBUG
						fmt.Printf("skipping %s\n", f)
						return filepath.SkipDir
					}
				}
				// always copy source folder even if its hidden.
				if path != dmove.sourceFolder && (!dmove.copyHiddenFiles && dest[:1] == ".") {
					// skip hidden folders
					return filepath.SkipDir
				}
				destPath := filepath.Join(dmove.destinationFolder, dest)
				if er = os.Mkdir(destPath, info.Mode()); er != nil {
					if os.IsExist(er) {
						// the folder exists.
						// not an error
						return nil
					}
					return er
				}
			} else {
				destPath := filepath.Join(dmove.destinationFolder, dest)
				copyFile(path, destPath)
			}
			return nil
		})
	if err != nil {
		return
	}
	if err = os.RemoveAll(dmove.sourceFolder); err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
	}
	return
}
