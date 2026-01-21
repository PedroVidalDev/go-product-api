package services

import (
	"go-product-api/internal/models"
	"go-product-api/internal/repositories"
)

type ProductService struct {
	repo repositories.IProductRepository
}

func NewProductService(repo repositories.IProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.repo.GetProducts()
}

func  (s *ProductService) CreateProduct(name string, price float64, userId uint) (models.Product, error) {
	if price < 0 {
		price = 0
	}

	newProduct := models.Product {
		Name: name,
		Price: price,
		UserId: userId,
	}

	return s.repo.CreateProduct(newProduct)
}

func (s *ProductService) UpdateProduct(id string, name string, price float64, userId uint) (models.Product, error) {
	if (price < 0) {
		price = 0
	}

	updatedProduct := models.Product {
		Name: name,
		Price: price,
		UserId: userId,
	}

	return s.repo.UpdateProduct(id, updatedProduct)
}

func (s *ProductService) DeleteProduct(id string, userId uint) error {
	if err := s.repo.DeleteProduct(id, userId); err != nil {
		return err
	}

	return nil
}