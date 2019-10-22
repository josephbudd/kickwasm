package ftools

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// CopyFile copies a file.
func CopyFile(src, dest string) (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "CopyFile")
		}
	}()

	err = copyFile(src, dest)
	return
}

func copyFile(src, dest string) (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "copyFile")
		}
	}()

	var srcInfo os.FileInfo
	srcInfo, err = os.Stat(src)
	if err != nil {
		err = errors.WithMessage(err, `os.Stat(src)`)
		return
	}
	var bb []byte
	bb, err = readFile(src)
	if err != nil {
		return
	}
	srcDirInfo, _ := os.Stat(filepath.Dir(src))
	dstDirPath := filepath.Dir(dest)
	if err = os.MkdirAll(dstDirPath, srcDirInfo.Mode()); err != nil {
		err = errors.WithMessage(err, `os.MkdirAll(dstDirPath, srcDirInfo.Mode())`)
		return
	}
	err = writeFile(dest, srcInfo.Mode(), bb)
	return
}

func readFile(path string) (bb []byte, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "readFile")
		}
	}()
	// read
	f, err := os.Open(path)
	if err != nil {
		err = errors.WithMessage(err, `os.Open(path)`)
		return
	}
	var info os.FileInfo
	info, err = f.Stat()
	if err != nil {
		err = errors.WithMessage(err, `f.Stat()`)
	}
	size := info.Size()
	bb = make([]byte, size, size)
	_, err = f.Read(bb)
	if err != nil {
		err = errors.WithMessage(err, `f.Read(bb)`)
	}
	return
}

func writeFile(path string, mode os.FileMode, bb []byte) (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "writeFile")
		}
	}()

	var ofile *os.File
	ofile, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		err = errors.WithMessage(err, `os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)`)
		return
	}
	if _, err = ofile.Write(bb); err != nil {
		ofile.Close()
		err = errors.WithMessage(err, `ofile.Write(bb)`)
		return
	}
	if err = ofile.Close(); err != nil {
		err = errors.WithMessage(err, `ofile.Close()`)
		return
	}
	return
}
