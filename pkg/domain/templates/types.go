package templates

// SettingsGo is the domain/types/settings.go template.
const SettingsGo = `package types

// ApplicationSettings are the settings for this application.
type ApplicationSettings struct {
	Host string {{.BackTick}}yaml:"host"{{.BackTick}}
	Port uint64 {{.BackTick}}yaml:"port"{{.BackTick}}
}
`
