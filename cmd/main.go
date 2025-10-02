package main

import (
	"News-portal/internal/app"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("fail loading env", "err", err)
		return
	}

	cfg, err := app.Load("./config/config.toml")
	if err != nil {
		slog.Error("fail load config", "err", err)
		return
	}

	application, err := app.New(cfg)
	if err != nil {
		slog.Error("fail init app", "err", err)
		return
	}

	if err := application.Run(); err != nil {
		slog.Error("fail run app", "err", err)
		return
	}
}
