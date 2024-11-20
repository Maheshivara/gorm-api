package driver

import (
	"fmt"
	"gormCompose/src/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func config() error {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("APP_DB_USER")
	dbPassword := os.Getenv("APP_DB_PASSWORD")
	dbName := os.Getenv("APP_DB_NAME")
	dbPort := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Maceio", dbHost, dbUser, dbPassword, dbName, dbPort)
	newDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = newDb
	return nil
}

func Migrate() {
	err := config()
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.Food{})
}
func Get() *gorm.DB {
	return DB
}
