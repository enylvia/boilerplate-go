package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize PostgreSQL database

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=1sampai10 dbname=belajar port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	return db
}

// Initialize Redis connection
