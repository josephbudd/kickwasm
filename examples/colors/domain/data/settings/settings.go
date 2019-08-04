package settings

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/colors/domain/data/filepaths"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
	"github.com/josephbudd/kickwasm/examples/colorssitepack"
)

// NewApplicationSettings makes a new ApplicationSettings.
// Returns a pointer to the ApplicationSettings and the error.
func NewApplicationSettings() (settings *types.ApplicationSettings, err error) {
	var fpath string
	var contents []byte
	var found bool
	fpath = filepaths.GetShortSettingsPath()
	if contents, found = colorssitepack.Contents(fpath); !found {
		emsg := fmt.Sprintf("can't find %q", fpath)
		err = errors.New(emsg)
		return
	}
	settings = &types.ApplicationSettings{}
	err = yaml.Unmarshal(contents, settings)
	return
}
