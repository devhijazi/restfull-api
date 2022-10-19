package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBSources struct {
	DB *gorm.DB
}

func (DS *DBSources) CreateConnectionDB() {
	log.Println("Create Connection DB...")

	dbHost := os.Getenv("PG_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PG_PASS")
	dbName := os.Getenv("PG_DATABASE")
	sslMode := os.Getenv("PG_SSL_MODE")

	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)

	log.Println("Starting connection with postgres")

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Error to connect with database: %w", err)
		return
	}
	DS.DB = db

}
