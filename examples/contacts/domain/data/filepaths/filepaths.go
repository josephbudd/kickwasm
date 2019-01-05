package filepaths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var userHomeDataPath string

// applicationSitePath is where the application settings yaml file is.
var applicationSitePath string

// faviconPath is where the favicon is.
var faviconPath string

// templatePath is where the view is.
var templatePath string

// fmode is the applications mode for files.
var fmode = os.FileMode(0666)

// dmode is the applications mode for folders.
var dmode = os.FileMode(0775)

var initerr error

var initialized bool

var testing bool

var appSettingsPath string

// Testing sets testing to true so that the test db is used not the normal database.
// Returns if in using test db.
func Testing() bool {
	if !initialized {
		testing = true
	}
	return testing
}

func initialize() {
	initialized = true
	buildUserHomeDataPath()
	if initerr != nil {
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		initerr = fmt.Errorf("os.Getwd() error is %s", initerr.Error())
		return
	}
	appSettingsPath = filepath.Join(pwd, "http.yaml")
	applicationSitePath = filepath.Join(pwd, "site")
	faviconPath = filepath.Join(applicationSitePath, "favicon.ico")
	templatePath = filepath.Join(applicationSitePath, "templates")
}

// GetSettingsPath returns the settings yaml path.
func GetSettingsPath() string {
	if !initialized {
		initialize()
	}
	return appSettingsPath
}

// GetFaviconPath returns the path of the favicon.
func GetFaviconPath() string {
	if !initialized {
		initialize()
	}
	return faviconPath
}

// GetTemplatePath returns the path of application's markup.
func GetTemplatePath() string {
	if !initialized {
		initialize()
	}
	return templatePath
}

// GetFmode returns the file mode for files.
func GetFmode() os.FileMode {
	return fmode
}

// GetDmode returns the file mode for directories.
func GetDmode() os.FileMode {
	return dmode
}

// BuildUserSubFoldersPath builds a sub folder path in the user's home folder.
// It makes the path if neccessary.
// Param sfpath [in] is the subfolder path.
// Returns the folder path.
func BuildUserSubFoldersPath(sfpath string) (string, error) {
	if !initialized {
		initialize()
	}
	if initerr != nil {
		return userHomeDataPath, initerr
	}
	path := filepath.Join(userHomeDataPath, sfpath)
	err := os.MkdirAll(path, dmode)
	return path, err
}

// BuildRendererPath returns renderer path to src.
// Param src comes from main.html
// ex: <script type="text/javascript" src="js/messenger.js" />
func BuildRendererPath(src string) string {
	if !initialized {
		initialize()
	}
	return filepath.Join(applicationSitePath, src)
}

func buildUserHomeDataPath() {
	var home string
	switch runtime.GOOS {
	case "darwin":
		home = os.Getenv("HOME")
	case "windows":
		home = os.Getenv("LOCALAPPDATA")
	default:
		home = os.Getenv("HOME")
	}
	if testing {
		userHomeDataPath = filepath.Join(home, "contacts_kwfw_tests")
	} else {
		userHomeDataPath = filepath.Join(home, ".contacts_kwfw")
	}
	if err := os.MkdirAll(userHomeDataPath, dmode); err != nil {
		initerr = fmt.Errorf("os.MkdirAll(userHomeDataPath, dmode) error is %s", initerr.Error())
	}
}

