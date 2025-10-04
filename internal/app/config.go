package app

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host            string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type DatabaseConfig struct {
	Host            string
	Port            string
	DBName          string
	SSLMode         string
	MaxOpenCons     int
	MaxIdleCons     int
	ConnMaxLifetime time.Duration
}

func Load(path string) (*Config, error) {
	var conf Config

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, fmt.Errorf("failed to load Config: %w", err)
	}

	return &conf, nil
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
	return fmt.Sprintf("http://%s:%s/api", s.Host, s.Port)
}
