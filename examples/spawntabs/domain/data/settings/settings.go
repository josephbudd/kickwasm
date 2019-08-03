package settings

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/spawntabs/domain/data/filepaths"
	"github.com/josephbudd/kickwasm/examples/spawntabs/domain/types"
	"github.com/josephbudd/kickwasm/examples/spawntabssitepack"
)

// NewApplicationSettings makes a new ApplicationSettings.
// Returns a pointer to the ApplicationSettings and the error.
func NewApplicationSettings() (settings *types.ApplicationSettings, err error) {
	var fpath string
	var contents []byte
	var found bool
	fpath = filepaths.GetShortSettingsPath()
	if contents, found = spawntabssitepack.Contents(fpath); !found {
		err = errors.New(fmt.Sprintf("can't find %q", fpath))
		return
	}
	settings = &types.ApplicationSettings{}
	err = yaml.Unmarshal(contents, settings)
	return
}
