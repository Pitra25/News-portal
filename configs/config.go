package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
}

type ServerConfig struct {
	Host            string        `toml:"host"`
	Port            string        `toml:"port"`
	ReadTimeout     time.Duration `toml:"read_timeout"`
	WriteTimeout    time.Duration `toml:"write_timeout"`
	ShutdownTimeout time.Duration `toml:"shutdown_timeout"`
}

type DatabaseConfig struct {
	Host            string        `toml:"host"`
	Port            string        `toml:"port"`
	DBName          string        `toml:"dbName"`
	SSLMode         string        `toml:"sslMode"`
	MaxOpenConns    int           `toml:"max_open_conns"`
	MaxIdleConns    int           `toml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `toml:"conn_max_lifetime"`
}

func Load(path string) (*Config, error) {
	var config Config

	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &config, nil
}

func (d *DatabaseConfig) DatabaseURL() string {
	if d.Host == "" || d.Port == "" || d.DBName == "" {
		slog.Error("missing required database configuration",
			"host", d.Host,
			"port", d.Port,
			"dbname", d.DBName,
		)
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port,
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		d.DBName, d.SSLMode,
	)
}

func (s *ServerConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
