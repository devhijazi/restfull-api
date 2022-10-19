package database

import (
	"fmt"
	"log"
	"os"

	"github.com/devhijazi/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	log.Println("Create Connection DB...")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_SSL_MODE"),
	)

	var err error
  
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successful database connection.")

	db.AutoMigrate(&models.User{})

	defer CloseDb()
}

func GetDb() *gorm.DB {
	return db
}

func CloseDb() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	sqlDB.Close()
}
