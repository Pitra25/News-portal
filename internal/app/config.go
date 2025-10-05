package app

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-pg/pg/v10"
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

func (d *DatabaseConfig) DBOptions() *pg.Options {
	if d.Host == "" || d.Port == "" || d.DBName == "" {
		slog.Error("missing required database configuration",
			"host", d.Host,
			"port", d.Port,
			"dbname", d.DBName,
		)
		return nil
	}

	return &pg.Options{
		Addr:     d.Host + ":" + d.Port,
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: d.DBName,
	}
}

func (s *ServerConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
