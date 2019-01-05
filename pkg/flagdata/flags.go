package flagdata

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Flags is some of the flags used to create this application.
type Flags struct {
	FlagCC            bool
	YAMLStartFileName string
	KWVersionBreaking int // Backwards compatibility.
	KWVersionFeature  int // Added features. Still backwards compatible.
	KWVersionPatch    int // Bug fix. No added features. Still backwards compatible.
}

// SaveFlags save the kick flags.
// Param fPath is the flags file path.
// Param locationsFlag is the LocationsFlag.
func SaveFlags(fPath string, locationsFlag bool, vBreaking, vFeature, vPatch int, yamlStartFileName string) (err error) {
	flags := &Flags{
		locationsFlag,
		yamlStartFileName,
		vBreaking,
		vFeature,
		vPatch,
	}
	var bb []byte
	bb, err = yaml.Marshal(flags)
	if err != nil {
		return
	}
	var f *os.File
	f, err = os.Create(fPath)
	if err != nil {
		return
	}
	if _, err = f.Write(bb); err != nil {
		f.Close()
		return
	}
	err = f.Close()
	return
}

// GetFlags returns flags from a file saved with SaveFlags.
func GetFlags(fPath string) (appFlags *Flags, err error) {
	var f *os.File
	f, err = os.Open(fPath)
	if err != nil {
		return
	}
	var info os.FileInfo
	info, err = f.Stat()
	if err != nil {
		return
	}
	bb := make([]byte, info.Size())
	_, err = f.Read(bb)
	if err != nil {
		return
	}
	appFlags = &Flags{}
	yaml.Unmarshal(bb, appFlags)
	return
}
