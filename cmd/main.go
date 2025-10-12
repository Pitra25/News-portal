package main

import (
	"News-portal/internal/app"
	"News-portal/internal/db"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
)

var conFlag = flag.String("config", "./config/config.toml", "config file path")

// colgen@ai:readme
//
//go:generate colgen
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	var cfg *app.Config
	if _, err := toml.DecodeFile(*conFlag, &cfg); err != nil {
		slog.Error("failed to load config", "err", err)
	}

	conn, err := db.Connect(&cfg.Database)
	if err != nil {
		slog.Error("fail init db", "err", err)
		return
	}

	application := app.New(cfg, conn)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err = application.Run(); err != nil {
			slog.Error("fail run app", "err", err)
			return
		}
		quit <- syscall.SIGTERM
	}()

	<-quit

	if err = application.Shutdown(); err != nil {
		slog.Error("fail shutdown app", "err", err)
	}
}
