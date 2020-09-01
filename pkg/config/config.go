package config

import (
	"github.com/BurntSushi/toml"
)

// Settings is decoded configuration file.
type Settings struct {
	DB Database `toml:"database"`
}

// Database contains configuration details for database.
type Database struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"username"`
	Password string `toml:"password"`
	Name     string `toml:"database"`
	SSLMode  string `toml:"ssl_mode"`
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	config := &Settings{}
	if _, err := toml.DecodeFile(path, config); err != nil {
		return nil, err
	}

	return config, nil
}
