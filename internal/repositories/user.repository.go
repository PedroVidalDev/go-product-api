package repositories

import (
	"go-product-api/internal/models"

	"gorm.io/gorm"
)

type IUserRepository interface { // Interface ja eh aplicada no productRepository, pela mesma ter as funcoes previstas em IProductRepository
	FindByEmail(email string) (models.User, error)
	CreateUser(p models.User) (models.User, error)
}

type userRepository struct { // 'Classe' da repository
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository { // Funcao que retorna uma nova instancia de repository
	return &userRepository { // & aponta para o ponteiro, original
		db: db,
	}
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	result := r.db.Where("email = ?", email).First(&user)

	if (result.Error != nil) {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) CreateUser(p models.User) (models.User, error) {
	result := r.db.Create(&p)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return p, nil
}