package db

import (
	config "News-portal/configs"
	"News-portal/internal/db/method"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	News       *method.NewsPG
	Tags       *method.TagsPG
	Categories *method.CategoriesPG
}

func New(db *sqlx.DB) *DB {
	return &DB{
		News:       method.NewNewsPG(db),
		Tags:       method.NewTagsPG(db),
		Categories: method.NewCategoriesPG(db),
	}
}

func NewPG(dsn string, cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	slog.Debug("db initialization")

	return db, nil
}
