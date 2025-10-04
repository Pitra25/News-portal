package db

import (
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
		News:       NewNews(db),
		Tags:       NewTags(db),
		Categories: NewCategory(db),
	}
}

func Connect(
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

	return db, nil
}
