package repository

import (
	"log"

	models "github.com/kwul0208/user/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user models.User) (models.User, error)
	FindByEmail(Email string) (models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (ar *authRepository) Create(user models.User) (models.User, error) {
	log.Print("ij")
	err := ar.db.Create(&user).Error

	return user, err
}

func (ar *authRepository) FindByEmail(Email string) (models.User, error) {
	var user models.User

	err := ar.db.Where("email = ?", Email).First(&user).Error

	return user, err
}
