package routes

import (
	"go-product-api/internal/controllers"
	"go-product-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, controller *controllers.ProductController ) {
	productGroup := r.Group("/products")
	productGroup.Use(middlewares.AuthMiddleware())

	productGroup.GET("", controller.GetProducts)
	productGroup.POST("", controller.CreateProduct)
	productGroup.PUT("/:id", controller.UpdateProduct)
	productGroup.DELETE("/:id", controller.DeleteProduct)
}