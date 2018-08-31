package flagdata

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Flags is some of the flags used to create this application.
type Flags struct {
	FlagAbout         bool
	FlagCC            bool
	YAMLStartFileName string
}

// SaveFlags save the kick flags.
// Param fPath is the flags file path.
// Param aboutFlag is the AboutFlag.
// Param locationsFlag is the LocationsFlag.
func SaveFlags(fPath string, aboutFlag, locationsFlag bool, yamlStartFileName string) error {
	flags := &Flags{
		aboutFlag,
		locationsFlag,
		yamlStartFileName,
	}
	bb, err := yaml.Marshal(flags)
	if err != nil {
		return err
	}
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	if _, err := f.Write(bb); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

// GetFlags returns flags from a file saved with SaveFlags.
func GetFlags(fPath string) (*Flags, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	bb := make([]byte, info.Size())
	_, err = f.Read(bb)
	if err != nil {
		return nil, err
	}
	appFlags := &Flags{}
	yaml.Unmarshal(bb, appFlags)
	return appFlags, nil
}
