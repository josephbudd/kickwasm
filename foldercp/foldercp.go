package foldercp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
)

// CopyYAML copies the yaml files.
// appPaths is the application paths.
// Param appYamlFilePath is the path of the application yaml file.
// Param panelFilePaths is the slice of each panel file's full path.
func CopyYAML(appPaths *paths.ApplicationPaths, appYamlFilePath string, panelFilePaths []string) (err error) {
	folderPaths := appPaths.GetPaths()
	srcFolder := filepath.Dir(appYamlFilePath)
	srcFile := filepath.Base(appYamlFilePath)
	// copy the application yaml file.
	if err = copyYaml(appPaths, appYamlFilePath, filepath.Join(folderPaths.OutputDotKickwasmYAML, srcFile)); err != nil {
		return
	}
	// copy the panel yaml files.
	var relPath string
	for _, srcPath := range panelFilePaths {
		relPath, err = filepath.Rel(srcFolder, srcPath)
		if err != nil {
			return
		}
		destPath := filepath.Join(folderPaths.OutputDotKickwasmYAML, relPath)
		if err = copyYaml(appPaths, srcPath, destPath); err != nil {
			return
		}
	}
	return
}

func copyYaml(appPaths *paths.ApplicationPaths, src, dest string) (err error) {
	if err = os.MkdirAll(filepath.Dir(dest), appPaths.DMode); err != nil {
		return fmt.Errorf(`os.MkdirAll(filepath.Dir(dest), appPaths.DMode) error is %s`, err.Error())
	}
	// read
	var ifile *os.File
	ifile, err = os.Open(src)
	if err != nil {
		err = fmt.Errorf(`os.Open(src) error is %s`, err.Error())
		return
	}
	var info os.FileInfo
	info, err = ifile.Stat()
	if err != nil {
		err = fmt.Errorf(`ifile.Stat() error is %s`, err.Error())
		return
	}
	size := info.Size()
	bb := make([]byte, size, size)
	_, err = ifile.Read(bb)
	if err != nil {
		err = fmt.Errorf(`ifile.Read(bb) error is %s`, err.Error())
		return
	}
	// write
	var ofile *os.File
	ofile, err = os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, appPaths.FMode)
	if err != nil {
		err = fmt.Errorf(`os.OpenFile(dest, os.O_WRONLY | os.O_CREATE, paths.FMode) error is %s`, err.Error())
		return
	}
	if _, err = ofile.Write(bb); err != nil {
		ofile.Close()
		err = fmt.Errorf(`ofile.Write(bb) error is %s`, err.Error())
		return
	}
	if err = ofile.Close(); err != nil {
		err = fmt.Errorf(`ofile.Close() error is %s`, err.Error())
		return
	}
	return
}
