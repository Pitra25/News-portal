package db

import (
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	News       *NewsRepo
	Tags       *TagRepo
	Categories *CategoryRepo
}

func Init(db *sqlx.DB) *DB {
	return &DB{
		News:       NewNewsPG(db),
		Tags:       NewTagsPG(db),
		Categories: NewCategoriesPG(db),
	}
}

func Connection(
	dsn string,
	maxOpenCons, maxIdleCons int,
	connMaxLifetime time.Duration,
) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenCons)
	db.SetMaxIdleConns(maxIdleCons)
	db.SetConnMaxLifetime(connMaxLifetime)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	slog.Debug("db initialization")

	return db, nil
}
