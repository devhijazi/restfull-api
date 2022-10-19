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

	_ "github.com/devhijazi/api/cmd/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	initDotEnv()
	initDatabase()
	initSwagger(&gin.RouterGroup{})
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

func initSwagger(rg *gin.RouterGroup) {
	serverPort := fmt.Sprintf(":%s", os.Getenv("PORT"))

	docsURL := fmt.Sprintf("http://localhost%s/docs/doc.json", serverPort)

	rg.GET("/cmd/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(docsURL)))

}

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

	user.POST("/", controllers.CreateUser)
	user.GET("/:id", controllers.GetUser)
}
