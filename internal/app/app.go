package app

import (
	"News-portal/internal/db"
	"News-portal/internal/newsportal"
	"News-portal/internal/rest"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
)

type App struct {
	cfg *config
	db  *sqlx.DB
	srv *http.Server
}

func New(cfg *config) (*App, error) {

	dbInit, err := db.NewPG(
		cfg.Database.DatabaseURL(),
		cfg.Database.MaxIdleCons,
		cfg.Database.MaxIdleCons,
		cfg.Database.ConnMaxLifetime,
	)
	if err != nil {
		return nil, err
	}

	db := db.New(dbInit)
	service := newsportal.New(db)
	rest := rest.New(service)

	srv := &http.Server{
		Addr:         cfg.Server.ServerAddress(),
		Handler:      rest.InitRoutes(),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	return &App{
		cfg: cfg,
		db:  dbInit,
		srv: srv,
	}, nil
}

func (a *App) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("fail start server", "err", err.Error())
		}
		slog.Info("Address server", "url", "http://"+a.cfg.Server.ServerAddress())
	}()

	<-quit

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
