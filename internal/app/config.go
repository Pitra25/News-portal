package app

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-pg/pg/v10"
)

type Config struct {
	Server   ServerConfig
	Database pg.Options
}

type ServerConfig struct {
	Host            string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func Load(path string) (*Config, error) {
	var conf Config

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &conf, nil
}

func (s *ServerConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
