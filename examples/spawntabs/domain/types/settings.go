package types

// ApplicationSettings are the settings for this application.
type ApplicationSettings struct {
	Host string `yaml:"host"`
	Port uint64 `yaml:"port"`
}
