package main

import (
	config "News-portal/configs"
	"News-portal/internal/app"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("loading env: " + err.Error())
		return
	}

	cfg, err := config.Load("./configs/config.toml")
	if err != nil {
		slog.Error("load config: " + err.Error())
		return
	}

	slog.Info("cfg: ", cfg)

	application, err := app.New(cfg)
	if err != nil {
		slog.Error("init application: " + err.Error())
		return
	}

	if err := application.Run(); err != nil {
		slog.Error("run application: " + err.Error())
		return
	}
}
