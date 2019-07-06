package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vayan/sisistay/src/api"
	"github.com/vayan/sisistay/src/model"
)

func main() {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=victoria dbname=godb password=secret sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config := api.Config{
		OrderStorage: model.OrderDatabase{
			Database: db,
		},
	}

	config.InitDB()

	err = config.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
