package main

import (
	"fmt"
	"go-product-api/internal/controllers"
	"go-product-api/internal/repositories"
	"go-product-api/internal/services"
	database "go-product-api/pkg"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
    }

	dbConnection, err := database.InitDB()
    if err != nil {
        log.Fatal(err)
    }
	
	productRepo := repositories.NewProductRepository(dbConnection)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	r := gin.Default()

	r.GET("/products", productController.GetProducts)
	r.POST("/products", productController.CreateProduct)

	port := fmt.Sprintf(
        ":%s",
        os.Getenv("PORT"),
    )
	r.Run(port)
}