package main

import (
	"go-product-api/internal/controllers"
	"go-product-api/internal/repositories"
	"go-product-api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	r := gin.Default()

	r.GET("/products", productController.GetProducts)
	r.POST("/products", productController.CreateProduct)

	r.Run(":8080")
}