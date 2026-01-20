package controllers

import (
	"go-product-api/internal/dtos/products"
	"go-product-api/internal/services"
	"net/http"
	_ "go-product-api/internal/models"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *services.ProductService
}

func NewProductController(s *services.ProductService) *ProductController {
	return &ProductController{service: s}
}

// GetProducts retrieves all products
// @Summary      Get Products
// @Description  Retrieves all products from the database
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Product
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /products [get]
func (c *ProductController) GetProducts(ctx *gin.Context) {
	products, err := c.service.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// CreateProduct creates a new product
// @Summary      Create Product
// @Description  Creates a product linked to the authenticated user
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request body dtos.CreateProductType true "Product Data"
// @Success      201  {object}  models.Product
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Security     BearerAuth
// @Router       /products [post]
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	userIdVal, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario nao autenticado"})
	}

	userId := uint(userIdVal.(float64)) // converting type any to float64, and after this to uint (to match the id type in database)

	var input dtos.CreateProductType

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
		return
	}

	product, err := c.service.CreateProduct(input.Name, input.Price, userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}