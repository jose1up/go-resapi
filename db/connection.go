package db

import (
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

// parameters for postgres connection

var DSN = "host=localhost user=postgres password=1234 dbname=gorm port=5432"
var DB *gorm.DB

// functions for postgres connection
func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connection success")
	}

}
