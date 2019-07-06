package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vayan/sisistay/src/api"
	"github.com/vayan/sisistay/src/model"
	"github.com/vayan/sisistay/src/service"
)

type Env struct {
	GoogleAPIKey string `env:"GOOGLE_API_KEY"`
	PGHost       string `env:"POSTGRES_HOST" envDefault:"postgres"`
	PGUser       string `env:"POSTGRES_USER" envDefault:"victoria"`
	PGPassword   string `env:"POSTGRES_PASSWORD" envDefault:"secret"`
	PGDb         string `env:"POSTGRES_DB" envDefault:"godb"`
	PGPort       string `env:"POSTGRES_PORT" envDefault:"5432"`
	Port         string `env:"PORT" envDefault:"8080"`
}

func main() {
	var e Env

	err := env.Parse(&e)
	if err != nil {
		log.Fatal("Please check your ENVs" + err.Error())
	}

	db, err := gorm.
		Open(
			"postgres",
			fmt.Sprintf(
				"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				e.PGHost, e.PGPort, e.PGUser, e.PGDb, e.PGPassword))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config := api.Config{
		Port: e.Port,
		OrderStorage: model.OrderDatabase{
			Database: db.Debug(),
		},
		RouteFetcher: service.GoogleRouteFetcher{
			APIKey: e.GoogleAPIKey,
		},
	}

	config.InitDB()

	err = config.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
