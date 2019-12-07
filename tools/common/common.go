package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/flagdata"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/slurp"
)

// Private and public data.
const (
	rekickwasmErrorFormat = `Oops! %s will not run because there is a rekickwasm/ folder open.
When you are finished using rekickwasm make sure to **rekickwasm -x** to remove the rekickwasm/ folder.

`
	goModErrorFormat = `Oops! %s will not run because there already is a go.mod file.

`

	wrongVersionFormat = `Oops!
  
This application framework was built with kickwasm version %d but %s only works with builds from kickwasm version %d on.
  `

	UseItAnyWhere = `  Use it anywhere in your frameworks's source code folder.`
	UseItInRoot   = `  Use it in your frameworks's root source code folder.`
)

// errors
var (
	ErrNoKickWasm       = errors.New("cant find kickwas.yaml")
	ErrRekickwasmExists = errors.New("rekickwasm exists")
	ErrWrongVersion     = errors.New("wrong version")
)

// PrintRekickwasmError prints the rekickwasm error message.
func PrintRekickwasmError(applicationName string) {
	fmt.Printf(rekickwasmErrorFormat, applicationName)
}

// PrintGoModFileError prints the go.mod file error message.
func PrintGoModFileError(applicationName string) {
	fmt.Printf(goModErrorFormat, applicationName)
}

// IsRootFolder returns if the current path is the application root folder.
func IsRootFolder() (root string, is bool, err error) {
	if root, err = os.Getwd(); err != nil {
		return
	}
	is = isRoot(root)
	return
}

// FindRoot searches for and changed the working directory to the application root folder.
func FindRoot() (root string, err error) {
	if root, err = os.Getwd(); err != nil {
		return
	}
	if isRoot(root) {
		return
	}
	volumeName := filepath.VolumeName(root)
	lastRoot := root
	for {
		root = filepath.Dir(root)
		if volumeName == root {
			err = errors.New("can't find root: volumeName == root")
			return
		}
		if lastRoot == root {
			err = errors.New("can't find root: lastRoot == root")
			return
		}
		if isRoot(root) {
			err = os.Chdir(root)
			return
		}
		lastRoot = root
	}
}

// Version returns the application version information.
func Version(applicationName string, versionBreaking, versionFeature, versionPatch uint64) (version string) {
	versions := []string{
		applicationName + ":",
		fmt.Sprintf(`  Version: %d.%d.%d`, versionBreaking, versionFeature, versionPatch),
		fmt.Sprintf(`  Compatible with kickwasm version: %d.%d.%d`, versionBreaking, versionFeature, versionPatch),
		fmt.Sprintf(`  Will only work with frameworks created with kickwasm version %d or greater.`, versionBreaking),
	}
	version = strings.Join(versions, "\n")
	return
}

// PrintWrongVersion prints the wrong version message.
func PrintWrongVersion(applicationName string, frameworkversion, minimumversion int) {
	fmt.Printf(wrongVersionFormat, frameworkversion, applicationName, minimumversion)
}

// ApplicationInfo is
func ApplicationInfo(rootFolderPath string) (appInfo *slurp.ApplicationInfo, err error) {
	infoPath := filepath.Join(rootFolderPath, ".kickwasm", "yaml", "kickwasm.yaml")
	appInfo, err = slurp.GetApplicationInfo(infoPath)
	return
}

// AppKickwasmVersion returns the kickwasm version that this application was created with.
func AppKickwasmVersion() (version int) {
	folderNames := paths.GetFolderNames()
	fileNames := paths.GetFileNames()
	path := filepath.Join(folderNames.DotKickwasm, fileNames.FlagDotYAML)
	var flags *flagdata.Flags
	var err error
	if flags, err = flagdata.GetFlags(path); err != nil {
		return
	}
	version = flags.KWVersionBreaking
	return
}

func isRoot(path string) (is bool) {
	if is = PathFound(filepath.Join(path, ".kickwasm")); !is {
		return
	}
	if is = PathFound(filepath.Join(path, ".kickstore")); !is {
		return
	}
	if is = PathFound(filepath.Join(path, "site")); !is {
		return
	}
	if is = PathFound(filepath.Join(path, "rendererprocess")); !is {
		return
	}
	if is = PathFound(filepath.Join(path, "mainprocess")); !is {
		return
	}
	return
}

// HaveRekickwasmFolder returns if this folder has a rekickwasm/ folder.
func HaveRekickwasmFolder(root string) (have bool) {
	have = PathFound(filepath.Join(root, "rekickwasm"))
	return
}

// HaveGoModFile returns if ./go.mod file exists.
func HaveGoModFile(root string) (have bool) {
	have = PathFound(filepath.Join(root, "go.mod"))
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
