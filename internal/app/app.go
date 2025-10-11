package app

import (
	"News-portal/internal/db"
	"News-portal/internal/newsportal"
	"News-portal/internal/rest"
	"News-portal/internal/rpc"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/vmkteam/zenrpc/v2"
)

type (
	App struct {
		cfg *Config
		db  *pg.DB
		srv *http.Server
	}

	Config struct {
		Server struct {
			Host            string
			Port            string
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
			ShutdownTimeout time.Duration
		}
		ServerCRP struct {
			ExposeSMD              bool
			AllowCORS              bool
			DisableTransportChecks bool
		}
		Database pg.Options
	}
)

func New(cfg *Config, dbInit *pg.DB) *App {
	conn := db.New(dbInit)
	manager := newsportal.NewManager(conn)
	router := rest.NewRouter(manager)

	srv := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
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
	go func() {
		rpc.Run(zenrpc.NewServer(zenrpc.Options{
			ExposeSMD:              a.cfg.ServerCRP.ExposeSMD,
			AllowCORS:              a.cfg.ServerCRP.AllowCORS,
			DisableTransportChecks: a.cfg.ServerCRP.DisableTransportChecks,
		}))
	}()

	if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("fail to shutdown server: %w", err)
	}

	if err := a.db.Close(); err != nil {
		return fmt.Errorf("fail to close database: %w", err)
	}

	return nil
}
