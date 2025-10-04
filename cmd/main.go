package main

import (
	"News-portal/internal/app"
	"News-portal/internal/db"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

var conFlag = flag.String("config", "./config/config.toml", "config file path")

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("fail loading env", "err", err)
		return
	}

	cfg, err := app.Load(*conFlag)
	if err != nil {
		slog.Error("fail load config", "err", err)
		return
	}

	conn, err := db.Connect(
		cfg.Database.DatabaseURL(),
		cfg.Database.MaxIdleCons,
		cfg.Database.MaxIdleCons,
		cfg.Database.ConnMaxLifetime,
	)
	if err != nil {
		slog.Error("fail init db", "err", err)
		return
	}

	application := app.New(cfg, conn)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := application.Run(); err != nil {
			slog.Error("fail run app", "err", err)
			return
		}
	}()

	<-quit

	if err := application.Shutdown(); err != nil {
		slog.Error("fail shutdown app", "err", err)
	}
}
