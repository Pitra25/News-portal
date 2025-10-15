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
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vmkteam/rpcgen/v2"
	"github.com/vmkteam/rpcgen/v2/golang"
	"github.com/vmkteam/zenrpc/v2"
)

type (
	App struct {
		cfg  *Config
		db   *db.DB
		echo *echo.Echo
	}

	Config struct {
		Server struct {
			Host            string
			Port            string
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
			ShutdownTimeout time.Duration
		}
		Database pg.Options
	}
)

func New(cfg *Config, db *db.DB) *App {
	return &App{
		cfg:  cfg,
		db:   db,
		echo: echo.New(),
	}
}

func (a *App) registerRoutes() {
	manager := newsportal.NewManager(a.db)
	router := rest.NewRouter(manager)

	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())

	router.AddRouter(a.echo)
}

func (a *App) registerRPC() {
	manager := newsportal.NewManager(a.db)
	srv := rpc.New(manager)
	gen := rpcgen.FromSMD(srv.SMD())

	a.echo.Any("/v1/rpc/", echo.WrapHandler(http.Handler(srv)))
	a.echo.Any("/v1/rpc/doc/", echo.WrapHandler(http.HandlerFunc(zenrpc.SMDBoxHandler)))
	a.echo.Any("/v1/rpc/client.ts", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.TSClient(nil)))))
	a.echo.Any("/v1/rpc/client.go", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.GoClient(golang.Settings{})))))
}

func (a *App) Run() error {
	a.registerRoutes()
	a.registerRPC()

	address := a.cfg.Server.Host + ":" + a.cfg.Server.Port
	if err := a.echo.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := a.echo.Shutdown(ctx); err != nil {
		return fmt.Errorf("fail to shutdown server: %w", err)
	}

	if err := a.db.Close(); err != nil {
		return fmt.Errorf("fail to close database: %w", err)
	}

	return nil
}
