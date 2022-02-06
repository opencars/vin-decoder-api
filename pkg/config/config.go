package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Server Server   `yaml:"server"`
	Log    Log      `yaml:"log"`
	DB     Database `yaml:"database"`
	NATS   NATS     `yaml:"nats"`
}

// Server represents settings for creating http server.
type Server struct {
	ShutdownTimeout Duration `yaml:"shutdown_timeout"`
	ReadTimeout     Duration `yaml:"read_timeout"`
	WriteTimeout    Duration `yaml:"write_timeout"`
	IdleTimeout     Duration `yaml:"idle_timeout"`
}

// Log represents settings for application logger.
type Log struct {
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

// NATS contains configuration details for application event API.
type NATS struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Address returns calculated address for connecting to NATS.
func (nats *NATS) Address() string {
	if nats.User != "" && nats.Password != "" {
		return fmt.Sprintf("nats://%s:%s@%s:%d", nats.User, nats.Password, nats.Host, nats.Port)
	}

	return fmt.Sprintf("nats://%s:%d", nats.Host, nats.Port)
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	var config Settings

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
