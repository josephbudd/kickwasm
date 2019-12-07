package common

import (
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
	// fmt.Printf("Write(%q, content string)\n", path)
	var f *os.File
	if f, err = os.Create(path); err != nil {
		return
	}
	defer func() {
		cErr := f.Close()
		if err == nil {
			err = cErr
		}
	}()

	if _, err = f.Write([]byte(content)); err != nil {
		return
	}
	return
}

// MkDir makes a dir.
func MkDir(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	return
}
