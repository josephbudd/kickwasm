package common

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Executable make a file name an executable file name.
func Executable(name string) (executable string) {
	switch runtime.GOOS {
	case "darwin":
		executable = name
	case "windows":
		executable = name + ".exe"
	default:
		executable = name
	}
	return
}

// shared vars
const (
	KickwasmDotYAML = "kickwasm.yaml"
)

// RekickwasmDotYAMLEditPath returns the path to the rekickwasm's editable kickwasm.yaml file.
func RekickwasmDotYAMLEditPath(folderPath string) (path string) {
	path = filepath.Join(folderPath, "rekickwasm", "edit", "yaml", KickwasmDotYAML)
	return
}

// PathFound returns if a path exists.
func PathFound(path string) (found bool) {
	if _, err := os.Stat(path); err != nil {
		return
	}
	found = true
	return
}

// Write writes and closes the file.
func Write(path, content string) (err error) {

	var f *os.File
	defer func() {
		if f != nil {
			cErr := f.Close()
			if err == nil {
				err = cErr
			}
		}
		if err != nil {
			err = fmt.Errorf("common.Write: %w", err)
		}
	}()

	if f, err = os.Create(path); err != nil {
		return
	}

	_, err = f.Write([]byte(content))
	return
}

// MkDir makes a dir.
func MkDir(path string) (err error) {
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		err = fmt.Errorf("common.MkDir: %w", err)
	}
	return
}
