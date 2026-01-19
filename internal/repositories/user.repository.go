package repositories

import (
	"go-product-api/internal/models"

	"gorm.io/gorm"
)

type IUserRepository interface { // Interface already applied to userRepository, as it has the functions defined in IUserRepository
	FindByEmail(email string) (models.User, error)
	CreateUser(p models.User) (models.User, error)
}

type userRepository struct { // Repository 'class'
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository { // Function that returns a new repository instance
	return &userRepository { // & points to the pointer, original
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