package db

import (
	"context"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
)

type DB struct {
	Repo *Repo
}

func New(db *pg.DB) *DB {
	return &DB{Repo: NewRepo(db)}
}

func Connect(opt *pg.Options) (*pg.DB, error) {
	db := pg.Connect(opt)

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	db.AddQueryHook(pgdebug.DebugHook{
		Verbose: true,
	})

	return db, nil
}
