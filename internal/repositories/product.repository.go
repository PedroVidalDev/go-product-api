package repositories

import (
	"go-product-api/internal/models"

	"gorm.io/gorm"
)

type IProductRepository interface { // Interface already applied to productRepository, as it has the functions defined in IProductRepository
	GetProducts() ([]models.Product, error)

	CreateProduct(product models.Product) (models.Product, error)

}

type productRepository struct { // Repository 'class'
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository { // Function that returns a new repository instance
	return &productRepository { // & points to the pointer, original
		db: db,
	}
}

func (r *productRepository) GetProducts() ([]models.Product, error) { // Place that creates the getProducts function and associates it with the repository struct, to fulfill the contract
	var products []models.Product

	result := r.db.Preload("User").Find(&products)

	if (result.Error != nil) {
		return nil, result.Error
	}

	return products, nil
}

func (r *productRepository) CreateProduct(p models.Product) (models.Product, error) { // Place that creates the createProduct function and associates it with the repository struct, to fulfill the contract
	result := r.db.Create(&p)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return p, nil
}