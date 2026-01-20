package routes

import (
	"go-product-api/internal/controllers"
	"go-product-api/internal/repositories"
	"go-product-api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	authRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	RegisterAuthRoutes(r, authController)
	RegisterProductRoutes(r, productController)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}