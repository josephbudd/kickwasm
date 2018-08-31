package foldercp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
)

// CopyYAML copies the yaml files.
// appPaths is the application paths.
// Param appYamlFilePath is the path of the application yaml file.
// Param panelFilePaths is the slice of each panel file's full path.
func CopyYAML(appPaths *paths.ApplicationPaths, appYamlFilePath string, panelFilePaths []string) error {
	folderPaths := appPaths.GetPaths()
	srcFolder := filepath.Dir(appYamlFilePath)
	srcFile := filepath.Base(appYamlFilePath)
	// copy the application yaml file.
	if err := copyYaml(appPaths, appYamlFilePath, filepath.Join(folderPaths.OutputDotKickYAML, srcFile)); err != nil {
		return err
	}
	// copy the panel yaml files.
	for _, srcPath := range panelFilePaths {
		relPath, err := filepath.Rel(srcFolder, srcPath)
		if err != nil {
			return err
		}
		destPath := filepath.Join(folderPaths.OutputDotKickYAML, relPath)
		if err := copyYaml(appPaths, srcPath, destPath); err != nil {
			return err
		}
	}
	return nil
}

func copyYaml(appPaths *paths.ApplicationPaths, src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), appPaths.DMode); err != nil {
		return fmt.Errorf(`os.MkdirAll(filepath.Dir(dest), appPaths.DMode) error is %s`, err.Error())
	}
	// read
	ifile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf(`os.Open(src) error is %s`, err.Error())
	}
	info, err := ifile.Stat()
	if err != nil {
		return fmt.Errorf(`ifile.Stat() error is %s`, err.Error())
	}
	size := info.Size()
	bb := make([]byte, size, size)
	_, err = ifile.Read(bb)
	if err != nil {
		return fmt.Errorf(`ifile.Read(bb) error is %s`, err.Error())
	}
	// write
	ofile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, appPaths.FMode)
	if err != nil {
		return fmt.Errorf(`os.OpenFile(dest, os.O_WRONLY | os.O_CREATE, paths.FMode) error is %s`, err.Error())
	}
	if _, err = ofile.Write(bb); err != nil {
		ofile.Close()
		return fmt.Errorf(`ofile.Write(bb) error is %s`, err.Error())
	}
	if err := ofile.Close(); err != nil {
		return fmt.Errorf(`ofile.Close() error is %s`, err.Error())
	}
	return nil
}
