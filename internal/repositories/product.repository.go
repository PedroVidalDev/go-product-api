package repositories

import "go-product-api/internal/models"

type IProductRepository interface { // Interface ja eh aplicada no productRepository, pela mesma ter as funcoes previstas em IProductRepository
	GetProducts() ([]models.Product, error)

	CreateProduct(product models.Product) (models.Product, error)

}

type productRepository struct { // 'Classe' da repository
	db []models.Product
	nextId int
}

func NewProductRepository() IProductRepository { // Funcao que retorna uma nova instancia de repository
	return &productRepository { // & aponta para o ponteiro, original
		db: []models.Product{},
		nextId: 1,
	}
}

func (r *productRepository) GetProducts() ([]models.Product, error) { // Lugar que cria a funcao getProducts e associa a struct de repository, para cumprir o contrato
	return r.db, nil
}

func (r *productRepository) CreateProduct(p models.Product) (models.Product, error) { // Lugar que cria a funcao createProduct e associa a struct de repository, para cumprir o contrato
	p.Id = r.nextId
	r.nextId++
	r.db = append(r.db, p)
	return p, nil
}