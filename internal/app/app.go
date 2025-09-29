package app

import (
	"News-portal/internal/app/server"
	"News-portal/internal/db"
	"News-portal/internal/handler"
	"News-portal/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Run() {
	initConfig()

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env: ", err.Error())
	}

	dbInit := initDB()

	db := db.New(dbInit)
	service := service.New(*db)
	handler := handler.New(*service)

	srv := startApp(handler)
	stopApp(srv, dbInit)

	log.Print("Server stop")
}

func startApp(h *handler.Handler) *server.Server {
	srv := new(server.Server)

	go func() {
		srv.Start(viper.GetString("server.port"), h.InitRoutes())
	}()

	log.Print("App Started")

	return srv
}

func stopApp(srv *server.Server, db *sqlx.DB) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("App shutting down")

	srv.Stop()

	if err := db.Close(); err != nil {
		log.Fatal("An error occurred while closing the database connection: ", err.Error())
	}
}
