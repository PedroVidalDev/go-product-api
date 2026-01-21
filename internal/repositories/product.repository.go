package repositories

import (
	"go-product-api/internal/models"

	"gorm.io/gorm"
)

type IProductRepository interface { // Interface already applied to productRepository, as it has the functions defined in IProductRepository
	GetProducts() ([]models.Product, error)

	CreateProduct(product models.Product) (models.Product, error)

	UpdateProduct(id string, product models.Product) (models.Product, error)

	DeleteProduct(id string, userId uint) error
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

func (r *productRepository) UpdateProduct(id string, p models.Product) (models.Product, error) {
	var existingProduct models.Product

	if err := r.db.Where("id = ? AND user_id = ?", id, p.UserId).First(&existingProduct).Error; err != nil {
		return models.Product{}, err
	}

	existingProduct.Name = p.Name
	existingProduct.Price = p.Price

	if err := r.db.Save(&existingProduct).Error; err != nil {
		return models.Product{}, err
	}

	return existingProduct, nil
}

func (r *productRepository) DeleteProduct(id string, userId uint) error {
	var product models.Product

	if err := r.db.Where("id = ? AND user_id = ?", id, userId).First(&product).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}