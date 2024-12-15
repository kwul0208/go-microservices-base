package repository

import (
	"log"

	"github.com/kwul0208/user/pkg/db"
	models "github.com/kwul0208/user/pkg/model"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	FindByEmail(Email string) (models.User, error)
}

type authRepository struct {
	db db.Handler
}

func NewAuthRepository(db db.Handler) *authRepository {
	return &authRepository{db}
}

func (ar *authRepository) Register(user models.User) (models.User, error) {
	log.Print("ij")
	err := ar.db.DB.Create(&user).Error

	return user, err
}

func (ar *authRepository) FindByEmail(Email string) (models.User, error) {
	var user models.User

	err := ar.db.DB.Where("email = ?", Email).First(&user).Error

	return user, err
}
