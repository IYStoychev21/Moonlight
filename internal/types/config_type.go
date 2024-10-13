package types

type Config struct {
	From string
	To   string
}

type ConfigFile struct {
	ConfigEntreis []Config `toml:"config_locations"`
}
