package app

import (
	"News-portal/internal/db"
	"News-portal/internal/newsportal"
	"News-portal/internal/rest"
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-pg/pg/v10"
)

type App struct {
	cfg *Config
	db  *pg.DB
	srv *http.Server
}

func New(cfg *Config, dbInit *pg.DB) *App {

	conn := db.NewDB(dbInit)
	manager := newsportal.NewManager(conn)
	router := rest.NewRouter(manager)

	srv := &http.Server{
		Addr:         cfg.Server.ServerAddress(),
		Handler:      router.InitRoutes(),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	return &App{
		cfg: cfg,
		db:  dbInit,
		srv: srv,
	}
}

func (a *App) Run() error {

	if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("fail start server", "err", err.Error())
	}

	return nil
}

func (a *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		slog.Error("forced shutdown", "err", err.Error())
	}

	if err := a.db.Close(); err != nil {
		slog.Error("database connection close failed", "err", err.Error())
	}
	return nil
}
