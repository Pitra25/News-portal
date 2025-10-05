package db

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
)

var (
	connDB *pg.DB
	opt    = pg.Options{
		Addr:     host + ":" + port,
		User:     user,
		Password: password,
		Database: dbname,
	}
)

func TestMain(m *testing.M) {
	conn, err := Connect(&opt)
	if err != nil {
		panic(err)
	}
	connDB = conn

	exitCode := m.Run()
	err = connDB.Close()
	if err != nil {
		slog.Error("Error closing DB connection.", "err", err)
		return
	}
	os.Exit(exitCode)
}

func TestConnect(t *testing.T) {

	tests := []struct {
		name    string
		args    *pg.Options
		want    *pg.DB
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Test new postrgesql",
			args:    &opt,
			want:    &pg.DB{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq, err := Connect(tt.args)
			if !tt.wantErr(t, err, fmt.Sprint("error", err)) {
				return
			}

			assert.NotNil(t, pq)
		})
	}
}
