package settings

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btestsitepack"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/domain/data/filepaths"
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
	if contents, found = tp2btestsitepack.Contents(fpath); !found {
		err = fmt.Errorf("can't find %q", fpath)
		return
	}
	settings = &ApplicationSettings{}
	err = yaml.Unmarshal(contents, settings)
	return
}
