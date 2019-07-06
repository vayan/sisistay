package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vayan/sisistay/src/api"
	"github.com/vayan/sisistay/src/model"
	"github.com/vayan/sisistay/src/service"
)

func main() {
	var googleApiKey = os.Getenv("GOOGLE_API_KEY")

	if googleApiKey == "" {
		log.Fatal("Please set the GOOGLE_API_KEY env in a .env file at the root of the project")
	}

	db, err := gorm.Open("postgres", "host=postgres port=5432 user=victoria dbname=godb password=secret sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config := api.Config{
		OrderStorage: model.OrderDatabase{
			Database: db,
		},
		RouteFetcher: service.GoogleRouteFetcher{
			APIKey: googleApiKey,
		},
	}

	config.InitDB()

	err = config.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
