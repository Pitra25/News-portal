package db

import (
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
)

var connDB *sqlx.DB

var dsn = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	host, port, user, password, dbname, sslmode,
)

func TestMain(m *testing.M) {
	conn, err := Connect(dsn, 5, 15, 5)
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
	type args struct {
		dsn             string
		maxOpenCons     int
		maxIdleCons     int
		connMaxLifetime time.Duration
	}

	tests := []struct {
		name    string
		args    args
		want    *sqlx.DB
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test new postrgesql",
			args: args{
				dsn:             dsn,
				maxOpenCons:     1,
				maxIdleCons:     1,
				connMaxLifetime: time.Minute,
			},
			want:    &sqlx.DB{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq, err := Connect(tt.args.dsn, tt.args.maxOpenCons, tt.args.maxIdleCons, tt.args.connMaxLifetime)
			if !tt.wantErr(t, err, fmt.Sprint("error", err)) {
				return
			}

			assert.NotNil(t, pq)
		})
	}
}
