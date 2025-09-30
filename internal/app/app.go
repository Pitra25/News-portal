package app

import (
	config "News-portal/configs"
	"News-portal/internal/db"
	"News-portal/internal/handler"
	"News-portal/internal/service"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
)

type App struct {
	cfg *config.Config
	db  *sqlx.DB
	srv *http.Server
}

func New(cfg *config.Config) (*App, error) {
	slog.Debug("db url: " + cfg.Database.DatabaseURL())

	dbInit, err := db.NewPG(cfg.Database.DatabaseURL(), cfg)
	if err != nil {
		return nil, err
	}

	db := db.New(dbInit)
	service := service.New(db)
	handler := handler.New(service)

	srv := &http.Server{
		Addr:         cfg.Server.ServerAddress(),
		Handler:      handler.InitRoutes(),
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
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("start server: " + err.Error())
		}
		slog.Info(a.cfg.Server.ServerAddress())
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		slog.Error("forced shutdown: " + err.Error())
	}

	if err := a.db.Close(); err != nil {
		slog.Error("database connection close failed: " + err.Error())
	}

	return nil
}
