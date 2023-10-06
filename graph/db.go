// db.go

package graph

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() *gorm.DB {
	// get .env variables pass into dsn variable
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Exec("SELECT 1").Error; err != nil {
		log.Fatal(err)
	}

	return db
}
