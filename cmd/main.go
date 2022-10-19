package main

import (
	"log"
	"net/http"
	"os"

	"github.com/devhijazi/api/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting server...")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to loading .env file\n", err)
	}

	// Inicializa a DB
	ds := &database.DBSources{}

	// if err := ds.CreateConnectionDB(); err != nil {
	// 	log.Fatalf("Unable to create database connection: %v\n", err)
	// }

	// if err != nil {
	// 	log.Fatalf("Unable to initialize database connection: %v\n", err)
	// }
	ds.CreateConnectionDB()
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
}
