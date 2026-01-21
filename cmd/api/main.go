package main

import (
	"fmt"
	"go-product-api/internal/routes"
	database "go-product-api/pkg/db"
	"log"
	"os"
	docs "go-product-api/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           Product API
// @version         1.0
// @description     This is a study project developed to learn the Go (Golang) programming language by building a simple REST API for product management.
// @termsOfService  http://swagger.io/terms/

// @contact.name    Pedro Vidal
// @contact.email   pedrohvidals@gmail.com

// @host            localhost:8080
// @BasePath        /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
    }

	dbConnection, err := database.InitDB()
    if err != nil {
        log.Fatal(err)
    }

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	routes.SetupRoutes(r, dbConnection)

	port := fmt.Sprintf(
        ":%s",
        os.Getenv("PORT"),
    )
	r.Run(port)
}