package app

import (
	"News-portal/internal/db"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error initialization config")
	}
	log.Print("config ok")
}

func initDB() *sqlx.DB {
	log.Print("db initialization start")

	db, err := db.NewPG(
		&db.PostgresConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			UserName: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	)
	if err != nil {
		log.Fatal("failed to initialization db: ", err.Error())
		return nil
	}
	log.Print("db ok")

	return db
}
