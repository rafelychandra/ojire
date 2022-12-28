package config

type InterfaceConfig interface {
	LoadConfig() (*Config, error)
}
