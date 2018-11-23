package settings

import (
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/josephbudd/kickwasm/examples/colors/domain/data/filepaths"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
)

func NewApplicationSettings() (*types.ApplicationSettings, error) {
	fpath := filepaths.GetSettingsPath()
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	l := stat.Size()
	yamlbb := make([]byte, l, l)
	_, err = f.Read(yamlbb)
	if err != nil {
		return nil, err
	}
	v := &types.ApplicationSettings{}
	if err := yaml.Unmarshal(yamlbb, v); err != nil {
		return nil, err
	}
	return v, nil
}

