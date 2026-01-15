package repositories

import (
	"go-product-api/internal/models"

	"gorm.io/gorm"
)

type IProductRepository interface { // Interface ja eh aplicada no productRepository, pela mesma ter as funcoes previstas em IProductRepository
	GetProducts() ([]models.Product, error)

	CreateProduct(product models.Product) (models.Product, error)

}

type productRepository struct { // 'Classe' da repository
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository { // Funcao que retorna uma nova instancia de repository
	return &productRepository { // & aponta para o ponteiro, original
		db: db,
	}
}

func (r *productRepository) GetProducts() ([]models.Product, error) { // Lugar que cria a funcao getProducts e associa a struct de repository, para cumprir o contrato
	var products []models.Product

	result := r.db.Find(&products)

	if (result.Error != nil) {
		return nil, result.Error
	}

	return products, nil
}

func (r *productRepository) CreateProduct(p models.Product) (models.Product, error) { // Lugar que cria a funcao createProduct e associa a struct de repository, para cumprir o contrato
	result := r.db.Create(&p)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return p, nil
}