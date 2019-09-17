package settings

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/push/domain/data/filepaths"
	"github.com/josephbudd/kickwasm/examples/pushsitepack"
)

// ApplicationSettings are the settings for this application.
type ApplicationSettings struct {
	Host string `yaml:"host"`
	Port uint64 `yaml:"port"`
}

// NewApplicationSettings makes a new ApplicationSettings.
// Returns a pointer to the ApplicationSettings and the error.
func NewApplicationSettings() (settings *ApplicationSettings, err error) {
	var fpath string
	var contents []byte
	var found bool
	fpath = filepaths.GetShortSettingsPath()
	if contents, found = pushsitepack.Contents(fpath); !found {
		emsg := fmt.Sprintf("can't find %q", fpath)
		err = errors.New(emsg)
		return
	}
	settings = &ApplicationSettings{}
	err = yaml.Unmarshal(contents, settings)
	return
}
