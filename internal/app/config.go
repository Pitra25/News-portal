package app

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type config struct {
	Server   serverConfig   `toml:"server"`
	Database databaseConfig `toml:"database"`
}

type serverConfig struct {
	Host            string        `toml:"host"`
	Port            string        `toml:"port"`
	ReadTimeout     time.Duration `toml:"read_timeout"`
	WriteTimeout    time.Duration `toml:"write_timeout"`
	ShutdownTimeout time.Duration `toml:"shutdown_timeout"`
}

type databaseConfig struct {
	Host            string        `toml:"host"`
	Port            string        `toml:"port"`
	DBName          string        `toml:"dbName"`
	SSLMode         string        `toml:"sslMode"`
	MaxOpenCons     int           `toml:"max_open_cons"`
	MaxIdleCons     int           `toml:"max_idle_cons"`
	ConnMaxLifetime time.Duration `toml:"conn_max_lifetime"`
}

func Load(path string) (*config, error) {
	var conf config

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &conf, nil
}

func (d *databaseConfig) DatabaseURL() string {
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

func (s *serverConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
