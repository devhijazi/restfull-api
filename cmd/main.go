package main

import (
	"fmt"
	"log"
	"os"

	"github.com/devhijazi/api/internal/controllers"
	"github.com/devhijazi/api/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	initDotEnv()
	initDatabase()
	initSwagger()
	initServer()
}

func initDotEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to loading .env file\n", err)
	}
}

func initDatabase() {
	database.Init()
}

func initSwagger() {}

func initServer() {
	gin.SetMode(gin.ReleaseMode)

	log.Println("Starting server...")

	r := gin.New()

	r.Use(cors.New(cors.Config{AllowAllOrigins: true}))

	initRoutes(r)

	serverPort := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Printf("Api started on \"http://localhost%s\"", serverPort)

	r.Run(serverPort)
}

func initRoutes(r *gin.Engine) {
  user := r.Group("/users")
  
	user.GET("/:id", controllers.GetUser)
}
