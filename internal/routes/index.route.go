package routes

import (
	"go-product-api/internal/controllers"
	"go-product-api/internal/repositories"
	"go-product-api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	RegisterProductRoutes(r, productController)
}