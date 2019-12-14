package templates

// DataLogLevelsGo is the domain/data/loglevels/loglevels.go template
const DataLogLevelsGo = `package loglevels

// Log levels
const (
	LogLevelNil uint64 = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)
`

// DataFilePathsGo is the domain/data/filepaths.go template.
const DataFilePathsGo = `package filepaths

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
var shortTemplatePath string
var shortSitePath string

// spawnTemplates
var spawnTemplatePath string
var shortSpawnTemplatePath string

// fmode is the applications mode for files.
var fmode = os.FileMode(0666)

// dmode is the applications mode for folders.
var dmode = os.FileMode(0775)

var initerr error

var initialized bool

var testing bool

var appSettingsPath string
var shortAppSettingsPath string

// Testing sets testing to true so that the test db is used not the normal database.
// Returns if in using test db.
func Testing() bool {
	if !initialized {
		testing = true
	}
	return testing
}

func initialize() {
	buildUserHomeDataPath()
	if initerr != nil {
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		initerr = fmt.Errorf("os.Getwd() error is %s", initerr.Error())
		return
	}
	appSettingsPath = filepath.Join(pwd, "{{.FileNames.HTTPDotYAML}}")
	shortAppSettingsPath = "{{.FileNames.HTTPDotYAML}}"
	applicationSitePath = filepath.Join(pwd, "{{.FolderNames.RendererSite}}")
	faviconPath = filepath.Join(applicationSitePath, "{{.FileNames.FavIconDotICO}}")
	templatePath = filepath.Join(applicationSitePath, "{{.FolderNames.Templates}}")
	spawnTemplatePath = filepath.Join(applicationSitePath, "{{.FolderNames.SpawnTemplates}}")
	shortSitePath = "site"
	shortTemplatePath = filepath.Join(shortSitePath, "{{.FolderNames.Templates}}")
	shortSpawnTemplatePath = filepath.Join(shortSitePath, "{{.FolderNames.SpawnTemplates}}")
	initialized = true
}

// GetSettingsPath returns the settings yaml path.
func GetSettingsPath() string {
	if !initialized {
		initialize()
	}
	return appSettingsPath
}

// GetShortSettingsPath returns the settings yaml path.
func GetShortSettingsPath() string {
	if !initialized {
		initialize()
	}
	return shortAppSettingsPath
}

// GetFaviconPath returns the path of the favicon.
func GetFaviconPath() string {
	if !initialized {
		initialize()
	}
	return faviconPath
}

// GetShortSitePath returns the short path to the applications site folder.
func GetShortSitePath() string {
	if !initialized {
		initialize()
	}
	return shortSitePath
}

// GetTemplatePath returns the path of application's markup.
func GetTemplatePath() string {
	if !initialized {
		initialize()
	}
	return templatePath
}

// GetShortTemplatePath returns the short path to the application's template folder.
func GetShortTemplatePath() string {
	if !initialized {
		initialize()
	}
	return shortTemplatePath
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
// It makes the path if necessary.
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
// ex: <script type="text/javascript" src="js/mycode.js" />
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
		userHomeDataPath = filepath.Join(home, "{{.ApplicationName}}_kwfw_tests")
	} else {
		userHomeDataPath = filepath.Join(home, ".{{.ApplicationName}}_kwfw")
	}
}
`

// DataSettingsGo is the /domain/data/settings.go file.
const DataSettingsGo = `package settings

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"{{.ApplicationGitPath}}{{.ImportDomainDataFilepaths}}"
	"{{.SitePackImportPath}}"
)

// ApplicationSettings are the settings for this application.
type ApplicationSettings struct {
	Host string {{.BackTick}}yaml:"host"{{.BackTick}}
	Port uint64 {{.BackTick}}yaml:"port"{{.BackTick}}
}

// NewApplicationSettings makes a new ApplicationSettings.
// Returns a pointer to the ApplicationSettings and the error.
func NewApplicationSettings() (settings *ApplicationSettings, err error) {
	var fpath string
	var contents []byte
	var found bool
	fpath = filepaths.GetShortSettingsPath()
	if contents, found = {{.SitePackPackage}}.Contents(fpath); !found {
		err = fmt.Errorf("can't find %q", fpath)
		return
	}
	settings = &ApplicationSettings{}
	err = yaml.Unmarshal(contents, settings)
	return
}
`
