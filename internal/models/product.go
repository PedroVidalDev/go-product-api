package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name  string  `json:"name"`
	Price float64 `json:"price"`

	UserId uint `json:"user_id"`
	User User `json:"user,omitempty" gorm:"foreignKey:UserId"`
}