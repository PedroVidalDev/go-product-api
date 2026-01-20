package controllers

import (
	dtos "go-product-api/internal/dtos/user"
	"go-product-api/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "go-product-api/internal/models"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(s *services.AuthService) *AuthController {
	return &AuthController{service: s}
}

// Register creates a new user account
// @Summary      Register User
// @Description  Creates a new user account in the system
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dtos.CreateUserType true "User Registration Data"
// @Success      201  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var input dtos.CreateUserType

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
		return
	}

	user, err := c.service.Register(input.Name, input.Email, input.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// Login authenticates a user
// @Summary      Login User
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dtos.LoginType true "User Login Credentials"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var input dtos.LoginType

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
		return
	}

	token, err := c.service.Login(input.Email, input.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}