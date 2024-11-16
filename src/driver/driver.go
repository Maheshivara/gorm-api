package driver

import (
	"fmt"
	"gormCompose/src/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("APP_DB_USER")
	dbPassword := os.Getenv("APP_DB_PASSWORD")
	dbName := os.Getenv("APP_DB_NAME")
	dbPort := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Maceio", dbHost, dbUser, dbPassword, dbName, dbPort)
	newDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return newDb, nil
}

func Migrate() {
	driver, err := New()
	if err != nil {
		log.Fatal(err)
	}
	driver.AutoMigrate(&models.Food{})
}
