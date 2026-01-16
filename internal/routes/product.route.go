package routes

import (
	"go-product-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, controller *controllers.ProductController ) {
	productGroup := r.Group("/products")

	{
		productGroup.GET("", controller.GetProducts)
		productGroup.POST("", controller.CreateProduct)

	}
}