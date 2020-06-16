package ftools

import (
	"fmt"
	"os"
	"path/filepath"
)

// PathExists returns if a path exists.
func PathExists(path string) (exists bool) {
	if _, err := os.Stat(path); err != nil {
		exists = !os.IsNotExist(err)
		return
	}
	exists = true
	return
}

// CopyFile copies a file.
func CopyFile(src, dest string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("CopyFile: %w", err)
		}
	}()

	err = copyFile(src, dest)
	return
}

func copyFile(src, dest string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("copyFile: %w", err)
		}
	}()

	var srcInfo os.FileInfo
	srcInfo, err = os.Stat(src)
	if err != nil {
		err = fmt.Errorf(`os.Stat(src): %w`, err)
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
		err = fmt.Errorf(`os.MkdirAll(dstDirPath, srcDirInfo.Mode()): %w`, err)
		return
	}
	err = writeFile(dest, srcInfo.Mode(), bb)
	return
}

func readFile(path string) (bb []byte, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("readFile: %w", err)
		}
	}()

	// read
	f, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf(`os.Open(path): %w`, err)
		return
	}
	var info os.FileInfo
	info, err = f.Stat()
	if err != nil {
		err = fmt.Errorf(`f.Stat(): %w`, err)
	}
	size := info.Size()
	bb = make([]byte, size, size)
	_, err = f.Read(bb)
	if err != nil {
		err = fmt.Errorf(`f.Read(bb): %w`, err)
	}
	return
}

func writeFile(path string, mode os.FileMode, bb []byte) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("writeFile: %w", err)
		}
	}()

	var ofile *os.File
	ofile, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		err = fmt.Errorf(`os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode): %w`, err)
		return
	}
	if _, err = ofile.Write(bb); err != nil {
		ofile.Close()
		err = fmt.Errorf(`ofile.Write(bb): %w`, err)
		return
	}
	if err = ofile.Close(); err != nil {
		err = fmt.Errorf(`ofile.Close(): %w`, err)
		return
	}
	return
}
