package config

type Config struct {
	Proxy Proxy `yaml:"proxy"`
	API   API   `yaml:"api"`
}
