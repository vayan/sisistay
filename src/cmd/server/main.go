package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vayan/sisistay/src/model"
)

type Config struct {
	DbUrl      string
	ServerPort string
}

func main() {
	config := Config{
		DbUrl:      "host=postgres port=5432 user=victoria dbname=godb password=secret sslmode=disable",
		ServerPort: "8080",
	}

	db, err := gorm.Open("postgres", config.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Order{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	err = http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
